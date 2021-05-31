package packages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func Create() (getOne, error) {
	var dataStruct getOne
	name := viper.GetString("moldyPackageInfo.name")
	author := viper.GetString("moldyPackageInfo.author")
	url := viper.GetString("moldyPackageInfo.url")
	description := viper.GetString("moldyPackageInfo.description")
	version := viper.GetString("moldyPackageInfo.version")
	colors.Info("All information successfully obtained from MoldyFile.toml")
	password := terminal.PasswordPrompt("Password for the package save in a safe place")

	if name == "" || author == "" || url == "" || description == "" || version == "" || password == "" {
		colors.Warn("Fill all blanks")
		return dataStruct, errors.New("fill all blanks")
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
		colors.Error(dataStruct.Message)
		return dataStruct, errors.New(dataStruct.Message)
	}
	colors.Success("\nCreated successfully")
	fmt.Printf("\nThis is the data of the package created:\n\nID: %v\nName: %v\nAuthor: %v\nDescription: %v\nURL: %v\nVersion: %v\n", dataStruct.Data.Id, dataStruct.Data.Name, dataStruct.Data.Author, dataStruct.Data.Description, dataStruct.Data.Url, dataStruct.Data.Version)
	fmt.Println("\nIf some of this data not is correct you can update it using\n\nmoldy update <id>\n\nGo to the docs to know more about this command")
	return dataStruct, nil
}
