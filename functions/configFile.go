package functions

import (
	"encoding/json"
	"fmt"

	"github.com/Moldy-Community/CLI/utils"
	"github.com/manifoldco/promptui"
	vp "github.com/spf13/viper"
)

/* Add the default values and paths */
var (
	defaults = map[string]interface{}{
		"githubUser":     "none",
		"githubPassword": "none",
		"progressBar":    true,
		"asciiMoldy":     true,
	}
	configName  = "MoldyConfig"
	configPaths = []string{
		"$HOME/.config/moldy",
		"%USERPROFILE%/.config/moldy",
	}
)

/* Creating the config file */
func CreateConfigFile() {
	vp.SetConfigType("toml")
	for k, v := range defaults {
		vp.SetDefault(k, v)
	}
	for _, p := range configPaths {
		vp.AddConfigPath(p)
	}
	vp.SetConfigName(configName)
	vp.SafeWriteConfig()
}

/* Setting the values */
func SetConfigFile() {
	prompt := promptui.Prompt{
		Label:   "Username",
		Default: "none",
	}
	result, err := prompt.Run()
	utils.CheckErrors(err, "2", "Error in the input of the user")
	vp.Set("githubUser", result)
	promptPass := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	resultPass, err := promptPass.Run()
	utils.CheckErrors(err, "2", "Error in the input of the password")
	fmt.Printf("Your username is %q\n", result)
	fmt.Printf("Your password is %q\n", resultPass)

	byteText := []byte(resultPass)
	key := utils.GetDotEnv()
	txtEncrypt, err := utils.Encrypt(byteText, []byte(key))
	utils.CheckErrors(err, "2", "Error iSn the encryption")
	vp.Set("githubPassword", txtEncrypt)
}

/* Reading the config file :D */

func ReadConfigFile() ([]byte, interface{}) {
	githubUser := vp.Get("githubUser")
	githubPassInterface := vp.Get("githubPassword")
	var txt []byte
	json.Unmarshal(txt, &githubPassInterface)
	key := utils.GetDotEnv()
	val, err := utils.Decrypt(txt, []byte(key))
	utils.CheckErrors(err, "2", "Error in de decryption")
	return val, githubUser
}
