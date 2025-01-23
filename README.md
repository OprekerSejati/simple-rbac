Database

Use file on sql/rbac.sql to import to your database

make .env file and put it in your project root folder
DB_HOST=localhost
DB_PORT=3306
DB_USER=YOUR-USER
DB_PASSWORD=YOUR-PASSWORD
DB_NAME=rbac
JWT_SECRET_KEY=YOUR-JWT-SECRET
JWT_EXPIRE_HOURS=24
REFRESH_TOKEN_EXPIRE_DAYS=7 

run
go run main.go

if success, you'll see similar result
2025/01/23 14:44:22 Server starting on :8080

for more detail please refer to files:
1. docs/specification.md
2. docs/docs.md

