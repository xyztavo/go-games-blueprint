# Project Go Games Blueprint
Project made using GO and [go-blueprint](https://github.com/Melkeydev/go-blueprint) from [Melkeydev](https://github.com/Melkeydev) that uses Fiber and Postgres to CRUD some games in JSON.

## Getting Started
You can check the routes at [./internal/server/routes.go](./internal/server/routes.go)
You can start up a dev server using [make](#makefile) commands

## MakeFile
### These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

(You can install make on windows using chocolatey)


run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```


## TODO
- [ ] add validation to req body.
- [ ] write tests for everything (pain).