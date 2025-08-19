# Event Booking REST API

This is a RESTful API for an event booking system built with Go. It allows users to manage events, register for them, and provides secure, role-based access for administrators.

## 🚀 Features

* **User Authentication**: Secure user registration and login with hashed passwords and JSON Web Tokens (JWT) for authentication.
* **Role-Based Access Control (RBAC)**:
    * **User Role**: Can create, update, delete their own events and register for any event.
    * **Admin Role**: Can manage all users and view all event registrations.
* **Event Management**: CRUD (Create, Read, Update, Delete) operations for events. Users can only modify/delete events they created.
* **Event Registration**: Users can register for and cancel their registration for events.
* **Rate Limiting**: Protects against abuse by limiting the number of requests a user can make to the API.
* **Centralized Error Handling**: Ensures consistent and clear error responses across all endpoints.
* **SQLite Database**: Uses SQLite for a lightweight, file-based database.

## Technologies Used

* **Go**: The core programming language.
* **Gin Gonic**: A high-performance web framework for Go.
* **SQLite**: A C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine.
* **`bcrypt`**: For secure password hashing.
* **`golang-jwt`**: For creating and verifying JWTs.
