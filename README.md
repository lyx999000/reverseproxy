# About
Custom reverse proxy using Golang and subject to the limitations as following:
1. The reverse proxy will be used to mirror content of https://history.nasa.gov/SP-4206/contents.htm onto its
   port 8080.
2. Proxy must not let content of Chapter 8 through.

# Reqirements
- Golang
Or
- Docker

# Usage

## Compiling the application locally

```
//To build
$ go mod tidy
$ go build -o reverseproxy .
//To run
$ ./reverseproxy

```
## Docker build using docker
```
//To build
$ docker build . -t reverseproxy:latest
// To run
$ docker run -p 8080:8080 reverseproxy:latest
```
