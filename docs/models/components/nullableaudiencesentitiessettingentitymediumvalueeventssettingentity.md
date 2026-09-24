# NullableAudiencesEntitiesSettingEntityMediumValueEventsSettingEntity


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ParentChanged`                                                         | `*bool`                                                                 | :heavy_minus_sign:                                                      | Include any events where the incident's parent changes in the timeline  |
| `ChildChanged`                                                          | `*bool`                                                                 | :heavy_minus_sign:                                                      | Include any events where the incident's children change in the timeline |
| `NewRelatedChangeEvent`                                                 | `*bool`                                                                 | :heavy_minus_sign:                                                      | Include any related change events in the timeline                       |
| `RunbookStepExecutionUpdate`                                            | `*bool`                                                                 | :heavy_minus_sign:                                                      | Include any runbook step updates in the timeline                        |
| `ChangeType`                                                            | `*bool`                                                                 | :heavy_minus_sign:                                                      | Include any incident type change events in the timeline                 |