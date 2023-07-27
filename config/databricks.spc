connection "databricks" {
  plugin = "databricks"

  # The target Databricks account SCIM token.
  # See: https://docs.databricks.com/administration-guide/account-settings/index.html#generate-a-scim-token
  # This can also be set via the DATABRICKS_TOKEN environment variable
  account_token = "dsapi5c72c067b40d88f73ccb6be3b085d3ba"

  # The target Databricks account console URL, which is typically https://accounts.cloud.databricks.com.
  # This can also be set via the DATABRICKS_HOST environment variable.
  account_host = "https://accounts.cloud.databricks.com/"

  # The target Databricks workspace Personal Access Token.
  # This can also be set via the DATABRICKS_TOKEN environment variable
  # See: https://docs.databricks.com/dev-tools/auth.html#databricks-personal-access-tokens-for-users
  workspace_token = "dsapi5c72c067b40d88f73ccb6be3b085d3ba"

  # The target Databricks workspace URL.
  # See https://docs.databricks.com/workspace/workspace-details.html#workspace-url
  # This can also be set via the DATABRICKS_HOST environment variable.
  workspace_host = "https://dbc-a1b2345c-d6e7.cloud.databricks.com"

  # The target Databricks account ID.
  # See Locate your account ID: https://docs.databricks.com/administration-guide/account-settings/index.html#account-id.
  account_id = "d26d0f81-9be0-4425-9e29-3a7d96782373"

  # A connection profile specified within .databrickscfg to use instead of DEFAULT.
  config_profile = "databricks-dev"

  # A non-default location of the Databricks CLI credentials file.
  config_file = "/Users/username/.databrickscfg"

  # The Databricks username part of basic authentication. Only possible when Host is *.cloud.databricks.com (AWS).
  data_username = "user@turbot.com"

  # The Databricks password part of basic authentication. Only possible when Host is *.cloud.databricks.com (AWS).
  data_password = "password"
}
