# SonarQube MCP Server

This repository provides model context protocol server for SonarQube server.

## API Version

* v9.9

## Build

Build binary.

```sh
go build -o bin/sonarqube-mcp-server ./cmd/sonarqube-mcp-server/main.go
```

Or build docker image.

```sh
docker build -t sonarqube-mcp-server .
```

## Usage

Run application.

Specify token or username and password.
If use user type token and SonarQube server 9.x or earlier, Specify token to `--user`.

```sh
$ ./bin/sonarqube-mcp-server -h
SonarQube MCP Server

Usage:
  sonarqube-mcp-server [flags]

Flags:
  -h, --help              help for sonarqube-mcp-server
      --password string   SonarQube server password.
      --readonly          HTTP GET method only. (default true)
      --token string      SonarQube server user type token.
      --url string        SonarQube server URL. (default "http://127.0.0.1:9000")
      --user string       SonarQube server username.
```

Set environment variable instead of arguments.

| Argument   | Environment Variable |
| :--------- | :------------------- |
| --url      | SONARQUBE_URL        |
| --user     | SONARQUBE_USER       |
| --password | SONARQUBE_PASSWORD   |
| --token    | SONARQUBE_TOKEN      |
| --readonly | SONARQUBE_READONLY   |

Or run container.

```sh
docker run --rm -i -e SONARQUBE_URL=<URL> -e SONARQUBE_TOKEN=<TOKEN> sonarqube-mcp-server
```

## Tools

