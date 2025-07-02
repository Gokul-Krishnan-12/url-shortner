# 🔗 Go URL Shortener

A simple and fast URL shortening service built in Go. Supports:

- 🔒 Custom short codes
- ⏰ Expiration times
- 📈 Stats tracking
- 💾 Local persistence via JSON file
- 🌐 Frontend form to create short URLs

---

## 📁 Project Structure

```
go-url-shortener/
├── main.go            # Starts the server, registers routes
├── handlers.go        # Business logic for shorten, redirect, stats
├── storage.go         # JSON file storage (save/load)
├── utils/
│   └── utils.go       # Helper functions for validation, random codes
├── static/
│   └── index.html     # Frontend UI
├── urls.json          # Saved URL data (auto-created)
```

---

## 🚀 Getting Started

### 1. ✅ Clone and Setup

```bash
git clone https://github.com/YOUR_USERNAME/go-url-shortener.git
cd go-url-shortener
go mod init go-url-shortener
go mod tidy
```

### 2. ▶️ Run the server

```bash
go run main.go handlers.go storage.go
```

Server will start at: [http://localhost:9000](http://localhost:9000)

---

## 🌐 Usage

### 🔗 Shorten a URL (API)

**POST** `/shorten`

```json
{
  "url": "https://example.com",
  "custom_code": "mycode",      // optional
  "expiry_mins": 60              // optional (default: 24h)
}
```

**Response:**
```json
{
  "short_url": "http://localhost:9000/mycode"
}
```

---

### 📥 Redirect

Visit: `http://localhost:9000/mycode`

---

### 📊 View Stats

**GET** `/stats/mycode`

```json
{
  "original_url": "https://example.com",
  "hit_count": 5,
  "expires_at": "2025-07-02T15:00:00Z",
  "expired": false
}
```

---

## 💻 Frontend UI

Visit: [http://localhost:9000/ui/](http://localhost:9000/ui/)

Features:
- Input original URL
- Custom short code
- Expiry duration (in minutes)
- Result with clickable link and copy button

---

## 🔐 File Permissions
The file `urls.json` is saved with `0644` permissions.

- Owner can read/write
- Group and Others can read

---

## 🛡️ Security Notes
- `0644` is safe for local file use
- For real-world apps, use `600` for secrets
- Don't use `chmod 777` in production

---

## 📌 To-Do / Future Ideas

- [ ] Add frontend stats checker
- [ ] Add database backend (PostgreSQL or MySQL)
- [ ] Add rate-limiting or user auth

---

## 👨‍💻 Author

Built with ❤️ by [Goku](https://github.com/Gokul-Krishnan-12)

---

