package helper

import "net/http"

// Sendresponse : This function sends the final response
// to the http request based on the input parameters.
func Sendresponse(statuscode int, paylod []byte, resp http.ResponseWriter) {

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(statuscode)
	resp.Write(paylod)
}
