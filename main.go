package main

import (
	"os"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/http"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/infrastructure/repository"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/infrastructure/repository/mssql"
	mssqlconfig "github.com/Ubidy/Ubidy_AgencyNotificationAPI/infrastructure/repository/mssql/config"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

var (
	telemetryClient         appinsights.TelemetryClient
	tokenInteractor         *usecase.TokenInteractor
	settingsInteractor      *usecase.SettingsInteractor
	applicationInteractor   *usecase.ApplicationInteractor
	clarificationInteractor *usecase.ClarificationInteractor
	biddingInteractor       *usecase.BiddingInteractor
	fulfillmentInteractor   *usecase.FulfillmentInteractor
	activityInteractor      *usecase.ActivityInteractor
	notificationInteractor  *usecase.NotificationInteractor
	commentInteractor       *usecase.CommentInteractor
)

func init() {
	client := createActivityStreamClient()
	telemetryClient = appinsights.NewTelemetryClient("4e6c52a9-21db-4f8f-8bf6-b1b28639b1e3")

	employerDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLEmployerConfig())
	agencyDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLAgencyConfig())
	auctionDB, _ := mssql.NewDBConnection(mssqlconfig.NewSQLAuctionConfig())

	tokenRepository := repository.NewTokenRepository(client, telemetryClient)
	settingsRepository := repository.NewSettingsRepository(employerDB, agencyDB, telemetryClient)
	applicationRepository := repository.NewApplicationRepository(client, employerDB, auctionDB, telemetryClient)
	clarificationRepository := repository.NewClarificationRepository(client, employerDB, auctionDB, telemetryClient)
	biddingRepository := repository.NewBiddingRepository(client, employerDB, auctionDB, telemetryClient)
	fulfillmentRepository := repository.NewFulfillmentRepository(client, employerDB, auctionDB, telemetryClient)
	activityRepository := repository.NewActivityRepository(client)
	notificationRepository := repository.NewNotificationRepository(client)
	commentRepository := repository.NewCommentRepository(client, employerDB, auctionDB, telemetryClient)

	tokenInteractor = usecase.NewTokenInteractor(tokenRepository)
	settingsInteractor = usecase.NewSettingsInteractor(settingsRepository)
	applicationInteractor = usecase.NewApplicationInteractor(applicationRepository)
	clarificationInteractor = usecase.NewClarificationInteractor(clarificationRepository)
	biddingInteractor = usecase.NewBiddingInteractor(biddingRepository)
	fulfillmentInteractor = usecase.NewFulfillmentInteractor(fulfillmentRepository)
	activityInteractor = usecase.NewActivityInteractor(activityRepository)
	notificationInteractor = usecase.NewNotificationInteractor(notificationRepository)
	commentInteractor = usecase.NewCommentInteractor(commentRepository)
}

func main() {
	http := http.New(
		tokenInteractor,
		settingsInteractor,
		applicationInteractor,
		clarificationInteractor,
		biddingInteractor,
		fulfillmentInteractor,
		activityInteractor,
		notificationInteractor,
		commentInteractor)
	http.Run(":5021")
}

func createActivityStreamClient() *stream.Client {
	switch env := os.Getenv("ENV"); env {
	case "staging", "live":
		client, _ := stream.NewClientFromEnv()
		return client
	default:
		key := "3qvb6784xdde"
		secret := "c77yxbqmssdctscbzeckfpv96yk6nmmnyrkq4bab9xr7f4hgm85nkcjsg4gz59jd"
		client, _ := stream.NewClient(key, secret)
		return client
	}
}
