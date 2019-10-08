
# WIP: Microservice structure (Study Case)

En este caso de estudio vamos a crear una tienda online con microservicios, la arquitectura montada va a ser la siguiente:

- Product Management
- Order Management
- Delivery Management

```
                       +--------------------+
                +------> Product Management |
                |      +--------------------+
+-----------+   |
|           |   |
|           +---+      +------------------+
|  GraphQL  |----------> Order Management |
|           +---+      +------------------+
|           |   |
+-----------+   |
                |      +--------------------+
                +------> Delivery Management|
                       +--------------------+

```

## Microservice presentation layer

### GraphQL Microservice Frontend

```
Graphql Microservice

| - client/
	 | - client.go 	# GRPC client
| - server/ 		
	 | - server.go # GRPC server
| - schema/ 		# graphsql schema
| - resolver.go
| - models_gen.go
| - gqlgen.yml
| - generated.go

```

## Product management Domain
```
product_management/
	| - api/  		# Use Cases implementation and interface
		 | - product_management.go
	| - cmd 		# CMD (Presenter layer)
		 | - product_management/
		 	  | - main.go
	| - grpc/ 		# GRPC Service (Presenter layer)
		 |- client/
		 	 |- client.go #GRP client
		 |- server/
		 	 |- server.go #GRP server
	     |- product_management_pb
		 	 |- product_management.pb.go
			 |- product_management.proto
	| - postgres/ 	# Repository postgres implementation
		 |- product/
		     |- product_repository.go
    | - redis/ 		# Repository redis implementation
		 |- product/
		     |- product_repository.go
  	| - cache/ 		# MicroserviceRepository CacheLRU implementation
	     |- product/
		     |- product_repository.go
    | - tracing/ 	# Opentracing service implementation
    | - metrics/ 	# Prometheus service implementation
	| - product/	# Entity DDD
		 product.go
		 product_repository.go
		 product_service.go
	| - offer/		# Entity DDD
		 offer.go
		 offer_repository.go
		 offer_service.go
    | - Makefile
	| - CONTRIBUTING.md
	| - docker-compose.yml
	| - Dockerfile
	| - go.mod 				
	| - go.sum
	| - LICENSE					# License 
	| - README.md 				# Readme
```