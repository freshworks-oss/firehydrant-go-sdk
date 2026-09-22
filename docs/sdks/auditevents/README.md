# AuditEvents

## Overview

Operations about Audit Events

### Available Operations

* [ListAuditEvents](#listauditevents) - List audit events
* [GetAuditEvent](#getauditevent) - Get a single audit event

## ListAuditEvents

List audit events

### Example Usage

<!-- UsageSnippet language="go" operationID="list_audit_events" method="get" path="/v1/audit_events" -->
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

    err := s.AuditEvents.ListAuditEvents(ctx, nil, nil, firehydrantgosdk.Pointer[int](20))
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `cursor`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Cursor for pagination.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `filter`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Query string to filter audit events, concatenated with AND keyword.<br/>Available filters with example:<br/>  - event.occurred_at < 2023-01-01T00:00:00Z<br/>  - event.key = signals.on_call_rotation.generate<br/>  - event.key = ticketing.sync<br/>  - event.actor.kind = user<br/>  - event.actor.id = 00000000-0000-0000-0000-000000000000<br/>  - resource.kind = incident<br/>  - resource.kind = ticketing.ticket<br/>  - resource.id = 00000000-0000-0000-0000-000000000000<br/>  - resource.response = FAILURE<br/>  - parent_id = 00000000-0000-0000-0000-000000000000<br/>Valid query looks like (without quotes):<br/>  event.occurred_at < 2023-01-01T00:00:00Z AND event.key = signals.on_call_rotation.generate<br/> |
| `limit`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | `*int`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Number of records to display in a single page, maximum is 100 entries. Smaller number is recommended for better performance.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAuditEvent

Get a single audit event

### Example Usage

<!-- UsageSnippet language="go" operationID="get_audit_event" method="get" path="/v1/audit_events/{id}" -->
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

    err := s.AuditEvents.GetAuditEvent(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |