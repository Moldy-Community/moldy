package packages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-resty/resty/v2"
)

func Delete(id, password string) (getOne, error) {
	var dataStruct getOne

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"password":"%v"}`, password)).
		Delete("https://moldy-api.herokuapp.com/api/v1/packages/delete/" + id)

	functions.CheckErrors(err, "Code 5", "The fetch to the API failed", "Unknown solution, API Internal server error")
	if err := json.Unmarshal(resp.Body(), &dataStruct); err != nil {
		panic(err)
	}

	if dataStruct.Error {
		colors.Error(dataStruct.Message)
		return dataStruct, errors.New(dataStruct.Message)
	}

	colors.Success("Package deleted successfully with the id:\n" + id)
	return dataStruct, nil
}
