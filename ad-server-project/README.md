# 목차

1. [프로젝트명](#프로젝트명)
2. [프로젝트 설명](#프로젝트-설명)
3. [프로젝트 실행 방법](#프로젝트-실행-방법)

(과제관련 문서)
- [새로운 광고 송출 및 트래픽 대응 방법 - 과제3](#새로운-광고-송출-및-트래픽-대응-방법---과제3)
- [리워드 부정 수급 차단 방법 - 과제6](#리워드-부정-수급-차단-방법---과제6)

* * *

# 프로젝트명
 > ad-server-project

# 프로젝트 설명

### 개발
- 개발언어: Golang:1.18
- web 프레임워크: [Gin](https://github.com/gin-gonic/gin)
- docker-compose:3.8

### 데이터베이스
- MySQL:8.0.31
- 관계형 데이터베이스로 가장 많이 쓰이는 MySQL 사용했습니다.

### 설계
- Clean Architecture
- 계층이 분리되어 있어 domain 확장이 쉽고 간단한 서버 구현에 적합한 Clean Architecture로 설계했습니다.
```bash
ad-server-project
├── src
    └── adapter
    └── domain
      └── model
    └── repository
    └── usecase
    └── main.go
├── initdb.d

```
- model: 개체의 구조와 메서드 정의합니다.
- repository: 데이터베이스와 연결, 데이터 처리를 담당합니다.
- usecase: 데이터를 가공, 비지니스 로직을 처리합니다.
- adapter: usecase의 output을 가져와 표시합니다.


# 프로젝트 실행 방법
도커가 설치되어있어야 합니다.

[Docker](https://www.docker.com/get-started) 설치 & 로그인 (tested on v4.3.0)


프로젝트를 다운받아 프로젝트 폴더로 이동합니다.
```bash
git clone https://github.com/buzzvil-assignments/sooutt-naver.com.git

cd ad-server-project
```

서버를 실행합니다.
```bash
make up
```
- 해당 명령어를 실행 하면, 프로젝트 이미지를 build 하여 docker-compose 로 띄우게 됩니다.
- MySQL 데이터베이스를 완전하게 띄운 후에 서버를 실행하도록 했습니다.
- 데이터베이스를 docker-compose 로 구성하면서, 필요한 리소스 데이터를 initdb sql 파일을 이용하여 import 합니다.

http://localhost:8080/ 혹은 http://0.0.0.0:8080/ 접속했을 때, "Hello world"가 출력된다면 서버가 정상적으로 실행된 것입니다.

## API 명세서
https://documenter.getpostman.com/view/19629582/2s9YeBfEc8

# 테스트 코드 실행
## repository 테스트
`sqlmock`을 사용하여 데이터베이스 의존하지 않고 쿼리문을 확인하는 테스트를 구현하였습니다.
- `sqlmock`: 데이터베이스와 상호작용하는 코드를 테스트하는데 사용되는 가짜 SQL 드라이버
- 
repository 테스트를 실행하고, coverage를 출력하는 명령어는 아래와 같습니다.
```bash
make td-repository
``` 

## usecase 테스트
`mockery`, `testify`를 사용하여 DB와의 영속성 보다는 비지니스 로직을 검증합니다.
- 데이터를 Mocking하여 가상의 데이터를 생성하고 비지니스 로직이 잘 작동하는지 확인합니다.
- `mockery`: 명령어를 사용하여 특정 인터페이스(repository_interface)에 대한 mock을 자동으로 생성합니다.
- `testify`: Go언어 테스트 라이브러리로, suite 패키지를 사용하여 여러가지의 테스트 케이스를 그룹화하였습니다.

usecase 테스트를 실행하고, coverage를 출력하는 명령어는 아래와 같습니다.
```bash
make td-usecase
```


* * *

# 새로운 광고 송출 및 트래픽 대응 방법 - 과제3

### 새로운 광고 송출 정책을 추가하는 방법

- DB에서 광고 정보의 country, gender 조건에 맞는 광고를 먼저 리스트로 받아옵니다.
- 조건에 맞는 광고 중 송출되는 광고를 선택하는 메서드를 정책에 따라 구현합니다.
- 정책을 한 개의 함수 단위로 구현하기 때문에 정책 추가가 쉽습니다.
- 또한 적용하는 정책 메서드만 바꿔주면 되기 때문에 정책 변경 적용도 용이합니다.

### 트래픽의 증가, 데이터의 증가에 따른 성능 문제에 대응하기 위해 적용한 부분
- 이중 For문을 만들지 않습니다.
  - for문을 이중으로 사용하게 되면 시간복잡도가 증가하여 응답시간이 지연되기 때문입니다.
  - map, sort 등을 사용하였습니다.

# 리워드 부정 수급 차단 방법 - 과제6

### 적용한 방법
- 기본적으로 광고의 reward값을 조작하여 적립하지 못하도록, 해당 광고와 리워드 정보가 일치한지 확인해봅니다.

### 추가적으로 고안한 방법
- 중복된 계정, 혹은 기기에서의 리워드 요청을 차단합니다.
  - 같은 광고로 연속 두번 이상 리워드 적립을 요청하는 경우 db를 조회하여 에러 처리합니다.