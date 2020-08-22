FROM golang:alpine AS build
WORKDIR $GOPATH/src/github.com/lyx999000/reverseproxy
COPY go.mod go.sum $GOPATH/src/github.com/lyx999000/reverseproxy/
RUN go mod tidy
COPY . .
RUN go build -o /reverseproxy

FROM alpine:latest
WORKDIR /reverseproxy
COPY --from=build /reverseproxy reverseproxy
EXPOSE 8080
ENTRYPOINT ["./reverseproxy"]
