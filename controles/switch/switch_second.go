package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Printf("Day")
	case t.Hour() < 18:
		fmt.Printf("Night")
	}
}
