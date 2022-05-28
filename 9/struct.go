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
	eit                  []animal
}

func main() {
	//Banana struct tipida instance yaratish
	secondBanana := banana{
		Color:      "Yellow",
		Make:       "India",
		Experience: time.Now(),
		MakeTime:   time.Now(),
		Width:      100,
		eit:        []animal{animal{"Monkey"}, animal{"Human"}},
	}

	fmt.Println(secondBanana)
}
