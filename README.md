# GoMall - Microservices E-commerce Platform

## 中文版

### 概述
GoMall 是一个基于 Go 语言和 CloudWeGo 框架（Hertz 和 Kitex）构建的微服务电商平台，旨在提供一个功能完善、性能优越的在线购物系统。项目采用前后端分离架构，包含用户服务、商品服务、购物车、订单处理、支付结算、通知服务等核心功能，支持分布式服务发现（Consul）、数据库优化（MySQL、Redis）、链路追踪（OpenTelemetry）以及容器化部署（Docker/Kubernetes）。


### 功能特性
- **用户管理**：
    - 用户注册/登录（支持 JWT 认证和权限控制）。
    - 个人中心：查看/编辑用户信息、订单历史。
- **商品服务**：
    - 商品浏览、分类和搜索。
    - 基于 AI 的个性化商品推荐（使用 Gorgonia 实现简单机器学习）。
- **购物车与订单**：
    - 实时购物车管理（添加、删除、修改数量）。
    - 分布式订单处理（事件驱动，NATS/Kafka + Saga 模式）。
    - 库存管理与高并发秒杀（分布式锁、Redis 缓存）。
- **支付与结算**：
    - 支付集成（支持模拟支付或第三方支付 API）。
    - 基于区块链的支付日志防篡改（Hyperledger Fabric + IPFS）。
- **通知服务**：
    - 实时消息推送（WebSocket + Redis Pub/Sub）。
    - 低库存警报和订单状态更新。
- **可观测性**：
    - 全链路追踪（OpenTelemetry）。
    - 监控与告警（Prometheus、Grafana、Alertmanager）。
- **部署与运维**：
    - 容器化部署（Docker Compose/Kubernetes）。
    - CI/CD 流水线（GitHub Actions/Jenkins）。
    - 黑盒监控与网站访问分析。

### 系统架构
- **微服务架构**：
    - `app/frontend`：Hertz 框架提供 HTTP API 和静态页面。
    - `app/user`：Kitex 框架提供用户相关 RPC 服务（注册、登录）。
    - `app/product`：商品服务，管理商品数据和推荐。
    - `app/cart`：购物车服务，处理商品添加/删除。
    - `app/order`：订单服务，处理订单创建、支付和状态跟踪。
    - `app/payment`：支付服务，集成支付接口和区块链日志。
    - `app/notification`：通知服务，推送实时消息。
    - `app/task`：分布式任务调度，管理库存和促销。
- **服务发现与通信**：
    - 使用 Consul 进行服务注册和发现，Kitex 实现 RPC 通信。
- **数据库与缓存**：
    - MySQL 存储核心数据（用户、商品、订单）。
    - Redis 缓存用户会话、库存状态、推荐结果。
- **可观测性**：
    - OpenTelemetry 记录全链路日志，Prometheus 收集指标，Grafana 展示仪表盘。

### 技术栈
- **语言与框架**：Go（1.23），CloudWeGo（Hertz、Kitex）。
- **服务发现**：Consul。
- **数据库**：MySQL（GORM），Redis。
- **消息队列**：NATS/Kafka。
- **可观测性**：OpenTelemetry，Prometheus，Grafana，Alertmanager。
- **容器化与部署**：Docker，Kubernetes。
- **CI/CD**：GitHub Actions，Jenkins。
- **其他**：区块链（Hyperledger Fabric），AI（Gorgonia），WebSocket（gorilla/websocket）。

### 安装与运行
#### 1. 环境准备
- Go 1.23
- Docker 和 Docker Compose
- MySQL 8.0
- Redis 7.0
- Consul 1.15

#### 2. 克隆项目
```bash
git clone https://github.com/meilingluolingluo/gomall.git
cd gomall
```

#### 3. 配置环境
- 复制 `.env.example` 为 `.env`，修改数据库、Redis 和 Consul 配置：
  ```bash
  cp .env.example .env
  ```
  编辑 `.env`：
  ```
  MYSQL_DSN="root:root@tcp(localhost:3306)/gomall?charset=utf8mb4&parseTime=True&loc=Local"
  REDIS_ADDRESS="localhost:6379"
  CONSUL_ADDRESS="localhost:8500"
  SESSION_SECRET="your-secret-key"
  ```

#### 4. 启动服务
- 启动 Docker 容器：
  ```bash
  docker-compose up -d
  ```
- 安装依赖并运行服务：
  ```bash
  cd app/frontend
  go mod tidy
  go run .

  cd ../../app/user
  go mod tidy
  go run .

  cd ../../app/product
  go mod tidy
  go run .
  # 依此类推运行其他服务
  ```

#### 5. 访问服务
- 打开浏览器，访问 `http://localhost:8080`，查看前端页面。
- 使用 Postman 测试 API（如 `POST /auth/register`）。

### 贡献指南
1. Fork 本仓库。
2. 创建新分支（`feature/your-feature`）。
3. 提交代码并推送至远程仓库。
4. 创建 Pull Request，描述更改和测试结果。

### 联系方式
- **作者**：Meiling Luo Lingluo
- **邮箱**：meilingluo@hrbeu.edu.com
- **GitHub**：https://github.com/meilingluolingluo

