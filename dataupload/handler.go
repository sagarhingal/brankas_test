package dataupload

import (
	"brankas_test/helper"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Uploadfile(resp http.ResponseWriter, req *http.Request) {

	// Generate a unique request ID
	requestID := helper.Generatenumber()
	// Add the log statement for the handler invocation
	log.Println("[", requestID, "] | Uploadfile() - Request received by IP:", req.RemoteAddr)

	if req.Method == "POST" {

		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the Size of the file
		inputfile, handler, err := req.FormFile("myfile")

		if err != nil {
			log.Println("[", requestID, "] | Uploadfile() - Unable to retrieve the file. | ", err)
			helper.Sendresponse(http.StatusForbidden, []byte("Unable to retrieve the given file."), resp)
			return
		}

		// Size validation
		if handler.Size > 8000000 {
			log.Println("[", requestID, "] | Uploadfile() - File size exceeded. | Given file and size:", handler.Filename, " ", handler.Size/1000000, "MB.")
			helper.Sendresponse(http.StatusForbidden, []byte("File size greater than 8 MB! Please use a smaller file."), resp)
			return
		}
		// Close the file in the end
		defer inputfile.Close()

		// Log the stats of the file
		if !helper.Checktype(handler.Header.Get("Content-Type")) {
			log.Println("[", requestID, "] | Uploadfile() - Unsupported file type. | ", handler.Header.Get("Content-Type"), "Given file: ", handler.Filename)
			helper.Sendresponse(http.StatusForbidden, []byte("Unsupported file type, only image-files are allowed."), resp)
			return
		}

		// Prepare the metadata object after all the validations
		mdata := Metadata{}
		mdata.Agent = req.UserAgent()
		mdata.ClientIP = req.RemoteAddr
		mdata.Contenttype = handler.Header.Get("Content-Type")
		mdata.Originalname = handler.Filename
		mdata.Filesize = handler.Size

		// Create a temp file to save the received image
		tempfile, err := ioutil.TempFile("files", "image-*.png")
		if err != nil {
			log.Println("[", requestID, "] | Uploadfile() - Temp file creation error. | ", err)
			helper.Sendresponse(http.StatusInternalServerError, []byte("Something went wrong with the file upload. Please try again later."), resp)
			return
		}

		defer tempfile.Close()

		// Decode the received file into byte array
		filebytes, err := ioutil.ReadAll(inputfile)
		if err != nil {
			log.Println("[", requestID, "] | Uploadfile() - Unable to read the file: ", handler.Filename, " | ", err)
			helper.Sendresponse(http.StatusBadRequest, []byte("Unable to read the given file."), resp)
			return
		}

		// Now write this byte array to our temp file
		tempfile.Write(filebytes)

		// Save the metadata in the database
		// Splitting the filename from the directory name
		mdata.Newname = strings.Split(tempfile.Name(), "/")[1]
		err = Dholder.iUploadSQLservice.Savemetadata(requestID, mdata)
		if err != nil {
			helper.Sendresponse(http.StatusInternalServerError, []byte("Something went wrong with the file upload. Please try again later."), resp)
			return
		}

		// Now send the final response
		log.Println("[", requestID, "] | Uploadfile() - File [", handler.Filename, "] uploaded successfully!")
		helper.Sendresponse(http.StatusOK, []byte("File "+handler.Filename+" uploaded successfully!"), resp)

		// return
	} else {
		helper.Sendresponse(http.StatusBadRequest, []byte("Wrong method detected. Only POST supported. ClientIP: "+req.RemoteAddr), resp)
	}
}

func Getdata(resp http.ResponseWriter, req *http.Request) {
	// Add the database retrieval part
}
