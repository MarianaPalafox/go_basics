package main

import (
	"fmt"
	"time"
)

// everything that starts with upper case is accesible from outside the package.
type Budget struct {
	CampaignID string
	Balance    float64 // USD
	Expires    time.Time
}

func main() {
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)

	fmt.Printf("%#v\n", b1) // main.Budget{CampaignID:"Kittens", Balance:22.3, Expires:time.Date(2022, time.August, 8, 16, 55, 16, 578889000, time.Local)}

	fmt.Println(b1.CampaignID)

	b2 := Budget{
		Balance:    19.3,
		CampaignID: "Puppies",
	}
	fmt.Printf("%#v\n", b2)

	var b3 Budget // it creates the struct with the zero values of each type
	fmt.Printf("%#v\n", b3)
}
