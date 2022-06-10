# Brankas GOLANG Test
.
.
# Tasks for the server:
1. Create a HTTP server which has 2 handlers - /upload && /getdata (optional)
2. A simple HTML form which has 2 inputs - auth (hidden, value provided by server) && file (takes a file uploaded by the user)
3. Save the received file to a temporary file in the directory
4. Write the file metadata to any database
.
.
# Pre-requisites for the data / file
1. File should be an image file
2. File should not exceed than 8 MB
3. Handle any error scenario for the upload on the server

# An overview of the code
1. This code uses a mixed pattern of "dependency-injection" and domain model.
2. Each functionality uses its own package, i.e any code related to the use-case will be under it's own package. e.g dataupload - it will contain all the services, models related to the functionality of uploading the data.
3. The "dependency" package holds all the common modules and variables that could be used by other modules. They all the initialised in the main.
4. "Main.go" is also responsible for initialising depedency object for the "dataupload" module for it's handler's business requirements.
5. To interact with the data on the go, this code uses "SQLite3" as the database source for persistent storage. 
6. The code deletes and creates a fresh database file everytime it is run fresh. (This behaviour is intentional and is only for POC purposes.)
7. YAML language is used for adding configuration variables for the service to make it more production-level and customisable. 
8. For logging, go's native "log" package is used. To track each request, a unique random number is assigned to each request which is passed down to each subroutine call till the endpoint.
9. Sufficient error handling is implemented right from entry of type of request(POST, GET) till the end response .
10. API response time is calculated and printed at the end of every response in the console log for understaing basic performance metric.

# Steps to run
1. go get .
2. go run main.go

# URLs of the service
1. http://localhost:3000/getdata [GET]
2. http://localhost:3000/upload [POST]
3. http://localhost:3000/index.html [GET]
