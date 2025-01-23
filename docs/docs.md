Simple Role-Based Access Control (RBAC) dengan Golang

Docs

1. Authentication Endpoints:

# Register User

POST http://localhost:8080/api/users
{
"username": "admin",
"password": "admin123",
"roles": ["admin"]
}

# Login

POST http://localhost:8080/api/login
{
"username": "admin",
"password": "admin123"
}

# Refresh Token

POST http://localhost:8080/api/refresh
Headers:
X-Refresh-Token: <your_refresh_token>

2. User Management Endpoints:

# Get All Users

GET http://localhost:8080/api/users
Headers:
Authorization: Bearer <your_access_token>

# Get Single User

GET http://localhost:8080/api/users/1
Headers:
Authorization: Bearer <your_access_token>

# Update User

PUT http://localhost:8080/api/users/1
Headers:
Authorization: Bearer <your_access_token>
{
"username": "updated_admin",
"roles": ["admin", "user"]
}

# Delete User

DELETE http://localhost:8080/api/users/1
Headers:
Authorization: Bearer <your_access_token>

3.Role Management Endpoints:

# Get All Roles

GET http://localhost:8080/api/roles
Headers:
Authorization: Bearer <your_access_token>

# Create New Role

POST http://localhost:8080/api/roles
Headers:
Authorization: Bearer <your_access_token>
{
"name": "editor",
"permissions": ["create_post", "edit_post", "view_post"]
}

# Get Single Role

GET http://localhost:8080/api/roles/1
Headers:
Authorization: Bearer <your_access_token>

# Update Role

PUT http://localhost:8080/api/roles/1
Headers:
Authorization: Bearer <your_access_token>
{
"name": "editor",
"permissions": ["create_post", "edit_post", "view_post", "delete_post"]
}

# Delete Role

DELETE http://localhost:8080/api/roles/1
Headers:
Authorization: Bearer <your_access_token>

Example Response Formats:

Successful Login Response:
{
"access_token": "eyJhbGciOiJIUzI1NiIs...",
"refresh_token": "eyJhbGciOiJIUzI1NiIs...",
"user": {
"id": 1,
"username": "admin",
"roles": ["admin"]
}
}

Get Users Response:
[
{
"id": 1,
"username": "admin",
"roles": ["admin"]
},
{
"id": 2,
"username": "user1",
"roles": ["user"]
}
]

Get Role Response:
{
"id": 1,
"name": "editor",
"permissions": [
"create_post",
"edit_post",
"view_post",
"delete_post"
]
}

Important Notes:

1. Always include the Authorization header with Bearer token for protected routes
2. The token format is: Bearer <your_access_token>
3. All request bodies should be in JSON format
4. Set Content-Type header to application/json
5. HTTP status codes:
   200: Success
   201: Created
   400: Bad Request
   401: Unauthorized
   403: Forbidden
   404: Not Found
   500: Internal Server Error
