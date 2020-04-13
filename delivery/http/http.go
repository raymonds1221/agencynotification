package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/http/controller"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/auth0-community/auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

const apiVersion = "v2.1"

// App for creating http
type App struct {
	Router *gin.Engine
}

// New create new instance of Http app
func New(
	tokenInteractor usecase.TokenInteractor,
	settingsInteractor usecase.SettingsInteractor,
	applicationInteractor usecase.ApplicationInteractor,
	clarificationInteractor usecase.ClarificationInteractor,
	biddingInterator usecase.BiddingInteractor,
	fulfillmentInteractor usecase.FulfillmentInteractor,
	activityInteractor usecase.ActivityInteractor,
	notificationInteractor usecase.NotificationInteractor,
	commentInteractor usecase.CommentInteractor,
) *App {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,Authorization,X-Requested-With,X-PINGOTHER,Access-Control-Allow-Origin,Accept,x-http-method-override")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Request-Headers", "authorization,content-type,x-requested-with")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		c.Next()
	})
	router.OPTIONS("/*cors", func(c *gin.Context) {
	})

	tokenController := controller.TokenController{}
	settingsController := controller.SettingsController{}
	applicationController := controller.ApplicationController{}
	clarificationController := controller.ClarificationController{}
	biddingController := controller.BiddingController{}
	fulfillmentController := controller.FulfillmentController{}
	activityController := controller.ActivityController{}
	notificationController := controller.NotificationController{}
	commentController := controller.CommentController{}

	v2 := router.Group(fmt.Sprintf("/api/%s", apiVersion)) //.Use(AuthMiddleware())
	{
		v2.GET("/activities/tokens/:userID", AuthMiddleware(tokenController.GetToken(tokenInteractor)))
		v2.GET("/activities/settings/clients/:clientID", AuthMiddleware(settingsController.GetSettingsByClientID(settingsInteractor)))
		v2.GET("/activities/settings/suppliers/:supplierID", AuthMiddleware(settingsController.GetSettingsBySupplierID(settingsInteractor)))
		v2.PATCH("/activities/settings", AuthMiddleware(settingsController.CreateOrUpdateSettings(settingsInteractor)))

		v2.POST("/activities/applications/apply", AuthMiddleware(applicationController.AddSubmitApplicationActivity(applicationInteractor)))
		v2.POST("/activities/applications/withdraw", AuthMiddleware(applicationController.AddWithdrawApplicationActivity(applicationInteractor)))
		v2.POST("/activities/applications/apply/successfee", AuthMiddleware(applicationController.AddSubmitApplicationSuccessFeeActivity(applicationInteractor)))
		v2.POST("/activities/applications/withdraw/successfee", AuthMiddleware(applicationController.AddWithdrawApplicationSuccessFeeActivity(applicationInteractor)))

		v2.POST("/activities/clarifications/topic", AuthMiddleware(clarificationController.AddPostTopicActivity(clarificationInteractor)))
		v2.POST("/activities/clarifications/question", AuthMiddleware(clarificationController.AddPostQuestionActivity(clarificationInteractor)))
		v2.POST("/activities/clarifications/topic/successfee", AuthMiddleware(clarificationController.AddPostTopicSuccessFeeActivity(clarificationInteractor)))
		v2.POST("/activities/clarifications/question/successfee", AuthMiddleware(clarificationController.AddPostQuestionSuccessFeeActivity(clarificationInteractor)))

		v2.POST("/activities/bidding/placebid", AuthMiddleware(biddingController.AddPlaceBidActivity(biddingInterator)))

		v2.POST("/activities/fulfillment/newcandidate", AuthMiddleware(fulfillmentController.AddNewCandidateActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/updatecandidate", AuthMiddleware(fulfillmentController.AddUpdateCandidateActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/candidatesubmission/3daysidle", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateSubmission3DaysIdleActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/candidatesubmission/10daysidle", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateSubmission10DaysIdleActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/candieatesubmission/14daysidle", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateSubmission14DaysIdleActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/newcandidate/successfee", fulfillmentController.AddNewCandidateSuccessFeeActivity(fulfillmentInteractor))
		v2.POST("/activities/fulfillment/updatecandidate/successfee", fulfillmentController.AddUpdateCandidateSuccessFeeActivity(fulfillmentInteractor))
		v2.POST("/activities/fulfillment/candidatesubmission/3daysidle/successfee", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateSubmission3DaysIdleSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/candidatesubmission/10daysidle/successfee", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateSubmission10DaysIdleSuccessFeeActivity(fulfillmentInteractor)))
		v2.POST("/activities/fulfillment/candidatesubmission/14daysidle/successfee", AuthMiddlewareServiceToService(fulfillmentController.AddCandidateSubmission14DaysIdleSuccessFeeActivity(fulfillmentInteractor)))

		v2.GET("/suppliers/activities/:supplierID", AuthMiddleware(activityController.GetActivities(activityInteractor)))
		v2.GET("/suppliers/notifications/:userID", AuthMiddleware(notificationController.GetNotifications(notificationInteractor)))
		v2.PUT("/suppliers/notifications/archive", AuthMiddleware(notificationController.UpdateNotificationArchive(notificationInteractor)))
		v2.PUT("/suppliers/notifications/viewed", AuthMiddleware(notificationController.UpdateNotificationViewed(notificationInteractor)))

		v2.POST("/activities/comment", AuthMiddleware(commentController.AddCommentToCandidateActivity(commentInteractor)))
		v2.POST("/activities/comment/successfee", AuthMiddleware(commentController.AddCommentToCandidateSuccessFeeActivity(commentInteractor)))
	}

	router.GET("/health/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	app := App{
		Router: router,
	}

	return &app
}

// Run start the server
func (a *App) Run(addr ...string) error {
	requireTLS := os.Getenv("GO_ENV") == "production" || os.Getenv("GO_ENV") == "uat"

	if requireTLS {
		certFile := "./cert.pem"
		privateKeyFile := "./server.key"
		return a.Router.RunTLS(addr[0], certFile, privateKeyFile)
	}

	return a.Router.Run(addr[0])
}

// AuthMiddleware apply middleware for the Auth0
func AuthMiddleware(f func(*gin.Context)) func(*gin.Context) {
	t := func(c *gin.Context) {
		const JWKSURI = "https://accounts.ubidyapp.com/.well-known/jwks.json"
		const issuer = "https://accounts.ubidyapp.com/"
		var audience = []string{"https://ubidy-api-endpoint/"}

		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKSURI}, nil)
		config := auth0.NewConfiguration(client, audience, issuer, jose.RS256)

		validator := auth0.NewValidator(config, nil)
		token, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println("Token is not valid or missing token: ", token, err.Error())
			c.JSON(http.StatusForbidden, gin.H{
				"errors:": map[string]string{
					"userMessage":     "Token is not valid or missing token",
					"internalMessage": err.Error(),
				},
			})
			c.Abort()
		} else {
			c.Next()
			f(c)
		}
	}

	return t
}

// AuthMiddlewareServiceToService apply middleware for the Auth0 on service to service microservices
func AuthMiddlewareServiceToService(f func(*gin.Context)) func(*gin.Context) {
	t := func(c *gin.Context) {
		const JWKSURI = "https://accounts.ubidyapp.com/.well-known/jwks.json"
		const issuer = "https://ubidy.au.auth0.com/"
		var audience = []string{"https://ubidy-api-endpoint/"}

		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKSURI}, nil)
		config := auth0.NewConfiguration(client, audience, issuer, jose.RS256)

		validator := auth0.NewValidator(config, nil)
		token, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println("Token is not valid or missing token: ", token, err.Error())
			c.JSON(http.StatusForbidden, gin.H{
				"errors:": map[string]string{
					"userMessage":     "Token is not valid or missing token",
					"internalMessage": err.Error(),
				},
			})
			c.Abort()
		} else {
			c.Next()
			f(c)
		}
	}

	return t
}
