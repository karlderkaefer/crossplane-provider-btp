package main

import (
	"context"
	"crypto/tls"
	"os"
	"path/filepath"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"flag"

	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/sap/crossplane-provider-btp/btp"
	"github.com/sap/crossplane-provider-btp/config"
	"github.com/sap/crossplane-provider-btp/internal/clients/tfclient"
	"github.com/sap/crossplane-provider-btp/internal/features"
	"gopkg.in/alecthomas/kingpin.v2"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/feature"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/sap/crossplane-provider-btp/apis"
	// "github.com/sap/crossplane-provider-btp/apis/account/v1beta1"
	"github.com/sap/crossplane-provider-btp/apis/v1alpha1"
	template "github.com/sap/crossplane-provider-btp/internal/controller"

	v1alpha1env "github.com/sap/crossplane-provider-btp/apis/environment/v1alpha1"
	// v1beta1env "github.com/sap/crossplane-provider-btp/apis/environment/v1beta1"
	// v1alpha1webhook "github.com/sap/crossplane-provider-btp/internal/webhook/v1alpha1"
	// v1beta1webhook "github.com/sap/crossplane-provider-btp/internal/webhook/v1beta1"
)

func main() {
	var (
		app            = kingpin.New(filepath.Base(os.Args[0]), "SAP BTP Account Management support for Crossplane.").DefaultEnvars()
		debug          = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		leaderElection = app.Flag(
			"leader-election",
			"Use leader election for the controller manager.",
		).Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()

		syncInterval = app.Flag(
			"sync",
			"How often all resources will be double-checked for drift from the desired state.",
		).Short('s').Default("1h").Duration()
		pollInterval = app.Flag(
			"poll",
			"How often individual resources will be checked for drift from the desired state",
		).Default("1m").Duration()
		maxReconcileRate = app.Flag(
			"max-reconcile-rate",
			"The global maximum rate per second at which resources may checked for drift from the desired state.",
		).Default("3").Int()

		namespace = app.Flag(
			"namespace",
			"Namespace used to set as default scope in default secret store config.",
		).Default("crossplane-system").Envar("POD_NAMESPACE").String()
		enableExternalSecretStores = app.Flag(
			"enable-external-secret-stores",
			"Enable support for ExternalSecretStores.",
		).Default("false").Envar("ENABLE_EXTERNAL_SECRET_STORES").Bool()
		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()

		terraformVersion = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()
	)

	tfclient.TF_VERSION_CALLBACK = func() tfclient.TfEnvVersion {
		return tfclient.TfEnvVersion{
			Version:         *terraformVersion,
			Providerversion: *providerVersion,
			ProviderSource:  *providerSource,
			DebugLogs:       *debug,
		}
	}

	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("crossplane-provider-btp"))
	ctrl.SetLogger(zl)
	btp.SetLogger(log)
	btp.SetDebug(*debug)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	var webhookCertPath, webhookCertName, webhookCertKey string
	flag.StringVar(&webhookCertPath, "webhook-cert-path", "/tmp/certs", "The directory that contains the webhook certificate.")
	flag.StringVar(&webhookCertName, "webhook-cert-name", "webhook.crt", "The name of the webhook certificate file.")
	flag.StringVar(&webhookCertKey, "webhook-cert-key", "webhook.key", "The name of the webhook key file.")

	var tlsOpts []func(*tls.Config)
	webhookTLSOpts := tlsOpts
	var webhookCertWatcher *certwatcher.CertWatcher

	if len(webhookCertPath) > 0 {

		var err error
		webhookCertWatcher, err = certwatcher.New(
			filepath.Join(webhookCertPath, webhookCertName),
			filepath.Join(webhookCertPath, webhookCertKey),
		)
		if err != nil {
			os.Exit(1)
		}

		webhookTLSOpts = append(webhookTLSOpts, func(config *tls.Config) {
			config.GetCertificate = webhookCertWatcher.GetCertificate
		})
		
	}

	// webhookTLSOpts = []func(*tls.Config){
	// 	func(cfg *tls.Config) {
	// 		// Load the certificate and key files
	// 		cert, err := tls.LoadX509KeyPair("/tmp/certs/webhook.crt", "/tmp/certs/webhook.key")
	// 		if err != nil {
	// 			kingpin.FatalIfError(err, "Failed to load TLS certificate and key")
	// 		}
	// 		cfg.Certificates = []tls.Certificate{cert}
	// 	},
	// }

	webhookServer := webhook.NewServer(webhook.Options{
		TLSOpts: webhookTLSOpts,
	})
	

	mgr, err := ctrl.NewManager(
		ratelimiter.LimitRESTConfig(cfg, *maxReconcileRate), ctrl.Options{
			Cache: cache.Options{SyncPeriod: syncInterval},

			// controller-runtime uses both ConfigMaps and Leases for leader
			// election by default. Leases expire after 15 seconds, with a
			// 10 second renewal deadline. We've observed leader loss due to
			// renewal deadlines being exceeded when under high load - i.e.
			// hundreds of reconciles per second and ~200rps to the API
			// server. Switching to Leases only and longer leases appears to
			// alleviate this.
			LeaderElection:             *leaderElection,
			LeaderElectionID:           "crossplane-leader-election-crossplane-provider-btp",
			LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
			LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
			RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
			WebhookServer:              webhookServer,
		},
	)


	// ctrl.NewWebhookManagedBy(mgr).For(&v1alpha1env.CloudFoundryEnvironment{}).Complete()
	// ctrl.NewWebhookManagedBy(mgr).For(&v1beta1env.CloudFoundryEnvironment{}).Complete()
    
	
		// if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		// 	kingpin.FatalIfError(err, "Cannot run controller manager")
		// 	os.Exit(1)
		// }
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apis.AddToScheme(mgr.GetScheme()), "Cannot add Template APIs to scheme")
	gvk := schema.GroupVersionKind{
		Group:   "environment.btp.sap.crossplane.io",
		Version: "v1beta1",
		Kind:    "CloudFoundryEnvironment",
	}
	//have to add the template.go to register v1beta1
	if _, err := mgr.GetScheme().New(gvk); err != nil {
		kingpin.FatalIfError(err, "not registered")
	} else {
		log.Info("registered")
	}
	if err := ctrl.NewWebhookManagedBy(mgr).For(&v1alpha1env.CloudFoundryEnvironment{}).Complete(); err != nil {
		kingpin.FatalIfError(err, "Cannot run controller manager")
	}
	// if err := v1alpha1webhook.SetupCronJobWebhookWithManager(mgr); err != nil {
	// 	kingpin.FatalIfError(err, "Failed to register webhook")
	// }
	// v1beta1webhook.SetupCronJobWebhookWithManager(mgr)
	setupTerraformControllers(mgr, log, maxReconcileRate, *pollInterval, enableManagementPolicies, enableExternalSecretStores, namespace, terraformVersion, providerSource, providerVersion)
	setupNativeControllers(mgr, log, maxReconcileRate, pollInterval, enableManagementPolicies, enableExternalSecretStores, namespace)

	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

