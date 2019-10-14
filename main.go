package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Product for modeling data dummy
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

// database dummy
var (
	database = make(map[string]Product)
)

// SetJSONResp function for return JSON
func SetJSONResp(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func main() {

	// init db
	database["001"] = Product{ID: "001", Name: "Samsung Galaxy S10", Qty: 10}
	database["002"] = Product{ID: "002", Name: "Iphone X", Qty: 5}

	// route goes here

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

	// product
	http.HandleFunc("/products", func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Invalid http method",
				"code": 405
			}`)

			SetJSONResp(res, http.StatusMethodNotAllowed, resMessage)
			return
		}

		var products []Product

		for _, product := range database {
			products = append(products, product)
		}

		// parsing data into json
		productJSON, err := json.Marshal(&products)
		if err != nil {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Error when parsing data",
				"code": 500
			}`)

			SetJSONResp(res, http.StatusInternalServerError, resMessage)
			return
		}

		// must be fixing later
		// resMessage := []byte(`{
		// 	"success": true,
		// 	"data": productJSON,
		// 	"message": "Your request has been process"
		// 	"code": 200
		// }`)

		SetJSONResp(res, http.StatusOK, productJSON)
	})

	http.HandleFunc("/product", func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Invalid http method",
				"code": 405
			}`)

			SetJSONResp(res, http.StatusMethodNotAllowed, resMessage)
			return
		}

		if _, ok := req.URL.Query()["id"]; !ok {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Required product id",
				"code": 500
			}`)

			SetJSONResp(res, http.StatusInternalServerError, resMessage)
			return
		}

		id := req.URL.Query()["id"][0]

		product, ok := database[id]
		if !ok {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Product not found",
				"code": 404
			}`)

			SetJSONResp(res, http.StatusNotFound, resMessage)
			return
		}

		productJSON, err := json.Marshal(&product)
		if err != nil {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Error when parsing data",
				"code": 500
			}`)

			SetJSONResp(res, http.StatusInternalServerError, resMessage)
			return
		}

		SetJSONResp(res, http.StatusOK, productJSON)
	})

	http.HandleFunc("/product/add", func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "POST" {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Invalid http method",
				"code": 405
			}`)

			SetJSONResp(res, http.StatusMethodNotAllowed, resMessage)
			return
		}

		var product Product

		payload := req.Body

		defer req.Body.Close()

		err := json.NewDecoder(payload).Decode(&product)
		if err != nil {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Error when parsing data",
				"code": 500
			}`)

			SetJSONResp(res, http.StatusInternalServerError, resMessage)
			return
		}

		database[product.ID] = product

		// must be fixing later
		resMessage := []byte(`{
			"success": true,
			"data": null,
			"message": "Success create product",
			"code": 201
		}`)

		SetJSONResp(res, http.StatusCreated, resMessage)
	})

	http.HandleFunc("/product/delete", func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "DELETE" {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Invalid http method",
				"code": 405
			}`)

			SetJSONResp(res, http.StatusMethodNotAllowed, resMessage)
			return
		}

		if _, ok := req.URL.Query()["id"]; !ok {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Required product id",
				"code": 500
			}`)

			SetJSONResp(res, http.StatusInternalServerError, resMessage)
			return
		}

		id := req.URL.Query()["id"][0]
		product, ok := database[id]
		if !ok {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Product not found",
				"code": 404
			}`)

			SetJSONResp(res, http.StatusNotFound, resMessage)
			return
		}

		productJSON, err := json.Marshal(&product)
		if err != nil {

			resMessage := []byte(`{
				"success": false,
				"data": null,
				"message": "Error when parsing data",
				"code": 500
			}`)

			SetJSONResp(res, http.StatusInternalServerError, resMessage)
			return
		}

		// must be fixing later
		// resMessage := []byte(`{
		// 	"success": true,
		// 	"data": productJSON,
		// 	"message": "Success delete product"
		// 	"code": 200
		// }`)

		SetJSONResp(res, http.StatusOK, productJSON)
	})

	// listen port
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
