package main

import (
	"fmt"
	"time"
)

func main()  {
	hourOfDay := time.Now().Hour()
	greeting := getGreeting(hourOfDay)
	fmt.Println(greeting)
	i := 0
	s := true
	for s  {
		fmt.Println(i)
		i++
		if i > 100 {
			s = false
		}

	}

}

func getGreeting(hour int) string  {
	if hour < 6 {
		return "Sleep....zzzz"
	}else if hour < 12 {
		return "Good Morning"
	}else if hour < 18 {
		return "Good Afternoon"
	}else if hour < 24{
		return "Good Evening"
	}else {
		return "Keyword Incorect"
	}


}
