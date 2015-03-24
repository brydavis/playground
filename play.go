package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type Output struct {
	Code []byte
	Text []byte
}

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/static/", Static)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		filename := "temp" + strconv.Itoa(int(time.Now().Unix()))
		filepath := "temp/" + filename + ".go"
		ioutil.WriteFile(filepath, []byte(r.FormValue("data")), 0700)

		exec.Command("go", "fmt", filepath).Run()
		exec.Command("goimports", "-w", filepath).Run()
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

func Templify(view, code, out []byte, w http.ResponseWriter) {
	t := template.New("output")
	t, err := t.Parse(string(view))
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(w, Output{code, out})
	if err != nil {
		fmt.Println(err)
	}
}

func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
