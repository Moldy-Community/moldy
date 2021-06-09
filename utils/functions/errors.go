/*
Copyright © 2021 Moldy-Community

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Moldy-Community/moldy/utils/colors"
)

func CheckErrors(err error, code, msg, solution string) {
	if err != nil {
		//Start the time
		dt := time.Now()
		//Write out the error detect
		colors.Warn("NEW ERROR DETECTED \n")

		//Parse as interface
		dataLog := map[string]interface{}{
			"Message":                  msg,
			"Code Error":               code,
			"Solution of error":        solution,
			"Time of the error create": dt.Format("01-02-2006 15:04:05 Monday"),
			"Error Complete":           err,
		}
		//Print with for loop
		for k, v := range dataLog {
			fmt.Println(k, v)
		}
		//Parse as json
		jsonParsed, err := json.Marshal(dataLog)
		if err != nil {
			log.Fatalf("Error in parse the map to json error: %s", err)
		}
		fmt.Print(jsonParsed)
		//Create the temp file
		tmpFile, err := ioutil.TempFile(os.TempDir(), "moldy-log-")
		if err != nil {
			log.Fatalf("Error in create the temp file :C %s", err)
		}
		defer os.Remove(tmpFile.Name())
		//Write with the content
		if _, err = tmpFile.Write(jsonParsed); err != nil {
			log.Fatal("Failed to write to temporary file", err)
		}

		// Close the file
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
