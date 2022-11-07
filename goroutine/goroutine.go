package goroutine

import (
	"fmt"
	"time"
)

func Goroutine() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	evilNinjas := []string{"Tommy", "Jonny", "Bobby", "Andy"}

	for _, evilNinja := range evilNinjas {
		attack(evilNinja)
	}

	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second)
}
