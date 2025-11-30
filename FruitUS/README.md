# FruitUS

FruitUS is a multi-tenant SaaS application designed to manage tenants and users efficiently. This application follows Clean Architecture principles, ensuring a clear separation of concerns and maintainability.

## Project Structure

```
FruitUS
├── cmd
│   └── fruitus
│       └── main.go          # Entry point of the application
├── internal
│   ├── domain
│   │   ├── tenant.go        # Defines the Tenant struct and methods
│   │   ├── user.go          # Defines the User struct and methods
│   │   ├── repository
│   │   │   └── tenant_repository.go # TenantRepository interface
│   │   └── service
│   │       └── tenant_service.go    # Business logic for tenants
│   ├── usecase
│   │   ├── tenant
│   │   │   ├── create_tenant.go # Logic for creating a new tenant
│   │   │   └── get_tenant.go    # Logic for retrieving tenant information
│   │   └── auth
│   │       └── auth_usecase.go   # Authentication-related logic
│   ├── interfaces
│   │   ├── http
│   │   │   ├── handler.go        # HTTP handlers for the application
│   │   │   └── middleware.go     # Middleware functions
│   │   └── repository
│   │       └── tenant_repo_pg.go # PostgreSQL implementation of TenantRepository
│   └── infrastructure
│       ├── db
│       │   └── postgres.go       # Database connection logic for PostgreSQL
│       └── config
│           └── config.go         # Configuration structures and methods
├── pkg
│   └── logger
│       └── logger.go            # Logger for application events
├── api
│   └── openapi.yaml             # OpenAPI specification for the API
├── migrations
│   └── 0001_init.up.sql         # SQL migration script for initializing the database schema
├── configs
│   └── config.example.yaml       # Example configuration file
├── scripts
│   └── run.sh                   # Shell script for running the application
├── go.mod                        # Module and dependencies
├── Makefile                      # Build instructions and commands
└── README.md                     # Documentation for the project
```

## Getting Started

To get started with the FruitUS application, clone the repository and follow the instructions in the `scripts/run.sh` file to run the application.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.