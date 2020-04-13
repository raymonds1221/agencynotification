export TF_VAR_az_storage_access_key=$(az storage account keys list -g MC_Ubidy.IT.Kubernetes.AustraliaEast.Production_ubidy-kube-prod_australiaeast -n ubidykubestorage | jq -r '.[0].value') 
export TF_VAR_storage_account_name="ubidykubestorage"
export TF_VAR_container_name="stagingtfstate"
export TF_VAR_key="stage.terraform.agencynotification.tfstate"

terraform init \
-backend-config "storage_account_name=$TF_VAR_storage_account_name" \
-backend-config "container_name=$TF_VAR_container_name" \
-backend-config "key=$TF_VAR_key" \
-backend-config "access_key=$TF_VAR_az_storage_access_key" 
