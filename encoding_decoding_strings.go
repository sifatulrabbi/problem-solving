package main

import (
	"encoding/base64"
	"fmt"
)

func RunEncodingDecodingStrings() {
	str := "TFMwdExTMUNSVWRKVGlCUVVrbFdRVlJGSUV0RldTMHRMUzB0Q2sxSlNVVjJRVWxDUVVSQlRrSm5hM0ZvYTJsSE9YY3dRa0ZSUlVaQlFWTkRRa3RaZDJkblUybEJaMFZCUVc5SlFrRlJReXRrWTBFd01qaFFRMWxHU25NS1lqQmhURVZ5U0dkMmRYZHZRakZQUjNSRVNIUnRSRmN6VFhScmJGZzFUbkpxV2psd2FXc3pTbEUwUTFKNWJXcEhibEJGVDBKRk1HeGtXR3RCYkc1c2FncEZTME5MZG5sQ056bFlWMU5xUjNsUWJrNHZUbnA1U0RKUWRqZExUV2xwY25wbWJWRklaekZ6U25kd2NpdHNTVms0WVdSc2VtUTRXVGRuYkVsNVEyVTNDbE42Ym1keGRYaHROSFpsY1dWRFkwUm1RazVvUjBKVFNqZDJaR0ZUZGl0RU1WTk1SRzlOVkM5SmFWVkJWRXhxWTFBNVEwRm9aVXBUY2pWMlRTc3lLMWtLVjI1eGR6aHBVV1IyYkVWT2MwcGxaa1ZQYlhkMVFWWmlNWGRVZGtRMVJrWlhWalZQTnpseEwydG5ZM2cwWVhwMVkxRTBUMlJLYmxkcVZXeGxWV3RUZFFveGNrSTBaSGRsVEU5dGFIc"
	fmt.Println(b64Decode(b64Decode(str)))
}

func b64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func b64Decode(str string) string {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(b)
}
