package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to redis.
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "redis:6379",
	})
	defer client.Close()

	// Create a new lock client.
	locker := redislock.New(client)

	ctx := context.Background()
	fmt.Println("ctx:", ctx)
	fmt.Printf("ctx:%#v\n", ctx)
	fmt.Printf("ctx:%#v\n", ctx)

	// Try to obtain lock.
	lock, err := locker.Obtain(ctx, "my-key", 10*time.Second, nil)
	if err == redislock.ErrNotObtained {
		log.Fatalln("Could not obtain lock!")
	} else if err != nil {
		log.Fatalln(err)
	}

	// Don't forget to defer Release.
	defer lock.Release(ctx)
	fmt.Println("I have a lock!")

	// Sleep and check the remaining TTL.
	time.Sleep(5 * time.Second)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl > 0 {
		fmt.Println("Yay, I still have my lock!")
	}

	// Extend my lock.
	if err := lock.Refresh(ctx, 10*time.Second, nil); err != nil {
		log.Fatalln(err)
	}

	// Sleep a little longer, then check.
	time.Sleep(10 * time.Second)
	if ttl, err := lock.TTL(ctx); err != nil {
		log.Fatalln(err)
	} else if ttl == 0 {
		fmt.Println("Now, my lock has expired!")
	}

}
