package modules

//package main

import (
	"log"
	"net/http"
)

// AuthDevAPIURL (Development)
const AuthDevAPIURL string = "http://127.0.0.1:5000"

// CheckAccess verification JWT access token
func CheckAccess(accessToken string) http.Response {
	url := AuthDevAPIURL + "/protected"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", accessToken)
	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return *response
}

// func main() {
// 	access := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2MDk1Nzg5NzIsIm5iZiI6MTYwOTU3ODk3MiwianRpIjoiMzdhNDFmZjItOWM0YS00MDlmLWEyNWEtMzQyOTMwYmNlZmYxIiwiZXhwIjoxNjA5NTgwNzcyLCJpZGVudGl0eSI6IjlhYzFhMTFmLWIwZGItNDNmNS1hODM2LTQ0ZjM3YTMwYWU1YiIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.tbYQFPU8yzpl40OOVP8vLH40XjhJl_q3_QiPXjilntA"
// 	token := "Bearer " + access
// 	result := CheckAccess(token)
// 	println(result.Status)
// }
