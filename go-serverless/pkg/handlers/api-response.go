package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"

)

func ApiResponse(status int, body interface{})(*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string["Content-Type":"application/json"]}
	resp.statusCode = status

	stringBody, _:=json.Marshal(body)

	resp.Body = string(stringBody)
	return &resp, nil
}
