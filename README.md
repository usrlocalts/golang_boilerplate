# Golang Boilerplate
This is a sample golang boilerplate which actually creates posts given a topic and information.

#About this project
This project supports the following integrations:

- New Relic for application monitoring
- Sentry for pushing warn+ level logs.
- Swagger - For API documentation

New Relic monitoring has been done for both API and Postgres DB Response times in this project.

Both New Relic and Sentry can be toggled off by setting the `sentry_enabled` / `new_relic_enabled` to false in application.yml.

## Technical Components Used
Glide - Package Manager
DB - Postgres
Log Integration - Sentry
App Monitoring - New Relic
API Documentation - Swagger/API Blueprint

## Migrations
Migrations can be found in `migrations` directory. This uses `github.com/mattes/migrate` for running migrations

A Migration can be also rolled back by `./golang_boilerplate rollback`

##Project Structure
- This project uses a feature based package management system. So, all the components related to posts creation viz handlers, services, request, response, repositories will be found in a single package.
- Appcontext, Logger, DB are initialized in main.go, and injected to different components.
- This project doesn't use any framework for dependency injection
- Common Handlers such as `ping` and `notfound` can be found in `handler` package 
- This project also has a framework for custom errors and corresponding error codes which can be found in `errors` package. 
- Test related DB configs, logs, and servers can be found in `testutil` package
- AppContext holds config and logger.
- DB related details are found in `db` package
- This project also uses `https://github.com/gojek-engineering/goconfig` which abstracts the way configs are read from `application.yml`

## Setup

This service runs on go.

- Install go
    - On OSX run `brew install go`.
    - Follow instructions on https://golang.org/doc/install for other OSes.
- Setup go
    - Make sure that the executable `go` is in your shell's path.
    - Add the following in your .zshrc or .bashrc: (where `<workspace_dir>` is the directory in
        which you'll checkout your code)

```
GOPATH=<workspace_dir>
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

- Checkout the code, run `make setup`

This does the following things:

- Go get lint
- Go get Datadog
- glide install
- creates a boilerplate postgres super user. Skip if already exists
- Creates application.yml
- Creates golang_boilerplate_dev
- Creates Out folder if doesn't exist
- Builds the binary file
- Runs Migrations

This task is idempotent
## Test
`make test`

## Start Server
`./golang_boilerplate start`

## Test API
`curl http://localhost:3000/ping`

Other API details can be found in `api_blueprints` directory

#Tasks Glossary
`make setup` - Sets up the project
`make test` - Recreates Test Database, Runs Migration against the Test DB, Runs Test Case
`make compile` - Builds the Codebase
`make test-coverage` - Runs test cases and computes coverage
`make db.migrate` - Runs Migration
`make copy-config` - Copies application.yml.sample to application.yml

Please refer the Makefile for the entire tasks list