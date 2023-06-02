
# ------------------------------------------------------------
# 		Environment variables
# ------------------------------------------------------------

# the name of the mysql docker image
mysql_image = cust-mysql-img
# the name of the mysql docker container
mysql_container = mysql-container
# the name of the php docker image
php_image = cust-php-image
# the name of the php docker container
php_container = php-container
# the name of the database
db = pwpc
# the port the database will listen on
db_port = 3306
# the db username
db_username=root
# the db password
db_password=i-am-a-bad-password
# the port php will listen on
php_port = 80
# name of the docker network.
network = lamp-net

# ------------------------------------------------------------
# 		Build All - Combination Commands
# ------------------------------------------------------------

# Make the mysql and start it, make php and start it, make the network and connect both.
all: db php network
	echo Containers running, head to http://localhost:80 to see the php
# Make the mysql image, then start the container.
db: db.image db.run
# Make the php image, then start the container.
php: php.image php.run 
# Make the network, then connect the containers.
network: network.create connect.php connect.db
# If you rebuilt containers individually, without removing network, use this command to reconnect them.
connect: connect.php connect.db
# Remove the containers, then remove the network.
remove: db.rm php.rm network.rm
# Remove the containers and the network, then rebuild them.
rebuild: remove all
# Only rebuild the php.
rebuild.php: php.rm php connect.php

# ------------------------------------------------------------
# 		MYSQL Database Image and Container Build Commands
# ------------------------------------------------------------


# 'database build image' Builds the mysql docker image from Dockerfile.Mysql, setting the shared environment variables according to this makefile.
db.image:
	docker build -t $(mysql_image) \
	--build-arg username=$(db_username) \
	--build-arg password=$(db_password) \
	--build-arg database=$(db) \
	--build-arg date=$(db_date) -f ./Dockerfile.mysql .

# 'database run' runs the mysql docker container from the image that was built.
db.run:
	docker run -p $(db_port):$(db_port) --name $(mysql_container) -d $(mysql_image)

# 'database remove' removes the mysql docker container.
db.rm:
	docker rm $(mysql_container) --force

# 'database rebuild' removes the mysql docker container, then rebuilds the image, then runs the container again with the new image.
db.rb: db.rm db.image db.run

# 'databash bash' Runs /bin/bash on the database container, allowing access to the database filesystem. Useful for debugging.
db.bash:
	docker exec -it $(mysql_container) /bin/bash

# 'database mysql' Runs mysql on the database container, allowing access to the interactive mysql repl. Useful for debugging. Some useful commands: "show tables;", "show databases;" "use betacusc_typ_livedb;" "select * from tablename;" etc.
db.mysql:
	docker exec -it ${mysql_container} mysql -u${db_username} -p${db_password}

# 'database image help' Prints help to the command line about the mysql image. You could > this into a file to make it more readable. 
db.image.help:
	docker run -it --rm ${mysql_image} --verbose --help

# ------------------------------------------------------------
# 		PHP Server Image and Container Build Commands
# ------------------------------------------------------------

# 'php build image' Builds the php server image from Dockerfile.php, setting the shared environment variables according to this makefile.
php.image:
	docker build -t $(php_image) \
	--build-arg db_container=$(mysql_container) \
	--build-arg db_user=$(db_username) \
	--build-arg db_password=$(db_password) \
	--build-arg db=$(db) \
	-f Dockerfile.php .

# 'php run' runs a php docker container from the image that was built. It also mounts .docker/php as a volume.
php.run:
	docker run -p $(php_port):80 --na6me $(php_container) -d $(php_image)

# 'php remove' removes the php docker container
php.rm:
	docker rm $(php_container) --force

# 'php rebuild' removes the php docker container, then rebuilds the image, then runs the container again with the new image.
php.rb: php.rm php.image php.run

# 'php shell' access the shell (bin/bash) on the php container for debug, checking installed modules, etc. Try 'php -m' to list installed modules, for instance.
php.bash:
	docker exec -it $(php_container) /bin/bash

# This copies the files in .docker/php to /var/www/ on the php container. It's faster to just copy the changes over than to re-run the container.
php.copy:
	docker cp ./.docker/php/. $(php_container):/var/www/html

# ------------------------------------------------------------
# 		Network Build Commands
# ------------------------------------------------------------
# The docker containers must be connected on a docker network to be able to see each other.

# creates a docker network to link the containers
network.create:
	docker network create $(network)

# removes the docker network
network.rm:
	docker network remove $(network) --force

# connect the mysql container to the docker network
connect.db:
	docker network connect $(network) $(mysql_container)

# connect the php container to the docker network
connect.php:
	docker network connect $(network) $(php_container)

# disconnect the mysql container from the docker network
dc.mysql:
	docker network disconnect $(network) $(mysql_container)

# disconnect the php container from the docker network
dc.php:
	docker network connect $(network) $(php_container)