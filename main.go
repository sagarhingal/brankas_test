package main

import (
	"brankas_test/config"
	"brankas_test/dataupload"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	// Constants for the configuration
	const (
		configpath string = "env/config.yaml"
	)

	// Load configuration parameters
	var configdata config.Config
	configdata, err := config.Loadconfig(configpath)
	if err != nil {
		log.Fatalln("Loadconfig() - Unable to open the configuration file. | ", err)
	}

	// Add the routes
	http.HandleFunc("/", servetemplate)
	http.HandleFunc("/upload", dataupload.Uploadfile)
	http.HandleFunc("/getdata", dataupload.Getdata)

	// Start the server
	log.Fatalln(http.ListenAndServe(":"+configdata.Webserver.Port, nil))
}

// servetemplate : This function parses the templates and severs
// the html files in the templates dir.
func servetemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "index.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "index", dataupload.Auth{Token: config.Configdata.Webserver.Auth})
}
