# 🔐 Auth Service

A complete authentication service built with Go, featuring JWT tokens, password hashing, and session management.

## 🚀 Quick Start

```bash
# Clone and run
git clone https://github.com/luongtien872003/auth-service.git
cd auth-service
docker compose up --build -d

# Test health endpoint
curl http://localhost:8080/health
```

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go 1.24 |
| Framework | Gin |
| Database | PostgreSQL 15 |
| Cache | Redis 7 |
| Container | Docker Compose |

## 📁 Project Structure

```
auth-service/
├── cmd/server/main.go    # Entry point
├── docker-compose.yml    # Container orchestration
├── Dockerfile            # Build instructions
├── go.mod               # Dependencies
└── README.md
```

## 🔌 API Endpoints

| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| GET | `/health` | Health check | ✅ Done |
| POST | `/register` | User registration | 🔄 Coming |
| POST | `/login` | User login | 🔄 Coming |
| POST | `/refresh` | Refresh token | 🔄 Coming |
| POST | `/logout` | User logout | 🔄 Coming |

## 📝 Development Progress

- [x] Day 1: Docker setup + Health endpoint
- [ ] Day 2: Database + User model
- [ ] Day 3: Register API
- [ ] Day 4: Login + JWT
- [ ] Day 5: Refresh token
- [ ] Day 6: Security
- [ ] Day 7: Testing + Docs

## 📄 License

MIT
