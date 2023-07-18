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
2. Navigate to envserver-Diaa
```
cd envserver-Diaa
```
3. Install dependencies
```go
go mod download
```

4. build the service
```go
go build -o app cmd/main.go
```

5. Run the service.
```go
./app -p <port>
```

### Run using Docker
First make sure that docker and docker-compose is installed in your system


1. Run the container
```
docker-compose up -d
```
Now the app is running on port 8080


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
go test -v ./... 
```
