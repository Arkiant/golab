# Production grade API test

## Structure

- api: Microservice code
    - handler: Rest Interface
    - service: Business Logic
    - cmd: Initializer and binary
    - models: Business models

## Docker build command

```
$ docker build -t api_production_grade .
```

## Execute docker image

```
$ docker run -d --name api -p 5000:5000 api_production_grade
```

```
http://localhost:5000/v1/api/status/health
```