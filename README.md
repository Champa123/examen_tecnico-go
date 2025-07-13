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
    "user_mail": "afernandezbina@frba.utn.edu.ar"
}'

```

