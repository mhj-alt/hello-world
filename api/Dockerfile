
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


### 멀티 스테이징

## docker file 여러가지 베이스 이미지

FROM alpine:latest

RUN apk --no-cache add ca-certificates


WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]