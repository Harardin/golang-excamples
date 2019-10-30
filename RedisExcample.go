package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("Prg was started")
	DoSome()
	log.Println("exit")
}

// DoSome func to do some with Redis
func DoSome() {
	// Doing something with Redis
	// Connection
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// Adding some excample value to the Redis
	// Simplest excample
	/*_, err = conn.Do("SET", "key", "value") // SET is a command to set new value, if key is taken it will override its value
	if err != nil {
		fmt.Println(err)
	}*/
	resp, err := redis.String(conn.Do("GET", "key")) // Returns a value of a key from DB
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("This is responce with a val: " + resp)

	// Changing value of a key before return it

	_, err = conn.Do("SET", "key", "NewValue") // SET is a command to set new value, if key is taken it will override its value
	if err != nil {
		fmt.Println(err)
	}

	// Same thing but returns bytes array

	bt, _ := redis.Bytes(conn.Do("GET", "key"))
	fmt.Println("This is value from bytes: " + string(bt))
	// etc..
}
