# BareCMS API Documentation

## Overview

BareCMS provides a RESTful API for content management and public data access. All authenticated endpoints require a JWT token in the `Authorization` header.

## Base URL

```bash
http://localhost:8080
```

## Authentication

Include the JWT token in the `Authorization` header for all authenticated endpoints:

```bash
Authorization: Bearer <your-jwt-token>
```

---

## 🌐 Public Endpoints

### Health Check

Check if the API is running and healthy.

**Endpoint:** `GET /api/health`

**Description:** Returns the health status of the API. No authentication required.

**Example Request:**

```bash
curl -X GET http://localhost:8080/api/health
```

**Example Response:**

```json
{
  "status": "up"
}
```

### Get Site Data

Retrieve all site content publicly without authentication. This is the primary endpoint for headless usage.

**Endpoint:** `GET /api/:siteSlug/data`

**Description:** Returns all collections and entries for a site using its slug. BareCMS keeps it simple by organizing all content under a `data` field, where each collection is accessible by its slug containing an array of entries with their field values directly accessible.

**Parameters:**

- `siteSlug` (path) - The unique slug of the site

**Example Request:**

```bash
curl -X GET http://localhost:8080/api/myblog/data
```

**Example Response:**

```json
{
  "id": "44394f36-daa3-451c-970f-59238c46ce36",
  "name": "myblog",
  "slug": "myblog",
  "data": {
    "articles": [
      {
        "content": "this is my article post content",
        "draft": "false",
        "published": "2025-07-21",
        "title": "my sample article"
      }
    ],
    "products": [
      {
        "name": "Sample Product",
        "price": "29.99",
        "description": "A great product for everyone"
      }
    ]
  }
}
```

**Response Structure:**

- `id` - The unique identifier of the site
- `name` - The name of the site
- `slug` - The URL-friendly slug of the site
- `data` - Object containing all collections, where:
  - Each key is a collection slug (e.g., "articles", "products")
  - Each value is an array of entries for that collection
  - Entry objects contain field names as keys with their values directly accessible (no nested `data` object)

This simple structure makes it easy to consume in frontend applications - you can directly access `response.data.articles` to get all articles, or `response.data.products` for products, etc.

---

## 🔐 Authentication Endpoints

### Register User

Create a new user account.

**Endpoint:** `POST /api/auth/register`

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "username": "myuser"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "user123",
    "email": "user@example.com",
    "username": "myuser"
  },
  "message": "User created successfully"
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
    "id": "user123",
    "email": "user@example.com",
    "username": "myuser"
  }
}
```

### Logout User

Invalidate the current session.

**Endpoint:** `POST /api/auth/logout`

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "message": "User logged out successfully"
}
```

---

## 👤 User Management (Authenticated)

### Get Authenticated User

Retrieve details of the currently authenticated user. Requires authentication.

**Endpoint:** `GET /api/user`

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "id": "user123",
  "email": "user@example.com",
  "username": "myuser"
}
```

### Delete User

Delete the currently authenticated user. Requires authentication.

**Endpoint:** `DELETE /api/user/:userId`

**Parameters:**

- `userId` (path) - The ID of the user to delete (must match authenticated user)

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "message": "User deleted successfully"
}
```

---

## 🏢 Sites Management (Authenticated)

### List Sites

Get all sites for the authenticated user.

**Endpoint:** `GET /api/sites`

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "sites": [
    {
      "id": "site123",
      "userId": "user789",
      "name": "My Blog",
      "slug": "my-blog"
    },
    {
      "id": "site456",
      "userId": "user789",
      "name": "My Portfolio",
      "slug": "my-portfolio"
    }
  ]
}
```

### Create Site

Create a new site for the authenticated user.

**Endpoint:** `POST /api/sites`

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "name": "My New Site",
  "userId": "user789"
}
```

**Response:**

```json
{
  "message": "Site created!"
}
```

### Get Site Details

Retrieve details of a specific site by its ID.

**Endpoint:** `GET /api/sites/:id`

**Parameters:**

- `id` (path) - The ID of the site to retrieve

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "site": {
    "id": "site123",
    "userId": "user789",
    "name": "My Blog",
    "slug": "my-blog"
  }
}
```

### Get Site Details with Collections

Retrieve details of a specific site by its ID, including all its collections.

**Endpoint:** `GET /api/sites/:id/collections`

**Parameters:**

- `id` (path) - The ID of the site

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "site": {
    "id": "site123",
    "userId": "user789",
    "name": "My Blog",
    "slug": "my-blog"
  },
  "collections": [
    {
      "id": "col123",
      "siteId": "site123",
      "name": "Posts",
      "slug": "posts",
      "fields": [
        { "name": "title", "type": "string" },
        { "name": "content", "type": "string" }
      ],
      "entries": []
    }
  ]
}
```

