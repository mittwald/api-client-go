package httpclient_mock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ResponseOption func(*http.Response)

func WithStatus(status int) ResponseOption {
	return func(resp *http.Response) {
		resp.StatusCode = status
		resp.Status = http.StatusText(status)
	}
}

func WithJSONResponse(body any) ResponseOption {
	return func(resp *http.Response) {
		j, _ := json.Marshal(body)

		jsonReader := bytes.NewReader(j)
		jsonReadCloser := io.NopCloser(jsonReader)

		resp.Body = jsonReadCloser
		if resp.Header == nil {
			resp.Header = make(http.Header)
		}
		resp.Header.Set("Content-Type", "application/json")

		if resp.StatusCode == 0 {
			WithStatus(200)(resp)
		}
	}
}
