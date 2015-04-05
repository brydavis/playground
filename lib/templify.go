package play

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Templify(view, code, out []byte, w http.ResponseWriter) {
	workdir, _ := os.Getwd()
	t := template.New("output")
	t, err := t.Parse(string(view))
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(w, Output{
		code,
		out,
		workdir,
	})
	if err != nil {
		fmt.Println(err)
	}
}