func setupTerraformControllers(mgr manager.Manager, log logging.Logger, maxReconcileRate *int, pollInterval time.Duration, enableManagementPolicies *bool, enableExternalSecretStores *bool, namespace *string, terraformVersion *string, providerSource *string, providerVersion *string) {
	o := tjcontroller.Options{
		Options: controller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            pollInterval,
			MaxConcurrentReconciles: 1,
			Features:                &feature.Flags{},
		},
		Provider: config.GetProvider(),
		// use the following WorkspaceStoreOption to enable the shared gRPC mode
		// terraform.WithProviderRunner(terraform.NewSharedProvider(log, os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH"), terraform.WithNativeProviderArgs("-debuggable")))
		WorkspaceStore: terraform.NewWorkspaceStore(log),
		SetupFn:        tfclient.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion),
	}

	if *enableManagementPolicies {
		o.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	if *enableExternalSecretStores {
		o.Features.Enable(features.EnableAlphaExternalSecretStores)
		log.Info("Alpha feature enabled", "flag", features.EnableAlphaExternalSecretStores)

		// Ensure default store config exists.
		kingpin.FatalIfError(
			resource.Ignore(
				kerrors.IsAlreadyExists, mgr.GetClient().Create(
					context.Background(), &v1alpha1.StoreConfig{
						ObjectMeta: metav1.ObjectMeta{
							Name: "default",
						},
						Spec: v1alpha1.StoreConfigSpec{
							// NOTE(turkenh): We only set required spec and expect optional
							// ones to properly be initialized with CRD level default values.
							SecretStoreConfig: xpv1.SecretStoreConfig{
								DefaultScope: *namespace,
							},
						},
					},
				),
			), "cannot create default store config",
		)
	}

	kingpin.FatalIfError(template.Setup(mgr, o), "Cannot setup controllers")
	conversion.RegisterConversions(o.Provider)
	
}
func setupNativeControllers(mgr manager.Manager, log logging.Logger, maxReconcileRate *int, pollInterval *time.Duration, enableManagementPolicies *bool, enableExternalSecretStores *bool, namespace *string) {
	co := controller.Options{
		Logger:                  log,
		MaxConcurrentReconciles: *maxReconcileRate,
		PollInterval:            *pollInterval,
		GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
		Features:                &feature.Flags{},
	}

	if *enableManagementPolicies {
		co.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	if *enableExternalSecretStores {
		co.Features.Enable(features.EnableAlphaExternalSecretStores)
		log.Info("Alpha feature enabled", "flag", features.EnableAlphaExternalSecretStores)

		// Ensure default store config exists.
		kingpin.FatalIfError(
			resource.Ignore(
				kerrors.IsAlreadyExists, mgr.GetClient().Create(
					context.Background(), &v1alpha1.StoreConfig{
						ObjectMeta: metav1.ObjectMeta{
							Name: "default",
						},
						Spec: v1alpha1.StoreConfigSpec{
							// NOTE(turkenh): We only set required spec and expect optional
							// ones to properly be initialized with CRD level default values.
							SecretStoreConfig: xpv1.SecretStoreConfig{
								DefaultScope: *namespace,
							},
						},
					},
				),
			), "cannot create default store config",
		)
	}
	kingpin.FatalIfError(template.CustomSetup(mgr, co), "Cannot setup controllers")
}
