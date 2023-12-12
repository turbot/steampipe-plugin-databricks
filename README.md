![image](https://hub.steampipe.io/images/plugins/turbot/databricks-social-graphic.png)

# Databricks Plugin for Steampipe

Use SQL to query clusters, jobs, users, and more from Databricks.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/databricks)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/databricks/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-databricks/issues)

## Quick start

### Install

Download and install the latest Databricks plugin:

```bash
steampipe plugin install databricks
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/databricks#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/databricks#configuration).

Configure your account details in `~/.steampipe/config/databricks.spc`:

```hcl
connection "databricks" {
  plugin = "databricks"

  # A connection profile specified within .databrickscfg to use instead of DEFAULT.
  # This can also be set via the `DATABRICKS_CONFIG_PROFILE` environment variable.
  # profile = "databricks-dev"

  # The target Databricks account ID.
  # This can also be set via the `DATABRICKS_ACCOUNT_ID` environment variable.
  # See Locate your account ID: https://docs.databricks.com/administration-guide/account-settings/index.html#account-id.
  # account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"

  # The target Databricks account SCIM token.
  # See: https://docs.databricks.com/administration-guide/account-settings/index.html#generate-a-scim-token
  # This can also be set via the `DATABRICKS_TOKEN` environment variable.
  # account_token = "dsapi5c72c067b40df73ccb6be3b085d3ba"

  # The target Databricks account console URL, which is typically https://accounts.cloud.databricks.com.
  # This can also be set via the `DATABRICKS_HOST` environment variable.
  # account_host = "https://accounts.cloud.databricks.com/"

  # The target Databricks workspace Personal Access Token.
  # This can also be set via the `DATABRICKS_TOKEN` environment variable.
  # See: https://docs.databricks.com/dev-tools/auth.html#databricks-personal-access-tokens-for-users
  # workspace_token = "dapia865b9d1d41389ed883455032d090ee"

  # The target Databricks workspace URL.
  # See https://docs.databricks.com/workspace/workspace-details.html#workspace-url
  # This can also be set via the `DATABRICKS_HOST` environment variable.
  # workspace_host = "https://dbc-a1b2c3d4-e6f7.cloud.databricks.com"

  # The Databricks username part of basic authentication. Only possible when Host is *.cloud.databricks.com (AWS).
  # This can also be set via the `DATABRICKS_USERNAME` environment variable.
  # username = "user@turbot.com"

  # The Databricks password part of basic authentication. Only possible when Host is *.cloud.databricks.com (AWS).
  # This can also be set via the `DATABRICKS_PASSWORD` environment variable.
  # password = "password"

  # A non-default location of the Databricks CLI credentials file.
  # This can also be set via the `DATABRICKS_CONFIG_FILE` environment variable.
  # config_file_path = "/Users/username/.databrickscfg"
}
```

- **[Detailed configuration guide →](https://hub.steampipe.io/plugins/turbot/databricks#quick-start)**

Or through environment variables:

```sh
export DATABRICKS_CONFIG_PROFILE=user1-test
export DATABRICKS_TOKEN=dsapi5c72c067b40df73ccb6be3b085d3ba
export DATABRICKS_HOST=https://accounts.cloud.databricks.com
export DATABRICKS_ACCOUNT_ID=abcdd0f81-9be0-4425-9e29-3a7d96782373
export DATABRICKS_USERNAME=user@turbot.com
export DATABRICKS_PASSWORD=password
```

Run steampipe:

```shell
steampipe query
```

List details of your Databricks clusters:

```sql
select
  cluster_id,
  title,
  cluster_source,
  creator_user_name,
  driver_node_type_id,
  node_type_id,
  state,
  start_time
from
  databricks_compute_cluster;
```

```
+----------------------+--------------------------------+----------------+-------------------+---------------------+--------------+------------+---------------------------+
| cluster_id           | title                          | cluster_source | creator_user_name | driver_node_type_id | node_type_id | state      | start_time                |
+----------------------+--------------------------------+----------------+-------------------+---------------------+--------------+------------+---------------------------+
| 1234-141524-10b6dv2h | [default]basic-starter-cluster | "API"          | user@turbot.com   | i3.xlarge           | i3.xlarge    | TERMINATED | 2023-07-21T19:45:24+05:30 |
| 1234-061816-mvns8mxz | test-cluster-for-ml            | "UI"           | user@turbot.com   | i3.xlarge           | i3.xlarge    | TERMINATED | 2023-07-28T11:48:16+05:30 |
+----------------------+--------------------------------+----------------+-------------------+---------------------+--------------+------------+---------------------------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-databricks.git
cd steampipe-plugin-databricks
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/databricks.spc
```

Try it!

```
steampipe query
> .inspect databricks
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Databricks Plugin](https://github.com/turbot/steampipe-plugin-databricks/labels/help%20wanted)
