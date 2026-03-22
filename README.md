# Net-Cat Chat in Go

A Go implementation of a **NetCat-style TCP chat** with multiple clients. Supports group chat, message history, timestamps, and usernames.

## Project Structure

```
net-cat/
├─ dependencies/
│  ├─ Additional_tools.go
│  ├─ Broadcast.go
│  ├─ ClientManager.go
│  ├─ HandleConnection.go
│  └─ HistoryHandle.go
├─ main.go
├─ go.mod
├─ history.txt
├─ logo.txt
└─ README.md
```
## Features

* TCP server with multiple clients (max 10)
* Clients provide a **unique username**
* Messages include **timestamp and username**
* Broadcasts messages to all clients
* Sends **message history** to new clients
* Notifies when clients **join/leave**
* Ignores empty messages
* Default port: `8989`

## Usage

### Start Server

go run . 8989

After starting, the server listens on the specified port and accepts client connections.

## How It Works

* Server listens on a TCP port and accepts connections
* Each client connection runs in a **goroutine**
* **Broadcast.go** sends messages to all clients
* **HistoryHandle.go** stores previous messages for new clients
* **Mutexes** protect shared resources (client list and history)
* `net.Conn` handles reading and writing data

## Learning Outcomes

* Go concurrency (goroutines, mutexes, channels)
* TCP socket programming
* Real-time chat server architecture
* Working with `net.Conn`

## Team

* **Your Name:** bguitoni
* **Team Member:** aelyoussef

## License

This project is maintained by the team members above. All rights reserved by the team.