| Tool       | Description          |
| :--------- | :------------------- |
| alm_integrations_import_gitlab_project | Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br/>Requires the 'Create Projects' permission  |
| alm_integrations_list_azure_projects | List Azure projects<br/>Requires the 'Create Projects' permission  |
| alm_integrations_list_bitbucketserver_projects | List the Bitbucket Server projects<br/>Requires the 'Create Projects' permission  |
| alm_integrations_search_azure_repos | Search the Azure repositories<br/>Requires the 'Create Projects' permission  |
| alm_integrations_search_bitbucketcloud_repos | Search the Bitbucket Cloud repositories<br/>Requires the 'Create Projects' permission  |
| alm_integrations_search_bitbucketserver_repos | Search the Bitbucket Server repositories with REPO_ADMIN access<br/>Requires the 'Create Projects' permission  |
| alm_integrations_search_gitlab_repos | Search the GitLab projects.<br/>Requires the 'Create Projects' permission  |
| alm_integrations_set_pat | Set a Personal Access Token for the given DevOps Platform setting<br/>Only valid for Azure DevOps, Bitbucket Server, GitLab and Bitbucket Cloud Setting<br/>Requires the 'Create Projects' permission  |
| alm_settings_count_binding | Count number of project bound to an DevOps Platform setting.<br/>Requires the 'Administer System' permission  |
| alm_settings_create_azure | Create Azure instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_create_bitbucket | Create Bitbucket instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_create_bitbucketcloud | Configure a new instance of Bitbucket Cloud. <br/>Requires the 'Administer System' permission  |
| alm_settings_create_github | Create GitHub instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_create_gitlab | Create GitLab instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_delete | Delete an DevOps Platform Setting.<br/>Requires the 'Administer System' permission  |
| alm_settings_get_binding | Get DevOps Platform binding of a given project.<br/>Requires the 'Administer' permission on the project  |
| alm_settings_list | List DevOps Platform setting available for a given project, sorted by DevOps Platform key<br/>Requires the 'Administer project' permission if the 'project' parameter is provided, requires the 'Create Projects' permission otherwise.  |
| alm_settings_list_definitions | List DevOps Platform Settings, sorted by created date.<br/>Requires the 'Administer System' permission  |
| alm_settings_update_azure | Update Azure instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_update_bitbucket | Update Bitbucket instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_update_bitbucketcloud | Update Bitbucket Cloud Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_update_github | Update GitHub instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_update_gitlab | Update GitLab instance Setting. <br/>Requires the 'Administer System' permission  |
| alm_settings_validate | Validate an DevOps Platform Setting by checking connectivity and permissions<br/>Requires the 'Administer System' permission  |
| analysis_cache_get | Get the scanner's cached data for a branch. Requires scan permission on the project. Data is returned gzipped if the corresponding 'Accept-Encoding' header is set in the request.  |
| authentication_login | Authenticate a user.  |
| authentication_logout | Logout a user.  |
| authentication_validate | Check credentials.  |
| ce_activity | Search for tasks.<br> Either componentId or component can be provided, but not both.<br> Requires the system administration permission, or project administration permission if componentId or component is set.  |
| ce_activity_status | Returns CE activity related metrics.<br>Requires 'Administer System' permission or 'Administer' rights on the specified project.  |
| ce_component | Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component.  |
| ce_task | Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Administer System' or 'Execute Analysis' permission.<br/>Since 6.1, field "logs" is deprecated and its value is always false.  |
| components_search | Search for components  |
| components_show | Returns a component (file, directory, project, portfolio…) and its ancestors. The ancestors are ordered from the parent to the root project. Requires the following permission: 'Browse' on the project of the specified component.  |
| components_tree | Navigate through components based on the chosen strategy.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.  |
| duplications_show | Get duplications. Require Browse permission on file's project  |
| favorites_add | Add a component (project, file etc.) as favorite for the authenticated user.<br>Only 100 components by qualifier can be added as favorite.<br>Requires authentication and the following permission: 'Browse' on the project of the specified component.  |
| favorites_remove | Remove a component (project, portfolio, application etc.) as favorite for the authenticated user.<br>Requires authentication.  |
| favorites_search | Search for the authenticated user favorites.<br>Requires authentication.  |
| hotspots_search | Search for Security Hotpots. <br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects. <br>When issue indexation is in progress returns 503 service unavailable HTTP code.  |
| hotspots_show | Provides the details of a Security Hotspot.  |
| issues_add_comment | Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.  |
| issues_assign | Assign/Unassign an issue. Requires authentication and Browse permission on project  |
| issues_authors | Search SCM accounts which match a given query.<br/>Requires authentication.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code.  |
| issues_bulk_change | Bulk change on issues. Up to 500 issues can be updated. <br/>Requires authentication.  |
| issues_changelog | Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue.  |
| issues_delete_comment | Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.  |
| issues_do_transition | Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>The transitions 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.<br/>The transitions involving security hotspots require the permission 'Administer Security Hotspot'.  |
| issues_edit_comment | Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.  |
| issues_reindex | Reindex issues for a project.<br> Require 'Administer System' permission.  |
| issues_search | Search for issues.<br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code.  |
| issues_set_severity | Change severity.<br/>Requires the following permissions:<ul> <li>'Authentication'</li> <li>'Browse' rights on project of the specified issue</li> <li>'Administer Issues' rights on project of the specified issue</li></ul>  |
| issues_set_tags | Set tags on an issue. <br/>Requires authentication and Browse permission on project  |
| issues_set_type | Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul> <li>'Authentication'</li> <li>'Browse' rights on project of the specified issue</li> <li>'Administer Issues' rights on project of the specified issue</li></ul>  |
| issues_tags | List tags matching a given query  |
| languages_list | List supported programming languages  |
| measures_component | Return component with specified measures.<br>Requires the following permission: 'Browse' on the project of specified component.  |
| measures_component_tree | Navigate through components based on the chosen strategy with specified measures.<br>Requires the following permission: 'Browse' on the specified project.<br>For applications, it also requires 'Browse' permission on its child projects. <br>When limiting search with the q parameter, directories are not returned.  |
| measures_search_history | Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component. <br>For applications, it also requires 'Browse' permission on its child projects.  |
| metrics_search | Search for metrics  |
| metrics_types | List all available metric types.  |
| monitoring_metrics | Return monitoring metrics in Prometheus format. Support content type 'text/plain' (default) and 'application/openmetrics-text'. this endpoint can be access using a Bearer token, that needs to be defined in sonar.properties with the 'sonar.web.systemPasscode' key.  |
| new_code_periods_list | List the New Code Periods for all branches in a project.<br>Requires the permission to browse the project  |
| new_code_periods_set | Updates the setting for the New Code Period on different levels:<br><ul><li>Project key must be provided to update the value for a project</li><li>Both project and branch keys must be provided to update the value for a branch</li></ul>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights on the specified project to change the project setting</li></ul>  |
| new_code_periods_show | Shows a setting for the New Code Period.<br> If the component requested doesn't exist or if no new code period is set for it, a value is inherited from the project or from the global setting.Requires one of the following permissions if a component is specified: <ul><li>'Administer' rights on the specified component</li><li>'Execute analysis' rights on the specified component</li></ul>  |
| new_code_periods_unset | Unset the New Code Period setting for a branch, project or global.<br>Requires one of the following permissions: <ul><li>'Administer System' to change the global setting</li><li>'Administer' rights for a specified component</li></ul>  |
| notifications_add | Add a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li> <li>System administration if a login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li></ul>  |
| notifications_list | List notifications of the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided</li> <li>System administration if a login is provided</li></ul>  |
| notifications_remove | Remove a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided</li> <li>System administration if a login is provided</li></ul>  |
| permissions_add_group | Add a permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name or group id must be provided. <br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>  |
| permissions_add_group_to_template | Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.  |
| permissions_add_project_creator_to_template | Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'.  |
| permissions_add_user | Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>  |
| permissions_add_user_to_template | Add a user to a permission template.<br /> Requires the following permission: 'Administer System'.  |
| permissions_apply_template | Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'.  |
| permissions_bulk_apply_template | Apply a permission template to several projects.<br />The template id or name must be provided.<br />Requires the following permission: 'Administer System'.  |
| permissions_create_template | Create a permission template.<br />Requires the following permission: 'Administer System'.  |
| permissions_delete_template | Delete a permission template.<br />Requires the following permission: 'Administer System'.  |
| permissions_remove_group | Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>  |
| permissions_remove_group_from_template | Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.  |
| permissions_remove_project_creator_from_template | Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'.  |
| permissions_remove_user | Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>  |
| permissions_remove_user_from_template | Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'.  |
| permissions_search_templates | List permission templates.<br />Requires the following permission: 'Administer System'.  |
| permissions_set_default_template | Set a permission template as default.<br />Requires the following permission: 'Administer System'.  |
| permissions_update_template | Update a permission template.<br />Requires the following permission: 'Administer System'.  |
| plugins_available | Get the list of all the plugins available for installation on the SonarQube instance, sorted by plugin name.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: <ul><li>COMPATIBLE: plugin is compatible with current SonarQube instance.</li><li>INCOMPATIBLE: plugin is not compatible with current SonarQube instance.</li><li>REQUIRES_SYSTEM_UPGRADE: plugin requires SonarQube to be upgraded before being installed.</li><li>DEPS_REQUIRE_SYSTEM_UPGRADE: at least one plugin on which the plugin is dependent requires SonarQube to be upgraded.</li></ul>Require 'Administer System' permission.  |
| plugins_cancel_all | Cancels any operation pending on any plugin (install, update or uninstall)<br/>Requires user to be authenticated with Administer System permissions  |
| plugins_install | Installs the latest version of a plugin specified by its key.<br/>Plugin information is retrieved from Update Center.<br/>Fails if used on commercial editions or plugin risk consent has not been accepted.<br/>Requires user to be authenticated with Administer System permissions  |
| plugins_installed | Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.<br/>Requires authentication.  |
| plugins_pending | Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.<br/>Require 'Administer System' permission.  |
| plugins_uninstall | Uninstalls the plugin specified by its key.<br/>Requires user to be authenticated with Administer System permissions.  |
| plugins_update | Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions  |
| plugins_updates | Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.<br/>Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].<br/>Require 'Administer System' permission.  |
| project_analyses_create_event | Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>  |
| project_analyses_delete | Delete a project analysis.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the project of the specified analysis</li></ul>  |
| project_analyses_delete_event | Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>  |
| project_analyses_search | Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project. <br>For applications, it also requires 'Browse' permission on its child projects.  |
| project_analyses_set_baseline | Set an analysis as the baseline of the New Code Period on a project or a branch.<br/>This manually set baseline.<br/>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>  |
| project_analyses_unset_baseline | Unset any manually-set New Code Period baseline on a project or a branch.<br/>Unsetting a manual baseline restores the use of the default new code period setting.<br/>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>  |
| project_analyses_update_event | Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul> <li>'Administer System'</li> <li>'Administer' rights on the specified project</li></ul>  |
| project_badges_measure | Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project.  |
| project_badges_quality_gate | Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project.  |
| project_badges_renew_token | Creates new token replacing any existing token for project or application badge access for private projects and applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Administer' permission on the specified project or application.  |
| project_badges_token | Retrieve a token to use for project or application badge access for private projects or applications.<br/>This token can be used to authenticate with api/project_badges/quality_gate and api/project_badges/measure endpoints.<br/>Requires 'Browse' permission on the specified project or application.  |
| project_branches_delete | Delete a non-main branch of a project or application.<br/>Requires 'Administer' rights on the specified project or application.  |
| project_branches_list | List the branches of a project or application.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project or application.  |
| project_branches_rename | Rename the main branch of a project or application.<br/>Requires 'Administer' permission on the specified project or application.  |
| project_branches_set_automatic_deletion_protection | Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.<br/>Requires 'Administer' permission on the specified project or application.  |
| project_dump_export | Triggers project dump so that the project can be imported to another SonarQube server (see api/project_dump/import, available in Enterprise Edition). Requires the 'Administer' permission.  |
| project_links_create | Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.  |
| project_links_delete | Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.  |
| project_links_search | List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>  |
| project_tags_search | Search tags  |
| project_tags_set | Set tags on a project.<br>Requires the following permission: 'Administer' rights on the specified project  |
| projects_bulk_delete | Delete one or several projects.<br />Only the 1'000 first items in project filters are taken into account.<br />Requires 'Administer System' permission.<br />At least one parameter is required among analyzedBefore, projects and q  |
| projects_create | Create a project.<br/>Requires 'Create Projects' permission  |
| projects_delete | Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.  |
| projects_search | Search for projects or views to administrate them.<br>Requires 'Administer System' permission  |
| projects_update_key | Update a project all its sub-components keys.<br>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>  |
| projects_update_visibility | Updates visibility of a project or view.<br>Requires 'Project administer' permission on the specified project or view  |
| qualitygates_copy | Copy a Quality Gate.<br>Either 'sourceName' or 'id' must be provided. Requires the 'Administer Quality Gates' permission.  |
| qualitygates_create | Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.  |
| qualitygates_create_condition | Add a new condition to a quality gate.<br>Either 'gateId' or 'gateName' must be provided. Requires the 'Administer Quality Gates' permission.  |
| qualitygates_delete_condition | Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission.  |
| qualitygates_deselect | Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>  |
| qualitygates_destroy | Delete a Quality Gate.<br>Either 'id' or 'name' must be specified. Requires the 'Administer Quality Gates' permission.  |
| qualitygates_get_by_project | Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>  |
| qualitygates_list | Get a list of quality gates  |
| qualitygates_project_status | Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided <br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li><li>'Execute Analysis' on the specified project</li></ul>  |
| qualitygates_rename | Rename a Quality Gate.<br>Either 'id' or 'currentName' must be specified. Requires the 'Administer Quality Gates' permission.  |
| qualitygates_search | Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for the current user will be returned.  |
| qualitygates_select | Associate a project to a quality gate.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Gates'</li> <li>'Administer' right on the specified project</li></ul>  |
| qualitygates_set_as_default | Set a quality gate as the default quality gate.<br>Either 'id' or 'name' must be specified. Requires the 'Administer Quality Gates' permission.  |
| qualitygates_show | Display the details of a quality gate  |
| qualitygates_update_condition | Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission.  |
| qualityprofiles_activate_rule | Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_activate_rules | Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_add_project | Associate a project with a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li> <li>Administer right on the specified project</li></ul>  |
| qualityprofiles_backup | Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.  |
| qualityprofiles_change_parent | Change a quality profile's parent.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_changelog | Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first).  |
| qualityprofiles_copy | Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.  |
| qualityprofiles_create | Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission.  |
| qualityprofiles_deactivate_rule | Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_deactivate_rules | Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_delete | Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_export | Export a quality profile.  |
| qualityprofiles_exporters | Lists available profile export formats.  |
| qualityprofiles_importers | List supported importers.  |
| qualityprofiles_inheritance | Show a quality profile's ancestors and children.  |
| qualityprofiles_projects | List projects with their association status regarding a quality profile <br/>See api/qualityprofiles/search in order to get the Quality Profile Key  |
| qualityprofiles_remove_project | Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li> <li>Administer right on the specified project</li></ul>  |
| qualityprofiles_rename | Rename a quality profile.<br> Requires one of the following permissions:<ul> <li>'Administer Quality Profiles'</li> <li>Edit right on the specified quality profile</li></ul>  |
| qualityprofiles_restore | Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.  |
| qualityprofiles_search | Search quality profiles  |
| qualityprofiles_set_default | Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.  |
| rules_create | Create a custom rule.<br>Requires the 'Administer Quality Profiles' permission  |
| rules_delete | Delete custom rule.<br/>Requires the 'Administer Quality Profiles' permission  |
| rules_repositories | List available rule repositories  |
| rules_search | Search for a collection of relevant rules matching a specified query.<br/>  |
| rules_show | Get detailed information about a rule<br>  |
| rules_tags | List rule tags  |
| rules_update | Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission  |
| server_version | Version of SonarQube in plain text  |
| settings_list_definitions | List settings definitions.<br>Requires 'Browse' permission when a component is specified<br/>To access licensed settings, authentication is required<br/>To access secured settings, one of the following permissions is required: <ul><li>'Execute Analysis'</li><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>  |
| settings_reset | Remove a setting value.<br>The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>  |
| settings_set | Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>  |
| settings_values | List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>The settings from conf/sonar.properties are excluded from results.<br>Requires 'Browse' or 'Execute Analysis' permission when a component is specified.<br/>Secured settings values are not returned by the endpoint.<br/>  |
| sources_raw | Get source code as raw text. Require 'See Source Code' permission on file  |
| sources_scm | Get SCM information of source files. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Author of the commit</li><li>Datetime of the commit (before 5.2 it was only the Date)</li><li>Revision of the commit (added in 5.2)</li></ol>  |
| sources_show | Get source code. Requires See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Content of the line</li></ol>  |
| system_change_log_level | Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.  |
| system_db_migration_status | Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>  |
| system_health | Provide health status of SonarQube.<p>Although global health is calculated based on both application and search nodes, detailed information is returned only for application nodes.</p><p> <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p>  |
| system_info | Get detailed information about system configuration.<br/>Requires 'Administer' permissions.  |
| system_logs | Get system logs in plain-text format. Requires system administration permission.  |
| system_migrate_db | Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>  |
| system_ping | Answers "pong" as plain-text  |
| system_restart | Restarts server. Requires 'Administer System' permission. Performs a full restart of the Web, Search and Compute Engine Servers processes. Does not reload sonar.properties.  |
| system_status | Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>  |
| system_upgrades | Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.  |
| user_groups_add_user | Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.  |
| user_groups_create | Create a group.<br>Requires the following permission: 'Administer System'.  |
| user_groups_delete | Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.  |
| user_groups_remove_user | Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'.  |
| user_groups_search | Search for user groups.<br>Requires the following permission: 'Administer System'.  |
| user_groups_update | Update a group.<br>Requires the following permission: 'Administer System'.  |
| user_groups_users | Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'.  |
| user_tokens_generate | Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />It requires administration permissions to specify a 'login' and generate a token for another user. Otherwise, a token is generated for the current user.  |
| user_tokens_revoke | Revoke a user access token. <br/>It requires administration permissions to specify a 'login' and revoke a token for another user. Otherwise, the token for the current user is revoked.  |
| user_tokens_search | List the access tokens of a user.<br>The login must exist and active.<br>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user is using a token many times in less than one hour.<br> It requires administration permissions to specify a 'login' and list the tokens of another user. Otherwise, tokens for the current user are listed. <br> Authentication is required for this API endpoint  |
| users_anonymize | Anonymize a deactivated user. Requires Administer System permission  |
| users_change_password | Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password.  |
| users_create | Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission  |
| users_deactivate | Deactivate a user. Requires Administer System permission  |
| users_dismiss_sonarlint_ad | Dismiss SonarLint advertisement. Deprecated since 9.6, replaced api/users/dismiss_notice  |
| users_groups | Lists the groups a user belongs to. <br/>Requires Administer System permission.  |
| users_search | Get a list of users. By default, only active users are returned.<br/>The following fields are only returned when user has Administer System permission or for logged-in in user :<ul> <li>'email'</li> <li>'externalIdentity'</li> <li>'externalProvider'</li> <li>'groups'</li> <li>'lastConnectionDate'</li> <li>'tokensCount'</li></ul>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour.  |
| users_update | Update a user.<br/>Requires Administer System permission  |
| users_update_identity_provider | Update identity provider information. <br/>It's only possible to migrate to an installed identity provider. Be careful that as soon as this information has been updated for a user, the user will only be able to authenticate on the new identity provider. It is not possible to migrate external user to local one.<br/>Requires Administer System permission.  |
| users_update_login | Update a user login. A login can be updated many times.<br/>Requires Administer System permission  |
| webhooks_create | Create a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.  |
| webhooks_delete | Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.  |
| webhooks_deliveries | Get the recent deliveries for a specified project or Compute Engine task.<br/>Require 'Administer' permission on the related project.<br/>Note that additional information are returned by api/webhooks/delivery.  |
| webhooks_delivery | Get a webhook delivery by its id.<br/>Require 'Administer System' permission.<br/>Note that additional information are returned by api/webhooks/delivery.  |
| webhooks_list | Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.  |
| webhooks_update | Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.  |
| webservices_list | List web services  |
| webservices_response_example | Display web service response example  |

## Testing

TODO
