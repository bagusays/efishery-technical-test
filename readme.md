# Efishery Technical Test

Consist of 2 apps
- Auth Service (Typescript/Node)
- Fetch Service (Go)

## Helicopter View / Design Diagram
![Alt text](design-diagram.png?raw=true " ")

## Auth Service
Provides user and authentication functionality.

![Alt text](auth-context.png?raw=true " ")

## Fetch Service
As a resource orchestrator for getting all informations about fisheries data.

![Alt text](fetch-context.png?raw=true " ")

## Prerequisites
- Create new config in `fetch-service/config` and `auth-service/config` by copy `sample.json` and paste as new file with name `config.json`. Ensure JWTSecret is the same on both.
- Docker installed or
- Node (14 or newer) & Go (1.6 or newer) installed

## How To Run
- Manually
  - `cd auth-service && yarn dev`
  - `cd fetch-service && go run main.go`

- Via Docker
  - `docker-compose up`

- and open
  - auth service at http://localhost:8080
  - fetch service at http://localhost:8081


# API Documentations
## Auth Service
### Create User
- URL: `/api/user`
- Method: `POST`
- Authentication Required: NO
- Role Permission required: NONE
- Request Body Constraints
```
{
    "name": "string",
    "phone": "string", // should be unique
    "role": "ADMIN" | "BASIC",
    "userName": "string" // should be unique
}
```

Success Response (200 OK)
```
{
    "password": "string"
}
```

### Login
- URL: `/api/auth/login`
- Method: `POST`
- Authentication Required: NO
- Role Permission required: NONE

Request Body Constraints
```
{
    "phone": "string",
    "password": "string"
}
```

Success Response (200 OK)
```
{
    "token": "string"
}
```

### Validate
- URL: `/api/auth/validate`
- Method: `GET`
- Authentication Required: YES
- Role Permission required: NONE

Headers
```
Authorization: Bearer $TOKEN
```

Success Response (200 OK)
```
{
    "name": "string",
    "phone": "string",
    "role": "ADMIN" | "BASIC",
    "created_at": "2022-03-06T06:56:15.945+00:00",
    "userName": "string",
    "iat": number,
    "exp": number
}
```


## Fetch Service
### Validate
- URL: `/api/auth/validate`
- Method: `GET`
- Authentication Required: YES
- Role Permission required: NONE

Headers
```
Authorization: Bearer $TOKEN
```

Success Response (200 OK)
```
{
    "name": "string",
    "phone": "string",
    "role": "ADMIN" | "BASIC",
    "created_at": "2022-03-06T06:56:15.945+00:00",
    "userName": "string",
    "iat": number,
    "exp": number
}
```

### Fetch All Resources
- URL: `/api/resources`
- Method: `GET`
- Authentication Required: YES
- Role Permission required: NONE (ALL)

Headers
```
Authorization: Bearer $TOKEN
```

Success Response (200 OK)
```
[
    {
        "uuid": "string",
        "komoditas": "string",
        "area_provinsi": "string",
        "area_kota": " string",
        "size": "string",
        "price": "string",
        "priceInUsd": "string",
        "tgl_parsed": "2022-02-21T04:45:58Z",
        "timestamp": "string"
    }
]
```

### Fetch Statistic of Resources by Province & by Weekly
- URL: `/api/resources/statistics`
- Method: `GET`
- Authentication Required: YES
- Role Permission required: ADMIN

Headers
```
Authorization: Bearer $TOKEN
```

Success Response (200 OK)
```
WIP
```

## General Constraints (Applied in all services)
Error (400 Bad Request)
```
{
    "code": "string",
    "message": "string"
}
```

Error (500 Bad Request)
```
{
    "code": "string",
    "message": "string"
}
```

# TODO
Not implemented yet because the time is up
* [x] Logger
* [x] Unit Test
* [x] Deploy to Heroku
* [x] Validator for parameter