---

## English Version

### Overview
GoMall is a microservices-based e-commerce platform built with Go and the CloudWeGo framework (Hertz and Kitex), designed to provide a feature-rich, high-performance online shopping system. The project adopts a front-end and back-end separation architecture, including core functionalities such as user services, product services, shopping carts, order processing, payment settlement, notification services, and more. It supports distributed service discovery (Consul), database optimization (MySQL, Redis), distributed tracing (OpenTelemetry), and containerized deployment (Docker/Kubernetes).

This project serves as a learning and practice platform for backend development, suitable for showcasing in job applications, featuring technical highlights like distributed system design, high-concurrency handling, AI recommendations, and blockchain applications.

### Features
- **User Management**:
    - User registration/login (supporting JWT authentication and role-based access control).
    - Personal center: view/edit user info, order history.
- **Product Services**:
    - Product browsing, categorization, and search.
    - AI-driven personalized product recommendations (using Gorgonia for simple machine learning).
- **Shopping Cart & Orders**:
    - Real-time shopping cart management (add, remove, modify quantities).
    - Distributed order processing (event-driven, NATS/Kafka + Saga pattern).
    - Inventory management and high-concurrency flash sales (distributed locks, Redis caching).
- **Payment & Settlement**:
    - Payment integration (supporting mock payments or third-party APIs).
    - Blockchain-based tamper-proof payment logs (Hyperledger Fabric + IPFS).
- **Notification Services**:
    - Real-time message push (WebSocket + Redis Pub/Sub).
    - Low inventory alerts and order status updates.
- **Observability**:
    - Full-stack tracing (OpenTelemetry).
    - Monitoring and alerting (Prometheus, Grafana, Alertmanager).
- **Deployment & Operations**:
    - Containerized deployment (Docker Compose/Kubernetes).
    - CI/CD pipeline (GitHub Actions/Jenkins).
    - Black-box monitoring and website access analysis.

### System Architecture
- **Microservices Architecture**:
    - `app/frontend`: Hertz framework provides HTTP APIs and static pages.
    - `app/user`: Kitex framework provides user-related RPC services (registration, login).
    - `app/product`: Product service, managing product data and recommendations.
    - `app/cart`: Shopping cart service, handling add/remove operations.
    - `app/order`: Order service, processing order creation, payment, and status tracking.
    - `app/payment`: Payment service, integrating payment APIs and blockchain logs.
    - `app/notification`: Notification service, pushing real-time messages.
    - `app/task`: Distributed task scheduling, managing inventory and promotions.
- **Service Discovery & Communication**:
    - Uses Consul for service registration and discovery, with Kitex for RPC communication.
- **Database & Caching**:
    - MySQL stores core data (users, products, orders).
    - Redis caches user sessions, inventory status, and recommendation results.
- **Observability**:
    - OpenTelemetry records full-stack traces, Prometheus collects metrics, and Grafana displays dashboards.

### Tech Stack
- **Language & Frameworks**: Go (1.23), CloudWeGo (Hertz, Kitex).
- **Service Discovery**: Consul.
- **Databases**: MySQL (GORM), Redis.
- **Message Queue**: NATS/Kafka.
- **Observability**: OpenTelemetry, Prometheus, Grafana, Alertmanager.
- **Containerization & Deployment**: Docker, Kubernetes.
- **CI/CD**: GitHub Actions, Jenkins.
- **Others**: Blockchain (Hyperledger Fabric), AI (Gorgonia), WebSocket (gorilla/websocket).

### Installation & Running
#### 1. Prerequisites
- Go 1.23
- Docker and Docker Compose
- MySQL 8.0
- Redis 7.0
- Consul 1.15

#### 2. Clone the Repository
```bash
git clone https://github.com/meilingluolingluo/gomall.git
cd gomall
```

#### 3. Configure Environment
- Copy `.env.example` to `.env` and modify database, Redis, and Consul configurations:
  ```bash
  cp .env.example .env
  ```
  Edit `.env`:
  ```
  MYSQL_DSN="root:root@tcp(localhost:3306)/gomall?charset=utf8mb4&parseTime=True&loc=Local"
  REDIS_ADDRESS="localhost:6379"
  CONSUL_ADDRESS="localhost:8500"
  SESSION_SECRET="your-secret-key"
  ```

#### 4. Start Services
- Launch Docker containers:
  ```bash
  docker-compose up -d
  ```
- Install dependencies and run services:
  ```bash
  cd app/frontend
  go mod tidy
  go run .

  cd ../../app/user
  go mod tidy
  go run .

  cd ../../app/product
  go mod tidy
  go run .
  # Run other services similarly
  ```

#### 5. Access the Services
- Open a browser and visit `http://localhost:8080` to view the frontend.
- Use Postman to test APIs (e.g., `POST /auth/register`).

### Contribution Guide
1. Fork this repository.
2. Create a new branch (`feature/your-feature`).
3. Commit your changes and push to the remote repository.
4. Create a Pull Request, describing changes and test results.

### Contact
- **Author**: Meiling Luo Lingluo
- **Email**: meilingluolingluo@example.com
- **GitHub**: https://github.com/meilingluolingluo

---
