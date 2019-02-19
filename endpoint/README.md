# GMD proxy token based authentication

The token based authentication with Node.js, Express.js, MongoDB and Mongoose, Redis

## RESTful API endpoint

    http://18.223.124.20:5392

### POST `/v1/tokens`

Request token

+ Method: `POST`
+ URL: `/v1/tokens`
+ Body:

```js
{
    "gma-id": "victord",
    "username": "admin",
    "password": "admin"
}
```

#### CLI command:
curl -k -i -d '{"data": {"password": "admin", "gma-id": "victord", "username": "admin"}}' -X POST http://18.223.124.20:5392/v1/tokens

### POST `/v1/items`

Authenticate user.

+ Method: `POST`
+ URL: `/v1/items`
+ Body:

```js
{
    "startRow": 0, 
    "operationType": "fetch"
}
```

#### CLI command:
curl -k -i -H "X-Auth-Token:b8b2b634ed272bc74eb675a71bd3a41b"  -d '{"startRow": 0, "operationType": "fetch"}'-X POST http://18.223.124.20:5392/v1/arrays -v
curl -k -i -H "X-Auth-Token:b8b2b634ed272bc74eb675a71bd3a41b"  -X POST http://18.223.124.20:5392/v1/arrays -v
curl -k -i -H "X-Auth-Token:b8b2b634ed272bc74eb675a71bd3a41b"  -X POST http://18.223.124.20:5392/v1/arrays/09178950a90fb3359a000000000000000000000001 -v

### GET `/v1/items`

Get items as an authenticated user.

+ Method: `GET`
+ URL: `/v1/arrrays?X-Auth-Token:b8b2b634ed272bc74eb675a71bd3a41b`

Example of a token string: `b8b2b634ed272bc74eb675a71bd3a41b`

#### CLI command:
curl -k -i -H "X-Auth-Token:b8b2b634ed272bc74eb675a71bd3a41b" -X GET http://18.223.124.20:5392/v1/arrays/09178950a90fb3359a000000000000000000000001

## Install

`npm install`

## Run

`npm start`

## References

+ https://scotch.io/tutorials/authenticate-a-node-js-api-with-json-web-tokens
+ https://jwt.io
