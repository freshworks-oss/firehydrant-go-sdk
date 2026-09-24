# UpdateAudienceDetail


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Question`                                              | `*string`                                               | :heavy_minus_sign:                                      | The incident detail question (max 255 characters)       |
| `Prompt`                                                | `*string`                                               | :heavy_minus_sign:                                      | The prompt to display when collecting this detail       |
| `Slug`                                                  | `*string`                                               | :heavy_minus_sign:                                      | Optional unique identifier for this detail              |
| `Position`                                              | `*int`                                                  | :heavy_minus_sign:                                      | Position of the question in the list (1-based indexing) |