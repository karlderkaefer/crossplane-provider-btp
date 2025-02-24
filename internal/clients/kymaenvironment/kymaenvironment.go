package environments

import (
	"bytes"
	"context"
	"unicode"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	json "github.com/json-iterator/go"
	provisioningclient "github.com/sap/crossplane-provider-btp/internal/openapi_clients/btp-provisioning-service-api-go/pkg"
	"sigs.k8s.io/yaml"

	"github.com/sap/crossplane-provider-btp/apis/environment/v1alpha1"
	"github.com/sap/crossplane-provider-btp/btp"
)

const (
	errKymaInstanceCreateFailed = "Could not create KymaEnvironment"
	errKymaInstanceUpdateFailed = "Could not update KymaEnvironment"
	errInstanceIdNotFound       = "Could not update kyma instance .status.AtProvider.Id is empty"
)

type KymaEnvironments struct {
	btp btp.Client
}

func NewKymaEnvironments(btp btp.Client) *KymaEnvironments {
	return &KymaEnvironments{btp: btp}
}

func (c KymaEnvironments) DescribeInstance(
	ctx context.Context,
	cr v1alpha1.KymaEnvironment,
) (*provisioningclient.EnvironmentInstanceResponseObject, error) {
	environment, err := c.btp.GetEnvironmentByNameAndType(ctx, cr.Name, btp.KymaEnvironmentType())
	if err != nil {
		return nil, err
	}

	if environment == nil {
		return nil, nil
	}

	return environment, nil
}

func (c KymaEnvironments) CreateInstance(ctx context.Context, cr v1alpha1.KymaEnvironment) error {

	parameters, err := UnmarshalRawParameters(cr.Spec.ForProvider.Parameters.Raw)
	parameters = AddKymaDefaultParameters(parameters, cr.Name, string(cr.UID))
	if err != nil {
		return err
	}
	err = c.btp.CreateKymaEnvironment(
		ctx,
		cr.Name,
		cr.Spec.ForProvider.PlanName,
		parameters,
		string(cr.UID),
		c.btp.Credential.UserCredential.Email,
	)

	return errors.Wrap(err, errKymaInstanceCreateFailed)
}

func (c KymaEnvironments) DeleteInstance(ctx context.Context, cr v1alpha1.KymaEnvironment) error {
	if cr.Status.AtProvider.ID == nil {
		return errors.New(errInstanceIdNotFound)
	}
	return c.btp.DeleteEnvironmentById(ctx, *cr.Status.AtProvider.ID)
}

func (c KymaEnvironments) UpdateInstance(ctx context.Context, cr v1alpha1.KymaEnvironment) error {

	if cr.Status.AtProvider.ID == nil {
		return errors.New(errInstanceIdNotFound)
	}

	parameters, err := UnmarshalRawParameters(cr.Spec.ForProvider.Parameters.Raw)
	parameters = AddKymaDefaultParameters(parameters, cr.Name, string(cr.UID))
	if err != nil {
		return err
	}
	err = c.btp.UpdateKymaEnvironment(
		ctx,
		*cr.Status.AtProvider.ID,
		cr.Spec.ForProvider.PlanName,
		parameters,
		string(cr.UID),
	)

	return errors.Wrap(err, errKymaInstanceUpdateFailed)
}

// UnmarshalRawParameters produces a map structure from a given raw YAML/JSON input
func UnmarshalRawParameters(in []byte) (map[string]interface{}, error) {
	parameters := make(map[string]interface{})

	if len(in) == 0 {
		return parameters, nil

	}
	if hasJSONPrefix(in) {
		if err := json.Unmarshal(in, &parameters); err != nil {
			return parameters, err
		}
		return parameters, nil
	}

	err := yaml.Unmarshal(in, &parameters)
	return parameters, err

}

var jsonPrefix = []byte("{")

// hasJSONPrefix returns true if the provided buffer appears to start with
// a JSON open brace.
func hasJSONPrefix(buf []byte) bool {
	return hasPrefix(buf, jsonPrefix)
}

// Return true if the first non-whitespace bytes in buf is prefix.
func hasPrefix(buf []byte, prefix []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, prefix)
}

func AddKymaDefaultParameters(parameters btp.InstanceParameters, instanceName string, resourceUID string) btp.InstanceParameters {
	parameters[btp.KymaenvironmentParameterInstanceName] = instanceName
	return parameters
}