### Delete Site

Delete a site and all its associated collections and entries. This action is irreversible.

**Endpoint:** `DELETE /api/sites/:id`

**Parameters:**

- `id` (path) - The ID of the site to delete

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "message": "Site deleted!"
}
```

---

## 📚 Collections Management (Authenticated)

### Create Collection

Create a new collection.

**Endpoint:** `POST /api/collections`

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "siteId": "site123",
  "name": "Products",
  "fields": [
    { "name": "title", "type": "string" },
    { "name": "price", "type": "number" }
  ]
}
```

**Supported Field Types:**

| Type      | Description       | Input           |
| --------- | ----------------- | --------------- |
| `string`  | Single line text  | Text input      |
| `text`    | Multi-line text   | Textarea        |
| `number`  | Numeric values    | Number input    |
| `boolean` | True/false values | Select (Yes/No) |
| `date`    | Date picker       | Date input      |
| `image`   | Image URL         | URL input       |
| `url`     | Website links     | URL input       |

**Response:**

```json
{
  "message": "Collection created!"
}
```

### Get Collection Details

Retrieve details of a specific collection.

**Endpoint:** `GET /api/collections/:id`

**Parameters:**

- `id` (path) - The ID of the collection

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "id": "col123",
  "siteId": "site123",
  "name": "Posts",
  "slug": "posts",
  "fields": [
    { "name": "title", "type": "string" },
    { "name": "content", "type": "string" }
  ],
  "entries": [
    {
      "id": "entry1",
      "collection_id": "col123",
      "data": {
        "title": "Sample Post",
        "content": "Sample content"
      }
    }
  ]
}
```

### Get Collections by Site ID

Get all collections for a specific site.

**Endpoint:** `GET /api/collections/:siteId`

**Parameters:**

- `siteId` (path) - The ID of the site

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "collections": [
    {
      "id": "col123",
      "siteId": "site123",
      "name": "Posts",
      "slug": "posts",
      "fields": [{ "name": "title", "type": "string" }],
      "entries": []
    }
  ]
}
```

### Get Collection Entries

Get all entries for a specific collection.

**Endpoint:** `GET /api/collections/:id/entries`

**Parameters:**

- `id` (path) - The ID of the collection

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "id": "col123",
  "siteId": "site123",
  "name": "Posts",
  "slug": "posts",
  "fields": [{ "name": "title", "type": "string" }],
  "entries": [
    {
      "id": "entry1",
      "collection_id": "col123",
      "data": {
        "title": "Sample Post"
      }
    }
  ]
}
```

### Delete Collection

Delete a collection and all its entries.

**Endpoint:** `DELETE /api/collections/:id`

**Parameters:**

- `id` (path) - The ID of the collection

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "message": "Collection deleted!"
}
```

---

## 📝 Entries Management (Authenticated)

### Create Entry

Create a new entry.

**Endpoint:** `POST /api/entries`

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Request Body:**

```json
{
  "collectionId": "col123",
  "data": {
    "title": {
      "value": "New Blog Post",
      "type": "string"
    },
    "content": {
      "value": "Content of the blog post...",
      "type": "string"
    },
    "published": {
      "value": "true",
      "type": "boolean"
    }
  }
}
```

**Response:**

```json
{
  "message": "Collection created!"
}
```

### Get Entry Details

Retrieve details of a specific entry.

**Endpoint:** `GET /api/entries/:id`

**Parameters:**

- `id` (path) - The ID of the entry

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "id": "entry1",
  "collection_id": "col123",
  "data": {
    "title": "Blog Post Title",
    "content": "Blog post content...",
    "slug": "blog-post-title"
  }
}
```

### Delete Entry

Delete an entry.

**Endpoint:** `DELETE /api/entries/:id`

**Parameters:**

- `id` (path) - The ID of the entry

**Headers:**

```bash
Authorization: Bearer <your-jwt-token>
```

**Response:**

```json
{
  "message": "Entry deleted!"
}
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
     -d '{"email": "user@example.com", "password": "password123", "username": "myuser"}'

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
     -d '{"name": "My Blog", "userId": "user123"}'

   # Create collection
   curl -X POST http://localhost:8080/api/collections \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"siteId": "site123", "name": "Posts", "fields": [{"name":"title","type":"string"},{"name":"content","type":"string"}]}'

   # Create entry
   curl -X POST http://localhost:8080/api/entries \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"collectionId": "col123", "data": {"title": {"value": "Hello World", "type": "string"}, "content": {"value": "My first post!", "type": "text"}}}'
   ```

3. **Access Data Publicly**
   ```bash
   # Get all site data (no authentication needed)
   curl -X GET http://localhost:8080/api/my-blog/data
   ```

This workflow demonstrates the core concept: use the authenticated API to manage content, then access it publicly via the site slug for your frontend applications.
