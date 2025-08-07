# 🧠 System Design

## 🔗 Project: Basic URL Shortener (No DB, No Caching)

A minimal implementation of a URL shortener using only in-memory storage (e.g., Go `map[string]string`) to demonstrate basic system behavior and routing, as part of learning core system design principles.

---

## 📖 What is System Design?

**System Design** is the process of defining the architecture, components, data flow, and interactions in a software system to meet specific functional and non-functional requirements. It's used to ensure that systems are **scalable**, **reliable**, **maintainable**, and **cost-efficient**.

---

## 🧩 HLD vs LLD

| Aspect                  | HLD (High-Level Design)                            | LLD (Low-Level Design)                              |
|------------------------|----------------------------------------------------|-----------------------------------------------------|
| Purpose                | Architectural blueprint                            | Detailed logic and code structure                   |
| Focus                  | Modules, services, data flow                       | Classes, methods, API contracts                     |
| Audience               | Architects, senior engineers                       | Developers, integrators                             |
| Example (URL Shortener)| Service takes long URL and returns short version  | How service maps the short key to the full URL      |

---

## 🏗️ Key System Design Concepts

- **Scalability**: Ability to handle growing traffic or data.
    - Vertical: More CPU/RAM
    - Horizontal: More machines
- **Availability**: System remains operational most of the time.
    - Measured in uptime (e.g., 99.99%)
- **Reliability**: System functions correctly and consistently.
    - No data loss or crashes
- **Maintainability**: Easy to fix bugs, add features, and refactor.
    - Clean, modular codebase

---

## 🧱 Monolithic vs Microservices

| Monolithic                          | Microservices                               |
|------------------------------------|---------------------------------------------|
| Single codebase                    | Multiple independent services               |
| Shared memory and DB               | Each service can have its own DB            |
| Easier to develop and deploy early | Easier to scale and maintain long-term      |
| Tight coupling                     | Loose coupling via APIs                     |

---

## 🛠️ Task: Basic URL Shortener

### Objective:
Build a working URL shortener with:
- `/shorten` – accepts a long URL and returns a short one
- `/redirect/{shortURL}` – redirects to the original URL

### Constraints:
- No database or external caching
- Use an in-memory map
- Focus on routing and logic
- No concurrency handling needed for Day 1

### Sample Flow:
1. `POST /shorten` with JSON `{ "url": "https://example.com" }`
2. Response: `{ "short_url": "http://localhost:8080/abc123" }`
3. `GET /abc123` → Redirects to `https://example.com`

---

## 🧾 Execution
1. go run main.go router.go handler.go utils.go
2. curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url":"https://example.com"}'
3. curl -i http://localhost:8080/abc123  # Replace with the returned short key