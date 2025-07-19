# BareCMS API Documentation

## Overview

BareCMS provides a RESTful API for content management and public data access. All authenticated endpoints require a JWT token in the Authorization header.

## Base URL

```
http://localhost:8080
```

## Authentication

Include the JWT token in the Authorization header for all authenticated endpoints:

```
Authorization: Bearer <your-jwt-token>
```

---

## 🌐 Public Endpoints

### Get Site Data

Retrieve all site content publicly without authentication.

**Endpoint:** `GET /api/:siteSlug/data`

**Description:** Returns all collections and entries for a site using its slug. This is the primary endpoint for headless usage.

**Parameters:**

- `siteSlug` (path) - The unique slug of the site

**Example Request:**

```bash
curl -X GET http://localhost:8080/my-blog/data
```

**Example Response:**

```json
{
  "site": {
    "id": 1,
    "name": "My Blog",
    "slug": "my-blog",
    "description": "A simple blog site"
  },
  "collections": [
    {
      "id": 1,
      "name": "Posts",
      "slug": "posts",
      "description": "Blog posts collection",
      "entries": [
        {
          "id": 1,
          "title": "Welcome to BareCMS",
          "content": "This is my first blog post using BareCMS...",
          "slug": "welcome-to-barecms",
          "created_at": "2024-01-15T10:30:00Z",
          "updated_at": "2024-01-15T10:30:00Z"
        },
        {
          "id": 2,
          "title": "Getting Started Guide",
          "content": "Here's how to get started with BareCMS...",
          "slug": "getting-started-guide",
          "created_at": "2024-01-16T14:22:00Z",
          "updated_at": "2024-01-16T14:22:00Z"
        }
      ]
    },
    {
      "id": 2,
      "name": "Pages",
      "slug": "pages",
      "description": "Static pages collection",
      "entries": [
        {
          "id": 3,
          "title": "About",
          "content": "Learn more about our mission...",
          "slug": "about",
          "created_at": "2024-01-15T11:00:00Z",
          "updated_at": "2024-01-15T11:00:00Z"
        }
      ]
    }
  ]
}
```

---

## 🔐 Authentication Endpoints

### Register User

Create a new user account.

**Endpoint:** `POST /api/auth/register`

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**

```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

### Login User

Authenticate and receive a JWT token.

**Endpoint:** `POST /api/auth/login`

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

### Logout User

Invalidate the current session.

**Endpoint:** `POST /api/auth/logout`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "message": "Logged out successfully"
}
```

---

## 🏢 Sites Management

### List Sites

Get all sites for the authenticated user.

**Endpoint:** `GET /api/sites`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "sites": [
    {
      "id": 1,
      "name": "My Blog",
      "slug": "my-blog",
      "description": "A simple blog site",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  ]
}
```

### Create Site

Create a new site.

**Endpoint:** `POST /api/sites`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "name": "My Portfolio",
  "slug": "my-portfolio",
  "description": "Personal portfolio website"
}
```

**Response:**

```json
{
  "site": {
    "id": 2,
    "name": "My Portfolio",
    "slug": "my-portfolio",
    "description": "Personal portfolio website",
    "created_at": "2024-01-16T09:00:00Z",
    "updated_at": "2024-01-16T09:00:00Z"
  }
}
```

### Get Site Details

Retrieve details of a specific site.

**Endpoint:** `GET /api/sites/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "site": {
    "id": 1,
    "name": "My Blog",
    "slug": "my-blog",
    "description": "A simple blog site",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

### Update Site

Update an existing site.

**Endpoint:** `PUT /api/sites/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "name": "My Updated Blog",
  "description": "Updated description"
}
```

### Delete Site

Delete a site and all its collections/entries.

**Endpoint:** `DELETE /api/sites/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

---

## 📚 Collections Management

### List Collections

Get all collections for a specific site.

**Endpoint:** `GET /api/sites/:siteId/collections`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "collections": [
    {
      "id": 1,
      "name": "Posts",
      "slug": "posts",
      "description": "Blog posts collection",
      "site_id": 1,
      "created_at": "2024-01-15T10:15:00Z",
      "updated_at": "2024-01-15T10:15:00Z"
    }
  ]
}
```

### Create Collection

Create a new collection within a site.

**Endpoint:** `POST /api/sites/:siteId/collections`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "name": "Products",
  "slug": "products",
  "description": "Product catalog collection"
}
```

### Get Collection Details

Retrieve details of a specific collection.

**Endpoint:** `GET /api/collections/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

### Update Collection

Update an existing collection.

**Endpoint:** `PUT /api/collections/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

### Delete Collection

Delete a collection and all its entries.

**Endpoint:** `DELETE /api/collections/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

---

## 📝 Entries Management

### List Entries

Get all entries for a specific collection.

**Endpoint:** `GET /api/collections/:collectionId/entries`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "entries": [
    {
      "id": 1,
      "title": "Welcome to BareCMS",
      "content": "This is my first blog post...",
      "slug": "welcome-to-barecms",
      "collection_id": 1,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ]
}
```

### Create Entry

Create a new entry within a collection.

**Endpoint:** `POST /api/collections/:collectionId/entries`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "title": "New Blog Post",
  "content": "Content of the blog post...",
  "slug": "new-blog-post"
}
```

### Get Entry Details

Retrieve details of a specific entry.

**Endpoint:** `GET /api/entries/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

### Update Entry

Update an existing entry.

**Endpoint:** `PUT /api/entries/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "title": "Updated Blog Post Title",
  "content": "Updated content..."
}
```

### Delete Entry

Delete an entry.

**Endpoint:** `DELETE /api/entries/:id`

**Headers:**

```
Authorization: Bearer <your-jwt-token>
```

---

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request

```json
{
  "error": "Invalid request data"
}
```

### 401 Unauthorized

```json
{
  "error": "Authentication required"
}
```

### 403 Forbidden

```json
{
  "error": "Access denied"
}
```

### 404 Not Found

```json
{
  "error": "Resource not found"
}
```

### 500 Internal Server Error

```json
{
  "error": "Internal server error"
}
```

---

## Usage Examples

### Headless CMS Workflow

1. **Register and Login**

   ```bash
   # Register
   curl -X POST http://localhost:8080/api/auth/register \
     -H "Content-Type: application/json" \
     -d '{"email": "user@example.com", "password": "password123"}'

   # Login
   curl -X POST http://localhost:8080/api/auth/login \
     -H "Content-Type: application/json" \
     -d '{"email": "user@example.com", "password": "password123"}'
   ```

2. **Create Site Structure**

   ```bash
   # Create site
   curl -X POST http://localhost:8080/api/sites \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"name": "My Blog", "slug": "my-blog", "description": "Personal blog"}'

   # Create collection
   curl -X POST http://localhost:8080/api/sites/1/collections \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"name": "Posts", "slug": "posts", "description": "Blog posts"}'

   # Create entry
   curl -X POST http://localhost:8080/api/collections/1/entries \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"title": "Hello World", "content": "My first post!", "slug": "hello-world"}'
   ```

3. **Access Data Publicly**
   ```bash
   # Get all site data (no authentication needed)
   curl -X GET http://localhost:8080/my-blog/data
   ```

This workflow demonstrates the core concept: use the authenticated API to manage content, then access it publicly via the site slug for your frontend applications.
