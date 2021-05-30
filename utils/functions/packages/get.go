package packages

import (
	"encoding/json"
	"errors"

	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-resty/resty/v2"
)

type dataPackage struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type getMany struct {
	Data    []dataPackage `json:"data"`
	Error   bool          `json:"error"`
	Message string        `json:"message"`
}

type getOne struct {
	Data    dataPackage `json:"data"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
}

func GetSearch(key string) (getMany, error) {
	client := resty.New()
	resp, err := client.R().Get("https://moldy-api.herokuapp.com/api/v1/packages/search?key=" + key)
	functions.CheckErrors(err, "Code 5", "The fetch to the API failed", "Unknown solution, API Internal server error")

	var dataStruct getMany

	if err := json.Unmarshal(resp.Body(), &dataStruct); err != nil {
		panic(err)
	}

	if dataStruct.Data == nil {
		return dataStruct, errors.New("Not data found")
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
		return dataStruct, errors.New("Not data found")
	}

	return dataStruct, nil
}
