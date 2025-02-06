package serviceplan

import (
	"context"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/pkg/errors"
	"github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	"github.com/sap/crossplane-provider-btp/internal"
	"github.com/sap/crossplane-provider-btp/internal/clients/servicemanager"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
)

var (
	apiError = errors.New("apiError")
)

func TestObserve(t *testing.T) {
	type args struct {
		client servicemanager.PlanIdResolver
		cr     *v1alpha1.ServicePlan
	}

	type want struct {
		o   managed.ExternalObservation
		err error
		cr  *v1alpha1.ServicePlan
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Lookup Error": {
			reason: "API error should be returned to reconciler",
			args: args{
				client: &servicemanager.PlanIdResolverFake{
					PlanLookupMockFn: func() (string, error) {
						return "", apiError
					},
				},
				cr: servicePlan("offeringName", "planName", "", false),
			},
			want: want{
				cr:  servicePlan("offeringName", "planName", "", false),
				err: apiError,
			},
		},
		"SuccessfulObserve": {
			reason: "Lookup of Serviceplan should succeed",
			args: args{
				client: &servicemanager.PlanIdResolverFake{
					PlanLookupMockFn: func() (string, error) {
						return "123", nil
					},
				},
				cr: servicePlan("offeringName", "planName", "", false),
			},
			want: want{
				cr: servicePlan("offeringName", "planName", "123", true),
				o: managed.ExternalObservation{
					ResourceExists:    true,
					ResourceUpToDate:  true,
					ConnectionDetails: managed.ConnectionDetails{},
				},
				err: nil,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{client: tc.args.client}

			got, err := e.Observe(context.Background(), tc.args.cr)

			internal.VerifyTestError(t, tc.want.err, err)

			if diff := cmp.Diff(tc.want.o, got); diff != "" {
				t.Errorf("\n%s\ne.Observe(...): -want, +got:\n%s\n", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.cr, tc.args.cr); diff != "" {
				t.Errorf("\ne.Observe(): expected cr after operation -want, +got:\n%s\n", diff)
			}
		})
	}
}

// servicePlan creates a ServicePlan object for testing with passed offeringName, PlanName and observed servicePlanId
func servicePlan(offeringName, planName, servicePlanId string, available bool) *v1alpha1.ServicePlan {
	sp := &v1alpha1.ServicePlan{
		Spec: v1alpha1.ServicePlanSpec{
			ForProvider: v1alpha1.ServicePlanParameters{
				OfferingName: offeringName,
				PlanName:     planName,
			},
		},
		Status: v1alpha1.ServicePlanStatus{
			AtProvider: v1alpha1.ServicePlanObservation{
				ServicePlanId: servicePlanId,
			},
		},
	}
	if available {
		sp.Status.SetConditions(xpv1.Available())
	}
	return sp
}
