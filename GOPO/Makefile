
# ------------------------------------------------------------
# 		Environment variables
# ------------------------------------------------------------
# the name of the postgres image
pg_image=cust-pg-img
# the name of the postgres docker container
pg_container=pg-container
# the name of the go docker image
go_image=cust-go-image
# the name of the go docker container
go_container=go-container

# the name of the database
db=pwpc
# the port the database will listen on
db_port=3306
# the db username
db_username=root
# the db password
db_password=i-am-a-bad-password
# the port php will listen on
go_port=80
# name of the docker network.
network=gopo-net

# ------------------------------------------------------------
# 		Build All - Combination Commands
# ------------------------------------------------------------

# Build both images and contaienrs, then network them
all: db go network
# Make the postgres image, then start the container.
db: db.image db.run
# Make the go image, then start the container.
go: go.image go.run 
# Make the network, then connect the containers
network: network.create connect.go connect.db
# Remove all
remove: db.rm go.rm network.rm
# Remove all then rebuild
rebuild: remove all
# Rebuild only the go container and re-connect it to the network
rb.go: go.rb connect.go
	docker logs $(go_container) 30s


# ------------------------------------------------------------
# 		Postgres Database Image and Container Build Commands
# ------------------------------------------------------------

# 'database build image'. Builds the pg docker image from Dockerfile.mysql, setting the shared environment variables according to this makefile.
db.image:
	docker build -t $(pg_image) \
	--build-arg username=$(db_username) \
	--build-arg password=$(db_password) \
	--build-arg db=$(db) \
	-f ./Dockerfile.postgres .

# 'database run' run the container
db.run:
	docker run -p $(db_port):$(db_port) --name $(pg_container) -d $(pg_image)

# 'database remove' stop and remove the container
db.rm:
	docker rm $(pg_container) --force

# 'database rebuild' - rebuild the image and rerun the container
db.rb: db.rm db.image db.run

# 'database remove, rebuild and copy database' - stop the container, copy database init files into container. Useful when having trouble getting initial db scripts to work. Used in before db.rbd
db.rmc:
	docker stop $(pg_container)
	docker cp ./.docker/postgres/. $(pg_container):/docker-entrypoint-initdb.d
	docker start $(pg_container)
	docker logs $(pg_container)

# 'database psql cli' -  access the postgress interactive command line interface
db.psql:
	docker exec -it $(pg_container) psql -U $(db_username) $(db)

# ------------------------------------------------------------
#		Go Image and Container Build Commands
# ------------------------------------------------------------

# 'build go image'. Builds to go docker image from Dockerfile.go, setting the shared environment variables according to this makefile.
go.image:
	docker build -t $(go_image) \
	--build-arg username=$(db_username) \
	--build-arg password=$(db_password) \
	--build-arg host=$(pg_container) \
	--build-arg db=$(db) \
	--build-arg httpport=$(go_port) \
	-f ./Dockerfile.go .

# 'run go container': Runs the container from the go image we created.
go.run:
	docker run -p $(go_port):$(go_port) --name $(go_container) -d $(go_image)

# 'remove go container': Removes the go container.
go.rm:
	docker rm $(go_container) --force

# 'rebuild go container'. Rebuilds the image, then the container
go.rb: go.rm go.image go.run

# 'run go bash': Runs bash
go.bash:
	docker exec -it -w /usr/src $(go_container) /bin/bash

# ------------------------------------------------------------
#		Docker network build commands
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
	docker network connect $(network) $(pg_container)

# connect the php container to the docker network
connect.go:
	docker network connect $(network) $(go_container)

# disconnect the mysql container from the docker network
dc.pg:
	docker network disconnect $(network) $(pg_container)

# disconnect the php container from the docker network
dc.go:
	docker network connect $(network) $(go_container)

	




