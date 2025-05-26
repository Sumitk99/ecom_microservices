FROM golang:1.23.3-alpine3.19 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Sumitk99/ecom_microservices/gateway
COPY .env ./
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO111MODULE=on go build -o /build/gateway ./cmd/gateway

FROM alpine:3.19
ENV PORT=8080
ENV PAYMENT_URL=localhost:8085
WORKDIR /usr/bin
COPY --from=build /build/gateway .
COPY .env .
EXPOSE 8080
CMD ["./gateway"]