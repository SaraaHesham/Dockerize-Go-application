# Dockerize-Go-Application

[![Docker Pulls](https://img.shields.io/docker/pulls/sarahhesham1/my-go-app)](https://hub.docker.com/repository/docker/sarahhesham1/my-go-app)

A lightweight Go REST API for managing users, supporting JSON-based storage and fully containerized with Docker.

## Project Overview

This project is a simple REST API built with **Go** that allows:

- Adding a new user (`POST /users`)
- Retrieving a user by ID (`GET /users?id=1`)
- Persistent storage in text files (`users_saved` folder)
- Easy containerization with Docker

## Project Structure

```
get_post_go_lang/
├── controllers/  # HTTP handlers
├── models/       # User model and JSON storage
├── main.go       # Entry point
├── go.mod
└── Dockerfile    # Docker build instructions
```

## Requirements

- Go 1.25+
- Docker (optional)


## Quick Start (One Command)

Run the API immediately from Docker Hub with persistent storage:

```bash
docker run -d -v ./users_saved:/app/users_saved -p 3030:3030 sarahhesham1/my-go-app:latest
```

* Pull the latest image from Docker Hub
* Mount your local `users_saved` folder for persistent storage
* Expose port 3030 on your host so you can access the API at `http://localhost:3030`

[Docker Hub link](https://hub.docker.com/repository/docker/sarahhesham1/my-go-app/tags/latest/sha256-ba5bbd40f59865d823de3a522fc0d326679d732854cbd76840321c750813c085)

## With Docker (Build Locally)

```bash
docker build -t get-post-go .
docker run -d -v ./users:/app/users -p 3030:3030 get-post-go
```

* `-d` → run container in detached mode
* `-v ./users:/app/users` → mount local folder for persistent storage
* `-p 3030:3030` → map container port to host

## Running Locally

With Go

```bash
go run main.go
```

## API Examples

### POST /users

```bash
curl -X POST http://localhost:3030/users \
-H "Content-Type: application/json" \
-d '{"Name":"Sara","Age":22,"Address":{"Street":"Elnady","City":"Tanta","Country":"Egypt"}}'
```

### GET /users

```bash
curl http://localhost:3030/users?id=1
```

## Notes

* Ensure the host `users` folder exists to persist data.
* Default API port: `3030`
* User data is stored as `.txt` files in JSON format.
* Fully Docker-compatible for easy deployment.
