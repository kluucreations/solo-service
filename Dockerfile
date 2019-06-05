# create multi stage docker build
FROM golang:1.12-alpine AS builder
# download dependencies
RUN apk add git make g++
ENV GO111MODULE=on
WORKDIR /app
# copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
# build go applicadtion
RUN go build -o solo-service

# final container
FROM alpine
# copy over built app and start
COPY --from=builder /app/solo-service /app
CMD ["./app"]