package hooks_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/firehydrant/firehydrant-go-sdk/internal/hooks"
	"github.com/firehydrant/firehydrant-go-sdk/internal/utils"
)

type testCase struct {
	Name                string
	Body                []byte
	ResponseContentType string
	ExpectedContentType string
}

var testCases = []testCase{
	{
		Name:                "valid JSON",
		Body:                []byte(`{"message": "i like turtles"}`),
		ResponseContentType: "application/json",
		ExpectedContentType: "application/json",
	},
	{
		Name:                "invalid JSON",
		Body:                []byte(`i like turtles`),
		ResponseContentType: "application/json",
		ExpectedContentType: "text/plain",
	},
	{
		Name:                "no body",
		Body:                nil,
		ResponseContentType: "application/json",
		ExpectedContentType: "text/plain",
	},
	{
		Name:                "empty body",
		Body:                []byte(""),
		ResponseContentType: "application/json",
		ExpectedContentType: "text/plain",
	},
}

func getTestCases() []testCase {
	cases := make([]testCase, len(testCases))
	copy(cases, testCases)
	return cases
}

func TestContentTypeGuardHook_AfterSuccess(t *testing.T) {
	for _, c := range getTestCases() {
		t.Run(c.Name, func(t *testing.T) {
			// Assemble
			hook := &hooks.ContentTypeGuardHook{}

			var body io.ReadCloser
			if c.Body != nil {
				body = io.NopCloser(bytes.NewReader(c.Body))
			}

			res := &http.Response{
				StatusCode: 200,
				Header: http.Header{
					"Content-Type": []string{c.ResponseContentType},
				},
				Body: body,
			}

			// Act
			result, err := hook.AfterSuccess(hooks.AfterSuccessContext{}, res)

			// Assert
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if result.Header.Get("Content-Type") != c.ExpectedContentType {
				t.Errorf("expected Content-Type to be %s, got %s", c.ExpectedContentType, result.Header.Get("Content-Type"))
			}
		})
	}
}

func TestContentTypeGuardHook_AfterSuccess_BodyStillReadable(t *testing.T) {
	// Assemble
	responseBody := []byte(`{"message": "i like turtles"}`)
	hook := &hooks.ContentTypeGuardHook{}

	res := &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(responseBody)),
	}

	result, err := hook.AfterSuccess(hooks.AfterSuccessContext{}, res)

	// Act
	body, err := utils.ConsumeRawBody(result)

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if string(body) != string(responseBody) {
		t.Errorf("expected body to be %s, got %s", string(responseBody), string(body))
	}
}

func TestContentTypeGuardHook_AfterError(t *testing.T) {
	for _, c := range getTestCases() {
		t.Run(c.Name, func(t *testing.T) {
			// Assemble
			hook := &hooks.ContentTypeGuardHook{}

			var body io.ReadCloser
			if c.Body != nil {
				body = io.NopCloser(bytes.NewReader(c.Body))
			}

			res := &http.Response{
				StatusCode: 200,
				Header: http.Header{
					"Content-Type": []string{c.ResponseContentType},
				},
				Body: body,
			}

			// Act
			result, err := hook.AfterError(hooks.AfterErrorContext{}, res, fmt.Errorf("some error"))

			// Assert
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if result.Header.Get("Content-Type") != c.ExpectedContentType {
				t.Errorf("expected Content-Type to be %s, got %s", c.ExpectedContentType, result.Header.Get("Content-Type"))
			}
		})
	}
}

func TestContentTypeGuardHook_AfterError_BodyStillReadable(t *testing.T) {
	// Assemble
	responseBody := []byte(`{"message": "i like turtles"}`)
	hook := &hooks.ContentTypeGuardHook{}

	res := &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(responseBody)),
	}

	result, err := hook.AfterError(hooks.AfterErrorContext{}, res, fmt.Errorf("some error"))

	// Act
	body, err := utils.ConsumeRawBody(result)

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if string(body) != string(responseBody) {
		t.Errorf("expected body to be %s, got %s", string(responseBody), string(body))
	}
}
