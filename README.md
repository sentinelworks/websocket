# GMD proxy token based authentication

The GMP proxy accepts the websocket connections from GMA. In addition, it waits at port 7999

## RESTful API endpoint

    http://18.223.124.20:7999

## Steps

    1. start GMP : nodemon server.js
    2. start GMA : in docker ubuntu root@8ec670aae94a:~/work# python3 webc.py 

start to request token and get the array info

## POST `/v1/tokens`

request a token

curl -k -i -d '{"data": {"password": "admin", "gma-id": "victord", "username": "admin"}}' -X POST http://18.223.124.20:7999

## POST `/v1/blabla

use request token

curl -k -i -H "X-Auth-Token:b8b2b634ed272bc74eb675a71bd3a41b"  -d '{"startRow": 0, "operationType": "fetch"}'-X POST http://18.223.124.20:7999/v1/arrays -v


