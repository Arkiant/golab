
# WIP: Microservice structure (Study Case)

## Microservice presentation layer

### GraphQL Microservice Frontend

```
Graphql Microservice

| - client/
	 | - client.go 	# GRPC client
| - server/ 		# server endpoint - graph
	 | - server.go
| - schema/ 		# graphsql schema
| - resolver.go
| - models_gen.go
| - gqlgen.yml
| - generated.go

```

### API HTTP Microservice Frontend

```
Api Rest Microservice

| - client/
	 | - client.go 	# GRPC
| - server/ 		# server endpoint - routes
| - resolver.go 	# handlers with communication in grpc

```

## Microservice business logic

```
microservice_name/
	| - client/
		 | - client.go # GRPC client
    | - server/
    	 | - server.go # GRPC server
    | - service/	# Service interface
    | - redis/ 		# MicroserviceRepository redis implementation
    | - postgres/ 	# MicroserviceRepository postgres implementation
    | - inmem/ 		# MicroserviceRepository in memory implementation
    | - mock/ 		# MicroserviceRepository mock implmenetation
  	| - cache/ 		# MicroserviceRepository CacheLRU implementation
    | - tracing/ 	# Opentracing service implementation
    | - metrics/ 	# Prometheus service implementation
    | - microservice_name.go 	# Service interface implementation
	| - docker-compose.yml  	# Docker services implementation
	| - CONTRIBUTING.md 		# Contributing guideline
	| - Makefile 				# Make microservices commands
	| - go.mod 				
	| - go.sum
	| - Dockerfile 				# Dockerfile build configuration
	| - LICENSE					# License 
	| - README.md 				# Readme
```


## Microservice presentation + business

```
Microservice/
	| - cmd/
		 | - main.go
    | - service/	# Service interface
    | - graph/ 		# Graph Front
    	 | - schema/
    	 		| - schema.graphql
    	 | - resolver.go
    	 | - models_gen.go
    	 | - gqlgen.yml # Graphql configuration file
    	 | - generated.go
    	 | - server/
    	 	 	| - server.go
    | - api/ 		# Rest Front
    	| - server.go
    | - redis/ 		# MicroserviceRepository redis implementation
    | - postgres/ 	# MicroserviceRepository postgres implementation
    | - inmem/ 		# MicroserviceRepository in memory implementation
    | - mock/ 		# MicroserviceRepository mock implmenetation
  	| - cache/ 		# MicroserviceRepository CacheLRU implementation
    | - tracing/ 	# Opentracing service implementation
    | - metrics/ 	# Prometheus service implementation
    | - microservice.go 	# Service interface implementation
	| - docker-compose.yml  # Docker services implementation
	| - CONTRIBUTING.md 	# Contributing guideline
	| - Makefile 			# Make microservices commands
	| - go.mod 				
	| - go.sum
	| - Dockerfile 			# Dockerfile build configuration
	| - LICENSE				# License 
	| - README.md 			# Readme
```