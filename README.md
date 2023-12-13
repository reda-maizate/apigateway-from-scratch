# API Gateway from Scratch

This repository contains a microservices architecture developed using gRPC and Protobuf. The architecture restricts all microservices from direct external access and are only available via the API Gateway. Inter-microservice communication is handled through gRPC.

## Getting Started

### Prerequisites to run locally

- Docker ([installation page](https://docs.docker.com/get-docker/))

### Running locally

1. Clone the repository
2. Run the command: `make run` in the `api-gateway` directory. This will start the API Gateway and all the microservices.
3. To test the API Gateway, you can use Postman or any other REST client. The API Gateway is available at `localhost:8080`.

### Testing the API Gateway
#### User service
To create a new user, send a POST request to `localhost:8080/v1/users` with the following body:

```json
{
  "email": "<your_email>",
  "password": "<your_password>"
}
```

If the user is successfully created, you should get a response like this:

```json
{
  "auth_token": "<your_auth_token>"
}
```

---

To log in, send a POST request to `localhost:8080/v1/users/login` with the following body:

```json
{
  "email": "<your_email>",
  "password": "<your_password>"
}
```

If the user is successfully logged in, you should get a response like this:

```json
{
  "auth_token": "<your_auth_token>"
}
```

#### Notes service

To create a new note, send a POST request to `localhost:8080/v1/notes/create` with a header `Authorization` with the value `<your_auth_token>` and the following body:

```json
{
  "title": "<your_note_title>",
  "content": "<your_note_content>"
}
```

If the note is successfully created, you should get a response like this:

```json
{}
```
----
To get all notes, send a GET request to `localhost:8080/v1/notes` with a header `Authorization` with the value `<your_auth_token>`.

If the notes are successfully retrieved, you should get a response like this:

```json
{
  "notes": [
    {
      "id": "<note_id>",
      "title": "<note_title>",
      "content": "<note_content>"
    },
    {
      "id": "<note_id>",
      "title": "<note_title>",
      "content": "<note_content>"
    },
  ]
}
```

## Development

⚠️ If you make **any changes** to the **protobuf files**, you need to regenerate the gRPC stubs. To do this, run the command: `make buf` in the `api-gateway` directory.

⚠️ Same things for to the **sqlc files**, you need to regenerate the database stubs. To do this, run the command: `make sqlc` in the `api-gateway` directory.

### Prerequisites to develop

- Protobuf compiler ([installation page](https://grpc.io/docs/protoc-installation/))
- sqlc ([installation page](https://docs.sqlc.dev/en/stable/overview/install.html))

### Developing

1. Clone the repository
2. Make changes
3. Run the command: `make run` in the `api-gateway` directory. This will start the API Gateway and all the microservices.
4. To test the API Gateway, you can use Postman or any other REST client. The API Gateway is available at `localhost:8080`.

## Built With

- [Go](https://golang.org/) - Programming language
- [gRPC](https://grpc.io/) - RPC framework
- [Protobuf](https://developers.google.com/protocol-buffers) - Interface definition language
- [PostgreSQL](https://www.postgresql.org/) - Database
- [sqlc](https://docs.sqlc.dev/en/stable/) - Generate type safe Go code from SQL
- [Docker](https://www.docker.com/) - Containerization platform
- [Make](https://www.gnu.org/software/make/) - Build automation tool



Made with ❤️ by [Réda Maizate](https://www.linkedin.com/in/reda-maizate/) @ Seoul 🇰🇷