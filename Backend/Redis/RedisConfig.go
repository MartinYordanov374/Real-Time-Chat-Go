package Redis

import (
	"github.com/redis/go-redis/v9"
)


// TODO: Replace the values below with .ENV values
var Client = redis.NewClient(&redis.Options{
	Addr: "redis:6379",
	Password: "",
	DB: 0,
})
