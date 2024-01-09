# Go Project

Project with repository pattern and dependency injection written in Go.

## Project Structure

This section explains the key directories and files in this project and their roles:

```
.
├── cmd
│   ├── api                     # Entry point for the API application.
│   └── scheduler               # Entry point for the scheduler application.
├── internal                    # Private application and library code.
│   ├── dto                     # Data Transfer Objects.
│   ├── handler                 # Handlers respond to HTTP API requests.
│   ├── model                   # Domain models used across the application.
│   ├── repository              # Interfaces for interacting with the data storage layer.
│   │   └── mysql               # MySQL implementation of the repository interfaces.
│   └── service                 # Business logic and service layer implementation.
├── locale
│   ├── en                      # English localization files.
│   └── id                      # Indonesian localization files.
├── log                         # Log files.
├── migration                   # Database migration scripts.
│   └── mysql                   # MySQL specific migration scripts.
├── mocks                       # Mock files for testing.
├── pkg                         # Library code that's ok to use by external applications.
├── sdk                         # SDK for third-party.
├── go.mod                      # Go module file declaring module path and dependency requirements.
├── go.sum                      # Go checksum file for the dependencies.
└── README.md

```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go: Provide instructions on how to install Go on their platform ([Go Installation Guide](https://golang.org/doc/install)).
- Navigate to the specific application folder, for example, the **API application** at `/cmd/api`.
- Create a copy of `.env.example` and modify the configuration accordingly.
- Make sure initial database is created and accessible, if not use the following query for MySQL

```
CREATE DATABASE `goproject_db` DEFAULT CHARACTER SET = `utf8mb4` DEFAULT COLLATE = `utf8mb4_general_ci`;
```

- Finally, run the application using the command `go run main.go`.

## Tools

### Migration (MySQL)

Migration scripts are managed via [Goose](https://github.com/pressly/goose). Make sure you have install `goose` binary on your local machine.

To generate a new migration file, run the following command:

```
cd migration/mysql
goose create new_table sql
```

To up a migration

```
cd migration/mysql
goose mysql "user:password@tcp(host:port)/dbname?charset=utf8&parseTime=true" up
```

# API Guidelines

These instructions will guide you on how to utilize the REST API:

## Authentication

Authentication is achieved through a Bearer Token using JWT in the request header.

```
Authorization: Bearer <jwt_token>
```

## Localization

To ensure proper localization, include `Accept-Timezone` with the value of desired timezone and `Accept-Language` with the value of the available supported translation file in your request header. List of available Timezone value can be checked from [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)

```
Accept-Timezone: Asia/Jakarta
Accept-Language: en
```

# License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE.md) file for details.
