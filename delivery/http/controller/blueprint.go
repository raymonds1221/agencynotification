package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// BlueprintController implementation of Blueprint controller
type BlueprintController struct {
}

// GetBlueprint api for returning blueprint
func (bc *BlueprintController) GetBlueprint(bi usecase.BlueprintInteractor) func(*gin.Context) {
	b, err := bi.GetBlueprint()

	if err != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error occur",
			})
		}
	}

	return func(c *gin.Context) {
		key := "j9uzyqp6cyzq"
		secret := "5y485r8nq9jre4fk6anpu59sqdcpq8xdkuqbd5jxqpvw455gek3aw27ysx4uq7tz"

		client, _ := stream.NewClient(key, secret)

		notifFeed := client.NotificationFeed("agency", "125")

		actor, _ := json.Marshal(map[string]interface{}{
			"url":        "http://example.org/martin",
			"objectType": "person",
			"id":         "tag:example.org,2011:martin",
			"image": map[string]interface{}{
				"url":    "http://example.org/martin/image",
				"width":  250,
				"height": 250,
			},
			"displayName": "Martin Smith",
		})
		object, _ := json.Marshal(map[string]string{
			"url": "http://example.org/blog/2011/02/entry",
			"id":  "tag:example.org,2011:abc123/xyz",
		})
		target, _ := json.Marshal(map[string]string{
			"url":         "http://example.org/blog/",
			"objectType":  "blog",
			"id":          "tag:example.org,2011:abc123",
			"displayName": "Martin's Blog",
		})

		resp, err := notifFeed.AddActivity(stream.Activity{
			Actor:  string(actor),
			Verb:   "post",
			Object: string(object),
			Target: string(target),
		})

		log.Print(resp)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, b)
	}
}

// GetAllBlueprint api for returning blueprint
func (bc *BlueprintController) GetAllBlueprint(bi usecase.BlueprintInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		key := "j9uzyqp6cyzq"
		secret := "5y485r8nq9jre4fk6anpu59sqdcpq8xdkuqbd5jxqpvw455gek3aw27ysx4uq7tz"

		client, _ := stream.NewClient(key, secret)

		notifFeed := client.NotificationFeed("agency", "125")

		resp, _ := notifFeed.GetActivities()

		var activities []stream.Activity

		for _, res := range resp.Results {
			for _, activity := range res.Activities {
				activities = append(activities, activity)
			}
		}

		c.JSON(http.StatusOK, activities)
	}
}
