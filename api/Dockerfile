
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# compile -> 실행가능한 바이너리 파일 생성
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


### 멀티 스테이징이용
## 베이스 이미지에서 실행가능한 바이너리 파일 복사
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]