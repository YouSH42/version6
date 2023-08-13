#!/bin/bash

# MongoDB 서버 실행
mongod --config /etc/mongod.conf &

# MongoDB 서버가 완전히 초기화될 때까지 대기
sleep 5

# 애플리케이션 빌드
go build Node/main.go

# 애플리케이션 실행
./main 
