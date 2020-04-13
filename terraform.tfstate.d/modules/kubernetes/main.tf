resource "kubernetes_config_map" "this" {
  metadata {
    name      = "${var.config_map_name}"
    namespace = "${var.namespace}"
  }

  data {
    vault-agent-config.hcl     = "${file("../hcl/uat/vault-agent-config.hcl")}"
    consul-template-config.hcl = "${file("../hcl/uat/consul-template-config.hcl")}"
  }
}

resource "vault_kubernetes_auth_backend_config" "this" {
  backend            = "kubernetes"
  kubernetes_host    = "${var.k8s_host}"
  kubernetes_ca_cert = "${var.sa_ca_crt}"
  token_reviewer_jwt = "${var.sa_jwt_token}"
}

resource "vault_kubernetes_auth_backend_role" "this" {
  backend                          = "kubernetes"
  role_name                        = "${var.rolename}"
  bound_service_account_names      = ["${var.serviceaccount}"]
  bound_service_account_namespaces = ["${var.namespace}"]
  policies                         = ["${var.policyname}"]
  ttl                              = 86400
}
