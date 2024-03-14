# Simple Disbursement

## About the project
This project is simple disbursement implementation.
There have 4 Endpoints and also have some [MockApi](https://mockapi.io) endpoint to act like a bank.

```
Base URL : localhost:3000/v1
----

Name : Get Access Token
Method : GET
Endpoint : /payments/auth/token
Description :
  Is use to get access token for each request in endpoint (account verification and disbursement).
  The token is valid for 5 minutes and valid only one time use.

Name : Account Verification
Method : GET
Endpoint : /payments/account-verification
Description :
  Is use to verify the number of bank account.

Name : Disbursement
Method : POST
Endpoint : /payments/disbursements
  Is use to transfer/disbursement some amount into the bank.

Name : Disbursement Callback
Method : POST
Endpoint : /payments/disbursements/cb
  Is use to accept the result of the process from the bank.
```

## How To Run
- Clone this repository
- Make sure go already installed in your environment
- Copy file `.env.example` into the `.env` file
- Make sure to create database (Postgres), with the same name in your .env file
- Run `go get all` or `go mod tidy`
- Run `go run engine/rest/main.go`
- Then the server is running on port 3000

