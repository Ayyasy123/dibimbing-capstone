# Capstone Project: Service Booking API

This project is a backend API for a service booking platform. It allows users to book services, make payments, leave reviews, and manage their profiles. The API is built using Go (Golang) with the Gin framework and uses a PostgreSQL database with GORM as the ORM.

## Table of Contents

1. [Features](#features)
2. [API Endpoints](#api-endpoints)
   - [User Endpoints](#user-endpoints)
   - [Service Endpoints](#service-endpoints)
   - [Booking Endpoints](#booking-endpoints)
   - [Payment Endpoints](#payment-endpoints)
   - [Review Endpoints](#review-endpoints)
3. [Middleware](#middleware)
4. [Entities](#entities)

---

## Features

- **User Management**: Register, login, update, and delete users. Users can register as technicians or admins.
- **Service Management**: Create, update, delete, and search for services.
- **Booking Management**: Book services, update booking status, and view booking history.
- **Payment Management**: Make payments, update payment status, and view payment reports.
- **Review Management**: Leave reviews for services and view review reports.
- **Pagination**: All `GET` endpoints support pagination using `limit` and `offset` query parameters.
- **Authentication & Authorization**: JWT-based authentication and role-based access control.

---

## API Endpoints

### User Endpoints

| Method | Endpoint                     | Description                                  | Authentication Required |
| ------ | ---------------------------- | -------------------------------------------- | ----------------------- |
| POST   | `/register`                  | Register a new user                          | No                      |
| POST   | `/login`                     | Login and get JWT token                      | No                      |
| POST   | `/register-admin`            | Register a new admin                         | No                      |
| GET    | `/users/:id`                 | Get user details by ID                       | Yes                     |
| GET    | `/users`                     | Get all users (with pagination)              | Yes                     |
| PUT    | `/users`                     | Update user details                          | Yes                     |
| DELETE | `/users/:id`                 | Delete a user                                | Yes                     |
| POST   | `/users/register-technician` | Register as a technician                     | Yes                     |
| PUT    | `/users/update-technician`   | Update technician details (technician/admin) | Yes                     |

---

### Service Endpoints

| Method | Endpoint                  | Description                                    | Authentication Required |
| ------ | ------------------------- | ---------------------------------------------- | ----------------------- |
| POST   | `/services`               | Create a new service (technician only)         | Yes (Technician)        |
| GET    | `/services/:id`           | Get service details by ID                      | Yes                     |
| PUT    | `/services`               | Update service details (technician only)       | Yes (Technician)        |
| DELETE | `/services/:id`           | Delete a service (technician only)             | Yes (Technician)        |
| GET    | `/services`               | Get all services (with pagination)             | Yes                     |
| GET    | `/services/user/:user_id` | Get services by user ID                        | Yes                     |
| GET    | `/services/search`        | Search services by query, min_price, max_price | Yes                     |

---

### Booking Endpoints

| Method | Endpoint                        | Description                                                      | Authentication Required |
| ------ | ------------------------------- | ---------------------------------------------------------------- | ----------------------- |
| GET    | `/bookings`                     | Get all bookings (with pagination)                               | Yes                     |
| GET    | `/bookings/:id`                 | Get booking details by ID                                        | Yes                     |
| POST   | `/bookings`                     | Create a new booking                                             | Yes                     |
| PUT    | `/bookings`                     | Update booking details                                           | Yes                     |
| DELETE | `/bookings/:id`                 | Delete a booking                                                 | Yes                     |
| GET    | `/bookings/user/:user_id`       | Get bookings by user ID                                          | Yes                     |
| GET    | `/bookings/service/:service_id` | Get bookings by service ID                                       | Yes                     |
| PUT    | `/bookings/:id/status`          | Update booking status                                            | Yes                     |
| GET    | `/bookings/available-dates`     | Get available dates for a service (with service_id, year, month) | Yes                     |
| GET    | `/bookings/reports`             | Get booking reports (with start_date, end_date)                  | Yes                     |

---

### Payment Endpoints

| Method | Endpoint               | Description                                                 | Authentication Required |
| ------ | ---------------------- | ----------------------------------------------------------- | ----------------------- |
| GET    | `/payments`            | Get all payments (with pagination)                          | Yes                     |
| GET    | `/payments/:id`        | Get payment details by ID                                   | Yes                     |
| POST   | `/payments`            | Create a new payment                                        | Yes                     |
| PUT    | `/payments`            | Update payment details                                      | Yes                     |
| DELETE | `/payments/:id`        | Delete a payment                                            | Yes                     |
| PUT    | `/payments/:id/status` | Update payment status                                       | Yes                     |
| GET    | `/payments/reports`    | Get payment reports (with start_date, end_date, service_id) | Yes                     |

---

### Review Endpoints

| Method | Endpoint           | Description                                                | Authentication Required |
| ------ | ------------------ | ---------------------------------------------------------- | ----------------------- |
| GET    | `/reviews`         | Get all reviews (with pagination)                          | Yes                     |
| GET    | `/reviews/:id`     | Get review details by ID                                   | Yes                     |
| POST   | `/reviews`         | Create a new review                                        | Yes                     |
| PUT    | `/reviews`         | Update review details                                      | Yes                     |
| DELETE | `/reviews/:id`     | Delete a review                                            | Yes                     |
| GET    | `/reviews/reports` | Get review reports (with start_date, end_date, service_id) | Yes                     |

---

## Middleware

### JWT Authentication (`auth.go`)

- **Purpose**: Validates JWT tokens in the `Authorization` header.
- **Behavior**:
  - Checks for the presence of the `Authorization` header.
  - Validates the token and extracts user claims (e.g., `user_id`, `role`).
  - Aborts the request if the token is invalid or expired.

### Role-Based Authorization (`role.go`)

- **Purpose**: Restricts access to endpoints based on user roles.
- **Behavior**:
  - Checks the user's role from the JWT claims.
  - Allows access only if the user's role matches one of the allowed roles.
  - Returns a `403 Forbidden` error if the user does not have permission.

---

## Entities

### User (`users.go`)

- Represents a user in the system.
- Fields: `ID`, `Name`, `Email`, `Password`, `Role`, `Address`, `Phone`, `Expertise`, `Availability`, `CreatedAt`, `UpdatedAt`.
- Relationships:
  - Has many `Services`.
  - Has many `Bookings`.

### Service (`services.go`)

- Represents a service offered by a technician.
- Fields: `ID`, `UserID`, `Name`, `Description`, `Cost`, `CreatedAt`, `UpdatedAt`.
- Relationships:
  - Belongs to `User`.
  - Has many `Bookings`.

### Booking (`bookings.go`)

- Represents a booking made by a user for a service.
- Fields: `ID`, `UserID`, `ServiceID`, `Date`, `Status`, `Description`, `CreatedAt`, `UpdatedAt`.
- Relationships:
  - Belongs to `User`.
  - Belongs to `Service`.

### Payment (`payments.go`)

- Represents a payment made for a booking.
- Fields: `ID`, `BookingID`, `Amount`, `Status`, `CreatedAt`, `UpdatedAt`.
- Relationships:
  - Belongs to `Booking`.

### Review (`reviews.go`)

- Represents a review left by a user for a service.
- Fields: `ID`, `BookingID`, `Rating`, `Comment`, `CreatedAt`, `UpdatedAt`.
- Relationships:
  - Belongs to `Booking`.

---
