terraform {
  backend "azurerm" {}
}

data "terraform_remote_state" "state" {
  backend = "azurerm"

  config = {
    storage_account_name = "${var.storage_account_name}"
    container_name       = "${var.container_name}"
    key                  = "${var.key}"
    access_key           = "${var.az_storage_access_key}"
  }
}

provider "vault" {
  address = "http://localhost:8200"
}

module "vault" {
  source = "../modules/vault"

  vault_mount_path = "${var.vault_mount_path}"
  employerdbname   = "${var.employerdbname}"
  agencydbname     = "${var.agencydbname}"
  auctiondbname    = "${var.auctiondbname}"
  username         = "${var.username}"
  password         = "${var.password}"
  host             = "${var.host}"
  port             = "${var.port}"
  rolename         = "${var.rolename}"
  policyname       = "${var.policyname}"
}

provider "kubernetes" {
  config_context_auth_info = "${var.config_context_auth_info}"
  config_context_cluster   = "${var.config_context_cluster}"
}

module "kubernetes" {
  source = "../modules/kubernetes"

  namespace              = "${var.namespace}"
  config_map_name        = "${var.config_map_name}"
  serviceaccount         = "${var.serviceaccount}"
  clusterrolebindingname = "${var.clusterrolebindingname}"
  sa_jwt_token           = "${var.sa_jwt_token}"
  sa_ca_crt              = "${var.sa_ca_crt}"
  k8s_host               = "https://${var.k8s_host}"
  rolename               = "${var.rolename}"
  policyname             = "${var.policyname}"
}
