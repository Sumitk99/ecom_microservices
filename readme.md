
# E-commerce Microservices

This project is an advanced e-commerce platform designed using a microservices architecture ensuring **Scability**, **Reliability** and **Avoiding Single Point of Failure**. The services communicate with each other using **gRPC** which uses **Protobufs**, ensuring efficient and high-performance interactions. The primary services include account, catalog, order, cart, and an API gateway for routing requests to the appropriate services.


## Microservices Architecture
![Microservices Architecture](https://res.cloudinary.com/dwd3oedmz/image/upload/v1738262722/readme/ynfct0vd3kxbqyd8w5to.png)
- `account`: Manages user accounts and addresses.
- `catalog`: Manages product catalog.
- `order`: Manages orders and order processing.
- `cart`: Manages user shopping carts.
- `gateway`: API Gateway for routing requests to appropriate services.

## Technologies Used

- **Go**: Programming language used for implementing the services.
- **gRPC**: Communication protocol between microservices.
- **PostgreSQL**: Database used for storing data.
- **Protocol Buffers**: Used for defining gRPC services.
- **Elasticsearch**: Used for searching products.
- **Docker**: Containerization of services.
- **AWS EC2**: Deployment of services.
- **Nginx**: To Secure Server with HTTPS and Reverse proxy To API Gateway.

## Prerequisites

- Docker

## Setup
### Using Docker
1. **Copy Docker Compose File**:


2. **Pull Docker Images**:
    ```sh
    docker compose pull
    ```
3. **Run Docker Compose**:
    ```sh
    docker compose up
    ```

## Usage

### Deployed on : https://micro-scale.software/api

### Account Routes

- **Login**: `POST /login`
- **Signup**: `POST /signup`
- **Get Account**: `GET /account`

### Catalog Routes

- **Add Product**: `POST /product/add`
- **Get Product by ID**: `GET /product/{id}`
- **Get Products**: `GET /products`
- **Search Products**: `GET /products?search={keyword}`

### Order Routes

- **Place Order**: `POST /order`
- **Get Order**: `GET /user/order/{id}`
- **Get Orders for Account**: `GET /user/orders`

### Cart Routes

- **Add to Cart**: `POST /cart/add/{product_id}/{quantity}`
- **Get Cart**: `GET /cart/get`
- **Remove Item from Cart**: `DELETE /cart/remove/{product_id}`
- **Update Cart**: `/cart/update/{product_id}/{quantity}`
- **Delete Cart**: `DELETE /cart/delete`
- **Checkout**: `POST /cart/checkout

## gRPC Schemas

- **Account Service**: `account.proto`
- **Catalog Service**: `catalog.proto`
- **Order Service**: `order.proto`
- **Cart Service**: `cart.proto`

## CI/CD Pipeline
![Microservices Architecture](https://res.cloudinary.com/dwd3oedmz/image/upload/v1738262722/readme/fekxw9lt2mmetlr0p76z.png)

## Scalability Considerations

- **Microservices Architecture**: Each service can be scaled and deployed independently.
- **gRPC for Communication**: Efficient and high-performance communication.
- **Database**: PostgreSQL is robust and scalable.
- **Retry Logic**: Ensures resilience and scalability.

