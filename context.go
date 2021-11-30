package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Req struct {
	Number int `json:"number"`
}

type Res struct {
	Number int `json:"number"`
	Square int `json:"square"`
}

func main() {
	router := httprouter.New()
	router.GET("/square", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Add("Content-Type", "application/json")
		request := Req{}
		reqBuffer, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBuffer, &request)
		ans := square(request.Number)
		response := Res{request.Number, ans}
		json.NewEncoder(w).Encode(response)
	})
	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func square(n int) int {
	return n * n
}
