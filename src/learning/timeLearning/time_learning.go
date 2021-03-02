package main

import (
	"fmt"
	"time"
)

func main() {
	//getTimeNow()
	//getDateInPast()
	//formatLayout()
	setTimeInDate()
}

func getTimeNow() {
	now := time.Now()
	fmt.Println(now)                      // время сейчас
	fmt.Println(now.Format(time.RFC3339)) // вывод в определенном формате
}

// дата в прошлом
func getDateInPast() {
	now := time.Now()
	past := now.AddDate(-1, -2, 1)
	fmt.Println(past.Format(time.RFC3339))
}

func layout() {
	now := time.Now()
	year, month, day := now.Date()
	fmt.Println(year, "-", month, day)
	fmt.Println(now.Format("2006-01-02T"))
}

func formatLayout() {
	t := "2020-11-15T23:55:34Z"
	lay := "2006-01-02 15:04:05"
	tN, _ := time.Parse(time.RFC3339, t)
	fmt.Println(tN)
	fmt.Println(tN.Format(lay))
}

func setTimeInDate() {
	str := "11:00:00"
	now := time.Now()
	t1, err := time.Parse("15:04:05.99999999", str)

	fmt.Println(err, t1)

	t := time.Date(now.Year(), now.Month(), now.Day(), t1.Hour(), t1.Minute(), t1.Second(), 0, time.UTC)
	fmt.Println(t)

}
