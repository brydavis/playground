package main

import "net/http"

import "./lib"

func main() {
	port := ":8080"
	http.HandleFunc("/", play.RootHandler)
	http.HandleFunc("/static/", StaticHandler)
	// http.HandleFunc("/editor/", play.EditorHandler)
	http.HandleFunc("/save/", play.SaveHandler)

	http.ListenAndServe(port, nil)

	// AdminTerminal(port)

}

func StaticHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}
