package locks

import (
	"github.com/Moldy-Community/moldy/core/packages"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	darthdb "github.com/TeoDev1611/darth-db/json"
	"github.com/google/uuid"
)

func GetContents(pkgname string) map[string]interface{} {
	data, err := packages.GetSearch(pkgname)
	functions.CheckErrors(err, "Code 2", "Error in lock the package ", "Check if the package exisist or report the error on github")

	for i, val := range data.Data {
		if i == 0 {
			Info := map[string]interface{}{
				"name":    val.Name,
				"version": val.Version,
				"id":      val.Id,
				"url":     val.Url,
			}
			return Info
		}
        break
	}
	return map[string]interface{}{}
}

func WriteLock(name string) {
	data := GetContents(name)
	darthdb.WriteDB("./MoldyFile.lock", "  ", false, data)
	colors.Success("Succesfuly locked the package")
}

func WriteLockUrl(url string) {
	id := uuid.New().String()
	data := map[string]interface{}{
		"name": url,
		"type": "url",
		"url":  url,
		"id":   id,
	}
	darthdb.WriteDB("./MoldyFile.lock", "  ", false, data)
	colors.Success("Succesfuly locked the package")
}
