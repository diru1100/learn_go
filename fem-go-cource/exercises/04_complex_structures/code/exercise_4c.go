package main

import "fmt"

func appendToSlice(sampleGrocerySlice[] string, groceriesToAdd ... string) ([]string) {

	
	for _,val := range groceriesToAdd {
		sampleGrocerySlice = append(sampleGrocerySlice, val)

	}
	
	return sampleGrocerySlice
}
func namePresentInSampleMap(sampleMap map[string]string, givenString string) bool {

	_, ok := sampleMap[givenString]

	return ok
}
func averageFloat(arr[] float64) (average float64) {

	for _, value := range arr {
		average+=value
	}
	average /= float64(len(arr))
	return
}

func main() {

	sampleArray := []float64 {1.2, 4.5, 6.443, 7.34, 8.90}
	fmt.Println( averageFloat(sampleArray) ) 

	sampleMap:= map[string]string{

		"jacky": "dog",
		"lacy": "cat",
		"bubbly": "octopus"}

	namePresent := "jacky"
	nameNotPresent := "Riley"
	fmt.Println(namePresentInSampleMap(sampleMap, namePresent))
	fmt.Println(namePresentInSampleMap(sampleMap, nameNotPresent))

	groceryList := []string{"tomato", "potato", "gileto", "coolato"}
	fmt.Println(groceryList," \n Before")
	addedGroceryList := appendToSlice(groceryList, "cabbage", "spinach", "bread", "nutella")
	fmt.Println(addedGroceryList, "\n After")
}