# API Security

Authentication and other API security mechanisms.

## Overview

This project implements various authentication mechanisms in Golang, including Basic Auth, Session Cookies, and JSON Web Tokens (JWT). It allows users to choose their preferred authentication mode by specifying the authentication method in a configuration file.

## Configuration

The authentication method and other relevant information can be specified in the `config/app/local_jwt_config.yaml` file (for JWT) or `config/app/local_basic_config.yaml` file (for Basic Auth). Here's an example configuration for JWT:

Copy code

```yaml
server:
  port: 8080
  rate_limit: 2
  retry_after: 2

https:
  enable: false
  cert_path: "certs/server/server.crt"
  key_path: "certs/server/server.key"

auth:
  mode: "jwt"
  duration_minutes: 10
  refresh_duration_minutes: 60
  signing_algorithm: HS256
```

### HTTPS Support

This project supports HTTPS, but you need to generate SSL/TLS certificates for it. For testing purposes, we have provided a script to generate self-signed certificates. To generate the certificates, run the `generateTlsCerts.sh` script located in the `scripts/` directory.

### Environment Variables

Certain environment variables need to be defined for the project to run successfully. These are specified in the `.env.dev` file and include:


```
SALT_PASSWORD_ENCRYPT=
SERVER_CONFIG_PATH=
CASBIN_CONF_FILE_PATH=
CASBIN_POLICY_FILE_PATH=
JWT_SECRET_ACC=
JWT_SECRET_REFR=
JWT_ISSUER=
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
OAUTH2_USER_INFO_ENDPOINT="https://www.googleapis.com/oauth2/v3/userinfo"
OAUTH2_REDIRECT_URL="http://localhost:8080/api/google/callback"
OAUTH2_AUTH_URL="https://accounts.google.com/o/oauth2/auth"
OAUTH2_TOKEN_URL="https://accounts.google.com/o/oauth2/token"
```
## Features
This project incorporates the following components:

### Casbin for Access Control

Casbin is used for access control in this project. The Casbin configuration file is placed in `config/casbin` and allows you to define access control rules based on the roles and permissions of users. It provides a powerful and flexible way to manage authorization and permissions within the application.

### Login with Google

This project also implements a login with Google feature using OAuth2. You can access `localhost:8080/home` and click on the button `Login with Google` to try it out.

### Repository Pattern

The project follows the repository pattern to simplify data access and improve maintainability. For simplification purposes, the token repository and user repository were implemented in memory. However, in a real-world scenario, these repositories could be replaced with database implementations.

## Routes
There are few routes implemented. Checkout specification in [docs/routes.md]()

## Running the Project

To run the project, execute the following command:

```bash
go run cmd/main.go
``` 

## Contributing

Contributions to the project are welcome!

## License

This project is licensed under the []. For more details, see the LICENSE.md file.

For any questions or issues, feel free to contact us at mikaellaferreira0@gmail.com. Happy coding!

----------
