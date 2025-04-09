# 📡 Simple TCP Server in Go

This project is a **TCP server** built in **Go**, designed to handle multiple clients **concurrently** and respond with a simple **HTTP 200 OK** and **"HELLO, WORLD"** message.  
Clients like `curl`, `telnet`, or `netcat` can connect and communicate easily.

---

## 🚀 Features

- Handles multiple client connections using **goroutines** (multithreading).
- Sends an **HTTP-style** response to every incoming connection.
- Simple and lightweight server without any external frameworks.
- Created as a **Go module** for easy management.

---

## 🛠️ How to Run

1. **Clone the repository**:

```bash
git clone https://github.com/RakshitNotFound/TCP-SERVER-GO.git
cd TCP-SERVER-GO
```

2. **Run the server**:

```bash
go run main.go
```

3. **Test the server**:

You can connect using different tools:

- Using **curl**:

```bash
curl http://localhost:1729
```

Expected Output:

```
HTTP/1.1 200 OK

HELLO, WORLD
```

- Using **telnet**:

```bash
telnet localhost 1729
```

- Using **netcat (nc)**:

```bash
nc localhost 1729
```

Type anything and you'll receive:

```
HTTP/1.1 200 OK

HELLO, WORLD
```

---

## 📄 Project Structure

- `main.go`  
  - Listens on TCP port **1729**.
  - Waits for client connections.
  - Handles each client in a **new goroutine** to ensure concurrency.
  - Reads incoming data, processes it, and sends back an HTTP-like response.

---


## 📜 License

This project is licensed under the [MIT License](LICENSE).

---

# 🌟 Example Server Logs

```bash
Server is running on port 1729...
Waiting for client to connect
Client connected
Processing the request
```
