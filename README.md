
# Go API Server

Implemented a simple API server in Go that returns json



## Tech Stack

- **Go**
- **Docker**
- **Docker Compose**
- **PostgreSQL**

## Enpoints
- `GET /`: Return "Welcome to the HomePage!”
- `GET /users`: Return all users
- `GET /users/{id}`: Returns the user with the specified id
-  `POST /users/{id}`: Create a new user
-  `PUT /users/{id}`: Update the user with the specified ID
-  `DELETE /users/{id}`: Delete a user with a specified ID




## Getting Started
- Clone Repository
``` bash
git clone git@github.com:anti-15/Go_API.git
cd repo
```
- Download Dependencies
``` bash
go mod download
```
- Launch the Docker container
```bash
docker-compose up --build
```

```bash
# success
......
golang_1  | db connected!!
```
- Once the connection to the db is confirmed, try the URL `http://localhost:8080` in your browser.You should see something like "**Welcome to the HomePage!"**.
Terminal: 
```bash
# success
golang_1  | 2023/07/08 03:31:43 Endpoint Hit: /
```

### Get all users
- Open a terminal and run the curl command

```bash
curl http://localhost:8080/users
```
Response:

```bash
# log
golang_1  | 2023/07/08 07:23:46 Endpoint Hit: users
```

```json
[
 {
  "id": 1,
  "email": "user1@example.com",
  "password": "password1"
 },
 {
  "id": 2,
  "email": "user2@example.com",
  "password": "password2"
 }
]
```


You can see that the table is created at container startup by `build/db/init/create_table.sql` and initial data is inserted by `build/db/init/insert_testdata.sql`

### Get user by ID
```bash
curl http://localhost:8080/users/1
```
Response:

```bash
# log
golang_1  | 2023/07/08 07:23:56 Endpoint Hit: users/1
```

```json
{
 "id": 1,
 "email": "user1@example.com",
 "password": "password1"
}
```



### Create user
```bash
curl -X POST -H "Content-Type: application/json" -d '{"email": "user3@example.com", "password": "password3"}' http://localhost:8080/users
```

```bash
# log
golang_1  | 2023/07/08 07:34:13 User successfully created!
```
Let's check it with the `users` endpoint

Response:
```json
[
 {
  "id": 1,
  "email": "user1@example.com",
  "password": "password1"
 },
 {
  "id": 2,
  "email": "user2@example.com",
  "password": "password2"
 },
 {
  "id": 3,
  "email": "user3@example.com",
  "password": "password3"
 }
]
```


### Update user by ID
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"email": "Update@example.com", "password": "Update"}' http://localhost:8080/users/2
```

```bash
# log
golang_1  | 2023/07/08 07:40:40 User update succeeded!
```

Let's check it with the `users` endpoint

Response:
```json
[
 {
  "id": 1,
  "email": "user1@example.com",
  "password": "password1"
 },
 {
  "id": 2,
  "email": "Update@example.com",
  "password": "Update"
 },
 {
  "id": 3,
  "email": "user3@example.com",
  "password": "password3"
 }
]
```
If a non-existent user is specified, `User not found` is returned

```json
curl -X PUT -H "Content-Type: application/json" -d '{"email": "Update@example.com", "password": "Update"}' http://localhost:8080/users/10

User not found
```

### Delete user by ID
```bash
curl -X DELETE http://localhost:8080/user/3
```

```bash
# log
golang_1  | 2023/07/08 07:43:29 User delete succeeded!
```
Let's check it with the `users` endpoint

Response:
```json
[
 {
  "id": 1,
  "email": "user1@example.com",
  "password": "password1"
 },
 {
  "id": 2,
  "email": "Update@example.com",
  "password": "Update"
 }
]
```

If a non-existent user is specified, `User not found` is returned

```bash
curl -X DELETE http://localhost:8080/user/10

User not found
```
