# Brankas GOLANG Test
# -------------------

# Tasks for the server:
1. Create a HTTP server which has 2 handlers - /upload && /showdata (optional)
2. A simple HTML form which has 2 inputs - auth (hidden, value provided by server) && file (takes a file uploaded by the user)
3. Save the received file to a temporary file in the directory
4. Write the file metadata to any database

# Pre-requisites for the data / file
1. File should be an image file
2. File should not exceed than 8 MB
3. Handle any error scenario for the upload on the server