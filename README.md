# Razorpay Payment Integration using AWS Lambda (Go)

This project demonstrates how to integrate **Razorpay payment order creation**
using **AWS Lambda (Go runtime)** and **API Gateway (HTTP API)**.

---

## 🚀 Features
- AWS Lambda using Go
- API Gateway HTTP API
- Razorpay Order Creation API
- Secure usage of environment variables

---

## 🛠 Tech Stack
- Go
- AWS Lambda
- API Gateway
- Razorpay API

---

## 📌 API Endpoint

### Create Order
**POST** `/create-order`

#### Request Body
```json
{
  "amount": 50000
}
```
#### Response
```json
{
  "id": "order_xxxxx",
  "amount": 50000,
  "currency": "INR",
  "status": "created"
}
```
