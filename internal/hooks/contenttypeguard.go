package hooks

import (
	"encoding/json"
	"net/http"

	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
)

type ContentTypeGuardHook struct{}

func (h *ContentTypeGuardHook) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	return h.fixContentType(res), nil
}

func (h *ContentTypeGuardHook) AfterError(hookCtx AfterErrorContext, res *http.Response, err error) (*http.Response, error) {
	return h.fixContentType(res), nil
}

func (h *ContentTypeGuardHook) fixContentType(res *http.Response) *http.Response {
	if res.Header.Get("Content-Type") == "application/json" {
		body, err := utils.ConsumeRawBody(res)
		if err == nil && !json.Valid(body) {
			res.Header.Set("Content-Type", "text/plain")
		}
	}

	return res
}
