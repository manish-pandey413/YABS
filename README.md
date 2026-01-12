# YABS: Yet Another Blog Site.

This is a project I built to learn backend devevolment in Go.

## Why I Built This
Most of my previous experience was on the surface level. With this project, I wanted to learn

* Idiomatic Go: Writing clean, maintainable code using standard libraries where possible.
* Database Management: Moving past simple queries to handle relationships and migrations.
* Security: Implementing proper password hashing and JWT-based authentication.
* Middleware: Learning how to handle logging, recovery, and auth checks in the request pipeline.

## Tech Stack
* Language: Go (Golang)
* Router: Echo
* Database: PostgreSQL
* Auth: JWT (JSON Web Tokens)
* Config: Environment-based configuration

## Key Features
* Full CRUD for Posts: Create, Read, Update, and Delete blog entries.
* User Accounts: Secure registration and login flow.
* JWT Middleware: Restricted routes that ensure only the author can edit their content.
* JSON API: Completely decoupled backend, ready to be plugged into any frontend.
* Clean Shutdown: The server handles interrupts gracefully to ensure no data is corrupted during a restart.
