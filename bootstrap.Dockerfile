FROM sang:myUbuntu-go-mongodb

# 작업 디렉토리 설정
WORKDIR /go/src/Node

# ping 추가
# RUN apt-get update && apt-get install -y iputils-ping

# 호스트의 모듈 파일을 컨테이너로 복사
COPY go.mod go.sum ./

# Go 모듈 다운로드
RUN go mod download

# 나머지 소스 코드 복사
COPY ./bootstrapNode .

# 나머지 소스 코드 복사
RUN go build bootstrap.go

# 스크립트 실행
CMD ["./bootstrap"]