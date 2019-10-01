
# DRAW implementation idea example

## Vehicle service

```go
// Interface
type VehicleService interface {
	GetAll() []Vehicles
	GetByID(id int32) Vehicle
}


// Implementation

type vehicleService struct {
	VehicleRepository
}


func NewVehicleService(vehicleRepository VehicleRepository) *vehicleService {
	return &vehicleService {
		vehicleRepository,
	}
}

func (v *vehicleService) GetAll() []Vehicles {
	// Business logic

	// Here we can access to v.FindAll() method and v.FindByID(id) method
}

func (v *vehicleService) GetByID(id int32) Vehicle {
	// Business logic

	// Here we can access to v.FindAll() method and v.FindByID(id) method
}
```

## Vehicle Repository

```go
// Common interface for all implementations

type VehicleRepository interface {
	FindAll() []Vehicle
	FindByID(id int32) Vehicle
}
```

### Postgres implementation
```go
// postgres implementation
type vehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) *vehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (v *vehicleRepository) FindAll() []Vehicle {
	// postgres implementation
}

func (v *vehicleRepository) FindByID(id int32) Vehicle {
	// postgres implementation
}
```

### Redis implementation
```go

// redis implementation

type vehicleRepository struct {
	client *redis.Client
}

func NewVehicleRepository(r *redis.Client) *vehicleRepository {
	return &vehicleRepository {
		client: r,
	}
}

func (v *vehicleRepository) FindAll() []Vehicle {
	// redis implementation
}

func (v *vehicleRepository) FindByID(id int32) Vehicle {
	// redis implementation
}

```

## Geolocation service

```go
// Interface
type GeolotaionService interface {
	GetByPostalCode(postalCode string) []Town
}

// Implementation
type geolocationService struct {
	GeolocationRepository
}

func NewGeolotaionService(geolocationRepository GeolocationRepository) *geolocationService {
	return &geolocationRepository{
		geolocationRepository,
	}
}

```

## Handler implementation

```go
type handler struct {
	v VehicleService,
	g GeolocationService,
}

func NewHandler(vehicleService VehicleService, geolocationService GeolocationService) *handler {
	return &handler{
		v: vehicleService,
		g: geolocationService,
	}
}
```