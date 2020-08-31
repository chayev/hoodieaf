package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var apiKey = "" // Add in apiKey from https://developer.accuweather.com/user/me/apps

// Hoodie init
func Hoodie(input string) string {

	locationKey := getLocationKey(input)

	fmt.Println(locationKey)

	// locationKey = "1-324505_1_AL"

	url := "http://dataservice.accuweather.com/currentconditions/v1/" + locationKey + "?apikey=" + apiKey
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

	// fmt.Println(string(body))

	return string(body)
}

func getLocationKey(search string) string {

	url := "http://dataservice.accuweather.com/locations/v1/cities/search?apikey=" + apiKey + "&q=" + search
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

	output := strings.Split(strings.Split(string(body), ",")[1], ":")[1]
	output = output[1 : len(output)-1]

	return output
}
