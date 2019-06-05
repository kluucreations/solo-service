# SoLo Service

SoLo Service is an HTTP API server written in Go.

This service depends on a database for persistence storage. For developmental purposes, I'm using sqlite3.


## Installation

This service is tested on Ubuntu 19. It can also be built on OSX but it has not been confirmed.
You simply need to be able to run ```make``` , ```docker``` and ```git``` and you can start the server locally:


Clone the repo:
```bash
git clone https://github.com/kluucreations/solo-service.git && cd solo-service
```
Build docker container
```bash
make build
```
Run docker image and expose port
```bash
docker run -p 8080:8080 solo-service
```
You should see the following message on your screen:
```bash
2019/06/05 03:54:00 Shutting down... 
2019/06/05 03:54:00 Stopped serving solo service at http://[::]:8080
```
This means the server is now up and running locally and exposed through 8080 port.

You can visit the docs page at http://localhost:8080/docs
## Usage

All clients (mobile and web) can depend on the OpenAPI specs declared in this service. All possible requests and responses have been explicitly declared in the document including HTTP Status Response codes and optional query parameters.

There is currently only two endpoints declared: 
1) /v1/loan/{loan_id}/payments
2) /v1/loans

Endpoint /v1/loans has optional query parameters to specify pagination (page, perPage), ascending/descending order of results and the value to sort by (amount, end_at and score)

