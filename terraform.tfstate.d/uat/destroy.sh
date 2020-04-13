source variables.sh

read -p 'Username: ' username
read -sp 'Password: ' password
echo
read -p 'MongoDB Username: ' mongodb_username
read -sp 'MongoDB Password: ' mongodb_password
echo

terraform destroy \
 -var "employerdbname=UbidyServicesEmployersDatabase" \
 -var "agencydbname=UbidyServicesAgenciesDatabase" \
 -var "auctiondbname=UbidyServicesAuctionsDatabase" \
 -var "username=$username" \
 -var "password=$password" \
 -var "host=210.4.126.35" \
 -var "port=55107" \
 -var "key=uat.terraform.agencynotificationapi.tfstate" \
 -var "config_map_name=agencynotificationapi-config" \
 -var "rolename=agencynotificationapirole" \
 -var "policyname=agencynotificationapipolicy" \
 -var "namespace=jx-uat" \
 -var "serviceaccount=agencynotificationapi-vault" \
 -var "clusterrolebindingname=role-agencynotificationapi-binding" \
 -var "sa_jwt_token=$TF_VAR_sa_jwt_token" \
 -var "sa_ca_crt=$TF_VAR_sa_ca_crt" \
 -var "k8s_host=$TF_VAR_k8s_host" \
 -var "config_context_auth_info=clusterUser_Ubidy.IT.Kubernetes.AustraliaEast.Production_ubidy-kube-prod" \
 -var "config_context_cluster=ubidy-kube-prod" \
