
# Búsqueda en multiples clientes
## Use Case

Vamos a crear un ejemplo de recuperación de un usuario logeado en el sistema, para eso necesitaremos una estructura que defina un usuario

```go
package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	Name     string `json:"name"`
}
```

La estrategia que vamos a seguir es la siguiente:

- Primero buscamos en una caché en memoria el objeto
- Si lo encontramos lo devolvemos
- Si no lo encontramos vamos a buscar a redis
- Si lo encontramos lo serializamos y lo añadimos en memoria a la vez que lo devolvemos
- Si no lo encontramos lo vamos a buscar a postgresql
- Si lo encontramos lo cacheamos en redis, lo serializamos y lo guardamos en memoria.

Para ello necesitaremos una interfaz donde implementar la búsqueda en memoria, la búsqueda en redis y la búsqueda en postgresql.

```go
package batch

type UserRepository interface {
    GetUser(username string) user.User
}
```

Implementación en memoria:

```go
package cache

type UserRepository struct {
    c *cache.ClientLRU
    ur *batch.UserRepository
}

var _ batch.UserRepository = &UserRepository{}

func NewUserRepository(c *cache.ClientLRU, ur batch.UserRepository) *UserRepository {
	return &UserRepository{
		c: c,
		ur: ur,
	}
}

func(ur *UserRepository) GetUser(username string) user.User {
    // Funcion que se ejcutará si no encuentra el objeto en memoria y lo irá a buscar en redis
    onFetch := func() (interface{}, error) {
        return ur.ur.GetUser(username)
    }

    // Buscamos el usuario en la cache LRU si está lo devolvemos

    // Si el usuario no está lo vamos a buscar a redis con la función onFetch

    // Si el valor no es vacío Serializamos el resultado de la función onFetch

    // Guardamos el resultado serializado en la caché LRU

    // Si el valor es vacío devolvemos vacío
}
```

Implementación en redis:

```go
package redis

type UserRepository struct {
    c *redis.Client
    ur *batch.UserRepository
}

var _ batch.UserRepository = &UserRepository{}

func NewUserRepository(c *redis.Client, ur batch.UserRepository) *UserRepository {
	return &UserRepository{
		c: c,
		ur: ur,
	}
}

func(ur *UserRepository) GetUser(username string) user.User {
    // Funcion que se ejcutará si no encuentra el objeto en redis y hace la búsqueda en base de datos
    onFetch := func() (interface{}, error) {
        return ur.ur.GetUser(username)
    }

    // Buscamos en redis el usuario y si está lo devolvemos

    // Si no está en redis vamos a buscarlo con la función onFetch

    // Si el resultado no es vacío guardamos en redis el resultado

    // Si es vacío devolvemos vacío
}

```

Implementación en postgres:

```go

package postgres

type UserRepository struct {
    c *postgres.Client
}

var _ batch.UserRepository = &UserRepository{}

func NewUserRepository(c *postgres.Client) *UserRepository {
	return &UserRepository{
		c: c,
	}
}

func(ur *UserRepository) GetUser(username string) user.User {

    // Buscamos el resultado en base de datos y devolvemos el resultado

    // Si no está devolvemos vacío
}

```

Ahora ya podemos utilizarlo de la siguiente manera:

```go
package main


func main() {
    userRepository:= cache.NewUserRepository(
        NewCacheLRU(),
        redis.NewUserRepository(
            NewRedisClient(),
            postgres.NewUserRepositoru(
                NewPostgresClient(),
            ),
        ),
    )

    user := userRepository.GetUser("test")

    fmt.Println("%v", user)
}

```





