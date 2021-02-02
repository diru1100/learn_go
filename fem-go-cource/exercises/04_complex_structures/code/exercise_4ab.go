package main

import "fmt"

func average (f1, f2, f3 float64) (fAvg float64) {

	fAvg = (f1+f2+f3)/3

	return
}

// doing the same with spread thing

func average1 (fNums ... float64) (fAvg float64) {

	
	for _,fVal := range fNums {
		fAvg += fVal
	}
	fAvg /= float64(len(fNums))
	return
}

func main() {

	fmt.Println(average(1.5, 2.5, 2.0))
	fmt.Println(average1(1.5, 1.5))

}