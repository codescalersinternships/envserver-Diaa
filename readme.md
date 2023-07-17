# Env Server

A simple web service that exposes two endpoints to get environment variables.

## EndPoints
- `/env`: returns all the environment variables as a JSON object.
- `/env/<key>`: returns the value of the environment variable with the given key as a plain text.


## Usage
To run the service, you need to have Go installed and set up on your system.

1. Clone the repo
```
git clone https://github.com/codescalersinternships/envserver-Diaa.git
```

2. Create .env file with env var PORT

```.env
PORT = 8080
```
3. Install dependencies
```go
go get -d ./...
```

4. Run the service
```
go run main.go
```

The service will listen on the port specified by the PORT environment variable (like above is 8080). You can use a web browser or a tool like curl to make requests to the endpoints. For example:

``` cmd
# Get all the environment variables
curl http://localhost:8080/env

# Get the value of the HOME environment variable
curl http://localhost:8080/env/HOME

```

## Testing

To test the service, you can use the go test command in the project directory. This will run some unit tests for the endpoints and check their status codes and responses. For example:

``` go
# Run the tests
go test -v 
```

