package helper

import (
	"crypto/rand"
	"net/http"
)

var Validtypes map[string]bool

// Sendresponse : This function sends the final response
// to the http request based on the input parameters.
func Sendresponse(statuscode int, paylod []byte, resp http.ResponseWriter) {

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(statuscode)
	resp.Write(paylod)
}

// Checktype : This function returns a boolean value
// based on the presence of passed content-type in the map.
func Checktype(contenttype string) bool {

	if Validtypes[contenttype] {
		return true
	} else {
		return false
	}
}

// Generatenumber : This function returns a random number using
// Prime function from the crypto/rand package.
func Generatenumber() int64 {
	number, _ := rand.Prime(rand.Reader, 16)
	return number.Int64()
}
