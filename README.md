# examen_tecnico-go

## Requirements
- Docker environment (Colima, docker CLI)

## Build
- ./build.sh

## Run
- ./run.sh

## Usage

The application exposes 
POST /transactions/summary/email
The endpoint receives 
```json
{
    "path_to_file": "transactions.csv",
    "user_mail": "andres.fernandez.bina@gmail.com"
}
```
Execute
```
curl -X POST http://localhost:8080/transactions/summary/email \
  -H "Content-Type: application/json" \
  -d '{
    "path_to_file": "transactions.csv",
    "user_mail": "andres.fernandez.bina@gmail.com"
}'

```

## To Dos
- Implement DB, develop interfaces to dynamically load from csv or DB.
    - Email can access to user information such as email, name and add that data to email too.
- Add more test to guarantee API responds clear information to user.
- Add more tests to validate transaction_service.go has a correct error handling.
- Deploy app in cloud. 

