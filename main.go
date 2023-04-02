package main

import (
	"fmt"
	"time"
)

const (
	day = time.Hour * 24
)

const (
	first_review   = day
	second_review  = day * 2
	third_review   = day * 4
	fourth_review  = day * 8
	fifth_review   = day * 16
	sixth_review   = day * 32
	seventh_review = day * 64
	eigth_review   = day * 128
	nineth_review  = day * 256
)

func main() {
	today := time.Now()
	reviews := []time.Time{}

	reviews = append(reviews, today.Add(time.Duration(first_review)))
	reviews = append(reviews, today.Add(time.Duration(second_review)))
	reviews = append(reviews, today.Add(time.Duration(third_review)))
	reviews = append(reviews, today.Add(time.Duration(fourth_review)))
	reviews = append(reviews, today.Add(time.Duration(fifth_review)))
	reviews = append(reviews, today.Add(time.Duration(sixth_review)))
	reviews = append(reviews, today.Add(time.Duration(seventh_review)))
	reviews = append(reviews, today.Add(time.Duration(eigth_review)))
	reviews = append(reviews, today.Add(time.Duration(nineth_review)))

	fmt.Println("DAY OF REVIEWS")
	for _, review := range reviews {
		fmt.Println(review.Format("02:Mon, 01:Jan of 2006"))
	}
}
