package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

// Helper contains list of method] act as helper
type Helper struct {
	telemetryClient appinsights.TelemetryClient
}

// NewHelper create new instance of helper struct
func NewHelper(telemetryClient appinsights.TelemetryClient) *Helper {
	return &Helper{
		telemetryClient: telemetryClient,
	}
}

func (h *Helper) post(path string, formData url.Values) (int, string) {
	client := http.Client{}

	baseURL := "http://devmessagequeue.ubidyapi.com"

	switch env := os.Getenv("GO_ENV"); env {
	case "production":
		baseURL = "https://messagequeue.ubidyapi.com"
		break
	case "uat":
		baseURL = "https://uatmessagequeue.ubidyapi.com"
		break
	}

	u := fmt.Sprintf("%s%s", baseURL, path)

	log.Print("baseURL: ", u)
	log.Println(formData.Encode())

	req, err := http.NewRequest("POST", u, strings.NewReader(formData.Encode()))

	if err != nil {
		h.telemetryClient.TrackException(err)
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.getAccessToken()))
	resp, err := client.Do(req)

	if err != nil {
		h.telemetryClient.TrackException(err)
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	log.Printf("body: %s", string(body))

	return resp.StatusCode, string(body)
}

func (h *Helper) getAccessToken() string {
	values, _ := json.Marshal(map[string]interface{}{
		"client_id":     "JI4Vrbu8l5lUCuI9CLTo12pqqwI43scj",
		"client_secret": "KcTcvtLVlQBHzwfgpb3e_cnGJy7Izy3xYAHGfSS1JMFgXmb4hXmznM8CpwfzOxIs",
		"audience":      "https://ubidy-api-endpoint/",
		"grant_type":    "client_credentials",
	})
	client := http.Client{}

	req, err := http.NewRequest("POST", "https://ubidy.au.auth0.com/oauth/token", bytes.NewBuffer(values))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}{}

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &data)

	return data.AccessToken
}

func (h *Helper) createJSONMarshal(id string, name string, typ string) ([]byte, error) {
	return json.Marshal(map[string]string{
		"id":   id,
		"name": name,
		"type": typ,
	})
}

func (h *Helper) createJSONMarshalWithUUID(id string, name string, typ string, uuid string) ([]byte, error) {
	return json.Marshal(map[string]string{
		"id":   id,
		"name": name,
		"type": typ,
		"uuid": uuid,
	})
}
