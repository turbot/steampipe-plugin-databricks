![image](https://hub.steampipe.io/images/plugins/turbot/databricks-social-graphic.png)

# Databricks Plugin for Steampipe

Use SQL to query audiences, automation workflows, campaigns, and more from Databricks

- **[Get started →](https://hub.steampipe.io/plugins/turbot/databricks)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/databricks/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
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

  # Authentication information
  databricks_api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

Or through environment variables:

```sh
export DATABRICKS_API_KEY=08355689e3e6f9fd0f5630362b16b1b5-us21
```

Run steampipe:

```shell
steampipe query
```

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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-databricks/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Databricks Plugin](https://github.com/turbot/steampipe-plugin-databricks/labels/help%20wanted)
