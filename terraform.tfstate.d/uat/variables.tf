variable "storage_account_name" {}
variable "container_name" {}

variable "az_storage_access_key" {}
variable "key" {}
variable "employerdbname" {}
variable "agencydbname" {}
variable "auctiondbname" {}

variable "username" {}
variable "password" {}

variable "host" {
  default = "210.4.126.35"
}

variable "port" {
  default = 1433
}

variable "vault_mount_path" {
  default = "mssqluat"
}

variable "rolename" {}
variable "policyname" {}
variable "namespace" {}
variable "serviceaccount" {}
variable "clusterrolebindingname" {}
variable "sa_jwt_token" {}
variable "sa_ca_crt" {}
variable "k8s_host" {}
variable "config_map_name" {}
variable "config_context_auth_info" {}
variable "config_context_cluster" {}
