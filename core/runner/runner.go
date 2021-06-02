package runner

import (
	"fmt"

	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/spf13/viper"
)

func RunnerWorker() {
	runVal := viper.GetStringMapString("moldyRunner")
	for n, c := range runVal {
		colors.Info("Name Of the command: \n")
		fmt.Println(n)
		colors.Info("Value of the command: \n")
		fmt.Println(c)
	}
}
