package sonarqube

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
	if !readonly {
		registerAlmIntegrationsImportGitlabProject(s)
	}
	registerAlmIntegrationsListAzureProjects(s)
	registerAlmIntegrationsListBitbucketserverProjects(s)
	registerAlmIntegrationsSearchAzureRepos(s)
	registerAlmIntegrationsSearchBitbucketcloudRepos(s)
	registerAlmIntegrationsSearchBitbucketserverRepos(s)
	registerAlmIntegrationsSearchGitlabRepos(s)
	if !readonly {
		registerAlmIntegrationsSetPat(s)
	}
	registerAlmSettingsCountBinding(s)
	if !readonly {
		registerAlmSettingsCreateAzure(s)
	}
	if !readonly {
		registerAlmSettingsCreateBitbucket(s)
	}
	if !readonly {
		registerAlmSettingsCreateBitbucketcloud(s)
	}
	if !readonly {
		registerAlmSettingsCreateGithub(s)
	}
	if !readonly {
		registerAlmSettingsCreateGitlab(s)
	}
	if !readonly {
		registerAlmSettingsDelete(s)
	}
	registerAlmSettingsGetBinding(s)
	registerAlmSettingsList(s)
	registerAlmSettingsListDefinitions(s)
	if !readonly {
		registerAlmSettingsUpdateAzure(s)
	}
	if !readonly {
		registerAlmSettingsUpdateBitbucket(s)
	}
	if !readonly {
		registerAlmSettingsUpdateBitbucketcloud(s)
	}
	if !readonly {
		registerAlmSettingsUpdateGithub(s)
	}
	if !readonly {
		registerAlmSettingsUpdateGitlab(s)
	}
	registerAlmSettingsValidate(s)
	registerAnalysisCacheGet(s)
	if !readonly {
		registerAuthenticationLogin(s)
	}
	if !readonly {
		registerAuthenticationLogout(s)
	}
	registerAuthenticationValidate(s)
	registerCeActivity(s)
	registerCeActivityStatus(s)
	registerCeComponent(s)
	registerCeTask(s)
	registerComponentsSearch(s)
	registerComponentsShow(s)
	registerComponentsTree(s)
	registerDuplicationsShow(s)
	if !readonly {
		registerFavoritesAdd(s)
	}
	if !readonly {
		registerFavoritesRemove(s)
	}
	registerFavoritesSearch(s)
	registerHotspotsSearch(s)
	registerHotspotsShow(s)
	if !readonly {
		registerIssuesAddComment(s)
	}
	if !readonly {
		registerIssuesAssign(s)
	}
	registerIssuesAuthors(s)
	if !readonly {
		registerIssuesBulkChange(s)
	}
	registerIssuesChangelog(s)
	if !readonly {
		registerIssuesDeleteComment(s)
	}
	if !readonly {
		registerIssuesDoTransition(s)
	}
	if !readonly {
		registerIssuesEditComment(s)
	}
	if !readonly {
		registerIssuesReindex(s)
	}
	registerIssuesSearch(s)
	if !readonly {
		registerIssuesSetSeverity(s)
	}
	if !readonly {
		registerIssuesSetTags(s)
	}
	if !readonly {
		registerIssuesSetType(s)
	}
	registerIssuesTags(s)
	registerLanguagesList(s)
	registerMeasuresComponent(s)
	registerMeasuresComponentTree(s)
	registerMeasuresSearchHistory(s)
	registerMetricsSearch(s)
	registerMetricsTypes(s)
	registerMonitoringMetrics(s)
	registerNewCodePeriodsList(s)
	if !readonly {
		registerNewCodePeriodsSet(s)
	}
	registerNewCodePeriodsShow(s)
	if !readonly {
		registerNewCodePeriodsUnset(s)
	}
	if !readonly {
		registerNotificationsAdd(s)
	}
	registerNotificationsList(s)
	if !readonly {
		registerNotificationsRemove(s)
	}
	if !readonly {
		registerPermissionsAddGroup(s)
	}
	if !readonly {
		registerPermissionsAddGroupToTemplate(s)
	}
	if !readonly {
		registerPermissionsAddProjectCreatorToTemplate(s)
	}
	if !readonly {
		registerPermissionsAddUser(s)
	}
	if !readonly {
		registerPermissionsAddUserToTemplate(s)
	}
	if !readonly {
		registerPermissionsApplyTemplate(s)
	}
	if !readonly {
		registerPermissionsBulkApplyTemplate(s)
	}
	if !readonly {
		registerPermissionsCreateTemplate(s)
	}
	if !readonly {
		registerPermissionsDeleteTemplate(s)
	}
	if !readonly {
		registerPermissionsRemoveGroup(s)
	}
	if !readonly {
		registerPermissionsRemoveGroupFromTemplate(s)
	}
	if !readonly {
		registerPermissionsRemoveProjectCreatorFromTemplate(s)
	}
	if !readonly {
		registerPermissionsRemoveUser(s)
	}
	if !readonly {
		registerPermissionsRemoveUserFromTemplate(s)
	}
	registerPermissionsSearchTemplates(s)
	if !readonly {
		registerPermissionsSetDefaultTemplate(s)
	}
	if !readonly {
		registerPermissionsUpdateTemplate(s)
	}
	registerPluginsAvailable(s)
	if !readonly {
		registerPluginsCancelAll(s)
	}
	if !readonly {
		registerPluginsInstall(s)
	}
	registerPluginsInstalled(s)
	registerPluginsPending(s)
	if !readonly {
		registerPluginsUninstall(s)
	}
	if !readonly {
		registerPluginsUpdate(s)
	}
	registerPluginsUpdates(s)
	if !readonly {
		registerProjectAnalysesCreateEvent(s)
	}
	if !readonly {
		registerProjectAnalysesDelete(s)
	}
	if !readonly {
		registerProjectAnalysesDeleteEvent(s)
	}
	registerProjectAnalysesSearch(s)
	if !readonly {
		registerProjectAnalysesSetBaseline(s)
	}
	if !readonly {
		registerProjectAnalysesUnsetBaseline(s)
	}
	if !readonly {
		registerProjectAnalysesUpdateEvent(s)
	}
	registerProjectBadgesMeasure(s)
	registerProjectBadgesQualityGate(s)
	if !readonly {
		registerProjectBadgesRenewToken(s)
	}
	registerProjectBadgesToken(s)
	if !readonly {
		registerProjectBranchesDelete(s)
	}
	registerProjectBranchesList(s)
	if !readonly {
		registerProjectBranchesRename(s)
	}
	if !readonly {
		registerProjectBranchesSetAutomaticDeletionProtection(s)
	}
	if !readonly {
		registerProjectDumpExport(s)
	}
	if !readonly {
		registerProjectLinksCreate(s)
	}
	if !readonly {
		registerProjectLinksDelete(s)
	}
	registerProjectLinksSearch(s)
	registerProjectTagsSearch(s)
	if !readonly {
		registerProjectTagsSet(s)
	}
	if !readonly {
		registerProjectsBulkDelete(s)
	}
	if !readonly {
		registerProjectsCreate(s)
	}
	if !readonly {
		registerProjectsDelete(s)
	}
	registerProjectsSearch(s)
	if !readonly {
		registerProjectsUpdateKey(s)
	}
	if !readonly {
		registerProjectsUpdateVisibility(s)
	}
	if !readonly {
		registerQualitygatesCopy(s)
	}
	if !readonly {
		registerQualitygatesCreate(s)
	}
	if !readonly {
		registerQualitygatesCreateCondition(s)
	}
	if !readonly {
		registerQualitygatesDeleteCondition(s)
	}
	if !readonly {
		registerQualitygatesDeselect(s)
	}
	if !readonly {
		registerQualitygatesDestroy(s)
	}
	registerQualitygatesGetByProject(s)
	registerQualitygatesList(s)
	registerQualitygatesProjectStatus(s)
	if !readonly {
		registerQualitygatesRename(s)
	}
	registerQualitygatesSearch(s)
	if !readonly {
		registerQualitygatesSelect(s)
	}
	if !readonly {
		registerQualitygatesSetAsDefault(s)
	}
	registerQualitygatesShow(s)
	if !readonly {
		registerQualitygatesUpdateCondition(s)
	}
	if !readonly {
		registerQualityprofilesActivateRule(s)
	}
	if !readonly {
		registerQualityprofilesActivateRules(s)
	}
	if !readonly {
		registerQualityprofilesAddProject(s)
	}
	registerQualityprofilesBackup(s)
	if !readonly {
		registerQualityprofilesChangeParent(s)
	}
	registerQualityprofilesChangelog(s)
	if !readonly {
		registerQualityprofilesCopy(s)
	}
	if !readonly {
		registerQualityprofilesCreate(s)
	}
	if !readonly {
		registerQualityprofilesDeactivateRule(s)
	}
	if !readonly {
		registerQualityprofilesDeactivateRules(s)
	}
	if !readonly {
		registerQualityprofilesDelete(s)
	}
	registerQualityprofilesExport(s)
	registerQualityprofilesExporters(s)
	registerQualityprofilesImporters(s)
	registerQualityprofilesInheritance(s)
	registerQualityprofilesProjects(s)
	if !readonly {
		registerQualityprofilesRemoveProject(s)
	}
	if !readonly {
		registerQualityprofilesRename(s)
	}
	if !readonly {
		registerQualityprofilesRestore(s)
	}
	registerQualityprofilesSearch(s)
	if !readonly {
		registerQualityprofilesSetDefault(s)
	}
	if !readonly {
		registerRulesCreate(s)
	}
	if !readonly {
		registerRulesDelete(s)
	}
	registerRulesRepositories(s)
	registerRulesSearch(s)
	registerRulesShow(s)
	registerRulesTags(s)
	if !readonly {
		registerRulesUpdate(s)
	}
	registerServerVersion(s)
	registerSettingsListDefinitions(s)
	if !readonly {
		registerSettingsReset(s)
	}
	if !readonly {
		registerSettingsSet(s)
	}
	registerSettingsValues(s)
	registerSourcesRaw(s)
	registerSourcesScm(s)
	registerSourcesShow(s)
	if !readonly {
		registerSystemChangeLogLevel(s)
	}
	registerSystemDbMigrationStatus(s)
	registerSystemHealth(s)
	registerSystemInfo(s)
	registerSystemLogs(s)
	if !readonly {
		registerSystemMigrateDb(s)
	}
	registerSystemPing(s)
	if !readonly {
		registerSystemRestart(s)
	}
	registerSystemStatus(s)
	registerSystemUpgrades(s)
	if !readonly {
		registerUserGroupsAddUser(s)
	}
	if !readonly {
		registerUserGroupsCreate(s)
	}
	if !readonly {
		registerUserGroupsDelete(s)
	}
	if !readonly {
		registerUserGroupsRemoveUser(s)
	}
	registerUserGroupsSearch(s)
	if !readonly {
		registerUserGroupsUpdate(s)
	}
	registerUserGroupsUsers(s)
	if !readonly {
		registerUserTokensGenerate(s)
	}
	if !readonly {
		registerUserTokensRevoke(s)
	}
	registerUserTokensSearch(s)
	if !readonly {
		registerUsersAnonymize(s)
	}
	if !readonly {
		registerUsersChangePassword(s)
	}
	if !readonly {
		registerUsersCreate(s)
	}
	if !readonly {
		registerUsersDeactivate(s)
	}
	if !readonly {
		registerUsersDismissSonarlintAd(s)
	}
	registerUsersGroups(s)
	registerUsersSearch(s)
	if !readonly {
		registerUsersUpdate(s)
	}
	if !readonly {
		registerUsersUpdateIdentityProvider(s)
	}
	if !readonly {
		registerUsersUpdateLogin(s)
	}
	if !readonly {
		registerWebhooksCreate(s)
	}
	if !readonly {
		registerWebhooksDelete(s)
	}
	registerWebhooksDeliveries(s)
	registerWebhooksDelivery(s)
	registerWebhooksList(s)
	if !readonly {
		registerWebhooksUpdate(s)
	}
	registerWebservicesList(s)
	registerWebservicesResponseExample(s)
}
