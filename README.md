# Bill Aggregation System - Docker Compose Setup Guide

## Overview

This guide provides step-by-step instructions to set up and run the Bill Aggregation System using Docker Compose. The system consists of multiple services, including microservices for users, providers, and billing, as well as supporting services like PostgreSQL, Redis, Kong API Gateway, and mock utility providers.

## Prerequisites

Ensure you have the following installed on your system:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://go.dev/dl/) (for local development)

## Services Overview

The system consists of the following services:

### Core Services

1. **Kong API Gateway (`kong`)**

   - Acts as the API Gateway for routing requests to different microservices.
   - Runs in declarative mode using `kong.yml` configuration.
   - Ports:
     - Proxy: `${KONG_PROXY_PORT}`
     - Admin: `${KONG_ADMIN_PORT}`

2. **PostgreSQL Database (`postgres`)**

   - Database server used by all microservices.
   - Initializes multiple databases using `init-multiple-databases.sql`.
   - Stores data for user, provider, and bill services.

3. **Redis (`redis`)**
   - In-memory data store for caching and session management.
   - Requires authentication via `${REDIS_PASSWORD}`.

### Microservices

4. **User Service (`user-service`)**

   - Handles user authentication and management.
   - Interacts with the `provider-service` and `bill-service`.
   - Uses `user_db` PostgreSQL database.

5. **Provider Service (`provider-service`)**

   - Manages utility providers and their integration.
   - Uses `provider_db` PostgreSQL database.

6. **Bill Service (`bill-service`)**
   - Fetches and aggregates bills from different utility providers.
   - Uses `bill_db` PostgreSQL database.

### Mock Utility Provider Services

7. **Mock Electricity Provider (`mock-electricity_provider`)**

   - Simulates an electricity provider API.
   - Runs on port `5001`.

8. **Mock Water Provider (`mock-water_provider`)**

   - Simulates a water provider API.
   - Runs on port `5002`.

9. **Mock Internet Provider (`mock-internet_provider`)**
   - Simulates an internet provider API.
   - Runs on port `5003`.

## Setup Instructions

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/zelalem-12/bill-aggregation-system_onetab.git
cd bill-aggregation-system_onetab
```

## 2. Configure Environment Variables

Create a `.env` file in the root directory and define the necessary environment variables.

### Example `.env` file:

      # API Gateway Proxy
      KONG_PROXY_HOST=0.0.0.0
      KONG_PROXY_PORT=8000

      # API Gateway Admin
      KONG_ADMIN_HOST=0.0.0.0
      KONG_ADMIN_PORT=8001

      # User Service Variables
      USER_SERVER_PORT=8080
      USER_BASE_URL=http://user-service:${USER_SERVER_PORT}

      # Provider Service Variables
      PROVIDER_SERVER_PORT=8081
      PROVIDER_BASE_URL=http://provider-service:${PROVIDER_SERVER_PORT}

      # Bill Service Variables
      BILL_SERVER_PORT=8082
      BILL_BASE_URL=http://bill-service:${BILL_SERVER_PORT}

      # Global Database Variables (Shared Across All Services)
      POSTGRES_HOST=postgres
      POSTGRES_PORT=5432
      POSTGRES_DATABASE=postgres
      POSTGRES_USER=postgres
      POSTGRES_PASSWORD=123123

      # Global JWT Variables
      ACCESS_TOKEN_KEY=thisisaccesstokenkey
      ACCESS_TOKEN_EXPIRY=7d
      REFRESH_TOKEN_KEY=thisisrefreshtokenkey
      REFRESH_TOKEN_EXPIRY=4w

      # SMTP Variables
      SMTP_HOST=smtp.gmail.com
      SMTP_PORT=587
      SENDER_EMAIL=zelalem.antigegn12@gmail.com
      SENDER_PASSWORD=hclubkvjubyfixjb

      # Client Variables
      FRONTEND_URL=https://onetap.et

      # Redis Variables
      REDIS_HOST=redis
      REDIS_PORT=6379
      REDIS_PASSWORD=123123

### 3. Build and Start the Services

Run the following command to build and start all services in detached mode:

```bash
docker-compose up -d --build
```

This command will:

Build and start the containers for Kong, PostgreSQL, Redis, and all microservices.

Set up the necessary databases and configurations based on the .env file.

### 4.Verify Running Containers

After the services are up and running, verify that the containers are properly started by running:

```bash
docker ps
```

You should see the following containers running:

    kong

    postgres

    redis

    user-service

    provider-service

    bill-service

    mock-electricity_provider

    mock-water_provider

    mock-internet_provider

### 5. Stop Services

To stop all running services and remove the containers, networks, and volumes created by Docker Compose, use:

```bash
docker-compose down
```

### 6. View Logs

If you need to view logs for any of the services, use the following command:

```bash
docker-compose logs -f [service-name]
```

For example, to view logs for the user-service, run:

```bash
docker-compose logs -f user-service
```
