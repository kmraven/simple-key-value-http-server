# Simple Key-Value HTTP Server

This is a lightweight HTTP server with key-value storage functionality. It allows you to store and retrieve data using specific keys through HTTP PUT and GET requests.

## Features

- Listen for incoming connections on port 8000.
- Handle requests to `/objects/{key}` where `{key}` is a string of up to 10 alphanumeric characters.
- **PUT** Method:
  - Store arbitrary byte data in the request body under the specified key.
  - Responds with status code 200 upon successful storage.
- **GET** Method:
  - Retrieve previously stored data associated with the specified key.
  - Responds with status code 200 and the data if found; responds with status code 404 if the key does not exist.
- Other Methods:
  - Responds with status code 405 (Method Not Allowed) for unsupported methods.
- All other paths:
  - Responds with status code 404 (Not Found).

## Usage

1. Start the server
```
% CGO_ENABLED=0 go run ./main.go
```

2. Store data
```
% curl -X PUT -d "Hello, world!" http://localhost:8000/objects/mykey
```
3. Retrieve data:  
```
% curl http://localhost:8000/objects/mykey
```