package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type values struct {
	TempC string `json:"temp_C"`
}

type responseData struct {
	CurrentCondition []values `json:"current_condition"`
}

// GetTemp given a search returns temp in C
func GetTemp(search string) int {

	url := "http://wttr.in/" + search + "?format=j1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var jsonData responseData

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	tempResult, err := strconv.Atoi(jsonData.CurrentCondition[0].TempC)
	if err != nil {
		fmt.Println(err)
	}

	return tempResult
}
