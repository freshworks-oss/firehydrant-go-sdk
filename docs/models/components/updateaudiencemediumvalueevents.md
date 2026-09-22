# UpdateAudienceMediumValueEvents


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ParentChanged`                                                         | `bool`                                                                  | :heavy_check_mark:                                                      | Include any events where the incident's parent changes in the timeline  |
| `ChildChanged`                                                          | `bool`                                                                  | :heavy_check_mark:                                                      | Include any events where the incident's children change in the timeline |
| `NewRelatedChangeEvent`                                                 | `bool`                                                                  | :heavy_check_mark:                                                      | Include any related change events in the timeline                       |
| `RunbookStepExecutionUpdate`                                            | `bool`                                                                  | :heavy_check_mark:                                                      | Include any runbook step updates in the timeline                        |
| `ChangeType`                                                            | `bool`                                                                  | :heavy_check_mark:                                                      | Include any incident type change events in the timeline                 |