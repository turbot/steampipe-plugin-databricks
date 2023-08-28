---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/databricks.svg"
brand_color: "#FF3621"
display_name: "Databricks"
short_name: "databricks"
description: "Steampipe plugin to query clusters, jobs, users, and more from Databricks."
og_description: "Query Databricks with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/databricks-social-graphic.png"
---

# Databricks + Steampipe

[Databricks](https://databricks.com) is a unified set of tools for building, deploying, sharing, and maintaining enterprise-grade data solutions at scale.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/databricks/tables)**

## Quick start

### Install

Download and install the latest Databricks plugin:

```bash
steampipe plugin install databricks
```

### Credentials

| Item| Description|
| - | - |
| Credentials | For Databricks native authentication, Specify a named profile from .databrickscfg file with the `profile` argument.|
| Permissions | Grant the `READ` permissions to your user.|
| Radius      | Each connection represents a single Databricks Installation.|
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/databricks.spc`)<br />2. Credentials specified in environment variables, e.g., `DATABRICKS_TOKEN`.<br />3. Credentials in the credential file (`~/.databrickscfg`) for the profile specified in the `DATABRICKS_CONFIG_PROFILE` environment variable.|

### Configuration

Installing the latest databricks plugin will create a config file (`~/.steampipe/config/databricks.spc`) with a single connection named `databricks`:

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

By default, all options are commented out in the default connection, thus Steampipe will resolve your credentials using the same mechanism as the Databricks CLI (Databricks environment variables, DEFAULT profile, etc). This provides a quick way to get started with Steampipe, but you will probably want to customize your experience using configuration options for [querying multiple accounts](#multi-account-connections), [configuring credentials](#configuring-databricks-credentials) from your [Databricks Profiles](#databricks-profile-credentials).

## Multi-Account Connections

You may create multiple databricks connections:
```hcl
connection "databricks_dev" {
  plugin  = "databricks"
  profile = "databricks_dev"
  account_id = abcdd0f81-9be0-4425-9e29-3a7d96782373
}

connection "databricks_qa" {
  plugin  = "databricks"
  profile = "databricks_qa"
  account_id = wxyzd0f81-9be0-4425-9e29-3a7d96782373
}

connection "databricks_prod" {
  plugin  = "databricks"
  profile = "databricks_prod"
  account_id = pqrsd0f81-9be0-4425-9e29-3a7d96782373
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from databricks_dev.databricks_iam_account_user;
```

You can create a multi-account connection by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "databricks_all" {
  plugin      = "databricks"
  type        = "aggregator"
  connections = ["databricks_dev", "databricks_qa", "databricks_prod"]
}
```

Querying tables from this connection will return results from the `databricks_dev`, `databricks_qa`, and `databricks_prod` connections:
```sql
select * from databricks_all.databricks_iam_account_user;
```

Alternatively, can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphabetically, so that it is the first connection in the search path (i.e. `databricks_all` comes before `databricks_dev`):
```sql
select * from databricks_iam_account_user;
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the Databricks plugin connections whose names begin with `databricks_`:

```hcl
connection "databricks_all" {
  type        = "aggregator"
  plugin      = "databricks"
  connections = ["databricks_*"]
}
```

## Configuring Databricks Credentials

### Databricks Profile Credentials

You may specify a named profile from a Databricks credential file with the `profile` argument. A connection per profile, using named profiles is probably the most common configuration:

#### databricks credential file:

```ini
[user1-account]
host       = https://accounts.cloud.databricks.com
token      = dsapi5c72c067b40df73ccb6be3b085d3ba
account_id = abcdd0f81-9be0-4425-9e29-3a7d96782373

[user1-workspace]
host       = https://dbc-a1b2c3d4-e6f7.cloud.databricks.com/
token      = dapia865b9d1d41389ed883455032d090ee
account_id = abcdd0f81-9be0-4425-9e29-3a7d96782373

[user1-basic]
host       = https://accounts.cloud.databricks.com
username   = user1@turbot.com
password   = Pass****word
account_id = abcdd0f81-9be0-4425-9e29-3a7d96782373
```

#### databricks.spc:

```hcl
connection "databricks_user1-account" {
  plugin  = "databricks"
  profile = "user1-account"
  account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"
}

connection "databricks_user1-workspace" {
  plugin  = "databricks"
  profile = "user1-workspace"
  account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"
}

connection "databricks_user1-basic" {
  plugin  = "databricks"
  profile = "user1-basic"
  account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"
}
```

### Databricks Account Credentials

Configuration to query Databricks account.

#### databricks credential file:

```ini
[user1-account]
host       = https://accounts.cloud.databricks.com
token      = dsapi5c72c067b40df73ccb6be3b085d3ba
account_id = abcdd0f81-9be0-4425-9e29-3a7d96782373
```

#### databricks.spc:

```hcl
connection "databricks_user1-account" {
  plugin  = "databricks"
  profile = "user1-account"
  account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"
}
```

### Databricks Wokspace Credentials

Configuration to query Databricks workspace.

#### databricks credential file:

```ini
[user1-workspace]
host       = https://dbc-a1b2c3d4-e6f7.cloud.databricks.com/
token      = dapia865b9d1d41389ed883455032d090ee
account_id = abcdd0f81-9be0-4425-9e29-3a7d96782373
```

#### databricks.spc:

```hcl
connection "databricks_user1-workspace" {
  plugin  = "databricks"
  profile = "user1-workspace"
  account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"
}
```

### Databricks Account and Workspace Credentials

Configuration to query Databricks workspace and account using the same connection.

#### databricks.spc:

```hcl
connection "databricks_user1-workspace" {
  plugin  = "databricks"

  account_id = "abcdd0f81-9be0-4425-9e29-3a7d96782373"

  account_host = "https://accounts.cloud.databricks.com/"
  account_token = "dsapi5c72c067b40df73ccb6be3b085d3ba"

  workspace_host = "https://dbc-a1b2c3d4-e6f7.cloud.databricks.com/"
  workspace_token = "dapia865b9d1d41389ed883455032d090ee"
}
```

### Credentials from Environment Variables

Alternatively, you can also use the standard Databricks environment variables to obtain credentials **only if other argument (`profile`, `account_id`, `account_token`/`account_host`/`workspace_token`/`workspace_host`) is not specified** in the connection:

```sh
export DATABRICKS_CONFIG_PROFILE=user1-test
export DATABRICKS_TOKEN=dsapi5c72c067b40df73ccb6be3b085d3ba
export DATABRICKS_HOST=https://accounts.cloud.databricks.com
export DATABRICKS_ACCOUNT_ID=abcdd0f81-9be0-4425-9e29-3a7d96782373
export DATABRICKS_USERNAME=user@turbot.com
export DATABRICKS_PASSWORD=password
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-databricks
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
