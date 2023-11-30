package tools_scripts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RandomHttpTest() {
	// url := "https://api.helloscribe.ai/api/togai/feature-credits-balance"
	url := "https://agi.helloscribe.ai/health"
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	reqBody := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&reqBody); err != nil {
		panic(err)
	}
	fmt.Println(reqBody)
}
