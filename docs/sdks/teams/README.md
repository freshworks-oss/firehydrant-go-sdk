# Teams

## Overview

Operations related to Teams

### Available Operations

* [ListTeams](#listteams) - List teams
* [CreateTeam](#createteam) - Create a team
* [GetTeam](#getteam) - Get a team
* [DeleteTeam](#deleteteam) - Archive a team
* [UpdateTeam](#updateteam) - Update a team
* [ListSchedules](#listschedules) - List schedules

## ListTeams

List all of the teams in the organization

### Example Usage

<!-- UsageSnippet language="go" operationID="list_teams" method="get" path="/v1/teams" -->
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

    res, err := s.Teams.ListTeams(ctx, operations.ListTeamsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListTeamsRequest](../../models/operations/listteamsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*components.TeamEntityPaginated](../../models/components/teamentitypaginated.md), error**

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

## CreateTeam

Create a new team

### Example Usage

<!-- UsageSnippet language="go" operationID="create_team" method="post" path="/v1/teams" -->
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

    res, err := s.Teams.CreateTeam(ctx, components.CreateTeam{
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.CreateTeam](../../models/components/createteam.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*components.TeamEntity](../../models/components/teamentity.md), error**

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

## GetTeam

Retrieve a single team from its ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_team" method="get" path="/v1/teams/{team_id}" -->
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

    res, err := s.Teams.GetTeam(ctx, "<id>", nil)
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
| `teamID`                                                                       | `string`                                                                       | :heavy_check_mark:                                                             | Team UUID or slug                                                              |
| `lite`                                                                         | `*bool`                                                                        | :heavy_minus_sign:                                                             | Boolean to determine whether to return a slimified version of the teams object |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*components.TeamEntity](../../models/components/teamentity.md), error**

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

## DeleteTeam

Archives an team which will hide it from lists and metrics

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_team" method="delete" path="/v1/teams/{team_id}" -->
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

    err := s.Teams.DeleteTeam(ctx, "<id>")
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

## UpdateTeam

Update a single team from its ID

### Example Usage

<!-- UsageSnippet language="go" operationID="update_team" method="patch" path="/v1/teams/{team_id}" -->
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

    res, err := s.Teams.UpdateTeam(ctx, "<id>", components.UpdateTeam{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `teamID`                                                       | `string`                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `updateTeam`                                                   | [components.UpdateTeam](../../models/components/updateteam.md) | :heavy_check_mark:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*components.TeamEntity](../../models/components/teamentity.md), error**

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

## ListSchedules

List all known schedules in FireHydrant as pulled from external sources

### Example Usage

<!-- UsageSnippet language="go" operationID="list_schedules" method="get" path="/v1/schedules" -->
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

    res, err := s.Teams.ListSchedules(ctx, nil, nil, nil)
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
| `query`                                                  | `*string`                                                | :heavy_minus_sign:                                       | Filter schedules with a query on their name              |
| `page`                                                   | `*int`                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | `*int`                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.ScheduleEntityPaginated](../../models/components/scheduleentitypaginated.md), error**

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