# Issue 29: audience summaries response mismatch

This local patch demonstrates the fix for
https://github.com/freshworks-oss/firehydrant-go-sdk/issues/29.

The endpoint returns `{"data":{"incident_id":"...","audiences":{...}}}`.
The SDK currently decodes the root as the summary entity, silently dropping
`data`, and declares `audiences` as a string instead of an object.

The patch unwraps `data` while keeping the method's return type, and changes
`Audiences` and its getter to `map[string]any` to preserve dynamic audience
names and summary fields. Changing the field type is a source compatibility
change for callers using the old string field.

## Durable upstream fix

CONTRIBUTING.md says direct generated-code PRs are not accepted. The generation
workflow fetches https://api.firehydrant.io/v1/swagger_doc and converts it to
OpenAPI 3; the checked-in openapi.yaml does not contain this endpoint.
Maintainers need to correct the upstream Swagger response definition and
regenerate the SDK. The equivalent OpenAPI 3 response shape is:

```yaml
type: object
properties:
  data:
    $ref: '#/components/schemas/Audiences_Entities_AudienceSummariesEntity'
```

Within `Audiences_Entities_AudienceSummariesEntity`, replace the string
`audiences` property with:

```yaml
type: object
description: Map of audience names to summary objects
additionalProperties: true
```

If the API team has a canonical summary schema, use it for
`additionalProperties` instead of a free-form object. Preserve the unwrapped
method return value using the generator's response-unwrapping configuration
if supported, or document the regenerated return-type change.

`audiences_test.go` exercises the public SDK method against fixture responses
for populated and empty audience maps, including nested fields. It requires
no API credentials.
