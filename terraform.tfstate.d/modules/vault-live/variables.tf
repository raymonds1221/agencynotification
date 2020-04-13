variable "vault_mount_path" {}

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

variable "rolename" {}

variable "policyname" {}

variable "vault_kv_mount_path" {}
