# Signals

## Overview

Operations related to Signals

### Available Operations

* [GetSupportHoursSchedule](#getsupporthoursschedule) - Get support hours schedule
* [DeleteSupportHoursSchedule](#deletesupporthoursschedule) - Delete a specific support hours schedule
* [UpdateSupportHoursSchedule](#updatesupporthoursschedule) - Update support hours schedule
* [ListTeamEscalationPolicies](#listteamescalationpolicies) - List escalation policies for a team
* [CreateTeamEscalationPolicy](#createteamescalationpolicy) - Create an escalation policy for a team
* [GetTeamEscalationPolicy](#getteamescalationpolicy) - Get an escalation policy for a team
* [DeleteTeamEscalationPolicy](#deleteteamescalationpolicy) - Delete an escalation policy for a team
* [UpdateTeamEscalationPolicy](#updateteamescalationpolicy) - Update an escalation policy for a team
* [PreviewTeamOnCallSchedule](#previewteamoncallschedule) - Preview a new on-call schedule for a team
* [ListTeamOnCallSchedules](#listteamoncallschedules) - List on-call schedules for a team
* [CreateTeamOnCallSchedule](#createteamoncallschedule) - Create an on-call schedule for a team
* [GetTeamOnCallSchedule](#getteamoncallschedule) - Get an on-call schedule for a team
* [DeleteTeamOnCallSchedule](#deleteteamoncallschedule) - Delete an on-call schedule for a team
* [UpdateTeamOnCallSchedule](#updateteamoncallschedule) - Update an on-call schedule for a team
* [PreviewOnCallScheduleRotation](#previewoncallschedulerotation) - Preview an on-call rotation
* [CreateOnCallScheduleRotation](#createoncallschedulerotation) - Create a new on-call rotation
* [CopyOnCallScheduleRotation](#copyoncallschedulerotation) - Copy an on-call schedule's rotation
* [GetOnCallScheduleRotation](#getoncallschedulerotation) - Get an on-call rotation
* [DeleteOnCallScheduleRotation](#deleteoncallschedulerotation) - Delete an on-call schedule's rotation
* [UpdateOnCallScheduleRotation](#updateoncallschedulerotation) - Update an on-call schedule's rotation
* [CreateOnCallScheduleRotationOverride](#createoncallschedulerotationoverride) - Override one or more shifts in an on-call rotation
* [CreateOnCallShift](#createoncallshift) - [DEPRECATED] Create a shift for an on-call schedule
* [GetOnCallShift](#getoncallshift) - [DEPRECATED] Get an on-call shift for a team schedule
* [DeleteOnCallShift](#deleteoncallshift) - [DEPRECATED] Delete an on-call shift from a team schedule
* [UpdateOnCallShift](#updateoncallshift) - [DEPRECATED] Update an on-call shift for a team schedule
* [ListTeamSignalRules](#listteamsignalrules) - List Signals rules
* [CreateTeamSignalRule](#createteamsignalrule) - Create a Signals rule
* [GetTeamSignalRule](#getteamsignalrule) - Get a Signals rule
* [DeleteTeamSignalRule](#deleteteamsignalrule) - Delete a Signals rule
* [UpdateTeamSignalRule](#updateteamsignalrule) - Update a Signals rule
* [ListSignalsEventSources](#listsignalseventsources) - List event sources for Signals
* [CreateSignalsEventSource](#createsignalseventsource) - Create an event source for Signals
* [GetSignalsEventSource](#getsignalseventsource) - Get an event source for Signals
* [DeleteSignalsEventSource](#deletesignalseventsource) - Delete an event source for Signals
* [GetSignalsHackerMode](#getsignalshackermode) - Get hacker mode status
* [ListSignalsAlertGroupingConfigurations](#listsignalsalertgroupingconfigurations) - List alert grouping configurations.
* [CreateSignalsAlertGroupingConfiguration](#createsignalsalertgroupingconfiguration) - Create an alert grouping configuration.
* [GetSignalsAlertGroupingConfiguration](#getsignalsalertgroupingconfiguration) - Get an alert grouping configuration.
* [DeleteSignalsAlertGroupingConfiguration](#deletesignalsalertgroupingconfiguration) - Delete an alert grouping configuration.
* [UpdateSignalsAlertGroupingConfiguration](#updatesignalsalertgroupingconfiguration) - Update an alert grouping configuration.
* [ListSignalsEmailTargets](#listsignalsemailtargets) - List email targets for signals
* [CreateSignalsEmailTarget](#createsignalsemailtarget) - Create an email target for signals
* [GetSignalsEmailTarget](#getsignalsemailtarget) - Get a signal email target
* [DeleteSignalsEmailTarget](#deletesignalsemailtarget) - Delete a signal email target
* [UpdateSignalsEmailTarget](#updatesignalsemailtarget) - Update an email target
* [ListSignalsWebhookTargets](#listsignalswebhooktargets) - List webhook targets
* [CreateSignalsWebhookTarget](#createsignalswebhooktarget) - Create a webhook target
* [GetSignalsWebhookTarget](#getsignalswebhooktarget) - Get a webhook target
* [DeleteSignalsWebhookTarget](#deletesignalswebhooktarget) - Delete a webhook target
* [UpdateSignalsWebhookTarget](#updatesignalswebhooktarget) - Update a webhook target
* [ListSignalsHeartbeatEndpointConfigurations](#listsignalsheartbeatendpointconfigurations) - List heartbeat endpoint configurations
* [CreateSignalsHeartbeatEndpointConfiguration](#createsignalsheartbeatendpointconfiguration) - Create a heartbeat endpoint configuration
* [GetSignalsHeartbeatEndpointStatus](#getsignalsheartbeatendpointstatus) - Get heartbeat endpoint status
* [GetSignalsHeartbeatEndpointURL](#getsignalsheartbeatendpointurl) - Get heartbeat endpoint URL
* [GetSignalsHeartbeatEndpointConfiguration](#getsignalsheartbeatendpointconfiguration) - Get a heartbeat endpoint configuration
* [DeleteSignalsHeartbeatEndpointConfiguration](#deletesignalsheartbeatendpointconfiguration) - Delete a heartbeat endpoint configuration
* [UpdateSignalsHeartbeatEndpointConfiguration](#updatesignalsheartbeatendpointconfiguration) - Update a heartbeat endpoint configuration
* [ListNotificationPolicySettings](#listnotificationpolicysettings) - List notification policies
* [CreateNotificationPolicy](#createnotificationpolicy) - Create a notification policy
* [GetNotificationPolicy](#getnotificationpolicy) - Get a notification policy
* [DeleteNotificationPolicy](#deletenotificationpolicy) - Delete a notification policy
* [UpdateNotificationPolicy](#updatenotificationpolicy) - Update a notification policy
* [ListUserNotificationSettingsByUserID](#listusernotificationsettingsbyuserid) - List notification settings for a user
* [ListSignalsTransposers](#listsignalstransposers) - List signal transposers
* [GetSignalsIngestURL](#getsignalsingesturl) - Get the signals ingestion URL
* [DebugSignalsExpression](#debugsignalsexpression) - Debug Signals expressions
* [ListOrganizationOnCallSchedules](#listorganizationoncallschedules) - List who's on call for the organization

## GetSupportHoursSchedule

Get support hours schedule for the team

### Example Usage

<!-- UsageSnippet language="go" operationID="get_support_hours_schedule" method="get" path="/v1/teams/{team_id}/support_hours_schedule" -->
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

    res, err := s.Signals.GetSupportHoursSchedule(ctx, "<id>")
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
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SupportHoursScheduleEntity](../../models/components/supporthoursscheduleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSupportHoursSchedule

Delete a specific support hours schedule

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_support_hours_schedule" method="delete" path="/v1/teams/{team_id}/support_hours_schedule" -->
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

    err := s.Signals.DeleteSupportHoursSchedule(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateSupportHoursSchedule

Update the team's support hours schedule

### Example Usage

<!-- UsageSnippet language="go" operationID="update_support_hours_schedule" method="patch" path="/v1/teams/{team_id}/support_hours_schedule" -->
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

    res, err := s.Signals.UpdateSupportHoursSchedule(ctx, "<id>", components.UpdateSupportHoursSchedule{})
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
| `teamID`                                                                                       | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateSupportHoursSchedule`                                                                   | [components.UpdateSupportHoursSchedule](../../models/components/updatesupporthoursschedule.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SupportHoursScheduleEntity](../../models/components/supporthoursscheduleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTeamEscalationPolicies

List all Signals escalation policies for a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_team_escalation_policies" method="get" path="/v1/teams/{team_id}/escalation_policies" -->
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

    res, err := s.Signals.ListTeamEscalationPolicies(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `teamID`                                                              | `string`                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `query`                                                               | `*string`                                                             | :heavy_minus_sign:                                                    | A query string for searching through the list of escalation policies. |
| `page`                                                                | `*int`                                                                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `perPage`                                                             | `*int`                                                                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*components.SignalsAPIEscalationPolicyEntityPaginated](../../models/components/signalsapiescalationpolicyentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTeamEscalationPolicy

Create a Signals escalation policy for a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_team_escalation_policy" method="post" path="/v1/teams/{team_id}/escalation_policies" -->
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

    res, err := s.Signals.CreateTeamEscalationPolicy(ctx, "<id>", components.CreateTeamEscalationPolicy{
        Name: "<value>",
        Steps: []components.CreateTeamEscalationPolicyStep{},
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
| `teamID`                                                                                       | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `createTeamEscalationPolicy`                                                                   | [components.CreateTeamEscalationPolicy](../../models/components/createteamescalationpolicy.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SignalsAPIEscalationPolicyEntity](../../models/components/signalsapiescalationpolicyentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTeamEscalationPolicy

Get a Signals escalation policy by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_team_escalation_policy" method="get" path="/v1/teams/{team_id}/escalation_policies/{id}" -->
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

    res, err := s.Signals.GetTeamEscalationPolicy(ctx, "<id>", "<id>")
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
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIEscalationPolicyEntity](../../models/components/signalsapiescalationpolicyentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTeamEscalationPolicy

Delete a Signals escalation policy by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_team_escalation_policy" method="delete" path="/v1/teams/{team_id}/escalation_policies/{id}" -->
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

    err := s.Signals.DeleteTeamEscalationPolicy(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTeamEscalationPolicy

Update a Signals escalation policy by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_team_escalation_policy" method="patch" path="/v1/teams/{team_id}/escalation_policies/{id}" -->
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

    res, err := s.Signals.UpdateTeamEscalationPolicy(ctx, "<id>", "<id>", components.UpdateTeamEscalationPolicy{})
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
| `teamID`                                                                                       | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `id`                                                                                           | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateTeamEscalationPolicy`                                                                   | [components.UpdateTeamEscalationPolicy](../../models/components/updateteamescalationpolicy.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SignalsAPIEscalationPolicyEntity](../../models/components/signalsapiescalationpolicyentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PreviewTeamOnCallSchedule

Preview a new on-call schedule based on the provided rotations, allowing you to see how the schedule will look before saving it.

### Example Usage

<!-- UsageSnippet language="go" operationID="preview_team_on_call_schedule" method="post" path="/v1/teams/{team_id}/on_call_schedules/preview" -->
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

    res, err := s.Signals.PreviewTeamOnCallSchedule(ctx, "<id>", components.PreviewTeamOnCallSchedule{
        Name: "<value>",
        Rotations: []components.PreviewTeamOnCallScheduleRotation{
            components.PreviewTeamOnCallScheduleRotation{
                Name: "<value>",
                TimeZone: "Europe/Prague",
                Strategy: components.PreviewTeamOnCallScheduleStrategy{
                    Type: components.PreviewTeamOnCallScheduleTypeCustom,
                },
            },
        },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `teamID`                                                                                     | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `previewTeamOnCallSchedule`                                                                  | [components.PreviewTeamOnCallSchedule](../../models/components/previewteamoncallschedule.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*components.SignalsAPIOnCallSchedulePreviewEntity](../../models/components/signalsapioncallschedulepreviewentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTeamOnCallSchedules

List all Signals on-call schedules for a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_team_on_call_schedules" method="get" path="/v1/teams/{team_id}/on_call_schedules" -->
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

    res, err := s.Signals.ListTeamOnCallSchedules(ctx, operations.ListTeamOnCallSchedulesRequest{
        TeamID: "<id>",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListTeamOnCallSchedulesRequest](../../models/operations/listteamoncallschedulesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*components.SignalsAPIOnCallScheduleEntityPaginated](../../models/components/signalsapioncallscheduleentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTeamOnCallSchedule

Create a Signals on-call schedule for a team with a single rotation. More rotations can be created later.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_team_on_call_schedule" method="post" path="/v1/teams/{team_id}/on_call_schedules" -->
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

    res, err := s.Signals.CreateTeamOnCallSchedule(ctx, "<id>", components.CreateTeamOnCallSchedule{
        Name: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `teamID`                                                                                   | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `createTeamOnCallSchedule`                                                                 | [components.CreateTeamOnCallSchedule](../../models/components/createteamoncallschedule.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SignalsAPIOnCallScheduleEntity](../../models/components/signalsapioncallscheduleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTeamOnCallSchedule

Get a Signals on-call schedule by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_team_on_call_schedule" method="get" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}" -->
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

    res, err := s.Signals.GetTeamOnCallSchedule(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                                     |
| `teamID`                                                                                                                                                                                                                                | `string`                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                      | N/A                                                                                                                                                                                                                                     |
| `scheduleID`                                                                                                                                                                                                                            | `string`                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                      | N/A                                                                                                                                                                                                                                     |
| `shiftTimeWindowStart`                                                                                                                                                                                                                  | `*string`                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                      | An optional ISO8601 timestamp for filtering the shifts listed in each on-call schedule to only include shifts that overlap with the provided time window. If provided, only shifts that end at or after this time will be included.     |
| `shiftTimeWindowEnd`                                                                                                                                                                                                                    | `*string`                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                      | An optional ISO8601 timestamp for filtering the shifts listed in each on-call schedule to only include shifts that overlap with the provided time window.. If provided, only shifts that start at or before this time will be included. |
| `opts`                                                                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                      | The options for this request.                                                                                                                                                                                                           |

### Response

**[*components.SignalsAPIOnCallScheduleEntity](../../models/components/signalsapioncallscheduleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTeamOnCallSchedule

Delete a Signals on-call schedule by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_team_on_call_schedule" method="delete" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}" -->
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

    err := s.Signals.DeleteTeamOnCallSchedule(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTeamOnCallSchedule

Update a Signals on-call schedule by ID. For backwards compatibility, all parameters except for
`name` and `description` will be ignored if the schedule has more than one rotation. If the schedule
has only one rotation, you can continue to update that rotation using the rotation-specific parameters.


### Example Usage

<!-- UsageSnippet language="go" operationID="update_team_on_call_schedule" method="patch" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}" -->
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

    res, err := s.Signals.UpdateTeamOnCallSchedule(ctx, "<id>", "<id>", components.UpdateTeamOnCallSchedule{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `teamID`                                                                                   | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `scheduleID`                                                                               | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `updateTeamOnCallSchedule`                                                                 | [components.UpdateTeamOnCallSchedule](../../models/components/updateteamoncallschedule.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SignalsAPIOnCallScheduleEntity](../../models/components/signalsapioncallscheduleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PreviewOnCallScheduleRotation

Preview a new on-call rotation orchanges to an existing on-call rotation

### Example Usage

<!-- UsageSnippet language="go" operationID="preview_on_call_schedule_rotation" method="post" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations/preview" -->
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

    res, err := s.Signals.PreviewOnCallScheduleRotation(ctx, "<id>", "<id>", components.PreviewOnCallScheduleRotation{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `teamID`                                                                                             | `string`                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `scheduleID`                                                                                         | `string`                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `previewOnCallScheduleRotation`                                                                      | [components.PreviewOnCallScheduleRotation](../../models/components/previewoncallschedulerotation.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*components.SignalsAPIOnCallSchedulePreviewEntityRotationPreviewEntity](../../models/components/signalsapioncallschedulepreviewentityrotationpreviewentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateOnCallScheduleRotation

Add a new rotation to an existing on-call schedule

### Example Usage

<!-- UsageSnippet language="go" operationID="create_on_call_schedule_rotation" method="post" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations" -->
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

    res, err := s.Signals.CreateOnCallScheduleRotation(ctx, "<id>", "<id>", components.CreateOnCallScheduleRotation{
        Name: "<value>",
        TimeZone: "Antarctica/DumontDUrville",
        Strategy: components.CreateOnCallScheduleRotationStrategy{
            Type: components.CreateOnCallScheduleRotationTypeWeekly,
        },
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `teamID`                                                                                           | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `scheduleID`                                                                                       | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `createOnCallScheduleRotation`                                                                     | [components.CreateOnCallScheduleRotation](../../models/components/createoncallschedulerotation.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*components.SignalsAPIOnCallRotationEntity](../../models/components/signalsapioncallrotationentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CopyOnCallScheduleRotation

Copy an on-call rotation into a different schedule, allowing you to merge them together safely.

### Example Usage

<!-- UsageSnippet language="go" operationID="copy_on_call_schedule_rotation" method="post" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations/{rotation_id}/copy" -->
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

    res, err := s.Signals.CopyOnCallScheduleRotation(ctx, "<id>", "<id>", "<id>", components.CopyOnCallScheduleRotation{
        TargetScheduleID: "<id>",
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
| `rotationID`                                                                                   | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `teamID`                                                                                       | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `scheduleID`                                                                                   | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `copyOnCallScheduleRotation`                                                                   | [components.CopyOnCallScheduleRotation](../../models/components/copyoncallschedulerotation.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SignalsAPIOnCallRotationEntity](../../models/components/signalsapioncallrotationentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetOnCallScheduleRotation

Get an on-call rotation by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_on_call_schedule_rotation" method="get" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations/{rotation_id}" -->
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

    res, err := s.Signals.GetOnCallScheduleRotation(ctx, "<id>", "<id>", "<id>")
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
| `rotationID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIOnCallRotationEntity](../../models/components/signalsapioncallrotationentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteOnCallScheduleRotation

Delete an on-call schedule's rotation by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_on_call_schedule_rotation" method="delete" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations/{rotation_id}" -->
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

    err := s.Signals.DeleteOnCallScheduleRotation(ctx, "<id>", "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `rotationID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateOnCallScheduleRotation

Update an on-call schedule's rotation by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_on_call_schedule_rotation" method="patch" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations/{rotation_id}" -->
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

    res, err := s.Signals.UpdateOnCallScheduleRotation(ctx, "<id>", "<id>", "<id>", components.UpdateOnCallScheduleRotation{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `rotationID`                                                                                       | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `teamID`                                                                                           | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `scheduleID`                                                                                       | `string`                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `updateOnCallScheduleRotation`                                                                     | [components.UpdateOnCallScheduleRotation](../../models/components/updateoncallschedulerotation.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*components.SignalsAPIOnCallRotationEntity](../../models/components/signalsapioncallrotationentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateOnCallScheduleRotationOverride

Create an override covering a specific  time period in an on-call rotation, re-assigning that
period to a specific user, or leaving it unassigned and claimable by any user, or even creating
a purposeful gap in coverage if desired.


### Example Usage

<!-- UsageSnippet language="go" operationID="create_on_call_schedule_rotation_override" method="post" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/rotations/{rotation_id}/overrides" -->
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

    res, err := s.Signals.CreateOnCallScheduleRotationOverride(ctx, "<id>", "<id>", "<id>", components.CreateOnCallScheduleRotationOverride{
        StartTime: "<value>",
        EndTime: "<value>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `rotationID`                                                                                                       | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `teamID`                                                                                                           | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `scheduleID`                                                                                                       | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `createOnCallScheduleRotationOverride`                                                                             | [components.CreateOnCallScheduleRotationOverride](../../models/components/createoncallschedulerotationoverride.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*components.SignalsAPIOnCallOverrideEntity](../../models/components/signalsapioncalloverrideentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateOnCallShift

NOTE: This endpoint is deprecated and overrides are now the only way to modify a schedule's shifts.
For compatibility, this endpoint will simply create an override for the specified time window.


### Example Usage

<!-- UsageSnippet language="go" operationID="create_on_call_shift" method="post" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts" -->
```go
package main

import(
	"context"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := firehydrantgosdk.New(
        firehydrantgosdk.WithSecurity(components.Security{
            APIKey: "<YOUR_API_KEY_HERE>",
        }),
    )

    res, err := s.Signals.CreateOnCallShift(ctx, "<id>", "<id>", components.CreateOnCallShift{
        StartTime: types.MustTimeFromString("<value>"),
        EndTime: types.MustTimeFromString("<value>"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `teamID`                                                                     | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `scheduleID`                                                                 | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `createOnCallShift`                                                          | [components.CreateOnCallShift](../../models/components/createoncallshift.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.SignalsAPIOnCallOverrideEntity](../../models/components/signalsapioncalloverrideentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetOnCallShift

NOTE: This endpoint is deprecated and should not be used due to the fact that overrides may result in
shifts being split into multiple, separated shifts that still share the same ID. For compatibility, this
endpoint will still return the specified shift, but it may not reflect later fragments of the same shift
that were split off by overrides. You should instead request the rotation itself and inspect the shifts
that are returned for the time window you care about.


### Example Usage

<!-- UsageSnippet language="go" operationID="get_on_call_shift" method="get" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts/{id}" -->
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

    res, err := s.Signals.GetOnCallShift(ctx, "<id>", "<id>", "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIOnCallShiftEntity](../../models/components/signalsapioncallshiftentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteOnCallShift

NOTE: This endpoint is deprecated and overrides are now the only way to modify a schedule's shifts.
For compatibility, this endpoint will simply create a "gap" override for the specified time window,
which will still result in no shifts being present for that time window in the final schedule.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete_on_call_shift" method="delete" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts/{id}" -->
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

    err := s.Signals.DeleteOnCallShift(ctx, "<id>", "<id>", "<id>")
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
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `scheduleID`                                             | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateOnCallShift

NOTE: This endpoint is deprecated and overrides are now the only way to modify a schedule's shifts.
For compatibility, this endpoint will simply create an override for the specified time window.


### Example Usage

<!-- UsageSnippet language="go" operationID="update_on_call_shift" method="patch" path="/v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts/{id}" -->
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

    res, err := s.Signals.UpdateOnCallShift(ctx, "<id>", "<id>", "<id>", components.UpdateOnCallShift{})
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
| `id`                                                                         | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `teamID`                                                                     | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `scheduleID`                                                                 | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `updateOnCallShift`                                                          | [components.UpdateOnCallShift](../../models/components/updateoncallshift.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*components.SignalsAPIOnCallShiftEntity](../../models/components/signalsapioncallshiftentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTeamSignalRules

List all Signals rules for a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_team_signal_rules" method="get" path="/v1/teams/{team_id}/signal_rules" -->
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

    res, err := s.Signals.ListTeamSignalRules(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `teamID`                                                         | `string`                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `query`                                                          | `*string`                                                        | :heavy_minus_sign:                                               | A query string for searching through the list of alerting rules. |
| `page`                                                           | `*int`                                                           | :heavy_minus_sign:                                               | N/A                                                              |
| `perPage`                                                        | `*int`                                                           | :heavy_minus_sign:                                               | N/A                                                              |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*components.SignalsAPIRuleEntityPaginated](../../models/components/signalsapiruleentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTeamSignalRule

Create a Signals rule for a team. We support up to 2000 rules per organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_team_signal_rule" method="post" path="/v1/teams/{team_id}/signal_rules" -->
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

    res, err := s.Signals.CreateTeamSignalRule(ctx, "<id>", components.CreateTeamSignalRule{
        Name: "<value>",
        Expression: "<value>",
        TargetType: components.CreateTeamSignalRuleTargetTypeOnCallSchedule,
        TargetID: "<id>",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `teamID`                                                                           | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `createTeamSignalRule`                                                             | [components.CreateTeamSignalRule](../../models/components/createteamsignalrule.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*components.SignalsAPIRuleEntity](../../models/components/signalsapiruleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTeamSignalRule

Get a Signals rule by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_team_signal_rule" method="get" path="/v1/teams/{team_id}/signal_rules/{id}" -->
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

    res, err := s.Signals.GetTeamSignalRule(ctx, "<id>", "<id>")
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
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIRuleEntity](../../models/components/signalsapiruleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTeamSignalRule

Delete a Signals rule by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_team_signal_rule" method="delete" path="/v1/teams/{team_id}/signal_rules/{id}" -->
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

    err := s.Signals.DeleteTeamSignalRule(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `teamID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTeamSignalRule

Update a Signals rule by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_team_signal_rule" method="patch" path="/v1/teams/{team_id}/signal_rules/{id}" -->
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

    res, err := s.Signals.UpdateTeamSignalRule(ctx, "<id>", "<id>", components.UpdateTeamSignalRule{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `teamID`                                                                           | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `updateTeamSignalRule`                                                             | [components.UpdateTeamSignalRule](../../models/components/updateteamsignalrule.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*components.SignalsAPIRuleEntity](../../models/components/signalsapiruleentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSignalsEventSources

List all Signals event sources for the authenticated user.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_signals_event_sources" method="get" path="/v1/signals/event_sources" -->
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

    res, err := s.Signals.ListSignalsEventSources(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `teamID`                                                                                     | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Team ID to send signals to directly                                                          |
| `escalationPolicyID`                                                                         | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Escalation policy ID to send signals to directly. `team_id` is required if this is provided. |
| `onCallScheduleID`                                                                           | `*string`                                                                                    | :heavy_minus_sign:                                                                           | On-call schedule ID to send signals to directly. `team_id` is required if this is provided.  |
| `userID`                                                                                     | `*string`                                                                                    | :heavy_minus_sign:                                                                           | User ID to send signals to directly                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*components.SignalsAPITransposerListEntity](../../models/components/signalsapitransposerlistentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSignalsEventSource

Create a Signals event source for the authenticated user.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_signals_event_source" method="put" path="/v1/signals/event_sources" -->
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

    res, err := s.Signals.CreateSignalsEventSource(ctx, components.CreateSignalsEventSource{
        Name: "<value>",
        Slug: "<value>",
        ExamplePayload: components.CreateSignalsEventSourceExamplePayload{},
        Javascript: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateSignalsEventSource](../../models/components/createsignalseventsource.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SignalsAPITransposerEntity](../../models/components/signalsapitransposerentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsEventSource

Get a Signals event source by slug

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_event_source" method="get" path="/v1/signals/event_sources/{transposer_slug}" -->
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

    res, err := s.Signals.GetSignalsEventSource(ctx, "<value>")
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
| `transposerSlug`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPITransposerEntity](../../models/components/signalsapitransposerentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSignalsEventSource

Delete a Signals event source by slug

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_signals_event_source" method="delete" path="/v1/signals/event_sources/{transposer_slug}" -->
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

    err := s.Signals.DeleteSignalsEventSource(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `transposerSlug`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsHackerMode

Get the status of the hacker mode for the current user

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_hacker_mode" method="get" path="/v1/signals/hacker_mode" -->
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

    res, err := s.Signals.GetSignalsHackerMode(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIHackerModeEntity](../../models/components/signalsapihackermodeentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSignalsAlertGroupingConfigurations

List all Signals alert grouping rules for the organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_signals_alert_grouping_configurations" method="get" path="/v1/signals/grouping" -->
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

    res, err := s.Signals.ListSignalsAlertGroupingConfigurations(ctx, nil, nil)
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
| `page`                                                   | `*int`                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | `*int`                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIGroupingEntityPaginated](../../models/components/signalsapigroupingentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSignalsAlertGroupingConfiguration

Create a Signals alert grouping rule for the organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_signals_alert_grouping_configuration" method="post" path="/v1/signals/grouping" -->
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

    res, err := s.Signals.CreateSignalsAlertGroupingConfiguration(ctx, components.CreateSignalsAlertGroupingConfiguration{
        Strategy: components.CreateSignalsAlertGroupingConfigurationStrategy{},
        ReferenceAlertTimePeriod: "<value>",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [components.CreateSignalsAlertGroupingConfiguration](../../models/components/createsignalsalertgroupingconfiguration.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*components.SignalsAPIGroupingEntity](../../models/components/signalsapigroupingentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsAlertGroupingConfiguration

Get a Signals alert grouping rule by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_alert_grouping_configuration" method="get" path="/v1/signals/grouping/{id}" -->
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

    res, err := s.Signals.GetSignalsAlertGroupingConfiguration(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIGroupingEntity](../../models/components/signalsapigroupingentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSignalsAlertGroupingConfiguration

Delete a Signals alert grouping rule by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_signals_alert_grouping_configuration" method="delete" path="/v1/signals/grouping/{id}" -->
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

    err := s.Signals.DeleteSignalsAlertGroupingConfiguration(ctx, "<id>")
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

## UpdateSignalsAlertGroupingConfiguration

Update a Signals alert grouping rule for the organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_signals_alert_grouping_configuration" method="patch" path="/v1/signals/grouping/{id}" -->
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

    res, err := s.Signals.UpdateSignalsAlertGroupingConfiguration(ctx, "<id>", components.UpdateSignalsAlertGroupingConfiguration{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `id`                                                                                                                     | `string`                                                                                                                 | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `updateSignalsAlertGroupingConfiguration`                                                                                | [components.UpdateSignalsAlertGroupingConfiguration](../../models/components/updatesignalsalertgroupingconfiguration.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*components.SignalsAPIGroupingEntity](../../models/components/signalsapigroupingentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSignalsEmailTargets

List all Signals email targets for a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_signals_email_targets" method="get" path="/v1/signals/email_targets" -->
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

    res, err := s.Signals.ListSignalsEmailTargets(ctx, nil)
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
| `query`                                                  | `*string`                                                | :heavy_minus_sign:                                       | A query string to search the list of targets by.         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIEmailTargetEntityPaginated](../../models/components/signalsapiemailtargetentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSignalsEmailTarget

Create a Signals email target for a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_signals_email_target" method="post" path="/v1/signals/email_targets" -->
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

    res, err := s.Signals.CreateSignalsEmailTarget(ctx, components.CreateSignalsEmailTarget{
        Name: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.CreateSignalsEmailTarget](../../models/components/createsignalsemailtarget.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SignalsAPIEmailTargetEntity](../../models/components/signalsapiemailtargetentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsEmailTarget

Get a Signals email target by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_email_target" method="get" path="/v1/signals/email_targets/{id}" -->
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

    res, err := s.Signals.GetSignalsEmailTarget(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIEmailTargetEntity](../../models/components/signalsapiemailtargetentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSignalsEmailTarget

Delete a Signals email target by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_signals_email_target" method="delete" path="/v1/signals/email_targets/{id}" -->
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

    err := s.Signals.DeleteSignalsEmailTarget(ctx, "<id>")
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

## UpdateSignalsEmailTarget

Update a Signals email target by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_signals_email_target" method="patch" path="/v1/signals/email_targets/{id}" -->
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

    res, err := s.Signals.UpdateSignalsEmailTarget(ctx, "<id>", components.UpdateSignalsEmailTarget{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `updateSignalsEmailTarget`                                                                 | [components.UpdateSignalsEmailTarget](../../models/components/updatesignalsemailtarget.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*components.SignalsAPIEmailTargetEntity](../../models/components/signalsapiemailtargetentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSignalsWebhookTargets

List all Signals webhook targets.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_signals_webhook_targets" method="get" path="/v1/signals/webhook_targets" -->
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

    res, err := s.Signals.ListSignalsWebhookTargets(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `query`                                                           | `*string`                                                         | :heavy_minus_sign:                                                | A query string for searching through the list of webhook targets. |
| `page`                                                            | `*int`                                                            | :heavy_minus_sign:                                                | N/A                                                               |
| `perPage`                                                         | `*int`                                                            | :heavy_minus_sign:                                                | N/A                                                               |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*components.SignalsAPIWebhookTargetEntityPaginated](../../models/components/signalsapiwebhooktargetentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSignalsWebhookTarget

Create a Signals webhook target.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_signals_webhook_target" method="post" path="/v1/signals/webhook_targets" -->
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

    res, err := s.Signals.CreateSignalsWebhookTarget(ctx, components.CreateSignalsWebhookTarget{
        Name: "<value>",
        URL: "https://puny-hydrolyze.net/",
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
| `request`                                                                                      | [components.CreateSignalsWebhookTarget](../../models/components/createsignalswebhooktarget.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SignalsAPIWebhookTargetEntity](../../models/components/signalsapiwebhooktargetentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsWebhookTarget

Get a Signals webhook target by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_webhook_target" method="get" path="/v1/signals/webhook_targets/{id}" -->
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

    res, err := s.Signals.GetSignalsWebhookTarget(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPIWebhookTargetEntity](../../models/components/signalsapiwebhooktargetentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteSignalsWebhookTarget

Delete a Signals webhook target by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_signals_webhook_target" method="delete" path="/v1/signals/webhook_targets/{id}" -->
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

    err := s.Signals.DeleteSignalsWebhookTarget(ctx, "<id>")
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

## UpdateSignalsWebhookTarget

Update a Signals webhook target by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_signals_webhook_target" method="patch" path="/v1/signals/webhook_targets/{id}" -->
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

    res, err := s.Signals.UpdateSignalsWebhookTarget(ctx, "<id>", components.UpdateSignalsWebhookTarget{})
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
| `id`                                                                                           | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateSignalsWebhookTarget`                                                                   | [components.UpdateSignalsWebhookTarget](../../models/components/updatesignalswebhooktarget.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.SignalsAPIWebhookTargetEntity](../../models/components/signalsapiwebhooktargetentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSignalsHeartbeatEndpointConfigurations

Retrieve all heartbeat endpoint configurations for your organization

### Example Usage

<!-- UsageSnippet language="go" operationID="list_signals_heartbeat_endpoint_configurations" method="get" path="/v1/signals/heartbeat_endpoints" -->
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

    err := s.Signals.ListSignalsHeartbeatEndpointConfigurations(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `teamID`                                                 | `*string`                                                | :heavy_minus_sign:                                       | ID of the team the heartbeat belongs to                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateSignalsHeartbeatEndpointConfiguration

Create a new heartbeat endpoint configuration for your organization

### Example Usage

<!-- UsageSnippet language="go" operationID="create_signals_heartbeat_endpoint_configuration" method="post" path="/v1/signals/heartbeat_endpoints" -->
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

    err := s.Signals.CreateSignalsHeartbeatEndpointConfiguration(ctx, components.CreateSignalsHeartbeatEndpointConfiguration{
        Slug: "<value>",
        Enabled: false,
        ExpectInterval: "<value>",
        TemplateSignal: components.CreateSignalsHeartbeatEndpointConfigurationTemplateSignal{},
        Kind: components.CreateSignalsHeartbeatEndpointConfigurationKindEmail,
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [components.CreateSignalsHeartbeatEndpointConfiguration](../../models/components/createsignalsheartbeatendpointconfiguration.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsHeartbeatEndpointStatus

Check the current status of a heartbeat endpoint

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_heartbeat_endpoint_status" method="get" path="/v1/signals/heartbeat_endpoints/statuses" -->
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

    err := s.Signals.GetSignalsHeartbeatEndpointStatus(ctx, operations.GetSignalsHeartbeatEndpointStatusRequest{
        Ids: []string{},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetSignalsHeartbeatEndpointStatusRequest](../../models/operations/getsignalsheartbeatendpointstatusrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsHeartbeatEndpointURL

Retrieve the URL for a heartbeat endpoint

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_heartbeat_endpoint_url" method="get" path="/v1/signals/heartbeat_endpoints/addresses" -->
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

    err := s.Signals.GetSignalsHeartbeatEndpointURL(ctx, operations.GetSignalsHeartbeatEndpointURLRequest{
        Ids: []string{},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetSignalsHeartbeatEndpointURLRequest](../../models/operations/getsignalsheartbeatendpointurlrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsHeartbeatEndpointConfiguration

Retrieve a single heartbeat endpoint configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_heartbeat_endpoint_configuration" method="get" path="/v1/signals/heartbeat_endpoints/{id}" -->
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

    err := s.Signals.GetSignalsHeartbeatEndpointConfiguration(ctx, "<id>")
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

## DeleteSignalsHeartbeatEndpointConfiguration

Delete a heartbeat endpoint configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_signals_heartbeat_endpoint_configuration" method="delete" path="/v1/signals/heartbeat_endpoints/{id}" -->
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

    err := s.Signals.DeleteSignalsHeartbeatEndpointConfiguration(ctx, "<id>")
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

## UpdateSignalsHeartbeatEndpointConfiguration

Update an existing heartbeat endpoint configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="update_signals_heartbeat_endpoint_configuration" method="patch" path="/v1/signals/heartbeat_endpoints/{id}" -->
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

    err := s.Signals.UpdateSignalsHeartbeatEndpointConfiguration(ctx, "<id>", components.UpdateSignalsHeartbeatEndpointConfiguration{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `id`                                                                                                                             | `string`                                                                                                                         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `updateSignalsHeartbeatEndpointConfiguration`                                                                                    | [components.UpdateSignalsHeartbeatEndpointConfiguration](../../models/components/updatesignalsheartbeatendpointconfiguration.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListNotificationPolicySettings

List all Signals notification policies.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_notification_policy_settings" method="get" path="/v1/signals/notification_policy_items" -->
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

    res, err := s.Signals.ListNotificationPolicySettings(ctx, nil, nil)
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
| `page`                                                   | `*int`                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | `*int`                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPINotificationPolicyItemEntityPaginated](../../models/components/signalsapinotificationpolicyitementitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateNotificationPolicy

Create a Signals notification policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_notification_policy" method="post" path="/v1/signals/notification_policy_items" -->
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

    res, err := s.Signals.CreateNotificationPolicy(ctx, operations.CreateNotificationPolicyRequest{
        NotificationGroupMethod: operations.CreateNotificationPolicyNotificationGroupMethodChat,
        MaxDelay: "<value>",
        Priority: operations.CreateNotificationPolicyPriorityLow,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateNotificationPolicyRequest](../../models/operations/createnotificationpolicyrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*components.SignalsAPINotificationPolicyItemEntity](../../models/components/signalsapinotificationpolicyitementity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetNotificationPolicy

Get a Signals notification policy by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_notification_policy" method="get" path="/v1/signals/notification_policy_items/{id}" -->
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

    res, err := s.Signals.GetNotificationPolicy(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPINotificationPolicyItemEntity](../../models/components/signalsapinotificationpolicyitementity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteNotificationPolicy

Delete a Signals notification policy by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_notification_policy" method="delete" path="/v1/signals/notification_policy_items/{id}" -->
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

    err := s.Signals.DeleteNotificationPolicy(ctx, "<id>")
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

## UpdateNotificationPolicy

Update a Signals notification policy by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_notification_policy" method="patch" path="/v1/signals/notification_policy_items/{id}" -->
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

    err := s.Signals.UpdateNotificationPolicy(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |
| `id`                                                                                                              | `string`                                                                                                          | :heavy_check_mark:                                                                                                | N/A                                                                                                               |
| `requestBody`                                                                                                     | [*operations.UpdateNotificationPolicyRequestBody](../../models/operations/updatenotificationpolicyrequestbody.md) | :heavy_minus_sign:                                                                                                | N/A                                                                                                               |
| `opts`                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                          | :heavy_minus_sign:                                                                                                | The options for this request.                                                                                     |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUserNotificationSettingsByUserID

List all Signals notification settings for a specific user. Requires an API key with PII access enabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_user_notification_settings_by_user_id" method="get" path="/v1/signals/users/{user_id}/notification_settings" -->
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

    res, err := s.Signals.ListUserNotificationSettingsByUserID(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `userID`                                                                            | `string`                                                                            | :heavy_check_mark:                                                                  | N/A                                                                                 |
| `page`                                                                              | `*int`                                                                              | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `perPage`                                                                           | `*int`                                                                              | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `notificationPriority`                                                              | [*operations.NotificationPriority](../../models/operations/notificationpriority.md) | :heavy_minus_sign:                                                                  | The level of priority for the notification setting.                                 |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*components.SignalsAPIUserNotificationSettingEntityPaginated](../../models/components/signalsapiusernotificationsettingentitypaginated.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorEntity | 403                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## ListSignalsTransposers

List all transposers for your organization

### Example Usage

<!-- UsageSnippet language="go" operationID="list_signals_transposers" method="get" path="/v1/signals/transposers" -->
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

    res, err := s.Signals.ListSignalsTransposers(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.SignalsAPITransposerListEntity](../../models/components/signalsapitransposerlistentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSignalsIngestURL

Retrieve the url for ingesting signals for your organization

### Example Usage

<!-- UsageSnippet language="go" operationID="get_signals_ingest_url" method="get" path="/v1/signals/ingest_url" -->
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

    res, err := s.Signals.GetSignalsIngestURL(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `teamID`                                                                                     | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Team ID to send signals to directly                                                          |
| `escalationPolicyID`                                                                         | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Escalation policy ID to send signals to directly. `team_id` is required if this is provided. |
| `onCallScheduleID`                                                                           | `*string`                                                                                    | :heavy_minus_sign:                                                                           | On-call schedule ID to send signals to directly. `team_id` is required if this is provided.  |
| `userID`                                                                                     | `*string`                                                                                    | :heavy_minus_sign:                                                                           | User ID to send signals to directly                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*components.SignalsAPIIngestKeyEntity](../../models/components/signalsapiingestkeyentity.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DebugSignalsExpression

Debug Signals expressions

### Example Usage

<!-- UsageSnippet language="go" operationID="debug_signals_expression" method="post" path="/v1/signals/debugger" -->
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

    err := s.Signals.DebugSignalsExpression(ctx, components.DebugSignalsExpression{
        Expression: "<value>",
        Signals: []components.Signal{},
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.DebugSignalsExpression](../../models/components/debugsignalsexpression.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListOrganizationOnCallSchedules

List all users who are currently on-call across the entire organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_organization_on_call_schedules" method="get" path="/v1/signals_on_call" -->
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

    res, err := s.Signals.ListOrganizationOnCallSchedules(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `teamID`                                                                          | `*string`                                                                         | :heavy_minus_sign:                                                                | An optional comma separated list of team IDs to filter currently on-call users by |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*components.SignalsAPIOrganizationOnCallScheduleEntityPaginated](../../models/components/signalsapiorganizationoncallscheduleentitypaginated.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |