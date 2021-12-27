package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration){
	for {
		t := time.Now()
		hour := t.Hour()
		minute := t.Minute()
		time.Sleep(delay)
		fmt.Println("Klockan är", minute, "." , hour ,":", text)
	}
	
}

func main(){

	go Remind("Dags att äta", 3*time.Hour)
	go Remind("Dags att arbeta", 8*time.Hour)
	Remind("Dags att sova", 24*time.Hour)
	select{}
}