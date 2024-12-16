# 공식 Go 이미지를 Docker Hub에서 가져옴
FROM golang:1.23

# 컨테이너에서 /app 디렉터리를 컨테이너 내부의 작업 디렉터리로 설정
# 이후 모든 RUN, CMD, COPY는 이 디렉터리를 기준으로 실행
WORKDIR /app

# 로컬(빌드 명령(.) 실행 위치) 기준
# go.mod와 go.sum 파일을 복사
# /app 디렉터리에 두 파일이 복사된다. 
COPY ./go.mod ./go.sum ./

# 디버깅: 컨테이너 내부 파일 구조 확인
RUN ls -R /app

# go.mod와 go.sum 파일을 기반으로 Go 모듈을 다운로드
RUN go mod download

# cmd 디렉터리를 통째로 복사
COPY ./cmd ./cmd

# 컨테이너 내 /app/cmd/main.go를 빌드하여 실행 가능한 바이너리 파일 api-server 생성
RUN go build -o api-server ./cmd/main.go

# 서버가 수신 대기할 포트 번호 설정
EXPOSE 8080

# 컨테이너에서 실행파일을 실행
CMD ["./api-server"]