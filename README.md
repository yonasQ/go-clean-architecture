# Clean Architecture Showcase in Go

## Overview

This project serves as a demonstration of implementing Clean Architecture principles in a Go application. While the core functionality revolves around a simple CRUD (Create, Read, Update, Delete) application for user management, the primary goal is to showcase how Clean Architecture can be effectively applied in Go projects.

## Key Objectives

- **Clean Architecture Principles:** The project diligently follows the principles of Clean Architecture, emphasizing separation of concerns, dependency inversion, and abstraction of implementation details.

- **Modularity and Scalability:** The structure of the project is designed to showcase modularity, making it easy to scale and extend the application while maintaining a clear separation between the business logic and external components.

- **Abstraction of Details:** Clean Architecture promotes the abstraction of external details, and this project demonstrates the use of interfaces and abstraction layers to achieve this separation.

## Usage

This project is intended as a reference and learning resource for developers aiming to implement Clean Architecture in their Go applications. It provides insights into how to structure code, manage dependencies, and achieve a high level of maintainability.

Feel free to explore the codebase and documentation to understand how the principles of Clean Architecture can be practically applied in a real-world Go project.

# Project Structure

## cmd

- **main.go**: Entry point of the application.

## config

- **config.yaml**: Configuration file for general settings.
- **sqlc.yaml**: Configuration file for SQLC code generation.

## docs

- **docs.go**: Documentation generation script.
- **swagger.json**: Swagger API documentation in JSON format.
- **swagger.yaml**: Swagger API documentation in YAML format.

## initiator

This directory contains various initialization and setup files.

- **config.go**: Configuration setup.
- **db.go**: Database setup.
- **handler.go**: HTTP request handlers.
- **initiator.go**: Overall project initialization.
- **logger.go**: Logging setup.
- **migration.go**: Database migrations.
- **module.go**: Module initialization.
- **persistence.go**: Persistence layer setup.
- **routes.go**: HTTP route setup.

## internal

This directory contains the internal components of the application.

### constants

- **dbinstance**: Database instance related constants.
- **errors**: Custom error definitions.
- **model**: Data models.
  - **db**: Database models.
  - **dto**: Data transfer objects.
  - **response.go**: Common response structures.

### query

- **queries**: SQL queries for SQLC code generation.
- **schemas**: SQL migration files.

### glue

- **routing**: Routing related components.
  - **route.go**: Generic route setup.

### handler

- **middleware**: HTTP middleware.
  - **errorMiddleware.go**: Error handling middleware.
- **rest**: RESTful API handlers.
  - **rest.go**: Generic RESTful API handler.
  - **user**: User-related RESTful API handlers.

### module

- **module.go**: Module interfaces.
- **user**: User module.
  - **user.go**: User module implementation.

### storage

- **persistence**: Persistence layer.
  - **user**: User-related database operations.
    - **user.go**: User persistence implementation.
- **storage.go**: Storage interfaces.

## Makefile

A Makefile with common tasks for building, testing, and cleaning.

## platform

- **logger**: Platform-specific components.
  - **logger.go**: Logging related components.


## Contribution

Contributions and feedback are welcome! If you have ideas for improvements or would like to contribute to showcasing Clean Architecture in Go, please feel free to open issues or pull requests.

