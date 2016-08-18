package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // set a seed
	origin := rand.Intn(5) + 1 // 1-5
	number := 0
	for {
		fmt.Print("let's guess:")
		_, err := fmt.Scanln(&number)
		if err != nil {
			panic("Cannot scan a line of string." + err.Error())
		}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		if number == origin { // compare
			fmt.Println("ok")
			break
		} else {
			if number > origin {
				fmt.Println("too big")
			} else {
				fmt.Println("too small")
			}
		}
	}
}
