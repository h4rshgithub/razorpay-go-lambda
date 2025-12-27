package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type OrderRequest struct {
	Amount   int    `json:"amount"` // in paise
	Currency string `json:"currency"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var orderReq OrderRequest
	err := json.Unmarshal([]byte(request.Body), &orderReq)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	keyId := os.Getenv("rzp_test_RvMQLhUactsOMz")
	keySecret := os.Getenv("0HdSIFlvfsdAgJDbkE6rEF3P")

	payload := map[string]interface{}{
		"amount":   orderReq.Amount,
		"currency": orderReq.Currency,
	}

	jsonData, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.razorpay.com/v1/orders",
		bytes.NewBuffer(jsonData))
	req.SetBasicAuth(keyId, keySecret)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	response, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(response),
	}, nil
}

func main() {
	lambda.Start(handler)
}
