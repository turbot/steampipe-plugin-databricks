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
  
  # OAuth secret client ID of a service principal
  # This can also be set via the `DATABRICKS_CLIENT_ID` environment variable.
  # client_id = "123-456-789"

  # OAuth secret value of a service principal
  # This can also be set via the `DATABRICKS_CLIENT_SECRET` environment variable.
  # client_secret = "dose1234567789abcde"
}
