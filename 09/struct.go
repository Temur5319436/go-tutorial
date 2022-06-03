package main

import (
	"fmt"
	"time"
)

type animal struct {
	name string
}

type banana struct {
	Color, Make          string
	Experience, MakeTime time.Time
	Width                float64
	eit                  [2]animal
}

func main() {
	//Banana struct tipida instance yaratish
	secondBanana := banana{
		Color:      "Yellow",
		Make:       "India",
		Experience: time.Now(),
		MakeTime:   time.Now(),
		Width:      100,
		eit:        [2]animal{animal{"Monkey"}, animal{"Human"}},
	}

	fmt.Println(secondBanana)
}
