FROM golang:1.8 as build

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o app
EXPOSE 8000
CMD ["./app"]

# FROM ubuntu:18.04
# WORKDIR /app
# COPY --from=build /go/src/app/app .
# EXPOSE 8000
# CMD ["./app"]

