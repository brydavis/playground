package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Output struct {
	Code []byte
	Text []byte
}

func main() {
	http.HandleFunc("/", Root)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ioutil.WriteFile("compile/example.go", []byte(r.FormValue("data")), 0700)
		exec.Command("go", "fmt", "compile/example.go").Run()
		exec.Command("goimports", "-w", "compile/example.go").Run()
		err := exec.Command("go", "build", "compile/example.go").Run()
		if err != nil {
			fmt.Println(err)
		}

		out, err := exec.Command("./example").Output()
		if err != nil {
			// fmt.Println(err)
			out = []byte(err.Error())
		}

		exec.Command("rm", "example").Run()
		view, _ := ioutil.ReadFile("views/base.html")
		t := template.New("output")
		t, err = t.Parse(string(view))
		if err != nil {
			fmt.Println(err)
		}

		code, _ := ioutil.ReadFile("compile/example.go")
		err = t.Execute(w, Output{code, out})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		data, _ := ioutil.ReadFile("form/data.go")
		view, _ := ioutil.ReadFile("views/base.html")

		// w.Write(view)

		t := template.New("output")
		t, err := t.Parse(string(view))
		if err != nil {
			fmt.Println(err)
		}

		// w.Write(out)
		err = t.Execute(w, Output{[]byte(data), []byte("")})
		if err != nil {
			fmt.Println(err)
		}

	}
}
