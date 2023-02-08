package main

import "fmt"

func buy(job_one, job_two bool) (bool, bool, bool) {
	television_50 := job_one && job_two
	television_32 := job_one != job_two
	ice_cream := job_one || job_two

	return television_50, television_32, ice_cream
}

func main() {
	television_50, television_32, ice_cream := buy(true, true)
	fmt.Printf("T (50) %v, T (32) %v, Ice Cream %v", television_50, television_32, ice_cream)
}
