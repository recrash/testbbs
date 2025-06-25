# TestBBS - 게시판 시스템

TestBBS는 Go 백엔드와 React 프론트엔드로 구성된 사용자 인증 기반 게시판 시스템입니다.

## 🏗️ 프로젝트 구조

```
testbbs/
├── cmd/server/          # Go 서버 메인 애플리케이션
├── handlers/            # HTTP 핸들러
├── internal/            # 내부 패키지
│   ├── auth/           # JWT 인증 관련
│   ├── db/             # 데이터베이스 연결 및 리포지토리
│   ├── middleware/     # 미들웨어
│   ├── models/         # 데이터 모델
│   └── util/           # 유틸리티 함수
├── migration/          # 데이터베이스 마이그레이션
├── web/               # React 프론트엔드
│   ├── src/
│   │   ├── api/       # API 클라이언트
│   │   ├── contexts/  # React Context
│   │   └── pages/     # 페이지 컴포넌트
│   └── public/        # 정적 파일
└── README.md
```

## 🚀 기술 스택

### 백엔드 (Go)
- **언어**: Go 1.23.6
- **웹 프레임워크**: net/http (표준 라이브러리)
- **데이터베이스**: PostgreSQL
- **인증**: JWT (JSON Web Token)
- **패스워드 해싱**: bcrypt
- **CORS**: github.com/rs/cors
- **환경변수**: github.com/joho/godotenv

### 프론트엔드 (React)
- **언어**: TypeScript
- **프레임워크**: React 19.0.0
- **빌드 도구**: Vite 6.2.0
- **UI 라이브러리**: Material-UI (MUI) 7.0.1
- **라우팅**: React Router DOM 7.5.0
- **HTTP 클라이언트**: Axios 1.8.4
- **스타일링**: Emotion 11.14.0

## 📋 주요 기능

### 사용자 인증
- ✅ 사용자 등록 (회원가입)
- ✅ 사용자 로그인
- ✅ JWT 토큰 기반 인증
- ✅ 리프레시 토큰 지원
- ✅ 사용자 로그아웃
- ✅ 사용자 프로필 조회

### 보안 기능
- ✅ 패스워드 해싱 (bcrypt)
- ✅ CORS 설정
- ✅ JWT 토큰 검증 미들웨어
- ✅ 리프레시 토큰 만료 관리

## 🛠️ 설치 및 실행

### 1. 환경 설정

#### 백엔드 환경변수 설정
프로젝트 루트에 `.env` 파일을 생성하세요:

```env
DATABASE_URL=postgres://bbsuser:qwerasdf@localhost:5432/testbbs?sslmode=disable
SERVER_PORT=:8081
JWT_SECRET=your_jwt_secret_key_here
```

#### 데이터베이스 설정
PostgreSQL을 설치하고 다음 명령을 실행하세요:

```sql
-- migration/init.sql 실행
CREATE USER bbsuser PASSWORD 'qwerasdf' SUPERUSER;
CREATE DATABASE testbbs;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE refresh_tokens (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL,
  token TEXT NOT NULL,
  expires_at TIMESTAMP NOT NULL
);
```

### 2. 백엔드 실행

```bash
# Go 모듈 의존성 설치
go mod tidy

# 서버 실행
go run cmd/server/main.go
```

서버는 `http://localhost:8081`에서 실행됩니다.

### 3. 프론트엔드 실행

```bash
# web 디렉토리로 이동
cd web

# Node.js 의존성 설치
npm install

# 개발 서버 실행
npm run dev
```

프론트엔드는 `http://localhost:5173`에서 실행됩니다.

## 🔌 API 엔드포인트

| 메서드 | 엔드포인트 | 설명 | 인증 필요 |
|--------|------------|------|-----------|
| POST | `/register` | 사용자 등록 | ❌ |
| POST | `/login` | 사용자 로그인 | ❌ |
| GET | `/profile` | 사용자 프로필 조회 | ✅ |
| POST | `/refresh` | 토큰 갱신 | ❌ |
| POST | `/logout` | 사용자 로그아웃 | ❌ |

### 요청/응답 예시

#### 사용자 등록
```http
POST /register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

#### 사용자 로그인
```http
POST /login
Content-Type: application/json

{
  "email": "test@example.com",
  "password": "password123"
}
```

## 🏛️ 아키텍처

### 백엔드 아키텍처
- **레이어드 아키텍처**: 핸들러 → 서비스 → 리포지토리 → 데이터베이스
- **미들웨어 패턴**: CORS, 인증, 로깅
- **의존성 주입**: 데이터베이스 연결을 핸들러에 주입

### 프론트엔드 아키텍처
- **컴포넌트 기반**: 재사용 가능한 UI 컴포넌트
- **Context API**: 전역 상태 관리 (인증 상태)
- **라우팅**: React Router를 통한 SPA 구현

## 🔒 보안 고려사항

- 패스워드는 bcrypt로 해싱되어 저장됩니다
- JWT 토큰은 서명되어 변조를 방지합니다
- 리프레시 토큰은 만료 시간이 설정되어 있습니다
- CORS 설정으로 허용된 도메인에서만 접근 가능합니다

## 🧪 테스트

### 백엔드 테스트
```bash
# Go 테스트 실행
go test ./...
```

### 프론트엔드 테스트
```bash
cd web
npm run lint
```

## 📝 개발 가이드라인

### 코드 품질
- 가독성을 우선시합니다
- 중복 코드를 제거하고 DRY 원칙을 따릅니다
- 단일 책임 원칙을 준수합니다
- 하드코딩된 값을 상수나 환경변수로 추출합니다

### 성능 최적화
- 데이터베이스 쿼리를 최적화합니다
- 지연 로딩과 페이지네이션을 활용합니다
- 무거운 작업은 비동기로 처리합니다
- 캐싱을 적절히 활용합니다

## 🤝 기여하기

1. 이 저장소를 포크합니다
2. 기능 브랜치를 생성합니다 (`git checkout -b feature/amazing-feature`)
3. 변경사항을 커밋합니다 (`git commit -m 'Add some amazing feature'`)
4. 브랜치에 푸시합니다 (`git push origin feature/amazing-feature`)
5. Pull Request를 생성합니다

## 📄 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다.

## 📞 문의

프로젝트에 대한 문의사항이 있으시면 이슈를 생성해 주세요.