package hooks

import (
	"encoding/json"
	"fmt"
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
	if res.Body == nil {
		fmt.Printf("Body is nil\n")
		res.Header.Set("Content-Type", "text/plain")
	} else {
		fmt.Printf("Body is not nil\n")
		if res.Header.Get("Content-Type") == "application/json" {
			fmt.Printf("Content-Type is application/json\n")
			body, err := utils.ConsumeRawBody(res)
			fmt.Printf("Body: %s\n", string(body))
			if err == nil && !json.Valid(body) {
				fmt.Printf("Body is not valid JSON\n")
				res.Header.Set("Content-Type", "text/plain")
			}
		}
	}

	return res
}
