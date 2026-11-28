# UpdateSeverityMatrixItemsRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `DataSeverity`                                   | []`string`                                       | :heavy_check_mark:                               | Slug of a severity. Send null to remove the item |
| `DataImpactID`                                   | []`string`                                       | :heavy_check_mark:                               | Impact id                                        |
| `DataConditionID`                                | []`string`                                       | :heavy_check_mark:                               | Condition id                                     |