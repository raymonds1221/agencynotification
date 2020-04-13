package config

import (
	"fmt"
	"io/ioutil"
	"os"

	rep "github.com/Ubidy/Ubidy_AgencyNotificationAPI/infrastructure/repository"
)

// NewSQLAgencyConfig create new sql config based on ENV environment variable
func NewSQLAgencyConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.EmployerDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return connectionString
}

// NewSQLEmployerConfig create new sql config based on ENV environment variable
func NewSQLEmployerConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.EmployerDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return connectionString
}

// NewSQLAuctionConfig create new sql config based on ENV environment variable
func NewSQLAuctionConfig() string {
	configuration := rep.NewConfig()

	host := configuration.DatabaseInfo.Host
	db := configuration.DatabaseInfo.AuctionDatabaseName
	usr := configuration.DatabaseInfo.Username
	pwd := configuration.DatabaseInfo.Password

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return connectionString
}

func developmentAgencyConfig() string {
	host := os.Getenv("AZURE_DEVELOPMENT_DB_HOST_FILE")
	db := os.Getenv("AGENCIES_DEVELOPMENT_DB_NAME_FILE")
	usr := os.Getenv("AGENCIES_DEVELOPMENT_DB_USER_FILE")
	pwd := os.Getenv("AGENCIES_DEVELOPMENT_DB_PASSWORD_FILE")

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func uatAgencyConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_UAT_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AGENCIES_UAT_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AGENCIES_UAT_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AGENCIES_UAT_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s:55107?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return cs
}

func productionAgencyConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AGENCIES_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AGENCIES_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AGENCIES_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", string(usr), string(pwd), string(host), string(db))

	return cs
}

func developmentEmployerConfig() string {
	host := os.Getenv("AZURE_DEVELOPMENT_DB_HOST_FILE")
	db := os.Getenv("EMPLOYERS_DEVELOPMENT_DB_NAME_FILE")
	usr := os.Getenv("EMPLOYERS_DEVELOPMENT_DB_USER_FILE")
	pwd := os.Getenv("EMPLOYERS_DEVELOPMENT_DB_PASSWORD_FILE")

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func uatEmployerConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_UAT_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_UAT_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_UAT_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_UAT_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s:55107?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func productionEmployerConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("EMPLOYERS_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func developmentAuctionConfig() string {
	host := os.Getenv("AZURE_DEVELOPMENT_DB_HOST_FILE")
	db := os.Getenv("AUCTIONS_DEVELOPMENT_DB_NAME_FILE")
	usr := os.Getenv("AUCTIONS_DEVELOPMENT_DB_USER_FILE")
	pwd := os.Getenv("AUCTIONS_DEVELOPMENT_DB_PASSWORD_FILE")

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func uatAuctionConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_UAT_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_UAT_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_UAT_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_UAT_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s:55107?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}

func productionAuctionConfig() string {
	host, _ := ioutil.ReadFile(os.Getenv("AZURE_DB_HOST_FILE"))
	db, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_DB_NAME_FILE"))
	usr, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_DB_USER_FILE"))
	pwd, _ := ioutil.ReadFile(os.Getenv("AUCTIONS_DB_PASSWORD_FILE"))

	cs := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", usr, pwd, host, db)

	return cs
}
