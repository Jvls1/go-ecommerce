
# Go E-commerce

E-commerce project with microservice architecture using Go, Gin-Gonic, Postgres, Mongo and Docker.


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

