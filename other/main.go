package main

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

func main() {
	googleClientID := "301425239768-0o5svokv0dv7e6da5m94v9446d6hakqh.apps.googleusercontent.com"
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjkzNGE1ODE2NDY4Yjk1NzAzOTUzZDE0ZTlmMTVkZjVkMDlhNDAxZTQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIzMDE0MjUyMzk3NjgtMG81c3Zva3YwZHY3ZTZkYTVtOTR2OTQ0NmQ2aGFrcWguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIzMDE0MjUyMzk3NjgtMG81c3Zva3YwZHY3ZTZkYTVtOTR2OTQ0NmQ2aGFrcWguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDgyMDU5NjA5ODk5NjkyNTgwMTUiLCJub25jZSI6IjExMTExIiwibmJmIjoxNzEyMzMwNjA1LCJuYW1lIjoia2lub2tpa2kiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jS1dZYW85Y09FQzB2c0Z0c2NNN3dDaVVUV0NtdTB2bXNMajM3RmVpdXBvQTdBclJtQlU9czk2LWMiLCJnaXZlbl9uYW1lIjoia2lub2tpa2kiLCJpYXQiOjE3MTIzMzA5MDUsImV4cCI6MTcxMjMzNDUwNSwianRpIjoiZjAyMjdkYmMxZjg0NDEyYmVmNzBhZjRlNWYzMWU5OGQzMDgwMDllMyJ9.bwtoQK2ekNJn5iD1-73jzmD5Bov40zaSFcl1k1bo45gSTbrCy-BgiS3ZlwG3wo90e2-FwfBTlJvMspEX08EuZSEIuPOBs4K1-cEZ4Jlu9Ne4cnR1seoLmRiMLR59EgyPkvZIQNav0kFJ1rsc9x3LmYKcbGoqooNx9EOZVb816ozJVbzC27WrKb_j_WcmUm1Wmlw8lUM_ybFbrf9NOz1MK9-FD_8BKVSw5Zm-oX_i69KgTMKhUGdG_5J2-wGH50ltoZ-T58eptfeS3OrLwuixNY1gQtjqB4_cajvuJ2hO3PBdB8QHEikSXYJMHhd4vHbSu8vqSj9cD9j80I9AVJO8qw"

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}
