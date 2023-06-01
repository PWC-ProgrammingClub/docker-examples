# LAMP

LAMP stands for Linux-Apache-Mysql-PHP.

It's an old school stack.

It's a fun one for docker, because it can be kind of a pain to set up PHP on windows, so most students have probably not had many chances to write PHP code.

## .docker folder

This holds files that are copied into the containers at the time they are built.

### .docker/mysql

The Mysql folder holds data in the form of .sql files. When the MySQL docker image is built, these files are copied to the special folder in the image called `docker-entrypoint-initdb.d`. You can read about that [here](https://hub.docker.com/_/mysql/#:~:text=Initializing%20a%20fresh%20instance). When the container is started for the first time, these *.sql are executed in alphabetical order to realize the database.


### .docker/php

When the PHP-Apache container is created, it will serve whatever is located at `/var/www/`. The .docker/php folder is turned into a shared volume so the container servers what is located at .docker/php.

