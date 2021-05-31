package packages

import (
	"encoding/json"
	"errors"

	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-resty/resty/v2"
)

func GetSearch(key string) (getMany, error) {
	client := resty.New()
	resp, err := client.R().Get("https://moldy-api.herokuapp.com/api/v1/packages/search?key=" + key)
	functions.CheckErrors(err, "Code 5", "The fetch to the API failed", "Unknown solution, API Internal server error")

	var dataStruct getMany

	if err := json.Unmarshal(resp.Body(), &dataStruct); err != nil {
		panic(err)
	}

	if dataStruct.Error || dataStruct.Data == nil {
		return dataStruct, errors.New(dataStruct.Message)
	}
	return dataStruct, nil
}

func GetId(id string) (getOne, error) {
	client := resty.New()
	resp, err := client.R().Get("https://moldy-api.herokuapp.com/api/v1/packages/" + id)
	functions.CheckErrors(err, "Code 5", "The fetch to the API failed", "Unknown solution, API Internal server error")

	var dataStruct getOne
	if err := json.Unmarshal(resp.Body(), &dataStruct); err != nil {
		panic(err)
	}

	if dataStruct.Error {
		return dataStruct, errors.New("not data found")
	}

	return dataStruct, nil
}
