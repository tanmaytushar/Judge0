# 🧠 Online Judge Platform

A scalable and secure online judge system, inspired by platforms like Codeforces and LeetCode. It allows users to submit code solutions, which are compiled and run inside isolated Docker containers. Outputs are compared against test cases to determine verdicts such as `Accepted`, `Wrong Answer`, or `Time Limit Exceeded`.

---

## 🚀 Features

- 🧾 REST API to submit and query code submissions  
- 🔒 Secure Docker containers for sandboxed code execution  
- 🧠 Custom judging logic (compile, run, compare I/O)  
- 🔁 Redis-based job queue for asynchronous processing  
- ⚡ Parallel execution of submissions  
- 🧪 Test case comparison and verdict generation  
- 📦 Dockerized services for portability and isolation  

---

## 🧱 Tech Stack

| Component         | Technology         |
|------------------|--------------------|
| Backend API       | Go (`gin`)         |
| Task Queue        | Redis              |
| Containerization  | Docker             |
| Job Worker        | Go goroutines      |
| Code Execution    | Isolated Docker containers |

---

## 📁 Project Structure

online-judge/
├── cmd/ # Entry points
│ └── judge/ # Main API server
├── internal/ # Core application logic
│ ├── api/ # HTTP handlers
│ ├── core/ # Judging logic
│ ├── docker/ # Docker container runner
│ ├── redis/ # Redis client & queue
│ ├── models/ # Structs for submissions, results
│ └── utils/ # Helpers and shared logic
├── jobs/ # Background worker service
├── submissions/ # Temp code + I/O per submission
├── testcases/ # Test case files (input/output)
├── Dockerfile # Docker setup for backend
├── docker-compose.yml # Multi-service setup
├── go.mod / go.sum # Go modules
└── README.md # Project overview



## 🔧 Setup Instructions


