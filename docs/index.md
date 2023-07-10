---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/databricks.svg"
brand_color: "#FFE01B"
display_name: "Databricks"
short_name: "databricks"
description: "Steampipe plugin to query audiences, automation workflows, campaigns, and more from Databricks."
og_description: "Query Databricks with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/databricks-social-graphic.png"
---

# Databricks + Steampipe

[Databricks](https://databricks.com) is a marketing automation and email marketing platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List details of your Databricks campaign:

```sql
select
  id,
  title,
  content_type,
  create_time,
  emails_sent,
  send_time,
  status,
  type
from
  databricks_campaign;
```

```
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
| id         | title                              | content_type | create_time               | emails_sent | send_time | status | type             |
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
| f739729f66 | We're here to help you get started | template     | 2023-06-16T17:51:52+05:30 | <null>      | <null>    | save   | automation-email |
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
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

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Databricks requires an [API key](https://databricks.com/developer/marketing/guides/quick-start/#generate-your-api-key/) for all requests.                                                               |
| Permissions | API keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                             |
| Radius      | Each connection represents a single Databricks Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/databricks.spc`)<br />2. Credentials specified in environment variables, e.g., `DATABRICKS_API_KEY`.                     |

### Configuration

Installing the latest databricks plugin will create a config file (`~/.steampipe/config/databricks.spc`) with a single connection named `databricks`:

```hcl
connection "databricks" {
  plugin = "databricks"

  # Databricks API key for requests. Required.
  # Generate your API Key as per: https://databricks.com/developer/marketing/guides/quick-start/#generate-your-api-key/
  # This can also be set via the `DATABRICKS_API_KEY` environment variable.
  # databricks_api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

Alternatively, you can also use the standard Databricks environment variables to obtain credentials **only if other argument (`databricks_api_key`) is not specified** in the connection:

```sh
export DATABRICKS_API_KEY=q8355689e3e6f9fd0f5630362b16b1b5-us21
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-databricks
- Community: [Slack Channel](https://steampipe.io/community/join)
