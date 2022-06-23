# simple-webhook
A simple webhook server written in GO

https://hub.docker.com/repository/docker/tefaa89/simple-webhook

# Usage
Send a POST request on port 8080, should receive 200 by default with no delays.

```
curl --location --request POST 'http://127.0.0.1:8080/call' \
--header 'Content-Type: application/json' \
```
You could specfiy "waitInSenconds" for response delay and/or "statusCode" for the response status code
```
curl --location --request POST 'http://127.0.0.1:8080/call' \
--header 'Content-Type: application/json' \
--data-raw '{
    "waitInSeconds": 10,
    "statusCode": 400
}'
```
