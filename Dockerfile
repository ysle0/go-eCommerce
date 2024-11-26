FROM golang:1.23.3-alpine
LABEL authors="ysl"

WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-eCommerce

EXPOSE 8080
CMD ["/go-eCommerce"]