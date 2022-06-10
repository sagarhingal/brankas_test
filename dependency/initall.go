package dependency

import "brankas_test/dependency/sqlservice"

func Initall(filename string) {

	// Initialise SQLite3 database
	sqlservice.Initdatabase(filename)

}
