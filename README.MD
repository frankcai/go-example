# Performance Comparison of Monolith to Microservices using a prototype

## Table of Contents

- [About The Project](#about-the-project)
- [Running on Ubuntu](#running-on-ubuntu)
- [Running on Windows](#running-on-windows)
- [Development](#development)

## About The Project

This project is a simple implementation of a document management system (DMS) in two software architecture styles, monolith and microservices. The goal is to compare the performance of each style. To achieve a fair comparison, only unique architectural code will be different, code that offers functionality will be kept identical. The program will implement the following use case: A User calls on the program to index and query a collection of his digital documents. 

The program consists of two distinct services. The first one has the ability to ingest document files and uses OCR to extract text. This text is saved and persisted onto the hard drive using a database. In a second step another service sends queries onto the database in order to find documents that contain the queried text. 

Technologies used are apache tika for extracting text from documents and bleve index as the storage container.

The document files used as the input are pdf files from an open dataset from [Kaggle](https://www.kaggle.com/datasets/mcvishnu1/fake-w2-us-tax-form-dataset).  

The comparison is performed as a benchmark via a client program that simulates multiple runs of the main DMS process. The monolith and microservices versions are each run multiple times and the results are displayed in the log file.

## Running on Ubuntu

- On a barebone Ubuntu operating system run this command to install the required dependencies

  ```bash
  sudo apt-get install git-all docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
  ```

- Clone the code into a local directory

  ```bash
  git clone git@github.com:frankcai/go-example.git
  ```

- Run the benchmarks using docker

  ```bash
  docker-compose build && docker-compose up
  ```

## Running on Windows

- Install git and docker

- Open a command prompt to clone the code into a local directory

  ```bash
  git clone git@github.com:frankcai/go-example.git
  ```

- Run the benchmarks using docker

  ```bash
  docker-compose build && docker-compose up
  ```

## Development

In the parent directory call `make` after making changes. This will build all projects including updating the dockerfiles. If a protobuf declaration has changed, run `make proto` first. To build a subproject instead of everything, run make in the appropriate subfolder.

### Required

- go 1.18
- Protocol buffer compiler in version 3+
- Plugins for the protocol compiler:
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
- Update the `PATH` so that the `protoc` compiler can find the plugins
  `export PATH="$PATH:$(go env GOPATH)/bin"`

To run the application without docker, each invoked service must be started and run manually before starting the client.

### Useful docker commands

Remove and prune old images `docker image prune -a`
Remove unused containers, networks and dangling images `docker system prune`