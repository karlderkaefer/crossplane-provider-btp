/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"btp_subaccount_trust_configuration":    config.IdentifierFromProvider,
	"btp_globalaccount_trust_configuration": config.IdentifierFromProvider,
	"btp_directory_entitlement":             config.IdentifierFromProvider,
	"btp_subaccount_service_instance":       config.IdentifierFromProvider,
	"btp_subaccount_service_binding":        config.IdentifierFromProvider,
	"btp_subaccount_service_broker":         config.IdentifierFromProvider,
	"btp_subaccount_api_credential":         config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
