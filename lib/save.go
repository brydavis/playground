package play

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SaveHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		filepath := req.FormValue("filepath")
		textarea := req.FormValue("textarea")
		err := ioutil.WriteFile(filepath, []byte(textarea), 0700)
		if err != nil {
			fmt.Println(err)
		}
		
		FormatGo(filepath)
		file, _ := ioutil.ReadFile(filepath)

		res.Write(file)
	}
}
