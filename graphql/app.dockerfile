FROM golang:1.23.3-alpine3.19 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Sumitk99/ecom_microservices
COPY go.mod go.sum ./
COPY vendor vendor
COPY account account
COPY catalog catalog
COPY order order
COPY graphql graphql
RUN G0111MODULE=on go build -mod vendor -o /go/bin/app ./graphql

FROM alpine:3.19
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]