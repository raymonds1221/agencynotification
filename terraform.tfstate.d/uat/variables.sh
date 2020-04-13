export TF_VAR_az_storage_access_key=$(az storage account keys list -g MC_Ubidy.IT.Kubernetes.AustraliaEast.Production_ubidy-kube-prod_australiaeast -n ubidykubestorage | jq -r '.[0].value') 
export VAULT_SA_NAME=$(kubectl get sa -n jx-uat agencynotificationapi-vault -o jsonpath="{.secrets[*]['name']}")
export TF_VAR_sa_jwt_token=$(kubectl get secret -n jx-uat $VAULT_SA_NAME -o jsonpath="{.data.token}" | base64 --decode; echo)
export TF_VAR_sa_ca_crt=$(kubectl get secret -n jx-uat $VAULT_SA_NAME -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)
export TF_VAR_k8s_host=$(kubectl exec -n vault consul-0 -- sh -c 'echo $KUBERNETES_SERVICE_HOST')

export TF_VAR_storage_account_name="ubidykubestorage"
export TF_VAR_container_name="uattfstate"
export TF_VAR_key="uat.terraform.agencynotification.tfstate"

# echo "TF_VAR_az_storage_access_key=$TF_VAR_az_storage_access_key"
# echo "VAULT_SA_NAME=$VAULT_SA_NAME"
# echo "SA_JWT_TOKEN=$TF_VAR_sa_jwt_token"
# echo "SA_CA_CRT=$TF_VAR_sa_ca_crt"
# echo "K8S_HOST=$TF_VAR_k8s_host"
