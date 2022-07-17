# streamtodb

## Pre-requisites

The provided **.env** file has pre-populate data to work with dockerized version.

## Run with docker and docker-compose
execute the following:

~~~~
docker-compose build
docker-compose up
~~~~

### run without docker
In case you don't have docker installed, you need to install the following:
#### 1. Go
#### 2. PostgresSQL

Run the server first:

`make service`


## misc

there is a bunch of tools in `Makefile`

~~~
make install
make lint
make test
make service
~~~
