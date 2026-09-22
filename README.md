<div align="center">
    <picture>
        <img src="https://files.readme.io/7fb504d-small-FireHydrant-logo.png" width="300">
    </picture>
   <p>Manage incidents from ring to retro</p>
	<p>Developer-friendly & type-safe Go SDK specifically catered to leverage <strong>FireHydrant</strong> API.</p>
   <a href=""><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
  * [SDK Installation](#sdk-installation)
  * [Authentication](#authentication)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/firehydrant/firehydrant-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  |
| -------- | ------ | ------- |
| `APIKey` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.AccountSettings.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.AccountSettings.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountSettings](docs/sdks/accountsettings/README.md)

* [Ping](docs/sdks/accountsettings/README.md#ping) - Check API connectivity
* [ListEntitlements](docs/sdks/accountsettings/README.md#listentitlements) - List entitlements
* [PingNoauth](docs/sdks/accountsettings/README.md#pingnoauth) - Check API connectivity
* [GetBootstrap](docs/sdks/accountsettings/README.md#getbootstrap) - Get initial application configuration
* [GetAiPreferences](docs/sdks/accountsettings/README.md#getaipreferences) - Get AI preferences
* [UpdateAiPreferences](docs/sdks/accountsettings/README.md#updateaipreferences) - Update AI preferences

### [Alerts](docs/sdks/alerts/README.md)

* [ListIncidentAlerts](docs/sdks/alerts/README.md#listincidentalerts) - List alerts for an incident
* [CreateIncidentAlert](docs/sdks/alerts/README.md#createincidentalert) - Attach an alert to an incident
* [UpdateIncidentAlertPrimary](docs/sdks/alerts/README.md#updateincidentalertprimary) - Set an alert as primary for an incident
* [DeleteIncidentAlert](docs/sdks/alerts/README.md#deleteincidentalert) - Remove an alert from an incident
* [ListAlerts](docs/sdks/alerts/README.md#listalerts) - List alerts
* [GetAlert](docs/sdks/alerts/README.md#getalert) - Get an alert
* [ListProcessingLogEntries](docs/sdks/alerts/README.md#listprocessinglogentries) - List alert processing log entries
* [UpdateSignalsAlert](docs/sdks/alerts/README.md#updatesignalsalert) - Update a Signal alert

### [Audiences](docs/sdks/audiences/README.md)

* [ListAudiences](docs/sdks/audiences/README.md#listaudiences) - List audiences
* [CreateAudience](docs/sdks/audiences/README.md#createaudience) - Create audience
* [GetAudience](docs/sdks/audiences/README.md#getaudience) - Get audience
* [ArchiveAudience](docs/sdks/audiences/README.md#archiveaudience) - Archive audience
* [UpdateAudience](docs/sdks/audiences/README.md#updateaudience) - Update audience
* [RestoreAudience](docs/sdks/audiences/README.md#restoreaudience) - Restore audience
* [GetMemberDefaultAudience](docs/sdks/audiences/README.md#getmemberdefaultaudience) - Get default audience
* [SetMemberDefaultAudience](docs/sdks/audiences/README.md#setmemberdefaultaudience) - Set default audience
* [GetAudienceSummary](docs/sdks/audiences/README.md#getaudiencesummary) - Get latest summary
* [GenerateAudienceSummary](docs/sdks/audiences/README.md#generateaudiencesummary) - Generate summary (async)
* [ListAudienceSummaries](docs/sdks/audiences/README.md#listaudiencesummaries) - List audience summaries

### [AuditEvents](docs/sdks/auditevents/README.md)

* [ListAuditEvents](docs/sdks/auditevents/README.md#listauditevents) - List audit events
* [GetAuditEvent](docs/sdks/auditevents/README.md#getauditevent) - Get a single audit event

### [CallRoutes](docs/sdks/callroutes/README.md)

* [ListTeamCallRoutes](docs/sdks/callroutes/README.md#listteamcallroutes) - List call routes for a team
* [CreateTeamCallRoute](docs/sdks/callroutes/README.md#createteamcallroute) - Create a call route for a team
* [ListCallRoutes](docs/sdks/callroutes/README.md#listcallroutes) - List call routes
* [GetCallRoute](docs/sdks/callroutes/README.md#getcallroute) - Retrieve a call route
* [DeleteCallRoute](docs/sdks/callroutes/README.md#deletecallroute) - Delete a call route
* [UpdateCallRoute](docs/sdks/callroutes/README.md#updatecallroute) - Update a call route

### [CatalogEntries](docs/sdks/catalogentries/README.md)

* [ListEnvironments](docs/sdks/catalogentries/README.md#listenvironments) - List environments
* [CreateEnvironment](docs/sdks/catalogentries/README.md#createenvironment) - Create an environment
* [GetEnvironment](docs/sdks/catalogentries/README.md#getenvironment) - Get an environment
* [DeleteEnvironment](docs/sdks/catalogentries/README.md#deleteenvironment) - Archive an environment
* [UpdateEnvironment](docs/sdks/catalogentries/README.md#updateenvironment) - Update an environment
* [ListEnvironmentServices](docs/sdks/catalogentries/README.md#listenvironmentservices) - List services for an environment
* [ListEnvironmentFunctionalities](docs/sdks/catalogentries/README.md#listenvironmentfunctionalities) - List functionalities for an environment
* [ListServices](docs/sdks/catalogentries/README.md#listservices) - List services
* [CreateService](docs/sdks/catalogentries/README.md#createservice) - Create a service
* [CreateServiceLinks](docs/sdks/catalogentries/README.md#createservicelinks) - Create multiple services linked to external services
* [GetService](docs/sdks/catalogentries/README.md#getservice) - Get a service
* [DeleteService](docs/sdks/catalogentries/README.md#deleteservice) - Delete a service
* [UpdateService](docs/sdks/catalogentries/README.md#updateservice) - Update a service
* [ListServiceEnvironments](docs/sdks/catalogentries/README.md#listserviceenvironments) - List environments for a service
* [GetServiceDependencies](docs/sdks/catalogentries/README.md#getservicedependencies) - List dependencies for a service
* [ListServiceAvailableUpstreamDependencies](docs/sdks/catalogentries/README.md#listserviceavailableupstreamdependencies) - List available upstream service dependencies
* [ListServiceAvailableDownstreamDependencies](docs/sdks/catalogentries/README.md#listserviceavailabledownstreamdependencies) - List available downstream service dependencies
* [DeleteServiceLink](docs/sdks/catalogentries/README.md#deleteservicelink) - Delete a service link
* [CreateServiceChecklistResponse](docs/sdks/catalogentries/README.md#createservicechecklistresponse) - Record a response for a checklist item
* [CreateServiceDependency](docs/sdks/catalogentries/README.md#createservicedependency) - Create a service dependency
* [GetServiceDependency](docs/sdks/catalogentries/README.md#getservicedependency) - Get a service dependency
* [DeleteServiceDependency](docs/sdks/catalogentries/README.md#deleteservicedependency) - Delete a service dependency
* [UpdateServiceDependency](docs/sdks/catalogentries/README.md#updateservicedependency) - Update a service dependency
* [ListFunctionalities](docs/sdks/catalogentries/README.md#listfunctionalities) - List functionalities
* [CreateFunctionality](docs/sdks/catalogentries/README.md#createfunctionality) - Create a functionality
* [GetFunctionality](docs/sdks/catalogentries/README.md#getfunctionality) - Get a functionality
* [DeleteFunctionality](docs/sdks/catalogentries/README.md#deletefunctionality) - Archive a functionality
* [UpdateFunctionality](docs/sdks/catalogentries/README.md#updatefunctionality) - Update a functionality
* [ListFunctionalityEnvironments](docs/sdks/catalogentries/README.md#listfunctionalityenvironments) - List environments for a functionality
* [ListFunctionalityServices](docs/sdks/catalogentries/README.md#listfunctionalityservices) - List services for a functionality
* [ListUserOwnedServices](docs/sdks/catalogentries/README.md#listuserownedservices) - List services owned by a user's teams
* [ListInfrastructures](docs/sdks/catalogentries/README.md#listinfrastructures) - Lists functionality, service and environment objects
* [RefreshCatalog](docs/sdks/catalogentries/README.md#refreshcatalog) - Refresh a service catalog
* [IngestCatalogData](docs/sdks/catalogentries/README.md#ingestcatalogdata) - Ingest service catalog data

### [Changes](docs/sdks/changes/README.md)

* [ListChanges](docs/sdks/changes/README.md#listchanges) - List changes
* [CreateChange](docs/sdks/changes/README.md#createchange) - Create a new change entry
* [DeleteChange](docs/sdks/changes/README.md#deletechange) - Archive a change entry
* [UpdateChange](docs/sdks/changes/README.md#updatechange) - Update a change entry
* [ListChangeIdentities](docs/sdks/changes/README.md#listchangeidentities) - List identities for a change entry
* [CreateChangeIdentity](docs/sdks/changes/README.md#createchangeidentity) - Create an identity for a change entry
* [DeleteChangeIdentity](docs/sdks/changes/README.md#deletechangeidentity) - Delete an identity from a change entry
* [UpdateChangeIdentity](docs/sdks/changes/README.md#updatechangeidentity) - Update an identity for a change entry
* [ListChangeEvents](docs/sdks/changes/README.md#listchangeevents) - List change events
* [CreateChangeEvent](docs/sdks/changes/README.md#createchangeevent) - Create a change event
* [GetChangeEvent](docs/sdks/changes/README.md#getchangeevent) - Get a change event
* [DeleteChangeEvent](docs/sdks/changes/README.md#deletechangeevent) - Delete a change event
* [UpdateChangeEvent](docs/sdks/changes/README.md#updatechangeevent) - Update a change event
* [ListChangeTypes](docs/sdks/changes/README.md#listchangetypes) - List change types

### [Communication](docs/sdks/communication/README.md)

* [ListStatusUpdateTemplates](docs/sdks/communication/README.md#liststatusupdatetemplates) - List status update templates
* [CreateStatusUpdateTemplate](docs/sdks/communication/README.md#createstatusupdatetemplate) - Create a status update template
* [GetStatusUpdateTemplate](docs/sdks/communication/README.md#getstatusupdatetemplate) - Get a status update template
* [DeleteStatusUpdateTemplate](docs/sdks/communication/README.md#deletestatusupdatetemplate) - Delete a status update template
* [UpdateStatusUpdateTemplate](docs/sdks/communication/README.md#updatestatusupdatetemplate) - Update a status update template

### [Conversations](docs/sdks/conversations/README.md)

* [GetVoteStatus](docs/sdks/conversations/README.md#getvotestatus) - Get votes
* [UpdateVote](docs/sdks/conversations/README.md#updatevote) - Update votes
* [DeleteCommentReaction](docs/sdks/conversations/README.md#deletecommentreaction) - Delete a reaction from a conversation comment
* [ListCommentReactions](docs/sdks/conversations/README.md#listcommentreactions) - List reactions for a conversation comment
* [CreateCommentReaction](docs/sdks/conversations/README.md#createcommentreaction) - Create a reaction for a conversation comment
* [GetComment](docs/sdks/conversations/README.md#getcomment) - Get a conversation comment
* [DeleteComment](docs/sdks/conversations/README.md#deletecomment) - Archive a conversation comment
* [UpdateComment](docs/sdks/conversations/README.md#updatecomment) - Update a conversation comment
* [ListComments](docs/sdks/conversations/README.md#listcomments) - List comments for a conversation
* [CreateComment](docs/sdks/conversations/README.md#createcomment) - Create a conversation comment

### [IncidentSettings](docs/sdks/incidentsettings/README.md)

* [ListIncidentRoles](docs/sdks/incidentsettings/README.md#listincidentroles) - List incident roles
* [CreateIncidentRole](docs/sdks/incidentsettings/README.md#createincidentrole) - Create an incident role
* [GetIncidentRole](docs/sdks/incidentsettings/README.md#getincidentrole) - Get an incident role
* [DeleteIncidentRole](docs/sdks/incidentsettings/README.md#deleteincidentrole) - Archive an incident role
* [UpdateIncidentRole](docs/sdks/incidentsettings/README.md#updateincidentrole) - Update an incident role
* [ValidateIncidentTags](docs/sdks/incidentsettings/README.md#validateincidenttags) - Validate incident tags
* [ListIncidentTags](docs/sdks/incidentsettings/README.md#listincidenttags) - List incident tags
* [ListIncidentTypes](docs/sdks/incidentsettings/README.md#listincidenttypes) - List incident types
* [CreateIncidentType](docs/sdks/incidentsettings/README.md#createincidenttype) - Create an incident type
* [GetIncidentType](docs/sdks/incidentsettings/README.md#getincidenttype) - Get an incident type
* [DeleteIncidentType](docs/sdks/incidentsettings/README.md#deleteincidenttype) - Archive an incident type
* [UpdateIncidentType](docs/sdks/incidentsettings/README.md#updateincidenttype) - Update an incident type
* [ListLifecycleMeasurementDefinitions](docs/sdks/incidentsettings/README.md#listlifecyclemeasurementdefinitions) - List measurement definitions
* [CreateLifecycleMeasurementDefinition](docs/sdks/incidentsettings/README.md#createlifecyclemeasurementdefinition) - Create a measurement definition
* [GetLifecycleMeasurementDefinition](docs/sdks/incidentsettings/README.md#getlifecyclemeasurementdefinition) - Get a measurement definition
* [DeleteLifecycleMeasurementDefinition](docs/sdks/incidentsettings/README.md#deletelifecyclemeasurementdefinition) - Archive a measurement definition
* [UpdateLifecycleMeasurementDefinition](docs/sdks/incidentsettings/README.md#updatelifecyclemeasurementdefinition) - Update a measurement definition
* [ListLifecyclePhases](docs/sdks/incidentsettings/README.md#listlifecyclephases) - List phases and milestones
* [CreateLifecycleMilestone](docs/sdks/incidentsettings/README.md#createlifecyclemilestone) - Create a milestone
* [DeleteLifecycleMilestone](docs/sdks/incidentsettings/README.md#deletelifecyclemilestone) - Delete a milestone
* [UpdateLifecycleMilestone](docs/sdks/incidentsettings/README.md#updatelifecyclemilestone) - Update a milestone
* [ListPriorities](docs/sdks/incidentsettings/README.md#listpriorities) - List priorities
* [CreatePriority](docs/sdks/incidentsettings/README.md#createpriority) - Create a priority
* [GetPriority](docs/sdks/incidentsettings/README.md#getpriority) - Get a priority
* [DeletePriority](docs/sdks/incidentsettings/README.md#deletepriority) - Delete a priority
* [UpdatePriority](docs/sdks/incidentsettings/README.md#updatepriority) - Update a priority
* [ListSeverities](docs/sdks/incidentsettings/README.md#listseverities) - List severities
* [CreateSeverity](docs/sdks/incidentsettings/README.md#createseverity) - Create a severity
* [GetSeverity](docs/sdks/incidentsettings/README.md#getseverity) - Get a severity
* [DeleteSeverity](docs/sdks/incidentsettings/README.md#deleteseverity) - Delete a severity
* [UpdateSeverity](docs/sdks/incidentsettings/README.md#updateseverity) - Update a severity
* [GetSeverityMatrix](docs/sdks/incidentsettings/README.md#getseveritymatrix) - Get severity matrix
* [UpdateSeverityMatrix](docs/sdks/incidentsettings/README.md#updateseveritymatrix) - Update severity matrix
* [ListSeverityMatrixConditions](docs/sdks/incidentsettings/README.md#listseveritymatrixconditions) - List severity matrix conditions
* [CreateSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#createseveritymatrixcondition) - Create a severity matrix condition
* [GetSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#getseveritymatrixcondition) - Get a severity matrix condition
* [DeleteSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#deleteseveritymatrixcondition) - Delete a severity matrix condition
* [UpdateSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#updateseveritymatrixcondition) - Update a severity matrix condition
* [ListSeverityMatrixImpacts](docs/sdks/incidentsettings/README.md#listseveritymatriximpacts) - List severity matrix impacts
* [CreateSeverityMatrixImpact](docs/sdks/incidentsettings/README.md#createseveritymatriximpact) - Create a severity matrix impact
* [DeleteSeverityMatrixImpact](docs/sdks/incidentsettings/README.md#deleteseveritymatriximpact) - Delete a severity matrix impact
* [UpdateSeverityMatrixImpact](docs/sdks/incidentsettings/README.md#updateseveritymatriximpact) - Update a severity matrix impact
* [DeleteCustomFieldDefinition](docs/sdks/incidentsettings/README.md#deletecustomfielddefinition) - Delete a custom field definition
* [UpdateCustomFieldDefinition](docs/sdks/incidentsettings/README.md#updatecustomfielddefinition) - Update a custom field definition
* [ListCustomFieldDefinitions](docs/sdks/incidentsettings/README.md#listcustomfielddefinitions) - List custom field definitions
* [CreateCustomFieldDefinition](docs/sdks/incidentsettings/README.md#createcustomfielddefinition) - Create a custom field definition
* [ListCustomFieldSelectOptions](docs/sdks/incidentsettings/README.md#listcustomfieldselectoptions) - Get available values for a custom field
* [AppendFormDataOnSelectedValueGet](docs/sdks/incidentsettings/README.md#appendformdataonselectedvalueget) - Get data for a form field on select
* [GetFormConfiguration](docs/sdks/incidentsettings/README.md#getformconfiguration) - Get a form configuration

### [Incidents](docs/sdks/incidents/README.md)

* [ListIncidents](docs/sdks/incidents/README.md#listincidents) - List incidents
* [CreateIncident](docs/sdks/incidents/README.md#createincident) - Create an incident
* [GetIncidentChannel](docs/sdks/incidents/README.md#getincidentchannel) - Get chat channel information for an incident
* [CloseIncident](docs/sdks/incidents/README.md#closeincident) - Close an incident
* [ResolveIncident](docs/sdks/incidents/README.md#resolveincident) - Resolve an incident
* [GetIncident](docs/sdks/incidents/README.md#getincident) - Get an incident
* [DeleteIncident](docs/sdks/incidents/README.md#deleteincident) - Archive an incident
* [UpdateIncident](docs/sdks/incidents/README.md#updateincident) - Update an incident
* [UnarchiveIncident](docs/sdks/incidents/README.md#unarchiveincident) - Unarchive an incident
* [BulkUpdateIncidentMilestones](docs/sdks/incidents/README.md#bulkupdateincidentmilestones) - Update milestone times
* [ListIncidentMilestones](docs/sdks/incidents/README.md#listincidentmilestones) - List incident milestones
* [ListIncidentChangeEvents](docs/sdks/incidents/README.md#listincidentchangeevents) - List related changes on an incident
* [CreateIncidentChangeEvent](docs/sdks/incidents/README.md#createincidentchangeevent) - Add a related change to an incident
* [UpdateIncidentChangeEvent](docs/sdks/incidents/README.md#updateincidentchangeevent) - Update a change attached to an incident
* [ListIncidentStatusPages](docs/sdks/incidents/README.md#listincidentstatuspages) - List status pages for an incident
* [CreateIncidentStatusPage](docs/sdks/incidents/README.md#createincidentstatuspage) - Add a status page to an incident
* [ListIncidentLinks](docs/sdks/incidents/README.md#listincidentlinks) - List links on an incident
* [CreateIncidentLink](docs/sdks/incidents/README.md#createincidentlink) - Add a link to an incident
* [UpdateIncidentLink](docs/sdks/incidents/README.md#updateincidentlink) - Update the external incident link
* [DeleteIncidentLink](docs/sdks/incidents/README.md#deleteincidentlink) - Remove a link from an incident
* [UpdateTranscriptAttribution](docs/sdks/incidents/README.md#updatetranscriptattribution) - Update the attribution of a transcript
* [ListTranscriptEntries](docs/sdks/incidents/README.md#listtranscriptentries) - Lists all of the messages in the incident's transcript
* [DeleteTranscriptEntry](docs/sdks/incidents/README.md#deletetranscriptentry) - Delete a transcript from an incident
* [ListIncidentConferenceBridges](docs/sdks/incidents/README.md#listincidentconferencebridges) - Retrieve all conference bridges for an incident
* [GetConferenceBridgeTranslation](docs/sdks/incidents/README.md#getconferencebridgetranslation) - Retrieve the translations for a specific conference bridge
* [ListSimilarIncidents](docs/sdks/incidents/README.md#listsimilarincidents) - List similar incidents
* [ListIncidentAttachments](docs/sdks/incidents/README.md#listincidentattachments) - List attachments for an incident
* [CreateIncidentAttachment](docs/sdks/incidents/README.md#createincidentattachment) - Add an attachment to the incident timeline
* [ListIncidentEvents](docs/sdks/incidents/README.md#listincidentevents) - List events for an incident
* [GetIncidentEvent](docs/sdks/incidents/README.md#getincidentevent) - Get an incident event
* [DeleteIncidentEvent](docs/sdks/incidents/README.md#deleteincidentevent) - Delete an incident event
* [UpdateIncidentEvent](docs/sdks/incidents/README.md#updateincidentevent) - Update an incident event
* [UpdateIncidentImpactPut](docs/sdks/incidents/README.md#updateincidentimpactput) - Update impacts for an incident
* [UpdateIncidentImpactPatch](docs/sdks/incidents/README.md#updateincidentimpactpatch) - Update impacts for an incident
* [ListIncidentImpacts](docs/sdks/incidents/README.md#listincidentimpacts) - List impacted infrastructure for an incident
* [CreateIncidentImpact](docs/sdks/incidents/README.md#createincidentimpact) - Add impacted infrastructure to an incident
* [DeleteIncidentImpact](docs/sdks/incidents/README.md#deleteincidentimpact) - Remove impacted infrastructure from an incident
* [CreateIncidentNote](docs/sdks/incidents/README.md#createincidentnote) - Add a note to an incident
* [UpdateIncidentNote](docs/sdks/incidents/README.md#updateincidentnote) - Update a note
* [CreateIncidentChatMessage](docs/sdks/incidents/README.md#createincidentchatmessage) - Add a chat message to an incident
* [DeleteIncidentChatMessage](docs/sdks/incidents/README.md#deleteincidentchatmessage) - Delete a chat message from an incident
* [UpdateIncidentChatMessage](docs/sdks/incidents/README.md#updateincidentchatmessage) - Update a chat message on an incident
* [ListIncidentRoleAssignments](docs/sdks/incidents/README.md#listincidentroleassignments) - List incident assignees
* [CreateIncidentRoleAssignment](docs/sdks/incidents/README.md#createincidentroleassignment) - Assign a user to an incident
* [DeleteIncidentRoleAssignment](docs/sdks/incidents/README.md#deleteincidentroleassignment) - Unassign a user from an incident
* [CreateIncidentTeamAssignment](docs/sdks/incidents/README.md#createincidentteamassignment) - Assign a team to an incident
* [DeleteIncidentTeamAssignment](docs/sdks/incidents/README.md#deleteincidentteamassignment) - Unassign a team from an incident
* [GetIncidentUser](docs/sdks/incidents/README.md#getincidentuser) - Get the current user's incident role
* [GetIncidentRelationships](docs/sdks/incidents/README.md#getincidentrelationships) - List incident relationships
* [ListScheduledMaintenances](docs/sdks/incidents/README.md#listscheduledmaintenances) - List scheduled maintenance events
* [CreateScheduledMaintenance](docs/sdks/incidents/README.md#createscheduledmaintenance) - Create a scheduled maintenance event
* [GetScheduledMaintenance](docs/sdks/incidents/README.md#getscheduledmaintenance) - Get a scheduled maintenance event
* [DeleteScheduledMaintenance](docs/sdks/incidents/README.md#deletescheduledmaintenance) - Delete a scheduled maintenance event
* [UpdateScheduledMaintenance](docs/sdks/incidents/README.md#updatescheduledmaintenance) - Update a scheduled maintenance event
* [GetAiIncidentSummaryVoteStatus](docs/sdks/incidents/README.md#getaiincidentsummaryvotestatus) - Get the current user's vote status for an AI-generated incident summary
* [VoteAiIncidentSummary](docs/sdks/incidents/README.md#voteaiincidentsummary) - Vote on an AI-generated incident summary

### [Integrations](docs/sdks/integrations/README.md)

* [ListIntegrations](docs/sdks/integrations/README.md#listintegrations) - List integrations
* [GetIntegration](docs/sdks/integrations/README.md#getintegration) - Get an integration
* [UpdateFieldMap](docs/sdks/integrations/README.md#updatefieldmap) - Update field mapping
* [ListFieldMapAvailableFields](docs/sdks/integrations/README.md#listfieldmapavailablefields) - List available fields for field mapping
* [ListAuthedProviders](docs/sdks/integrations/README.md#listauthedproviders) - Lists the available and configured integrations
* [UpdateAuthedProvider](docs/sdks/integrations/README.md#updateauthedprovider) - Get an authed provider
* [ListConnections](docs/sdks/integrations/README.md#listconnections) - List integration connections
* [CreateConnection](docs/sdks/integrations/README.md#createconnection) - Create a new integration connection
* [RefreshConnection](docs/sdks/integrations/README.md#refreshconnection) - Refresh an integration connection's incident role schedules
* [UpdateConnection](docs/sdks/integrations/README.md#updateconnection) - Update an integration connection
* [ListConnectionStatuses](docs/sdks/integrations/README.md#listconnectionstatuses) - Get integration connection status
* [ListConnectionStatusesBySlug](docs/sdks/integrations/README.md#listconnectionstatusesbyslug) - Get an integration connection status
* [ListConnectionStatusesBySlugAndID](docs/sdks/integrations/README.md#listconnectionstatusesbyslugandid) - Get an integration connection status
* [ListAwsConnections](docs/sdks/integrations/README.md#listawsconnections) - List AWS connections
* [GetAwsConnection](docs/sdks/integrations/README.md#getawsconnection) - Get an AWS connection
* [UpdateAwsConnection](docs/sdks/integrations/README.md#updateawsconnection) - Update an AWS connection
* [ListAwsCloudtrailBatches](docs/sdks/integrations/README.md#listawscloudtrailbatches) - List CloudTrail batches
* [GetAwsCloudtrailBatch](docs/sdks/integrations/README.md#getawscloudtrailbatch) - Get a CloudTrail batch
* [UpdateAwsCloudtrailBatch](docs/sdks/integrations/README.md#updateawscloudtrailbatch) - Update a CloudTrail batch
* [ListAwsCloudtrailBatchEvents](docs/sdks/integrations/README.md#listawscloudtrailbatchevents) - List events for an AWS CloudTrail batch
* [SearchConfluenceSpaces](docs/sdks/integrations/README.md#searchconfluencespaces) - List Confluence spaces
* [ListSlackWorkspaces](docs/sdks/integrations/README.md#listslackworkspaces) - List Slack workspaces
* [SearchSlackChannels](docs/sdks/integrations/README.md#searchslackchannels) - List Slack channels
* [TestSlackChannel](docs/sdks/integrations/README.md#testslackchannel) - Test a Slack channel
* [ListSlackUsergroups](docs/sdks/integrations/README.md#listslackusergroups) - List Slack user groups
* [ListSlackEmojiActions](docs/sdks/integrations/README.md#listslackemojiactions) - List Slack emoji actions
* [CreateSlackEmojiAction](docs/sdks/integrations/README.md#createslackemojiaction) - Create a new Slack emoji action
* [GetSlackEmojiAction](docs/sdks/integrations/README.md#getslackemojiaction) - Get a Slack emoji action
* [DeleteSlackEmojiAction](docs/sdks/integrations/README.md#deleteslackemojiaction) - Delete a Slack emoji action
* [UpdateSlackEmojiAction](docs/sdks/integrations/README.md#updateslackemojiaction) - Update a Slack emoji action
* [ListStatuspageConnections](docs/sdks/integrations/README.md#liststatuspageconnections) - List Statuspage connections
* [GetStatuspageConnection](docs/sdks/integrations/README.md#getstatuspageconnection) - Get a Statuspage connection
* [DeleteStatuspageConnection](docs/sdks/integrations/README.md#deletestatuspageconnection) - Delete a Statuspage connection
* [UpdateStatuspageConnection](docs/sdks/integrations/README.md#updatestatuspageconnection) - Update a Statuspage connection
* [ListStatuspageConnectionPages](docs/sdks/integrations/README.md#liststatuspageconnectionpages) - List StatusPage pages for a connection
* [SearchZendeskTickets](docs/sdks/integrations/README.md#searchzendesktickets) - Search for Zendesk tickets
* [GetZendeskCustomerSupportIssue](docs/sdks/integrations/README.md#getzendeskcustomersupportissue) - Search for Zendesk tickets

### [MetricsReporting](docs/sdks/metricsreporting/README.md)

* [GetMeanTimeReport](docs/sdks/metricsreporting/README.md#getmeantimereport) - Get mean time metrics for incidents
* [ListTicketFunnelMetrics](docs/sdks/metricsreporting/README.md#listticketfunnelmetrics) - List ticket task and follow up creation and completion metrics
* [ListRetrospectiveMetrics](docs/sdks/metricsreporting/README.md#listretrospectivemetrics) - List retrospective metrics
* [ListMilestoneFunnelMetrics](docs/sdks/metricsreporting/README.md#listmilestonefunnelmetrics) - List milestone funnel metrics
* [ListUserInvolvementMetrics](docs/sdks/metricsreporting/README.md#listuserinvolvementmetrics) - List user metrics
* [ListIncidentMetrics](docs/sdks/metricsreporting/README.md#listincidentmetrics) - List incident metrics and analytics
* [ListMttxMetrics](docs/sdks/metricsreporting/README.md#listmttxmetrics) - Get infrastructure metrics
* [ListInfrastructureTypeMetrics](docs/sdks/metricsreporting/README.md#listinfrastructuretypemetrics) - List metrics for a component type
* [ListInfrastructureMetrics](docs/sdks/metricsreporting/README.md#listinfrastructuremetrics) - Get metrics for a component
* [GetSavedSearch](docs/sdks/metricsreporting/README.md#getsavedsearch) - Get a saved search
* [DeleteSavedSearch](docs/sdks/metricsreporting/README.md#deletesavedsearch) - Delete a saved search
* [UpdateSavedSearch](docs/sdks/metricsreporting/README.md#updatesavedsearch) - Update a saved search
* [ListSavedSearches](docs/sdks/metricsreporting/README.md#listsavedsearches) - List saved searches
* [CreateSavedSearch](docs/sdks/metricsreporting/README.md#createsavedsearch) - Create a saved search
* [GetSignalsTimeseriesAnalytics](docs/sdks/metricsreporting/README.md#getsignalstimeseriesanalytics) - Generate timeseries alert metrics
* [GetSignalsGroupedMetrics](docs/sdks/metricsreporting/README.md#getsignalsgroupedmetrics) - Generate grouped alert metrics
* [GetSignalsMttxAnalytics](docs/sdks/metricsreporting/README.md#getsignalsmttxanalytics) - Get MTTX analytics for signals
* [GetSignalsNoiseAnalytics](docs/sdks/metricsreporting/README.md#getsignalsnoiseanalytics) - Get noise analytics for signals
* [ExportSignalsShiftAnalytics](docs/sdks/metricsreporting/README.md#exportsignalsshiftanalytics) - Export on-call hours report

### [Pages](docs/sdks/pages/README.md)

* [CreateSignalsPage](docs/sdks/pages/README.md#createsignalspage) - Page a user, team, on-call schedule, or escalation policy

### [Permissions](docs/sdks/permissions/README.md)

* [ListPermissions](docs/sdks/permissions/README.md#listpermissions) - List permissions
* [ListCurrentUserPermissions](docs/sdks/permissions/README.md#listcurrentuserpermissions) - Get all permissions for the current user
* [ListTeamPermissions](docs/sdks/permissions/README.md#listteampermissions) - Get all permissions for a team

### [Retrospectives](docs/sdks/retrospectives/README.md)

* [ShareIncidentRetrospectives](docs/sdks/retrospectives/README.md#shareincidentretrospectives) - Share an incident's retrospective
* [ExportIncidentRetrospectivesMarkdown](docs/sdks/retrospectives/README.md#exportincidentretrospectivesmarkdown) - Export an incident's retrospective(s) as markdown
* [ExportIncidentRetrospectives](docs/sdks/retrospectives/README.md#exportincidentretrospectives) - Export an incident's retrospective(s)
* [ListIncidentRetrospectives](docs/sdks/retrospectives/README.md#listincidentretrospectives) - All attached retrospectives for an incident
* [CreateIncidentRetrospective](docs/sdks/retrospectives/README.md#createincidentretrospective) - Create a new retrospective on the incident using the template
* [UpdateIncidentRetrospective](docs/sdks/retrospectives/README.md#updateincidentretrospective) - Update a retrospective on the incident
* [CreateIncidentRetrospectiveField](docs/sdks/retrospectives/README.md#createincidentretrospectivefield) - Appends a new incident retrospective field to an incident retrospective
* [GetIncidentRetrospectiveField](docs/sdks/retrospectives/README.md#getincidentretrospectivefield) - Get a retrospective field
* [UpdateIncidentRetrospectiveField](docs/sdks/retrospectives/README.md#updateincidentretrospectivefield) - Update the value on a retrospective field
* [CreateIncidentRetrospectiveDynamicInput](docs/sdks/retrospectives/README.md#createincidentretrospectivedynamicinput) - Add a new dynamic input field to a retrospective's dynamic input group field
* [DeleteIncidentRetrospectiveDynamicInput](docs/sdks/retrospectives/README.md#deleteincidentretrospectivedynamicinput) - Removes a dynamic input from a retrospective's dynamic input group field
* [ListRetrospectives](docs/sdks/retrospectives/README.md#listretrospectives) - List retrospective reports
* [ListPostMortemReports](docs/sdks/retrospectives/README.md#listpostmortemreports) - List retrospective reports
* [CreatePostMortemReport](docs/sdks/retrospectives/README.md#createpostmortemreport) - Create a retrospective report
* [GetPostMortemReport](docs/sdks/retrospectives/README.md#getpostmortemreport) - Get a retrospective report
* [UpdatePostMortemReport](docs/sdks/retrospectives/README.md#updatepostmortemreport) - Update a retrospective report
* [ListPostMortemReasons](docs/sdks/retrospectives/README.md#listpostmortemreasons) - List contributing factors for a retrospective report
* [CreatePostMortemReason](docs/sdks/retrospectives/README.md#createpostmortemreason) - Create a contributing factor for a retrospective report
* [DeletePostMortemReason](docs/sdks/retrospectives/README.md#deletepostmortemreason) - Delete a contributing factor from a retrospective report
* [UpdatePostMortemReason](docs/sdks/retrospectives/README.md#updatepostmortemreason) - Update a contributing factor in a retrospective report
* [ReorderPostMortemReasons](docs/sdks/retrospectives/README.md#reorderpostmortemreasons) - Reorder a contributing factor for a retrospective report
* [PublishPostMortemReport](docs/sdks/retrospectives/README.md#publishpostmortemreport) - Publish a retrospective report
* [UpdatePostMortemField](docs/sdks/retrospectives/README.md#updatepostmortemfield) - Update a retrospective field
* [ListPostMortemQuestions](docs/sdks/retrospectives/README.md#listpostmortemquestions) - List retrospective questions
* [UpdatePostMortemQuestions](docs/sdks/retrospectives/README.md#updatepostmortemquestions) - Update retrospective questions
* [GetPostMortemQuestion](docs/sdks/retrospectives/README.md#getpostmortemquestion) - Get a retrospective question
* [ListRetrospectiveTemplates](docs/sdks/retrospectives/README.md#listretrospectivetemplates) - List retrospective templates
* [CreateRetrospectiveTemplate](docs/sdks/retrospectives/README.md#createretrospectivetemplate) - Create a retrospective template
* [GetRetrospectiveTemplate](docs/sdks/retrospectives/README.md#getretrospectivetemplate) - Get a retrospective template
* [DeleteRetrospectiveTemplate](docs/sdks/retrospectives/README.md#deleteretrospectivetemplate) - Delete a retrospective template
* [UpdateRetrospectiveTemplate](docs/sdks/retrospectives/README.md#updateretrospectivetemplate) - Update a retrospective template

### [Roles](docs/sdks/roles/README.md)

* [ListRoles](docs/sdks/roles/README.md#listroles) - Get all roles
* [CreateRole](docs/sdks/roles/README.md#createrole) - Create a role
* [GetRole](docs/sdks/roles/README.md#getrole) - Get a role
* [DeleteRole](docs/sdks/roles/README.md#deleterole) - Delete a role
* [UpdateRole](docs/sdks/roles/README.md#updaterole) - Update a role

### [Runbooks](docs/sdks/runbooks/README.md)

* [ListRunbookActions](docs/sdks/runbooks/README.md#listrunbookactions) - List runbook actions
* [ListRunbookExecutions](docs/sdks/runbooks/README.md#listrunbookexecutions) - List runbook executions
* [CreateRunbookExecution](docs/sdks/runbooks/README.md#createrunbookexecution) - Create a runbook execution
* [GetRunbookExecution](docs/sdks/runbooks/README.md#getrunbookexecution) - Get a runbook execution
* [DeleteRunbookExecution](docs/sdks/runbooks/README.md#deleterunbookexecution) - Terminate a runbook execution
* [UpdateRunbookExecutionStep](docs/sdks/runbooks/README.md#updaterunbookexecutionstep) - Update a runbook step execution
* [GetRunbookExecutionStepScript](docs/sdks/runbooks/README.md#getrunbookexecutionstepscript) - Get a step's bash script
* [UpdateRunbookExecutionStepScript](docs/sdks/runbooks/README.md#updaterunbookexecutionstepscript) - Update a script step's execution status
* [GetRunbookActionFieldOptions](docs/sdks/runbooks/README.md#getrunbookactionfieldoptions) - List select options for a runbook integration action field
* [ListRunbooks](docs/sdks/runbooks/README.md#listrunbooks) - List runbooks
* [CreateRunbook](docs/sdks/runbooks/README.md#createrunbook) - Create a runbook
* [GetRunbook](docs/sdks/runbooks/README.md#getrunbook) - Get a runbook
* [UpdateRunbook](docs/sdks/runbooks/README.md#updaterunbook) - Update a runbook
* [DeleteRunbook](docs/sdks/runbooks/README.md#deleterunbook) - Delete a runbook
* [ListRunbookAudits](docs/sdks/runbooks/README.md#listrunbookaudits) - List runbook audits

### [Scim](docs/sdks/scim/README.md)

* [GetScimGroup](docs/sdks/scim/README.md#getscimgroup) - Get a SCIM group
* [UpdateScimGroup](docs/sdks/scim/README.md#updatescimgroup) - Update a SCIM group and assign members
* [DeleteScimGroup](docs/sdks/scim/README.md#deletescimgroup) - Delete a SCIM group
* [PatchScimGroup](docs/sdks/scim/README.md#patchscimgroup) - Partially update a SCIM group
* [ListScimGroups](docs/sdks/scim/README.md#listscimgroups) - List SCIM groups
* [CreateScimGroup](docs/sdks/scim/README.md#createscimgroup) - Create a SCIM group and assign members
* [GetScimUser](docs/sdks/scim/README.md#getscimuser) - Get a SCIM user
* [UpdateScimUser](docs/sdks/scim/README.md#updatescimuser) - Update a User from SCIM data
* [DeleteScimUser](docs/sdks/scim/README.md#deletescimuser) - Delete a User matching SCIM data
* [PatchScimUser](docs/sdks/scim/README.md#patchscimuser) - Update a User from SCIM data
* [ListScimUsers](docs/sdks/scim/README.md#listscimusers) - List SCIM users
* [CreateScimUser](docs/sdks/scim/README.md#createscimuser) - Create a User from SCIM data

### [Signals](docs/sdks/signals/README.md)

* [GetSupportHoursSchedule](docs/sdks/signals/README.md#getsupporthoursschedule) - Get support hours schedule
* [DeleteSupportHoursSchedule](docs/sdks/signals/README.md#deletesupporthoursschedule) - Delete a specific support hours schedule
* [UpdateSupportHoursSchedule](docs/sdks/signals/README.md#updatesupporthoursschedule) - Update support hours schedule
* [ListTeamEscalationPolicies](docs/sdks/signals/README.md#listteamescalationpolicies) - List escalation policies for a team
* [CreateTeamEscalationPolicy](docs/sdks/signals/README.md#createteamescalationpolicy) - Create an escalation policy for a team
* [GetTeamEscalationPolicy](docs/sdks/signals/README.md#getteamescalationpolicy) - Get an escalation policy for a team
* [DeleteTeamEscalationPolicy](docs/sdks/signals/README.md#deleteteamescalationpolicy) - Delete an escalation policy for a team
* [UpdateTeamEscalationPolicy](docs/sdks/signals/README.md#updateteamescalationpolicy) - Update an escalation policy for a team
* [PreviewTeamOnCallSchedule](docs/sdks/signals/README.md#previewteamoncallschedule) - Preview a new on-call schedule for a team
* [ListTeamOnCallSchedules](docs/sdks/signals/README.md#listteamoncallschedules) - List on-call schedules for a team
* [CreateTeamOnCallSchedule](docs/sdks/signals/README.md#createteamoncallschedule) - Create an on-call schedule for a team
* [GetTeamOnCallSchedule](docs/sdks/signals/README.md#getteamoncallschedule) - Get an on-call schedule for a team
* [DeleteTeamOnCallSchedule](docs/sdks/signals/README.md#deleteteamoncallschedule) - Delete an on-call schedule for a team
* [UpdateTeamOnCallSchedule](docs/sdks/signals/README.md#updateteamoncallschedule) - Update an on-call schedule for a team
* [PreviewOnCallScheduleRotation](docs/sdks/signals/README.md#previewoncallschedulerotation) - Preview an on-call rotation
* [CreateOnCallScheduleRotation](docs/sdks/signals/README.md#createoncallschedulerotation) - Create a new on-call rotation
* [CopyOnCallScheduleRotation](docs/sdks/signals/README.md#copyoncallschedulerotation) - Copy an on-call schedule's rotation
* [GetOnCallScheduleRotation](docs/sdks/signals/README.md#getoncallschedulerotation) - Get an on-call rotation
* [DeleteOnCallScheduleRotation](docs/sdks/signals/README.md#deleteoncallschedulerotation) - Delete an on-call schedule's rotation
* [UpdateOnCallScheduleRotation](docs/sdks/signals/README.md#updateoncallschedulerotation) - Update an on-call schedule's rotation
* [CreateOnCallScheduleRotationOverride](docs/sdks/signals/README.md#createoncallschedulerotationoverride) - Override one or more shifts in an on-call rotation
* [CreateOnCallShift](docs/sdks/signals/README.md#createoncallshift) - [DEPRECATED] Create a shift for an on-call schedule
* [GetOnCallShift](docs/sdks/signals/README.md#getoncallshift) - [DEPRECATED] Get an on-call shift for a team schedule
* [DeleteOnCallShift](docs/sdks/signals/README.md#deleteoncallshift) - [DEPRECATED] Delete an on-call shift from a team schedule
* [UpdateOnCallShift](docs/sdks/signals/README.md#updateoncallshift) - [DEPRECATED] Update an on-call shift for a team schedule
* [ListTeamSignalRules](docs/sdks/signals/README.md#listteamsignalrules) - List Signals rules
* [CreateTeamSignalRule](docs/sdks/signals/README.md#createteamsignalrule) - Create a Signals rule
* [GetTeamSignalRule](docs/sdks/signals/README.md#getteamsignalrule) - Get a Signals rule
* [DeleteTeamSignalRule](docs/sdks/signals/README.md#deleteteamsignalrule) - Delete a Signals rule
* [UpdateTeamSignalRule](docs/sdks/signals/README.md#updateteamsignalrule) - Update a Signals rule
* [ListSignalsEventSources](docs/sdks/signals/README.md#listsignalseventsources) - List event sources for Signals
* [CreateSignalsEventSource](docs/sdks/signals/README.md#createsignalseventsource) - Create an event source for Signals
* [GetSignalsEventSource](docs/sdks/signals/README.md#getsignalseventsource) - Get an event source for Signals
* [DeleteSignalsEventSource](docs/sdks/signals/README.md#deletesignalseventsource) - Delete an event source for Signals
* [GetSignalsHackerMode](docs/sdks/signals/README.md#getsignalshackermode) - Get hacker mode status
* [ListSignalsAlertGroupingConfigurations](docs/sdks/signals/README.md#listsignalsalertgroupingconfigurations) - List alert grouping configurations.
* [CreateSignalsAlertGroupingConfiguration](docs/sdks/signals/README.md#createsignalsalertgroupingconfiguration) - Create an alert grouping configuration.
* [GetSignalsAlertGroupingConfiguration](docs/sdks/signals/README.md#getsignalsalertgroupingconfiguration) - Get an alert grouping configuration.
* [DeleteSignalsAlertGroupingConfiguration](docs/sdks/signals/README.md#deletesignalsalertgroupingconfiguration) - Delete an alert grouping configuration.
* [UpdateSignalsAlertGroupingConfiguration](docs/sdks/signals/README.md#updatesignalsalertgroupingconfiguration) - Update an alert grouping configuration.
* [ListSignalsEmailTargets](docs/sdks/signals/README.md#listsignalsemailtargets) - List email targets for signals
* [CreateSignalsEmailTarget](docs/sdks/signals/README.md#createsignalsemailtarget) - Create an email target for signals
* [GetSignalsEmailTarget](docs/sdks/signals/README.md#getsignalsemailtarget) - Get a signal email target
* [DeleteSignalsEmailTarget](docs/sdks/signals/README.md#deletesignalsemailtarget) - Delete a signal email target
* [UpdateSignalsEmailTarget](docs/sdks/signals/README.md#updatesignalsemailtarget) - Update an email target
* [ListSignalsWebhookTargets](docs/sdks/signals/README.md#listsignalswebhooktargets) - List webhook targets
* [CreateSignalsWebhookTarget](docs/sdks/signals/README.md#createsignalswebhooktarget) - Create a webhook target
* [GetSignalsWebhookTarget](docs/sdks/signals/README.md#getsignalswebhooktarget) - Get a webhook target
* [DeleteSignalsWebhookTarget](docs/sdks/signals/README.md#deletesignalswebhooktarget) - Delete a webhook target
* [UpdateSignalsWebhookTarget](docs/sdks/signals/README.md#updatesignalswebhooktarget) - Update a webhook target
* [ListSignalsHeartbeatEndpointConfigurations](docs/sdks/signals/README.md#listsignalsheartbeatendpointconfigurations) - List heartbeat endpoint configurations
* [CreateSignalsHeartbeatEndpointConfiguration](docs/sdks/signals/README.md#createsignalsheartbeatendpointconfiguration) - Create a heartbeat endpoint configuration
* [GetSignalsHeartbeatEndpointStatus](docs/sdks/signals/README.md#getsignalsheartbeatendpointstatus) - Get heartbeat endpoint status
* [GetSignalsHeartbeatEndpointURL](docs/sdks/signals/README.md#getsignalsheartbeatendpointurl) - Get heartbeat endpoint URL
* [GetSignalsHeartbeatEndpointConfiguration](docs/sdks/signals/README.md#getsignalsheartbeatendpointconfiguration) - Get a heartbeat endpoint configuration
* [DeleteSignalsHeartbeatEndpointConfiguration](docs/sdks/signals/README.md#deletesignalsheartbeatendpointconfiguration) - Delete a heartbeat endpoint configuration
* [UpdateSignalsHeartbeatEndpointConfiguration](docs/sdks/signals/README.md#updatesignalsheartbeatendpointconfiguration) - Update a heartbeat endpoint configuration
* [ListNotificationPolicySettings](docs/sdks/signals/README.md#listnotificationpolicysettings) - List notification policies
* [CreateNotificationPolicy](docs/sdks/signals/README.md#createnotificationpolicy) - Create a notification policy
* [GetNotificationPolicy](docs/sdks/signals/README.md#getnotificationpolicy) - Get a notification policy
* [DeleteNotificationPolicy](docs/sdks/signals/README.md#deletenotificationpolicy) - Delete a notification policy
* [UpdateNotificationPolicy](docs/sdks/signals/README.md#updatenotificationpolicy) - Update a notification policy
* [ListUserNotificationSettingsByUserID](docs/sdks/signals/README.md#listusernotificationsettingsbyuserid) - List notification settings for a user
* [ListSignalsTransposers](docs/sdks/signals/README.md#listsignalstransposers) - List signal transposers
* [GetSignalsIngestURL](docs/sdks/signals/README.md#getsignalsingesturl) - Get the signals ingestion URL
* [DebugSignalsExpression](docs/sdks/signals/README.md#debugsignalsexpression) - Debug Signals expressions
* [ListOrganizationOnCallSchedules](docs/sdks/signals/README.md#listorganizationoncallschedules) - List who's on call for the organization

### [StatusPages](docs/sdks/statuspages/README.md)

* [DeleteIncidentStatusPage](docs/sdks/statuspages/README.md#deleteincidentstatuspage) - Remove a status page from an incident
* [ListNuncConnections](docs/sdks/statuspages/README.md#listnuncconnections) - List status pages
* [CreateNuncConnection](docs/sdks/statuspages/README.md#createnuncconnection) - Create a status page
* [ListEmailSubscribers](docs/sdks/statuspages/README.md#listemailsubscribers) - List status page subscribers
* [CreateEmailSubscriber](docs/sdks/statuspages/README.md#createemailsubscriber) - Add subscribers to a status page
* [DeleteEmailSubscriber](docs/sdks/statuspages/README.md#deleteemailsubscriber) - Remove subscribers from a status page
* [GetNuncConnection](docs/sdks/statuspages/README.md#getnuncconnection) - Get a status page
* [UpdateNuncConnection](docs/sdks/statuspages/README.md#updatenuncconnection) - Update a status page
* [DeleteNuncConnection](docs/sdks/statuspages/README.md#deletenuncconnection) - Delete a status page
* [UnpublishNuncConnection](docs/sdks/statuspages/README.md#unpublishnuncconnection) - Unpublish a status page
* [PublishNuncConnection](docs/sdks/statuspages/README.md#publishnuncconnection) - Publish a status page
* [DeleteNuncComponentGroup](docs/sdks/statuspages/README.md#deletenunccomponentgroup) - Delete a status page component group
* [UpdateNuncComponentGroup](docs/sdks/statuspages/README.md#updatenunccomponentgroup) - Update a status page component group
* [CreateNuncComponentGroup](docs/sdks/statuspages/README.md#createnunccomponentgroup) - Create a component group for a status page
* [DeleteNuncLink](docs/sdks/statuspages/README.md#deletenunclink) - Delete a status page link
* [UpdateNuncLink](docs/sdks/statuspages/README.md#updatenunclink) - Update a status page link
* [CreateNuncLink](docs/sdks/statuspages/README.md#createnunclink) - Add link to a status page
* [UpdateNuncImage](docs/sdks/statuspages/README.md#updatenuncimage) - Upload an image for a status page
* [DeleteNuncImage](docs/sdks/statuspages/README.md#deletenuncimage) - Delete an image from a status page
* [DeleteNuncSubscription](docs/sdks/statuspages/README.md#deletenuncsubscription) - Unsubscribe from status page notifications
* [CreateNuncSubscription](docs/sdks/statuspages/README.md#createnuncsubscription) - Create a status page subscription

### [Tasks](docs/sdks/tasks/README.md)

* [CreateIncidentTaskList](docs/sdks/tasks/README.md#createincidenttasklist) - Add tasks from a task list to an incident
* [ListIncidentTasks](docs/sdks/tasks/README.md#listincidenttasks) - List tasks for an incident
* [CreateIncidentTask](docs/sdks/tasks/README.md#createincidenttask) - Create an incident task
* [GetIncidentTask](docs/sdks/tasks/README.md#getincidenttask) - Get an incident task
* [DeleteIncidentTask](docs/sdks/tasks/README.md#deleteincidenttask) - Delete an incident task
* [UpdateIncidentTask](docs/sdks/tasks/README.md#updateincidenttask) - Update an incident task
* [ConvertIncidentTask](docs/sdks/tasks/README.md#convertincidenttask) - Convert a task to a follow-up
* [ListTaskLists](docs/sdks/tasks/README.md#listtasklists) - List task lists
* [CreateTaskList](docs/sdks/tasks/README.md#createtasklist) - Create a task list
* [GetTaskList](docs/sdks/tasks/README.md#gettasklist) - Get a task list
* [DeleteTaskList](docs/sdks/tasks/README.md#deletetasklist) - Delete a task list
* [UpdateTaskList](docs/sdks/tasks/README.md#updatetasklist) - Update a task list
* [ListChecklistTemplates](docs/sdks/tasks/README.md#listchecklisttemplates) - List checklist templates
* [CreateChecklistTemplate](docs/sdks/tasks/README.md#createchecklisttemplate) - Create a checklist template
* [GetChecklistTemplate](docs/sdks/tasks/README.md#getchecklisttemplate) - Get a checklist template
* [DeleteChecklistTemplate](docs/sdks/tasks/README.md#deletechecklisttemplate) - Archive a checklist template
* [UpdateChecklistTemplate](docs/sdks/tasks/README.md#updatechecklisttemplate) - Update a checklist template

### [Teams](docs/sdks/teams/README.md)

* [ListTeams](docs/sdks/teams/README.md#listteams) - List teams
* [CreateTeam](docs/sdks/teams/README.md#createteam) - Create a team
* [GetTeam](docs/sdks/teams/README.md#getteam) - Get a team
* [DeleteTeam](docs/sdks/teams/README.md#deleteteam) - Archive a team
* [UpdateTeam](docs/sdks/teams/README.md#updateteam) - Update a team
* [ListSchedules](docs/sdks/teams/README.md#listschedules) - List schedules

### [Ticketing](docs/sdks/ticketing/README.md)

* [ListTickets](docs/sdks/ticketing/README.md#listtickets) - List tickets
* [CreateTicket](docs/sdks/ticketing/README.md#createticket) - Create a ticket
* [GetTicket](docs/sdks/ticketing/README.md#getticket) - Get a ticket
* [DeleteTicket](docs/sdks/ticketing/README.md#deleteticket) - Archive a ticket
* [UpdateTicket](docs/sdks/ticketing/README.md#updateticket) - Update a ticket
* [ListTicketingProjects](docs/sdks/ticketing/README.md#listticketingprojects) - List ticketing projects
* [GetTicketingProject](docs/sdks/ticketing/README.md#getticketingproject) - Get a ticketing project
* [GetConfigurationOptions](docs/sdks/ticketing/README.md#getconfigurationoptions) - List configuration options for a ticketing project
* [GetOptionsForField](docs/sdks/ticketing/README.md#getoptionsforfield) - List a field's configuration options for a ticketing project
* [ListAvailableTicketingFieldMaps](docs/sdks/ticketing/README.md#listavailableticketingfieldmaps) - List available fields for ticket field mapping
* [CreateTicketingFieldMap](docs/sdks/ticketing/README.md#createticketingfieldmap) - Create a field mapping for a ticketing project
* [GetTicketingFieldMap](docs/sdks/ticketing/README.md#getticketingfieldmap) - Get a field map for a ticketing project
* [DeleteTicketingFieldMap](docs/sdks/ticketing/README.md#deleteticketingfieldmap) - Archive a field map for a ticketing project
* [UpdateTicketingFieldMap](docs/sdks/ticketing/README.md#updateticketingfieldmap) - Update a field map for a ticketing project
* [ListAvailableInboundFieldMaps](docs/sdks/ticketing/README.md#listavailableinboundfieldmaps) - List available fields for ticket field mapping
* [ListInboundFieldMaps](docs/sdks/ticketing/README.md#listinboundfieldmaps) - List inbound field maps for a ticketing project
* [CreateInboundFieldMap](docs/sdks/ticketing/README.md#createinboundfieldmap) - Create inbound field map for a ticketing project
* [GetInboundFieldMap](docs/sdks/ticketing/README.md#getinboundfieldmap) - Get inbound field map for a ticketing project
* [UpdateInboundFieldMap](docs/sdks/ticketing/README.md#updateinboundfieldmap) - Update inbound field map for a ticketing project
* [DeleteInboundFieldMap](docs/sdks/ticketing/README.md#deleteinboundfieldmap) - Archive inbound field map for a ticketing project
* [CreateTicketingProjectConfig](docs/sdks/ticketing/README.md#createticketingprojectconfig) - Create a ticketing project configuration
* [GetTicketingProjectConfig](docs/sdks/ticketing/README.md#getticketingprojectconfig) - Get configuration for a ticketing project
* [DeleteTicketingProjectConfig](docs/sdks/ticketing/README.md#deleteticketingprojectconfig) - Archive a ticketing project configuration
* [UpdateTicketingProjectConfig](docs/sdks/ticketing/README.md#updateticketingprojectconfig) - Update configuration for a ticketing project
* [ListTicketingPriorities](docs/sdks/ticketing/README.md#listticketingpriorities) - List ticketing priorities
* [CreateTicketingPriority](docs/sdks/ticketing/README.md#createticketingpriority) - Create a ticketing priority
* [GetTicketingPriority](docs/sdks/ticketing/README.md#getticketingpriority) - Get a ticketing priority
* [DeleteTicketingPriority](docs/sdks/ticketing/README.md#deleteticketingpriority) - Delete a ticketing priority
* [UpdateTicketingPriority](docs/sdks/ticketing/README.md#updateticketingpriority) - Update a ticketing priority
* [ListTicketTags](docs/sdks/ticketing/README.md#listtickettags) - List ticket tags
* [GetTicketingFormConfiguration](docs/sdks/ticketing/README.md#getticketingformconfiguration) - Get the ticketing form configuration
* [ListTicketingCustomDefinitions](docs/sdks/ticketing/README.md#listticketingcustomdefinitions) - List ticketing custom fields
* [CreateTicketingCustomDefinition](docs/sdks/ticketing/README.md#createticketingcustomdefinition) - Create a ticketing custom field
* [DeleteTicketingCustomDefinition](docs/sdks/ticketing/README.md#deleteticketingcustomdefinition) - Delete a ticketing custom field
* [UpdateTicketingCustomDefinition](docs/sdks/ticketing/README.md#updateticketingcustomdefinition) - Update a ticketing custom field

### [Users](docs/sdks/users/README.md)

* [ListUsers](docs/sdks/users/README.md#listusers) - List users
* [GetUser](docs/sdks/users/README.md#getuser) - Get a user
* [GetCurrentUser](docs/sdks/users/README.md#getcurrentuser) - Get the currently authenticated user

### [Webhooks](docs/sdks/webhooks/README.md)

* [ListWebhooks](docs/sdks/webhooks/README.md#listwebhooks) - List webhooks
* [CreateWebhook](docs/sdks/webhooks/README.md#createwebhook) - Create a webhook
* [ListWebhookDeliveries](docs/sdks/webhooks/README.md#listwebhookdeliveries) - List webhook deliveries
* [GetWebhook](docs/sdks/webhooks/README.md#getwebhook) - Get a webhook
* [DeleteWebhook](docs/sdks/webhooks/README.md#deletewebhook) - Delete a webhook
* [UpdateWebhook](docs/sdks/webhooks/README.md#updatewebhook) - Update a webhook

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"github.com/firehydrant/firehydrant-go-sdk/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.AccountSettings.Ping(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"github.com/firehydrant/firehydrant-go-sdk/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.AccountSettings.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateService` function may return the following errors:

| Error Type            | Status Code | Content Type     |
| --------------------- | ----------- | ---------------- |
| sdkerrors.ErrorEntity | 400         | application/json |
| sdkerrors.SDKError    | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"github.com/firehydrant/firehydrant-go-sdk/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.CatalogEntries.CreateService(ctx, components.CreateService{
		Name: "<value>",
	})
	if err != nil {

		var e *sdkerrors.ErrorEntity
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	firehydrantgosdk "github.com/firehydrant/firehydrant-go-sdk"
	"github.com/firehydrant/firehydrant-go-sdk/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrantgosdk.New(
		firehydrantgosdk.WithServerURL("https://api.firehydrant.io/"),
		firehydrantgosdk.WithSecurity(components.Security{
			APIKey: "<YOUR_API_KEY_HERE>",
		}),
	)

	res, err := s.AccountSettings.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/firehydrant/firehydrant-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = firehydrantgosdk.New(firehydrantgosdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go)
