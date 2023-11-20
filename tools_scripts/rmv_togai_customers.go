package tools_scripts

import (
	"fmt"
	"net/http"
	"time"
)

func BulkRemoveTogaiCustomers() {
	customers := []string{}
	for _, cusId := range customers {
		removeTogaiCustomer(cusId)
	}
}

func removeTogaiCustomer(cusId string) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("https://sandbox-api.togai.com/accounts/%s", cusId), http.NoBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer ")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	time.Sleep(time.Second * 2)
}
