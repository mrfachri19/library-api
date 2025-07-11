# 📚 Digital Library Analytics Dashboard (Backend)

A RESTful backend service for managing a digital library system, including books, lending records, and analytics features for librarians. Built with **Go Fiber**, **PostgreSQL**, and implements **Clean Architecture**.

---

## 🚀 Features

- ✅ Book CRUD (title, author, ISBN, quantity, category)
- ✅ Lending management (borrow/return)
- ✅ Analytics endpoint (most borrowed books, total lendings, etc.)
- ✅ JWT-ready structure (can be extended)
- ✅ Clean architecture pattern
- ✅ PostgreSQL with foreign key relations

---

## 📦 Tech Stack

- **Language**: Go
- **Web Framework**: [Fiber](https://gofiber.io)
- **Database**: PostgreSQL
- **DB Driver**: pgx (PostgreSQL driver)
- **Architecture**: Clean Architecture (inspired by [go-clean-arch](https://github.com/bxcodec/go-clean-arch))

---

## 📁 Project Structure

internal/
├── config/ # DB config, .env loader, migration
├── entity/ # Structs for books, lendings
├── repository/ # Database layer
├── usecase/ # Business logic layer
├── handler/ # HTTP handlers
main.go # Application entry point


---

## 🛠️ Setup Instructions

### 1. Clone this repo
```bash
git clone https://github.com/yourusername/digital-library-backend.git
cd digital-library-backend
