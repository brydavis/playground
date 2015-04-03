package play

import "net/http"

func SaveHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		res.Write("Yay!")

	} else {

	}
}
