package packages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-resty/resty/v2"
)

func Update(id string) (getOne, error) {
	var dataStruct getOne
	name := terminal.BasicPrompt("Name", "")
	author := terminal.BasicPrompt("Author", "")
	url := terminal.BasicPrompt("URL", "")
	description := terminal.BasicPrompt("Description", "")
	version := terminal.BasicPrompt("Version", "")
	password := terminal.PasswordPrompt("Password")
	newPassword := terminal.PasswordPrompt("New password")

	if name == "" || author == "" || url == "" || description == "" || version == "" || password == "" || newPassword == "" {
		colors.Warn("Fill all blanks")
		return dataStruct, errors.New("fill all blanks")
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"name":"%v", "author":"%v", "url":"%v", "description":"%v", "version":"%v", "password":"%v", "newpassword":"%v"}`, name, author, url, description, version, password, newPassword)).
		Put("https://moldy-api.herokuapp.com/api/v1/packages/update/" + id)

	functions.CheckErrors(err, "Code 5", "The fetch to the API failed", "Unknown solution, API Internal server error")
	if err := json.Unmarshal(resp.Body(), &dataStruct); err != nil {
		panic(err)
	}

	if dataStruct.Error {
		colors.Error(dataStruct.Message)
		return dataStruct, errors.New(dataStruct.Message)
	}

	colors.Success("Updated successfully")
	fmt.Printf("\nThis is the data of the package now:\nID: %v\nName: %v\nAuthor: %v\nDescription: %v\nURL: %v\nVersion: %v\n", dataStruct.Data.Id, dataStruct.Data.Name, dataStruct.Data.Author, dataStruct.Data.Description, dataStruct.Data.Url, dataStruct.Data.Version)
	return dataStruct, nil
}
