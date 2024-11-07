package Trial

import (
	"fmt"
	"github.com/redis/go-redis"
)

//func connectToRedis() (redis.Conn, error) {
//	conn, err := redis.DialURL("redis://ibm_cloud_4c9e7cc5_2f91_465f_aebc_a29ec068f10c:6ebffd47568900c6d72a1b7a4eeab9213d3dc4711ad55fbcba953dec178e7e5c@a4cc9c6f-a0b6-48d2-ba94-5bf6e8aa4b99.co21lv7d0he2pp3gvq90.dev.databases.appdomain.cloud:31282") // Replace with your Redis connection string
//	if err != nil {
//		return nil, err
//	}
//	return conn, nil
//}

func connectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(client)
}

func main() {
	connectRedis()
}
