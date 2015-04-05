package main

import "net/http"

import . "./lib"

func main() {
	port := ":8080"
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/editor/", EditorHandler)
	http.HandleFunc("/save/", SaveHandler)

	go http.ListenAndServe(port, nil)

	AdminTerminal(port)

}

func StaticHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}
