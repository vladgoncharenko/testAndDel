package main

import (
	"fmt"
	"tips/src/learning/algoritm/recursia"
)

//DNA_strand ("ATTGC") # return "TAACG"
//
//DNA_strand ("GTAT") # return "CATA"

func main() {

	var data = []int{4, 5, 1}
	fmt.Println(data)
	fmt.Println(recursia.QuickSort(data))

}

//var operationsMap = map[string]bool{models.OperationStatus: false, models.OperationGetArnCode: false}
var operationsMap = map[string]bool{"1": false, "2": false}
