# golang_project

## - Running

###  Locally

1. Set environment values in .env.local

2. Migrate DB (https://github.com/golang-migrate/migrate)

    ```js
    migrate -path migrations -database "postgresql://user:password@localhost:5432/db" -verbose up
    ```
3. Seed DB with scripts/db/seed_data.sql

4. Run locally (port 8080 default)

   ```js
   GO_APP_ENV=local go run cmd/server/main.go
   ```
5. Usage and postman collection at /docs
    - Authorization Bearer required for protected endpoints