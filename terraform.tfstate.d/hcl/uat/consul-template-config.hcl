vault {
  renew_token            = false
  vault_agent_token_file = "/home/vault/.vault-token"

  retry {
    backoff = "1s"
  }
}

template {
  destination = "/etc/secrets/config.json"

  contents = <<EOF
  {
    "DatabaseInfo": {
    {{- with secret (printf "%s/creds/agencynotificationapirole" (env "VAULT_PATH")) }}
      "Host": "{{ env "DATABASE_HOST" }}",
      "EmployerDatabaseName": "{{ env "EMPLOYER_DATABASE_NAME" }}",
      "AgencyDatabaseName": "{{ env "AGENCY_DATABASE_NAME" }}",
      "AuctionDatabaseName": "{{ env "AUCTION_DATABASE_NAME" }}",
      "Port": "{{ env "DATABASE_PORT"}}",
      "Username":"{{ .Data.username }}",
      "Password":"{{ .Data.password }}"
    {{ end }}
    }
  }
  EOF
}
