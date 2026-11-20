package firehydrantgosdk_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	sdk "github.com/firehydrant/firehydrant-go-sdk"
)

type audienceSummaryClient struct {
	t    *testing.T
	body string
}

func (c audienceSummaryClient) Do(req *http.Request) (*http.Response, error) {
	c.t.Helper()
	if req.Method != http.MethodGet || req.URL.Path != "/v1/audiences/summaries/incident-123" {
		c.t.Fatalf("unexpected request: %s %s", req.Method, req.URL.Path)
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(strings.NewReader(c.body)),
		Request:    req,
	}, nil
}

func TestListAudienceSummaries(t *testing.T) {
	for _, tc := range []struct {
		name  string
		body  string
		count int
	}{
		{"populated", `{"data":{"incident_id":"incident-123","audiences":{"my-team":{"audience_id":"audience-1","content":"Investigating","extra":{"ready":true}},"another-team":{"audience_id":"audience-2","content":"Resolved"}}}}`, 2},
		{"empty", `{"data":{"incident_id":"incident-123","audiences":{}}}`, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			client := sdk.New(sdk.WithClient(audienceSummaryClient{t: t, body: tc.body}))
			result, err := client.Audiences.ListAudienceSummaries(context.Background(), "incident-123")
			if err != nil {
				t.Fatal(err)
			}
			if result == nil || result.GetIncidentID() == nil || *result.GetIncidentID() != "incident-123" {
				t.Fatalf("incident ID not decoded: %+v", result)
			}
			if result.GetAudiences() == nil || len(result.GetAudiences()) != tc.count {
				t.Fatalf("audiences not decoded: %+v", result.GetAudiences())
			}
			if tc.count > 0 {
				summary, ok := result.GetAudiences()["my-team"].(map[string]any)
				if !ok || summary["audience_id"] != "audience-1" || summary["content"] != "Investigating" {
					t.Fatalf("summary not preserved: %+v", summary)
				}
				extra, ok := summary["extra"].(map[string]any)
				if !ok || extra["ready"] != true {
					t.Fatalf("nested fields not preserved: %+v", summary)
				}
			}
		})
	}
}
