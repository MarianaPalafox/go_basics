package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("empty campaignID")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be bigger that 0")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}
	return &b, nil
}

func main() {
	expires := time.Now().Add(7 * 24 * time.Hour)
	b1, error := NewBudget("Puppies", 32.2, expires)
	if error != nil {
		fmt.Println("ERROR: ", error)
	} else {
		fmt.Printf("%#v\n", b1)
	}

	b2, error2 := NewBudget("Kittens", -32.2, expires)
	if error2 != nil {
		fmt.Println("ERROR: ", error2)
	} else {
		fmt.Printf("%#v\n", b2)
	}
}
