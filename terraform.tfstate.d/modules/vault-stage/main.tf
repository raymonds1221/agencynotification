resource "vault_generic_secret" "this" {
  path = "${var.vault_kv_mount_path}/${var.rolename}"

  data_json = <<EOT
  {
    "username": "${var.username}",
    "password": "${var.password}"
  }
  EOT
}

resource "vault_policy" "this" {
  name = "${var.policyname}"

  policy = <<EOT
path "${var.vault_kv_mount_path}/${var.rolename}" {
  capabilities = ["read", "create", "update"]
}
path "sys/leases/renew" {
  capabilities = ["create"]
}
path "sys/leases/revoke" {
  capabilities = ["update"]
}
  EOT
}
