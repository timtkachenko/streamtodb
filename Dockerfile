# syntax=docker/dockerfile:1
##
## BUILD
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /main ./cmd/

##
## RUN
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 9999

ENTRYPOINT ["/main"]