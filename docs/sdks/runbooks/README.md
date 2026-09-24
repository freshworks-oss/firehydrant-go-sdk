# Runbooks

## Overview

Operations related to Runbooks

### Available Operations

* [ListRunbookActions](#listrunbookactions) - List runbook actions
* [ListRunbookExecutions](#listrunbookexecutions) - List runbook executions
* [CreateRunbookExecution](#createrunbookexecution) - Create a runbook execution
* [GetRunbookExecution](#getrunbookexecution) - Get a runbook execution
* [DeleteRunbookExecution](#deleterunbookexecution) - Terminate a runbook execution
* [UpdateRunbookExecutionStep](#updaterunbookexecutionstep) - Update a runbook step execution
* [GetRunbookExecutionStepScript](#getrunbookexecutionstepscript) - Get a step's bash script
* [UpdateRunbookExecutionStepScript](#updaterunbookexecutionstepscript) - Update a script step's execution status
* [GetRunbookActionFieldOptions](#getrunbookactionfieldoptions) - List select options for a runbook integration action field
* [ListRunbooks](#listrunbooks) - List runbooks
* [CreateRunbook](#createrunbook) - Create a runbook
* [GetRunbook](#getrunbook) - Get a runbook
* [UpdateRunbook](#updaterunbook) - Update a runbook
* [DeleteRunbook](#deleterunbook) - Delete a runbook
* [ListRunbookAudits](#listrunbookaudits) - List runbook audits

## ListRunbookActions

List all runbook actions available through your connected integrations

### Example Usage

<!-- UsageSnippet language="go" operationID="list_runbook_actions" method="get" path="/v1/runbooks/actions" -->
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

    res, err := s.Runbooks.ListRunbookActions(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `page`                                                                                        | `*int`                                                                                        | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `perPage`                                                                                     | `*int`                                                                                        | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `type_`                                                                                       | `*string`                                                                                     | :heavy_minus_sign:                                                                            | List actions supporting this specific Runbook type                                            |
| `lite`                                                                                        | `*bool`                                                                                       | :heavy_minus_sign:                                                                            | Boolean to determine whether to return a slimified version of the action object's integration |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*components.RunbooksActionsEntityPaginated](../../models/components/runbooksactionsentitypaginated.md), error**

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

## ListRunbookExecutions

List all runbook executions across all runbooks

### Example Usage

<!-- UsageSnippet language="go" operationID="list_runbook_executions" method="get" path="/v1/runbooks/executions" -->
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

    res, err := s.Runbooks.ListRunbookExecutions(ctx, nil, nil)
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

**[*components.RunbooksExecutionEntityPaginated](../../models/components/runbooksexecutionentitypaginated.md), error**

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

## CreateRunbookExecution

Attaches a runbook to an incident and executes it

### Example Usage

<!-- UsageSnippet language="go" operationID="create_runbook_execution" method="post" path="/v1/runbooks/executions" -->
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

    res, err := s.Runbooks.CreateRunbookExecution(ctx, components.CreateRunbookExecution{
        ExecuteFor: "<value>",
        RunbookID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateRunbookExecution](../../models/components/createrunbookexecution.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*components.RunbooksExecutionEntity](../../models/components/runbooksexecutionentity.md), error**

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

## GetRunbookExecution

Retrieve a runbook execution by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="get_runbook_execution" method="get" path="/v1/runbooks/executions/{execution_id}" -->
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

    res, err := s.Runbooks.GetRunbookExecution(ctx, "<id>")
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
| `executionID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.RunbooksExecutionEntity](../../models/components/runbooksexecutionentity.md), error**

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

## DeleteRunbookExecution

Terminates a runbook execution, preventing any further steps from being executed

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_runbook_execution" method="delete" path="/v1/runbooks/executions/{execution_id}" -->
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

    err := s.Runbooks.DeleteRunbookExecution(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `executionID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `reason`                                                 | `string`                                                 | :heavy_check_mark:                                       | The reason for terminating the runbook execution         |
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

## UpdateRunbookExecutionStep

Updates a runbook step execution, especially for changing the state of a step execution.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_runbook_execution_step" method="put" path="/v1/runbooks/executions/{execution_id}/steps/{step_id}" -->
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

    res, err := s.Runbooks.UpdateRunbookExecutionStep(ctx, "<id>", "<id>", components.UpdateRunbookExecutionStep{
        State: "Kentucky",
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
| `executionID`                                                                                  | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `stepID`                                                                                       | `string`                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `updateRunbookExecutionStep`                                                                   | [components.UpdateRunbookExecutionStep](../../models/components/updaterunbookexecutionstep.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*components.RunbooksExecutionEntity](../../models/components/runbooksexecutionentity.md), error**

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

## GetRunbookExecutionStepScript

Retrieves the bash script from a "script" step.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_runbook_execution_step_script" method="get" path="/v1/runbooks/executions/{execution_id}/steps/{step_id}/script" -->
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

    res, err := s.Runbooks.GetRunbookExecutionStepScript(ctx, "<id>", "<id>")
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
| `executionID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `stepID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.RunbooksExecutionEntity](../../models/components/runbooksexecutionentity.md), error**

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

## UpdateRunbookExecutionStepScript

Updates the execution's step.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_runbook_execution_step_script" method="put" path="/v1/runbooks/executions/{execution_id}/steps/{step_id}/script/{state}" -->
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

    res, err := s.Runbooks.UpdateRunbookExecutionStepScript(ctx, "<id>", "<id>", "North Dakota")
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
| `executionID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `stepID`                                                 | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `state`                                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.RunbooksExecutionEntity](../../models/components/runbooksexecutionentity.md), error**

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

## GetRunbookActionFieldOptions

List select options for a runbook integration action field

### Example Usage

<!-- UsageSnippet language="go" operationID="get_runbook_action_field_options" method="get" path="/v1/runbooks/select_options/{integration_slug}/{action_slug}/{field}" -->
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

    err := s.Runbooks.GetRunbookActionFieldOptions(ctx, operations.GetRunbookActionFieldOptionsRequest{
        IntegrationSlug: "<value>",
        ActionSlug: "<value>",
        Field: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetRunbookActionFieldOptionsRequest](../../models/operations/getrunbookactionfieldoptionsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

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

## ListRunbooks

Lists all available runbooks.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_runbooks" method="get" path="/v1/runbooks" -->
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

    res, err := s.Runbooks.ListRunbooks(ctx, operations.ListRunbooksRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListRunbooksRequest](../../models/operations/listrunbooksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*components.SlimRunbookEntityPaginated](../../models/components/slimrunbookentitypaginated.md), error**

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

## CreateRunbook

Create a new runbook for use with incidents.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_runbook" method="post" path="/v1/runbooks" -->
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

    res, err := s.Runbooks.CreateRunbook(ctx, components.CreateRunbook{
        Name: "<value>",
        Type: components.CreateRunbookTypeInfrastructure,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.CreateRunbook](../../models/components/createrunbook.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*components.RunbookEntity](../../models/components/runbookentity.md), error**

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

## GetRunbook

Get a runbook and all its configuration

### Example Usage

<!-- UsageSnippet language="go" operationID="get_runbook" method="get" path="/v1/runbooks/{runbook_id}" -->
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

    res, err := s.Runbooks.GetRunbook(ctx, "<id>")
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
| `runbookID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.RunbookEntity](../../models/components/runbookentity.md), error**

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

## UpdateRunbook

Update a runbook and any attachment rules associated with it. This endpoint is used to configure nearly everything
about a runbook, including but not limited to the steps, environments, attachment rules, and severities.


### Example Usage

<!-- UsageSnippet language="go" operationID="update_runbook" method="put" path="/v1/runbooks/{runbook_id}" -->
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

    res, err := s.Runbooks.UpdateRunbook(ctx, "<id>", components.UpdateRunbook{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `runbookID`                                                          | `string`                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `updateRunbook`                                                      | [components.UpdateRunbook](../../models/components/updaterunbook.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*components.RunbookEntity](../../models/components/runbookentity.md), error**

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

## DeleteRunbook

Delete a runbook and make it unavailable for any future incidents.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_runbook" method="delete" path="/v1/runbooks/{runbook_id}" -->
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

    err := s.Runbooks.DeleteRunbook(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `runbookID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
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

## ListRunbookAudits

This endpoint is deprecated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_runbook_audits" method="get" path="/v1/runbook_audits" -->
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

    err := s.Runbooks.ListRunbookAudits(ctx, nil, nil, operations.AuditableTypeRunbooksStep.ToPointer(), nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `page`                                                                            | `*int`                                                                            | :heavy_minus_sign:                                                                | N/A                                                                               |
| `perPage`                                                                         | `*int`                                                                            | :heavy_minus_sign:                                                                | N/A                                                                               |
| `auditableType`                                                                   | [*operations.AuditableType](../../models/operations/auditabletype.md)             | :heavy_minus_sign:                                                                | A query to filter audits by type                                                  |
| `sort`                                                                            | `*string`                                                                         | :heavy_minus_sign:                                                                | A query to sort audits by their created_at timestamp. Options are 'asc' or 'desc' |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

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