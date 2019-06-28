# Schutzstreifen
A web application that displays a map with user-contributed hazards for cyclists.
Built with the [Buffalo framework](https://gobuffalo.io).

## Run locally
For development, having Go and Go dev tools installed on your machine is extremely helpful.

Instead of also installing Buffalo and PostgresSQL, you can simply run the app with Docker. Otherwise, please refer to the Buffalo documentation on how to run the framework natively.

Use the following make commands to start/stop the containers:
```
make start
make stop
```
When all containers are up, the application is available at http://localhost:22080/

### Display container log output
With the default ```start``` target, containers are started in detached mode. If you prefer running them in the foreground, use the ```make start-dev``` target instead.

## Installation
This will set up the database and generate CSS.
```
make install
```

## Run tests
```
make test
```

## Buffalo CLI
Use this shortcut to open a shell inside the dev container. You can then use the full Buffalo CLI, including Pop for setting up the database and running migrations:
```
make cli
```

## Force rebuilding containers
This will skip the Docker build cache.
```
make build-containers
```

## Compile CSS
Create CSS files from Sass sources:
```
make css
```

## Run database migrations
Execute all migrations (only works if the database already exists).
```
make migrate
```
