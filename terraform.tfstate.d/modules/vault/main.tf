resource "vault_database_secret_backend_connection" "this" {
  name          = "${var.employerdbname}"
  allowed_roles = ["${var.rolename}", "employerapirole", "agencyemailalertapirole", "employeremailalertapirole", "employernotificationapirole"]
  backend       = "${var.vault_mount_path}"

  mssql {
    connection_url = "sqlserver://${var.username}:${var.password}@${var.host}:${var.port}"
  }
}

resource "vault_database_secret_backend_role" "this" {
  backend = "${var.vault_mount_path}"
  name    = "${var.rolename}"
  db_name = "${var.employerdbname}"

  creation_statements = [
    "USE ${var.employerdbname};",
    "CREATE LOGIN [{{name}}] WITH PASSWORD = '{{password}}';",
    "CREATE USER [{{name}}] FOR LOGIN [{{name}}];",
    "GRANT UNMASK TO [{{name}}];",
    "GRANT INSERT,SELECT,UPDATE,DELETE ON SCHEMA::dbo TO [{{name}}];",
    "USE ${var.agencydbname};",
    "CREATE USER [{{name}}] FOR LOGIN [{{name}}];",
    "GRANT UNMASK TO [{{name}}];",
    "GRANT INSERT,SELECT,UPDATE,DELETE ON SCHEMA::dbo TO [{{name}}];",
    "USE ${var.auctiondbname};",
    "CREATE USER [{{name}}] FOR LOGIN [{{name}}];",
    "GRANT UNMASK TO [{{name}}];",
    "GRANT INSERT,SELECT,UPDATE,DELETE ON SCHEMA::dbo TO [{{name}}];",
  ]

  default_ttl = 3600
  max_ttl     = 86400
}

resource "vault_policy" "this" {
  name = "${var.policyname}"

  policy = <<EOT
path "${var.vault_mount_path}/creds/${var.rolename}" {
  capabilities = ["read"]
}
path "sys/leases/renew" {
  capabilities = ["create"]
}
path "sys/leases/revoke" {
  capabilities = ["update"]
}
  EOT
}
