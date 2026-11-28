# UpdateSeverityMatrix

Deprecated, use PATCH /v1/severity_matrix/items. Replaces the severity matrix with the items in the request; items that are not in the request are removed.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Summary`                                                                                    | `*string`                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Data`                                                                                       | [][components.UpdateSeverityMatrixData](../../models/components/updateseveritymatrixdata.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |