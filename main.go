package main

import (
	"fmt"
	"net/http"
	"os"
)

// SetJSONResp function for return JSON
func SetJSONResp(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func main() {

	// root
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		resMessage := []byte(`{
			"success": true,
			"data": "index",
			"message": "This service is running properly",
			"code": 200
		}`)

		SetJSONResp(res, http.StatusOK, resMessage)
	})

	// listen port
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
