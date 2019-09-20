# Production grade API test

## Docker build command

```
$ docker build -t api_production_grade .
```

## Execute docker image

```
$ docker run -d --name api -p 5000:5000 api_production_grade
```

```
http://localhost:5000/api/v1/status/health
```