# internal-pos-system-api

## Overview

A cloud-based, serverless API for a multi-tenant Point of Sale (POS) system. The POS system is meant to be used internally within the vicinity of each tenant and main consumers are their employees.

> _ie. imagine a large company where it has its own cafeteria at different sites. The POS will be placed in those cafeteria and the employees are the buyers/consumers themselves. The employees use their "internal credit" to pay._

### Tech stack

- AWS Cognito (Auth)
- AWS API Gateway (Rest)
- AWS Lambda
- AWS DynamoDB (Datastore)
- AWS S3
- AWS EventBridge (Scheduled task / cron)
- AWS SNS/SQS

### Features

- Multi-tenant
- Employee Credit
  - To be used for checkout payment
  - Credit gets replenish every cutoff (semi-monthly / monthly)
- Checkout
- Misc / Admin
  - User / Employees
  - Menu
  - Concessionaire

## High Level Flow Diagram

![high level flow diagram](docs/internal-pos-system-api-high-level-flow.png)
