
# Go E-commerce

E-commerce project with microservice architecture using Go, Gin-Gonic, Postgres, Mongo and Docker.


## Auth Service
The auth service handle the login with JWT Token and user creation, also create the roles and permissions for access control.

Techs: Go, Chi, Postgres

Header for all the POST routes except the /login:
| Header      | Value       |
| :---------- | :--------- |
| `Content-Type`      | `application/json`   |
| `Authorization`     | `Bearer jwt-token`   |

Header for the GET routes:
| Header      | Value       |
| :---------- | :--------- |
| `Authorization`     | `Bearer jwt-token`   |

#### Login

```http
POST /login
```
| Header      | Type       |
| :---------- | :--------- |
| `Content-Type`      | `application/json`   |

```json
{
  "email": "email@gmail.com",
  "password": "strongPassword123"
}
```

#### Create permission
```http
POST /permissions
```
```json
{
  "name": "read-permission",
  "description": "Grant access to create a permission"
}
```

#### Find permission by id
```http
GET /permissions/${permissionId}
```
| Param       | Type       |
| :---------- | :--------- |
| `permissionId`      | `uuid`   |

#### Create permission
```http
POST /permissions/
```
```json
{
  "name": "read-permission",
  "description": "Grant access to read a permission"
}
```

#### Find user by id
```http
GET /users/${userId}
```
| Param       | Type       |
| :---------- | :--------- |
| `userId`      | `uuid`   |

#### Create user
```http
POST /users/
```
```json
{
  "name": "John Doe",
  "password": "strongPassword123",
  "email": "test@gmail.com"
}
```

#### Add role to user
```http
POST /users/roles
```
```json
{
  "userId": "bd42b8fd-5ba1-482b-8b3e-13b77095a7d7",
  "roleId": "5658b58d-b749-4a74-bb71-483324eb0705"
}
```

#### Find role by id
```http
GET /roles/${roleId}
```
| Param       | Type       |
| :---------- | :--------- |
| `roleId`      | `uuid`   |

#### Create role
```http
POST /roles/
```
```json
{
  "name": "Admin1",
  "description": "Admin role"
}
```

#### Add permission to role
```http
POST /roles/permissions
```
```json
{
  "roleId": "5658b58d-b749-4a74-bb71-483324eb0705",
  "permissionId": "c523de53-8b8e-47fa-a751-15f0e678f8d9"
}
```

## Product Service

#### Create product

```http
POST /api/v1/products/
```

```json
{
  "name": "Product Test",
  "description": "This is a test product",
  "image_url": "https://example.com/image.jpg",
  "price": 19.99,
  "quantity": 10,
  "department_id": 1
}
```

#### Find all products (with pagination)

```http
GET /api/v1/products?page=0&pageSize=10
```
| Param       | Type       |
| :---------- | :--------- |
| `page`      | `number`   |
| `pageSize`  | `number`   |

#### Find product by id

```http
GET /api/v1/products/${id}
```

| Param       | Type       | 
| :---------- | :--------- | 
| `id`        | `uuid`     |

#### Find product by department id

```http
GET /api/v1/products/department/${departmentId}
```

| Param           | Type       | 
| :-------------- | :--------- | 
| `departmentId`  | `uuid`     |
