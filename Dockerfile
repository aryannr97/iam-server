FROM golang:alpine

WORKDIR /go/src/iam-server
COPY . .

RUN CGO_ENABLED=0 GOOS=${GOOS} go build -mod vendor cmd/main.go

EXPOSE 8080
CMD ["./main"]