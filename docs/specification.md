# RBAC (Role-Based Access Control) System Specification

## Overview

This document outlines the specifications for our Role-Based Access Control (RBAC) system implemented using Go, Gin Framework, JWT, and MySQL.

## Technologies Used

### Core Technologies

1. **Gin Framework**

   - Modern web framework for Go
   - Provides routing and middleware capabilities
   - Efficient request handling and response formatting

2. **JWT (JSON Web Tokens)**

   - Secure method for authentication
   - Contains encoded user information and claims
   - Split into access token (short-lived) and refresh token (long-lived)

3. **MySQL Database**
   - Stores user data, roles, and permissions
   - Maintains relationships between users, roles, and permissions
   - Uses foreign keys for data integrity

## System Architecture

### 1. Authentication System

#### Login Process

- User provides username/password
- System verifies credentials
- Generates JWT tokens upon successful authentication
- Returns user information and tokens

#### Token Refresh

- Uses refresh token to generate new access tokens
- Helps maintain user sessions securely
- Prevents frequent logins

### 2. Authorization System

#### Middleware Layer

- Validates JWT tokens on protected routes
- Checks user permissions
- Controls access to resources

#### Permission Checking

- Uses database queries to verify user permissions
- Implements role-based access control
- Prevents unauthorized access

## Logical Flow

### 1. User Authentication Flow

- User registers with username/password
- User logs in and receives tokens
- System stores user roles and permissions
- Tokens are used for subsequent requests

### 2. Request Authorization Flow

- User makes request with access token
- Middleware validates token
- System checks user permissions
- Grants or denies access based on roles

### 3. Role and Permission Management

- Admins can create/modify roles
- Permissions are assigned to roles
- Users are assigned roles
- Permissions cascade through role assignments

## Security Features

1. Password hashing using bcrypt
2. JWT token-based authentication
3. Role-based access control
4. Permission-level granular access
5. Secure session management
6. Token refresh mechanism

## Database Structure

### Tables

1. **Users**

   - Stores user information
   - Primary user data storage

2. **Roles**

   - Defines available roles
   - Role definitions and metadata

3. **Permissions**
   - Lists all permissions
   - System-wide permission definitions

### Junction Tables

1. **user_roles**

   - Links users to roles
   - Many-to-many relationship

2. **role_permissions**
   - Links roles to permissions
   - Many-to-many relationship

## API Endpoints

### Authentication Endpoints

1. `POST /api/users` - Register new user
2. `POST /api/login` - User login
3. `POST /api/refresh` - Refresh access token

### User Management Endpoints

1. `GET /api/users` - List all users
2. `GET /api/users/:id` - Get user details
3. `PUT /api/users/:id` - Update user
4. `DELETE /api/users/:id` - Delete user

### Role Management Endpoints

1. `GET /api/roles` - List all roles
2. `POST /api/roles` - Create new role
3. `GET /api/roles/:id` - Get role details
4. `PUT /api/roles/:id` - Update role
5. `DELETE /api/roles/:id` - Delete role

## Security Considerations

- All passwords must be hashed before storage
- Access tokens expire after 24 hours
- Refresh tokens expire after 7 days
- Protected routes require valid JWT
- Role-based permissions are strictly enforced
- Database queries use prepared statements

## Error Handling

- Standardized error responses
- Proper HTTP status codes
- Detailed error messages (in development)
- Sanitized error messages (in production)
