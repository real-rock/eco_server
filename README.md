# Economicus backend

> 이코노미쿠스 프로젝트에 대한 설명 (작성 예정) 

## Requirements

- `docker & docker-compose`
- `.env` `.env.mysql` `.env.mongo`
- 실행 방법

  ```shell
  git clone https://github.com/real-rock/eco_server.git

  docker-compose up -d --build
  ```

## Overview
총 2개의 서버로 나뉘며 각각은 gRPC 인터페이스로 데이터를 주고 받습니다. 
이 중 main 서버는 Golang으로 작성되어 전반적인 기능을 담당합니다.
quant 서버는 주가 데이터를 적재 및 가공하며 main 서버의 요청에 의해 계산된 결과를 반환합니다.
 

## Software Used
- Docker & docker-compose
- Golang
  - gin
  - gorm
- logrus
- jwt
- Swagger
- mysql DB
- mongo DB
- AWS S3

## More
