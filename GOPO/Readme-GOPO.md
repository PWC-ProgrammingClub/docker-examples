# GoPo

Maybe I made this acronymn up, but it stands for Go-Postgres.

It's a newer school stack. It uses a Go HTTP Server with Gin and a Postgres database.

Go can be a fun language to write in, but it doesn't like being a subdirectory. You'll have a better time opening that folder as its own VSCode workspace.

## .docker folder

This holds files that are copied into the containers at the time they are built.

### .docker/go

This holds Go source code. It's built by the docker container before it's served - no need to install Go on your machine.

### .docker/postgres

The Mysql folder holds data in the form of .sql files. When the Postgres docker image is built, these files are copied to the special folder in the image called `docker-entrypoint-initdb.d`. You can read about that [here](https://hub.docker.com/_/postgres/). When the container is started for the first time, these *.sql are executed in alphabetical order to realize the database.


### Further reading

[Postgres SQL Syntax](https://www.postgresql.org/docs/current/sql-syntax.html)
[HTML Rendering with Gin for Go](https://gin-gonic.com/docs/examples/html-rendering/)
