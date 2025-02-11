package wrappedserviceinstance

import (
	"context"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	"github.com/sap/crossplane-provider-btp/internal"
	"testing"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

var apiError = errors.New("handler error")

func TestObserve(t *testing.T) {
	type args struct {
		mg resource.Managed
	}

	type want struct {
		o   managed.ExternalObservation
		err error
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"NeedsCreation": {
			reason: "gracefully return a required creation",
			args: args{
				mg: &v1alpha1.WrappedServiceInstance{},
			},
			want: want{
				o: managed.ExternalObservation{
					ResourceExists: false,
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{}
			got, err := e.Observe(context.Background(), tc.args.mg)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\ne.Observe(...): -want error, +got error:\n%s\n", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.o, got); diff != "" {
				t.Errorf("\n%s\ne.Observe(...): -want, +got:\n%s\n", tc.reason, diff)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		mg      resource.Managed
		handler AsyncServiceInstanceHandler
	}

	type want struct {
		err error
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"handlerError": {
			reason: "return error from handler to reconciler",
			args: args{
				mg: &v1alpha1.WrappedServiceInstance{},
				handler: MockServiceInstanceHandler{
					err: apiError,
				},
			},
			want: want{
				err: apiError,
			},
		},
		"success": {
			reason: "no error in case of success",
			args: args{
				mg: &v1alpha1.WrappedServiceInstance{},
				handler: MockServiceInstanceHandler{
					err: nil,
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := external{
				instanceHandler: tc.args.handler,
			}
			_, err := e.Create(context.Background(), tc.args.mg)
			internal.VerifyTestError(t, tc.want.err, err)
		})
	}
}

var _ AsyncServiceInstanceHandler = &MockServiceInstanceHandler{}

type MockServiceInstanceHandler struct {
	err error
}

func (m MockServiceInstanceHandler) CreateResource() error {
	return m.err
}
