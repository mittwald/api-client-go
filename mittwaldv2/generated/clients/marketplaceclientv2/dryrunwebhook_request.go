package marketplaceclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// DryRunWebhookRequest models a request for the 'extension-dry-run-webhook'
// operation. See [1] for more information.
//
// Dry run a webhook with random or given values.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/marketplace/extension-dry-run-webhook
type DryRunWebhookRequest struct {
	ContributorID       string
	ExtensionID         string
	ExtensionInstanceID string
	WebhookKind         marketplacev2.WebhookKind
	ContextID           *string
	Scopes              []string
	InstanceDisabled    *bool
	CreatedAt           *time.Time
	Secret              *string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *DryRunWebhookRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *DryRunWebhookRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *DryRunWebhookRequest) url() string {
	u := url.URL{
		Path:     fmt.Sprintf("/v2/contributors/%s/extensions/%s/extension-instances/%s/actions/dry-run/%s", url.PathEscape(r.ContributorID), url.PathEscape(r.ExtensionID), url.PathEscape(r.ExtensionInstanceID), url.PathEscape(string(r.WebhookKind))),
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *DryRunWebhookRequest) query() url.Values {
	q := make(url.Values)
	if r.ContextID != nil {
		q.Set("contextId", *r.ContextID)
	}
	for _, val := range r.Scopes {
		q.Add("scopes", val)
	}
	if r.InstanceDisabled != nil {
		q.Set("instanceDisabled", strconv.FormatBool(*r.InstanceDisabled))
	}
	if r.CreatedAt != nil {
		q.Set("createdAt", r.CreatedAt.Format(time.RFC3339))
	}
	if r.Secret != nil {
		q.Set("secret", *r.Secret)
	}
	return q
}
