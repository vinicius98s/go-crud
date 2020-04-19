## GO CRUD

This is a simple repository for a simple REST API made with Go using the following stack:

- [Gin](https://github.com/gin-gonic/gin)
- [MongoDB driver](https://github.com/mongodb/mongo-go-driver)

### Getting started

You can run `make dev` if you are using a unix base os to get the project running.

If you are on Windows, you can try installing make with [chocolatey](https://chocolatey.org/install) and then running:

```sh
choco install make
make dev
```

### Routes

#### Users

- [GET] /users

  - Return all current users

- [POST] /users

  - Add a new user to database
  - Expected data:

  ```json
  {
    "name": "Example",
    "email": "example@test.com",
    "password": "example"
  }
  ```
