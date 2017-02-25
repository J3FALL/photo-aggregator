package main

import (
	"fmt"
	"photo-aggregator/src/domain"
)

func main() {
	ph := new(domain.Photographer)

	fmt.Println(ph.ID)
}
