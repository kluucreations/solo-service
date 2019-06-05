# SoLo Service

SoLo Service is an HTTP API server written in Go.

This service depends on a database for persistence storage. For developmental purposes, I'm using sqlite3.

# Assignment
You can view the SQL Queries here
1) https://github.com/kluucreations/solo-service/blob/master/internal/datastore/loan.go#L22
2) https://github.com/kluucreations/solo-service/blob/master/internal/datastore/payment.go

Explanation: 

I created a join table (lender_loans) to allow RDS to enforce data integrity. By creating a unique constraint on lender_loan.loan_id we can enforce that a loan can only ever be assigned to one lender.


# Installation

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
# Usage

All clients (mobile and web) can depend on the OpenAPI specs declared in this service. All possible requests and responses have been explicitly declared in the document including HTTP Status Response codes and optional query parameters.

There is currently only two endpoints declared: 
1) /v1/loan/{loan_id}/payments
2) /v1/loans

Endpoint /v1/loans has optional query parameters to specify pagination (page, perPage), ascending/descending order of results and the value to sort by (amount, end_at and score)

# Train of Thought
## Go
The first thing I wanted to have defined was the OpenAPI Specifications.
Once I got the json specs created, I was able to generate server code to re-use models defined.
You can view the specifications here:
https://github.com/kluucreations/solo-service/blob/master/swagger/api.json

http://localhost:8080/docs is rendering the same exact document but in a user friendly way.

I opted to go with Go for the following reasons:
1) Go is currently my most fluent programming language
2) Compiled langauge allows me to enforce response payload because values and objects are checked at compile time
3) The final container size is extremely small allowing this server to be quickly deployed (less download/upload time), and smaller container registry. The final size of this server is 20.2MB
```bash
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
solo-service        latest              9d2e725cc058        About an hour ago   20.2MB
```

## Database

I lacked a full-featured relational database when playing with this code.
Instead of spinning up a PostgreSQL/MySQL instance, I opted to use SQLite3.
You can spin up the local SQLite3 server by installing ```sqlite3``` and running the following:
```bash
sqlite3 internal/datastore/solo.db
```
I have also extracted the schema in CSV format so it can be easily inspected.
https://github.com/kluucreations/solo-service/tree/master/internal/schema

SQLite3 lacks Foreign key constraints and formal Time representation.
If I had more time, I would fully define all the relationship. I believe the FK columns are pretty self-explained as they all end in *_id .

## Currency
I wanted to allow customers to specify what kind of currency they wanted to accept.
The API specification is built out to handle iso_currency_code to be future proof. This feature is not fully implemented yet.
Since I wanted to accomodate for different currency types in the future, it did not make sense to represent currency amount as an integer (as cents). Conversion between currency codes will result in precision lost, so I opted to use float64 to capture as much precision as possible.

