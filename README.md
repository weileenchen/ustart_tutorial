# ustart_tutorial

Introduction to the organizational and operational structure of UStart back end architecture and git version control

## Example Case

This is a small backend implementation for a car rental service. It's split up into 3 microservices: car , customer, rental record. Cars and customers need to be searchable so they will be backed by ElasticSearch. Rental record is mostly relational so it will be backed by SQL (namely prostgres). All microservervices will have a dedicated REST or RPC api so they can be individually scaled as necessary.

## Task

Using the already completed `customer` microsverice and the framework defined in your `.proto` files in the `*pb` subfolders as a guideline, complete both the `car` and `record` subservices. You may choose one or both of either REST or RPC for the API's but you have to use each API at least once. That is to say if you use only REST for `car`, you must use at least RPC with `record"`. Each subservice needs to beveloped in it's own git branch until completion before submitting a pull request and merging back with origin.

## Dev-Environment Requirements

In order to get everything functional you must have the following installed and/or runnable:

- golang

- protobuf

- protoc

- tree

- PostgreSQL (it is recomended to also have pgAdmin installed to make monitoring data storage easier)

- ElasticSearch (it is recomended to also have Kibana installed to make monitoring data storage easier)

To test if you have sucessfully installed protobuf, protoc, and tree you should be able to run `gen-proto.sh` to without problem. By running `gen-proto.sh` you generate Go files based on the framework laid out in in your `.proto` files.

## About the Repo

DO NOT git clone this repository or any forks of this repo, always use

```bash
go get github.com/ACCOUNT/REPO
```

This way when both Go and your IDE will be able to easily find any local packages imports with the added benefit of also retrieving any external package dependancies.

Also note how `*.pb.go` and `*/config.json` are included in the `.gitignore`. Pushing anything that is relevant to only a local instance like versions of a framework and configurations should be avoided. If you wish to update the go version of the framework you are working after a modification has been made to the `.proto` simply rerun `gen-proto.sh`
