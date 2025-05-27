# Online Judge

A scalable and secure online judge, inspired by platforms like Codeforces and LeetCode. It allows users to submit code solutions, which are compiled and run inside isolated Docker containers. Outputs are compared against test cases to determine verdicts such as `Accepted`, `Wrong Answer`, or `Time Limit Exceeded`.

---

## 🚀 Features

- REST API to submit and query code submissions  
- Secure Docker containers for sandboxed code execution  
- Custom judging logic (compile, run, compare I/O)  
- Redis-based job queue for asynchronous processing  
- Parallel execution of submissions  
- Test case comparison and verdict generation  
- Dockerized services for portability and isolation  

---

## Tech Stack

| Component         | Technology         |
|------------------|--------------------|
| Backend API       | Go (`gin`)         |
| Task Queue        | Redis              |
| Containerization  | Docker             |
| Job Worker        | Go goroutines      |
| Code Execution    | Isolated Docker containers |


---

## 🔧 Setup Instructions



