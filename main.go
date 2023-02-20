package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env Files")
	}

	var TWILIO_ACCOUNT_SID string = os.Getenv("TWILIO_ACCOUNT_SID")
	var TWILIO_AUTH_TOKEN string = os.Getenv("TWILIO_AUTH_TOKEN")
	// var VERIFY_SERVICE_SID string = os.Getenv("VERFIY_SERVICE_SID")

	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+233209739442")
	params.SetFrom("+233558485290")
	params.SetBody("Hello, this is Tea-Code....")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
