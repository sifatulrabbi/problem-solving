package tools_scripts

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func BulkRemoveTogaiCustomers(prod bool) {
	customers := []string{}

	url := "https://sandbox-api.togai.com"
	togaiApiKey := ""

	if prod {
		url = "https://api.togai.com"

		togaiApiKey = os.Getenv("TOGAI_API_KEY_PROD")
		if togaiApiKey == "" {
			panic(errors.New("no togai api key for prod found among the env vars"))
		}
	} else {
		togaiApiKey = os.Getenv("TOGAI_API_KEY")
		if togaiApiKey == "" {
			panic(errors.New("no togai api key found among the env vars"))
		}
	}

	for _, cusId := range customers {
		removeTogaiAccount(fmt.Sprintf("%s/accounts/%s", url, cusId), togaiApiKey)
	}
	for _, cusId := range customers {
		removeTogaiCustomer(fmt.Sprintf("%s/accounts/%s", url, cusId), togaiApiKey)
	}
}

func removeTogaiAccount(cusId, togaiApiKey string) {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://sandbox-api.togai.com/accounts/%s", cusId),
		http.NoBody,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", togaiApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resBody := map[string]any{}
	json.NewDecoder(res.Body).Decode(&resBody)
	fmt.Println(res.StatusCode, resBody)
	time.Sleep(time.Second * 2)
}

func removeTogaiCustomer(cusId, togaiApiKey string) {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://sandbox-api.togai.com/customers/%s", cusId),
		http.NoBody,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", togaiApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	resBody := map[string]any{}
	json.NewDecoder(res.Body).Decode(&resBody)
	fmt.Println(res.StatusCode, resBody)
	time.Sleep(time.Second * 2)
}
