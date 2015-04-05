package play

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		filename := "temp" + strconv.Itoa(int(time.Now().Unix()))
		filepath := "temp/" + filename + ".go"
		ioutil.WriteFile(filepath, []byte(r.FormValue("data")), 0700)

		FormatGo(filepath)
		// exec.Command("go", "fmt", filepath).Run()
		// exec.Command("goimports", "-w", filepath).Run()
		err := exec.Command("go", "build", filepath).Run()
		if err != nil {
			fmt.Println(err)
		}

		out, err := exec.Command("./" + filename).Output()
		if err != nil {
			out = []byte(err.Error())
		}

		view, _ := ioutil.ReadFile("views/base.html")
		code, _ := ioutil.ReadFile(filepath)
		Templify(view, code, out, w)

		exec.Command("rm", filename).Run()
		exec.Command("rm", filepath).Run()

	} else {
		data, _ := ioutil.ReadFile("static/example.txt")
		view, _ := ioutil.ReadFile("views/base.html")
		Templify(view, data, []byte(""), w)
	}
}
