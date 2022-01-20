// utils.go is responsible for parsing the body sent in the form 
// of JSON in the HTTP Request and converting it into a structure which
// Golang can understand

package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ParseBody takes in a Golang readable struct and an io.Reader and 
// unmarshals the JSON data into the struct
func ParseBody(x interface{}, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		log.Fatal(err)
	}
}
