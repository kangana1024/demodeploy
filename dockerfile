FROM golang:1.16
WORKDIR /go/src/demodeploy
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /go/src/demodeploy/main ./
EXPOSE 3000
CMD ["./main"] 