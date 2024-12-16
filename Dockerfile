# 공식 Go 이미지를 Docker Hub에서 가져옴
FROM golang:1.23 as builder

# 컨테이너에서 /app 디렉터리를 컨테이너 내부의 작업 디렉터리로 설정
# 이후 모든 RUN, CMD, COPY는 이 디렉터리를 기준으로 실행
WORKDIR /app

# /app 디렉터리에 두 파일이 복사된다. 
COPY ./go.mod ./go.sum ./

# go.mod와 go.sum 파일을 기반으로 Go 모듈을 다운로드
RUN go mod download

# Copy source code
COPY . .

# 실행 가능한 바이너리 파일 api-server 생성
RUN go build -o api-server .

# 런타임 이미지 ???????????????????
FROM debian:buster-slim

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/api-server .

# 포트
EXPOSE 8080

# 컨테이너 실행되는 위치
CMD ["./api-server"]
