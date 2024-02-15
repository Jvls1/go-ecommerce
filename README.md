
# Product Service

Service to create and list products.


## Endpoints

#### Create product

```http
POST /api/products/
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
GET /api/products
```


#### Find product by UUID

```http
GET /api/products/${uuid}
```

| Param       | Type       | Description                         |
| :---------- | :--------- | :---------------------------------- |
| `id`        | `string`   | Product UUID                        |

