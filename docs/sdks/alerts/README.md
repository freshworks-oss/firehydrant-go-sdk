# Alerts

## Overview

Operations related to Alerts

### Available Operations

* [ListIncidentAlerts](#listincidentalerts) - List alerts for an incident
* [CreateIncidentAlert](#createincidentalert) - Attach an alert to an incident
* [UpdateIncidentAlertPrimary](#updateincidentalertprimary) - Set an alert as primary for an incident
* [DeleteIncidentAlert](#deleteincidentalert) - Remove an alert from an incident
* [ListAlerts](#listalerts) - List alerts
* [GetAlert](#getalert) - Get an alert
* [ListProcessingLogEntries](#listprocessinglogentries) - List alert processing log entries
* [UpdateSignalsAlert](#updatesignalsalert) - Update a Signal alert

## ListIncidentAlerts

List alerts that have been attached to an incident

### Example Usage

<!-- UsageSnippet language="go" operationID="list_incident_alerts" method="get" path="/v1/incidents/{incident_id}/alerts" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Alerts.ListIncidentAlerts(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.IncidentsAlertEntityPaginated](../../models/components/incidentsalertentitypaginated.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## CreateIncidentAlert

Add an alert to an incident. FireHydrant needs to have ingested the alert from a third party system in order to attach it to the incident.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_incident_alert" method="post" path="/v1/incidents/{incident_id}/alerts" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    err := s.Alerts.CreateIncidentAlert(ctx, "<id>", []string{
        "<value 1>",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `requestBody`                                            | []`string`                                               | :heavy_check_mark:                                       | Array of alert IDs to be assigned to the incident        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## UpdateIncidentAlertPrimary

Setting an alert as primary will overwrite milestone times in the FireHydrant incident with times included in the primary alert. Services attached to the primary alert will also be automatically added to the incident.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_incident_alert_primary" method="patch" path="/v1/incidents/{incident_id}/alerts/{incident_alert_id}/primary" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Alerts.UpdateIncidentAlertPrimary(ctx, "<id>", "<id>", components.UpdateIncidentAlertPrimary{
        Primary: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `incidentAlertID`                                                                              | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `incidentID`                                                                                   | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateIncidentAlertPrimary`                                                                   | [components.UpdateIncidentAlertPrimary](../../models/components/updateincidentalertprimary.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.IncidentsAlertEntity](../../models/components/incidentsalertentity.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## DeleteIncidentAlert

Remove an alert from an incident

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_incident_alert" method="delete" path="/v1/incidents/{incident_id}/alerts/{incident_alert_id}" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    err := s.Alerts.DeleteIncidentAlert(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentAlertID`                                        | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListAlerts

Retrieve all alerts, including Signals alerts and third-party

### Example Usage

<!-- UsageSnippet language="go" operationID="list_alerts" method="get" path="/v1/alerts" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Alerts.ListAlerts(ctx, operations.ListAlertsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListAlertsRequest](../../models/operations/listalertsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.AlertsAlertEntityPaginated](../../models/components/alertsalertentitypaginated.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## GetAlert

Retrieve a single alert

### Example Usage

<!-- UsageSnippet language="go" operationID="get_alert" method="get" path="/v1/alerts/{alert_id}" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Alerts.GetAlert(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `alertID`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.AlertsAlertEntity](../../models/components/alertsalertentity.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## ListProcessingLogEntries

Processing Log Entries for a specific alert

### Example Usage

<!-- UsageSnippet language="go" operationID="list_processing_log_entries" method="get" path="/v1/processing_log_entries" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Alerts.ListProcessingLogEntries(ctx, operations.ListProcessingLogEntriesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListProcessingLogEntriesRequest](../../models/operations/listprocessinglogentriesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*components.AlertsProcessingLogEntryEntityPaginated](../../models/components/alertsprocessinglogentryentitypaginated.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.NotFoundError       | 404                           | application/json              |
| sdkerrors.UnauthorizedError   | 401, 403, 407                 | application/json              |
| sdkerrors.TimeoutError        | 408                           | application/json              |
| sdkerrors.RateLimitedError    | 429                           | application/json              |
| sdkerrors.BadRequestError     | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.TimeoutError        | 504                           | application/json              |
| sdkerrors.NotFoundError       | 501, 505                      | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.BadRequestError     | 510                           | application/json              |
| sdkerrors.UnauthorizedError   | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |

## UpdateSignalsAlert

Update a Signal alert

### Example Usage

<!-- UsageSnippet language="go" operationID="update_signals_alert" method="patch" path="/v1/signals/alerts/{id}" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Alerts.UpdateSignalsAlert(ctx, "<id>", components.UpdateSignalsAlert{
        Noise: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `updateSignalsAlert`                                                           | [components.UpdateSignalsAlert](../../models/components/updatesignalsalert.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*components.AlertsSignalAlertEntity](../../models/components/alertssignalalertentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |