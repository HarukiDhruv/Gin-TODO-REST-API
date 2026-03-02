# Gin TODO REST API 

This project is a simple **RESTful Todo API built using Go (Golang) and the Gin web framework**. It was created as my first backend project in Go to understand how web servers, APIs, and data handling work in a statically typed and high-performance language.

The API provides basic **CRUD (Create, Read, Update, Delete)** functionality for managing todo items. Each todo is represented using a Go struct and stored in memory using slices, simulating how a backend service interacts with structured data.

---

## Overview

The application uses **Gin**, a lightweight and fast HTTP web framework for Go, to handle routing, request parsing, and JSON responses. It exposes REST endpoints that allow clients such as Postman, web applications, or terminal tools to interact with the todo data through standard HTTP methods.

This project demonstrates how Go can be used to build backend services by combining:

- Go structs for defining data models  
- JSON encoding and decoding for communication  
- HTTP routing and handlers using Gin  
- Slice operations to manage dynamic collections of data  
- REST API principles and proper HTTP status codes  

---

## Core Concepts Demonstrated

**REST Architecture**  
The API follows REST principles, where resources (todos) are accessed and manipulated using HTTP methods like GET, POST, PUT, and DELETE.

**Structured Data Modeling**  
Each todo is defined as a Go struct, showing how strongly typed models are used in backend systems.

**JSON Communication**  
The server converts Go structs into JSON responses and parses incoming JSON requests into Go objects.

**Routing and Request Handling**  
Gin manages incoming HTTP requests and routes them to specific handler functions responsible for business logic.

**In-Memory Data Management**  
Todo items are stored in a slice, demonstrating how Go manages collections dynamically. In production systems, this would typically be replaced with a database.

**Backend Server Fundamentals**  
The project shows how a Go application can function as an HTTP server, process requests, and return structured responses.

---

## Purpose

The goal of this project was to build a foundational understanding of backend development using Go, including API design, request handling, and data manipulation. It serves as a starting point for more advanced backend systems involving databases, authentication, concurrency, and scalable architectures.

---

## Author

Dhruba Hazarika  
Computer Science Student exploring Go, backend systems, and developer tooling
