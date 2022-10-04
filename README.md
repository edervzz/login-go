# login-go, stratplus

This API works with a stub for repository purpose.

## Create user
```
curl --location --request POST 'localhost:8000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "ederv",
    "email": "edervzz@gmail.com",
    "phoneNumber": "1122334455",
    "password": "Eder123$"
}'
```

## Login user
```
curl --location --request POST 'localhost:8000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "ederv",
    "password": "Eder123$"
}'
```