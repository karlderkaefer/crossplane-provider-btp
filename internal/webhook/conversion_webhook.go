package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sap/crossplane-provider-btp/apis/environment/v1alpha1"
	"github.com/sap/crossplane-provider-btp/apis/environment/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/convert,mutating=true,failurePolicy=fail,sideEffects=None,groups=environment.btp.sap.crossplane.io,resources=cloudfoundryenvironments,verbs=update;create,versions=v1alpha1;v1beta1,name=environment.btp.sap.xp.io,admissionReviewVersions=v1alpha1

// ConversionWebhook handles conversion requests for all versions.
type ConversionWebhook struct {
	decoder *admission.Decoder
}

// NewConversionWebhook creates a new ConversionWebhook.
//
//	func NewConversionWebhook() *ConversionWebhook {
//	    return &ConversionWebhook{}
//	}
//
// InjectDecoder injects the decoder into the webhook.
func (w *ConversionWebhook) InjectDecoder(d *admission.Decoder) error {
	w.decoder = d
	return nil
}

// Handle handles the conversion requests.
func (w *ConversionWebhook) Handle(ctx context.Context, req admission.Request) admission.Response {
	// Decode the incoming object
	src := &v1alpha1.CloudFoundryEnvironment{}
	if err := w.decoder.Decode(req, src); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("failed to decode object: %w", err))
	}

	// Perform the conversion
	dst := &v1beta1.CloudFoundryEnvironment{}
	if err := src.ConvertTo(dst); err != nil {
		return admission.Errored(http.StatusInternalServerError, fmt.Errorf("conversion failed: %w", err))
	}

	// Encode the converted object ba5ck to raw JSON
	raw, err := json.Marshal(dst)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, fmt.Errorf("failed to encode converted object: %w", err))
	}

	// Return the patch response
	return admission.PatchResponseFromRaw(req.Object.Raw, raw)

}
