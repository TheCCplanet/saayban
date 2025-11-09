# 🌌 Saayban — Decentralized Communication & Publishing Platform

> **Saayban** is a decentralized, privacy-first communication and publishing system.  
> It lets people **chat, share, and express ideas freely** — without relying on centralized servers or intermediaries.  
> Each user is their own node: encrypted, autonomous, and untouchable.

---

## 🧭 The Idea

Sayban was born from a simple belief:  
**no one should own your data except you.**

Every user has their own **encrypted local database** — fully isolated, password-protected, and under their complete control.  
This database can be **locked**, **unlocked**, and managed independently, ensuring that your private world remains truly yours.

Sayban forms the foundation for a network where people can:

- 🗣️ Chat and communicate privately  
- 📰 Publish weblogs or journals anonymously  
- 🔗 Connect directly without any central authority  
- 🧱 Store and own all data locally, encrypted by their keys  

---

## 🧩 Architecture Overview

| Layer | Purpose |
|-------|----------|
| **API Layer (Go + Chi)** | RESTful endpoints for registration, lock/unlock, and future messaging. |
| **Database (SQLite + SQLCipher)** | Each user has their own encrypted database file. |
| **Service Layer** | Handles user registration, encryption, and auto-lock logic. |
| **DB Manager** | Manages active databases, concurrency, and time-based locks. |
| **Modular Design** | Ready for extensions like P2P messaging, weblogs, and distributed syncing. |

---

## ⚙️ Features

- 🔐 **Encrypted Databases** — Each user’s data is stored in their own SQLCipher-encrypted file.  
- 🧱 **User Isolation** — No shared data layer; every user is their own secure instance.  
- ⏱️ **Auto-Lock Timer** — Databases automatically lock after inactivity for added security.  
- ⚙️ **Clean Service Architecture** — Decoupled, testable, and easy to extend.  
- 🌍 **Configurable & Lightweight** — Runs anywhere, perfect for local-first and decentralized apps.

---

## 🚧 Upcoming Features

- 💬 Real-time peer-to-peer messaging  
- 🪶 Anonymous weblog publishing  
- 🔑 Cryptographic user identity system  
- 🌐 Federated peer routing  
- 📦 Offline delivery and message caching  

---

## 🧠 Tech Stack

| Component | Technology |
|------------|-------------|
| **Language** | Go |
| **Database** | SQLite + SQLCipher |
| **Routing** | Chi |
| **Architecture** | Modular service-layer pattern |
| **Security** | Password-based encryption, auto-lock timers |

---

## 🚀 Getting Started

### 1️⃣ Clone & Run

```bash
git clone https://github.com/TheCCplanet/saayban.git
cd saayban
go run ./cmd/server
