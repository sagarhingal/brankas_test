package dependency

import "brankas_test/dependency/sqlservice"

func Initall(filename string, requestID int64) {

	// Initialise SQLite3 database
	sqlservice.Initdatabase(filename, requestID)

}
