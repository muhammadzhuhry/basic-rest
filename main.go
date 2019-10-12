package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{
			"success": true,
			"data": "index",
			"message": "This service is running properly",
			"code": 200
		}`))
	})

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
