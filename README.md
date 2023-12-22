# ms-user-auth
Microservice responsible for user authentication.

ms-user-auth exposes an endpoint which takes username and password verifies it against the database and creates a jwt token. This token can be utilized in other microservices to authorize the user.
