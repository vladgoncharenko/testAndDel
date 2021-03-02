package main

import (
	"fmt"
	"time"
)

//Тикер: сработает раз в заданное кол-во времени

func tic() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("do something")
		}
	}
}
