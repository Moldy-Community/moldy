package packages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-resty/resty/v2"
)

func Create() (getOne, error) {
	var dataStruct getOne
	name := terminal.BasicPrompt("Name", "")
	author := terminal.BasicPrompt("Author", "")
	url := terminal.BasicPrompt("URL", "")
	description := terminal.BasicPrompt("Description", "")
	version := terminal.BasicPrompt("Version", "")
	password := terminal.PasswordPrompt("Password")

	if name == "" || author == "" || url == "" || description == "" || version == "" || password == "" {
		fmt.Println("Please fill all blanks")
		return dataStruct, errors.New("Fill all blanks")
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"name":"%v", "author":"%v", "url":"%v", "description":"%v", "version":"%v", "password":"%v"}`, name, author, url, description, version, password)).
		Post("https://moldy-api.herokuapp.com/api/v1/packages/new")

	functions.CheckErrors(err, "Code 5", "The fetch to the API failed", "Unknown solution, API Internal server error")
	if err := json.Unmarshal(resp.Body(), &dataStruct); err != nil {
		panic(err)
	}

	if dataStruct.Error {
		fmt.Println(dataStruct.Message)
		return dataStruct, errors.New(dataStruct.Message)
	}

	fmt.Printf("Created successfully\nThis is the data of the package created:\nID: %v\nName: %v\nAuthor: %v\nDescription: %v\nURL: %v\nVersion: %v\n", dataStruct.Data.Id, dataStruct.Data.Name, dataStruct.Data.Author, dataStruct.Data.Description, dataStruct.Data.Url, dataStruct.Data.Version)
	fmt.Println("If some of this data not is correct you can update it using\nmoldy update <id>\nGo to the docs to know more about this command")
	return dataStruct, nil
}
