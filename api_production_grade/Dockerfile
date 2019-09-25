FROM golang:latest as builder

LABEL maintainer="Samuel Porras <samuelporras@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download

RUN cd cmd/ && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/main .

EXPOSE 5000

CMD ["./main"]