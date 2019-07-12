# Contributing

You can work on the ATC without a full Concourse deployment. This is useful for testing changes to the web UI. The only limitation is that any builds you run will fail with the `no workers` error. To test your local changes with a full deployment, follow the instructions in the [concourse/concourse CONTRIBUTING.md](https://github.com/concourse/concourse/blob/master/CONTRIBUTING.md) instead.

## Install development tools
Concourse is built with Go and Elm. You also need Node and few modules. Assuming you're using a mac:

- Install Elm 0.18 from https://guide.elm-lang.org/install.html
- Install homebrew from http://brew.sh/

Then use homebrew to install the following:

```
brew install go node postgres
```

Finally use Node to install the javascript tools:

```
npm install -g elm uglify-js less less-plugin-clean-css
```

## Setting up the database

You need a running postgres database named `atc`. The ATC itself takes care of creating and upgrading the schema, so you just need to create an empty database. If it's the first time you've installed postgres you need to run `initdb`

```
initdb /usr/local/var/postgres -E utf8
```

After that you can start the server and create the empty `atc` database:

```
postgres -D /usr/local/var/postgres/
createdb atc
```

## Building and running the code

```
go run cmd/atc/*.go run --add-local-user test:test --main-team-local-user test
```

Concourse should be live at http://127.0.0.1:8080
