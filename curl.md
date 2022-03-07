# Auth Service
## Create User
```
curl --location --request POST 'localhost:8080/api/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "1",
    "phone": "1",
    "role": "ADMIN",
    "userName": "a"
}'
```


## Login
```
curl --location --request POST 'localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phone": "4",
    "password": "u%hT"
}'
```


## Validate
```
curl --location --request GET 'localhost:8080/api/auth/validate' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYXNkYXNkYXNkIiwicGhvbmUiOiI0Iiwicm9sZSI6IkFETUlOIiwiY3JlYXRlZF9hdCI6IjIwMjItMDMtMDdUMTA6MDQ6NDkuODY1KzAwOjAwIiwidXNlck5hbWUiOiJhZG1pbiIsImlhdCI6MTY0NjY0NzQ5OSwiZXhwIjoxNjQ2NjUxMDk5fQ.NcRN7w_JziDuowqYTQj7nCWVdAAesnCB63OTGQ9adMQ'
```

# Fetch Service
## Validate
```
curl --location --request GET 'localhost:8081/api/auth/validate' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYXNkYXNkYXNkIiwicGhvbmUiOiI0Iiwicm9sZSI6IkFETUlOIiwiY3JlYXRlZF9hdCI6IjIwMjItMDMtMDdUMTA6MDQ6NDkuODY1KzAwOjAwIiwidXNlck5hbWUiOiJhZG1pbiIsImlhdCI6MTY0NjY0NzQ5OSwiZXhwIjoxNjQ2NjUxMDk5fQ.NcRN7w_JziDuowqYTQj7nCWVdAAesnCB63OTGQ9adMQ'
```


## Fetch Resources
```
curl --location --request GET 'localhost:8081/api/resources' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYmFzaWMiLCJwaG9uZSI6IjEiLCJyb2xlIjoiQkFTSUMiLCJjcmVhdGVkX2F0IjoiMjAyMi0wMy0wN1QwOTo1ODozMy40NTErMDA6MDAiLCJ1c2VyTmFtZSI6ImJhc2ljIiwiaWF0IjoxNjQ2NjQ3MjEzLCJleHAiOjE2NDY2NTA4MTN9.EeMoPJ1mfMrgFuc7bKN0z61I3DoWtmwjyIQbCZ8MgNE'
```


## Fetch Statistics of Resources
```
curl --location --request GET 'localhost:8081/api/resources/statistics' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiMSIsInBob25lIjoiNCIsInJvbGUiOiJBRE1JTiIsImNyZWF0ZWRfYXQiOiIyMDIyLTAzLTA3VDE3OjE4OjAyLjc3NSswNzowMCIsInVzZXJOYW1lIjoidiIsImlhdCI6MTY0NjY0ODI5MiwiZXhwIjoxNjQ2NjUxODkyfQ.HB67GZEMOQ5XEl7fVx8eMJxi3OvhH78pz0I--beOS5o'
```