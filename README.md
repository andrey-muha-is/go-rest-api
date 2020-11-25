# go-rest-api
Simple REST API implemented with Golang. To start application successfully make sure that you have your database running locally.

# App structure
* `./src/main.go` - entry point of application
* `./src/config` - application config
* `./src/handlers` - route handlers
* `./src/models` - models that are used across application
* `./src/repositories` - repositories for each model
* `./src/utils` - utils functions for different cases

# Starting app
* Create `.env` file in project folder and set values or copy them from `.env.example`

## Running application in docker
* `docker build -t go-rest .`
* `docker run --network="host" go-rest`

## Running application locally
* `go build -o app ./src`
* `./app`