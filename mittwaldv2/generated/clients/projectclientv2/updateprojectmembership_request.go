package projectclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// UpdateProjectMembershipRequest models a request for the
// 'project-update-project-membership' operation. See [1] for more information.
//
// Update a ProjectMembership.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/project/project-update-project-membership
type UpdateProjectMembershipRequest struct {
	Body                UpdateProjectMembershipRequestBody
	ProjectMembershipID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *UpdateProjectMembershipRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, r.url(), body)
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

func (r *UpdateProjectMembershipRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *UpdateProjectMembershipRequest) url() string {
	u := url.URL{
		Path: fmt.Sprintf("/v2/project-memberships/%s", url.PathEscape(r.ProjectMembershipID)),
	}
	return u.String()
}

func (r *UpdateProjectMembershipRequest) query() url.Values {
	return nil
}
