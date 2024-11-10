package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

type Body struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	MarksI   []int  `json:"marksI"`
	MarksII  []int  `json:"marksII"`
	MarksIII []int  `json:"marksIII"`
	MarksIV  []int  `json:"marksIV"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	jsonStr, _ := reader.ReadString('\n')

	b := []Body{}

	err := json.Unmarshal([]byte(jsonStr), &b)
	if err != nil {
		log.Fatalln(err)
	}

	list := make([]Body, 0)
	list = append(list, b...)

	//"name":"Джейсон", "surname":"Стетхем", "marksI":[5, 5, 5, 4], "marksII":[5, 3, 3, 5],
	//"marksIII":[4, 2, 5, 5, 5, 5], "marksIV":[5, 4, 4, 3, 5]
	//[[marksI], [marksII], [marksIII], [marksIV]]

	// [5 5 5 4 3 4 3 4 5] 4.2
	// 0.64 0.64 0.64 0.04 1.44 0.04 1.44 0.004 0.64
	//
	marksAverage := make([][]int, 4)
	for i := range list {
		marksAverage[0] = append(marksAverage[0], list[i].MarksI...)
		marksAverage[1] = append(marksAverage[1], list[i].MarksII...)
		marksAverage[2] = append(marksAverage[2], list[i].MarksIII...)
		marksAverage[3] = append(marksAverage[3], list[i].MarksIV...)
	}

	for i := range marksAverage {
		average := math.Round(SummInList(marksAverage[i])/float64(len(marksAverage[i]))*10) / 10
		differenceWithAverage := make([]float64, 0)
		for j := range marksAverage[i] {
			differenceWithAverage = append(differenceWithAverage, math.Pow((average-float64(marksAverage[i][j])), 2))
		}
		dispersia := math.Round(SummInListDouble(differenceWithAverage)/float64(len(marksAverage[i]))*10) / 10
		fmt.Printf("%.1f %.1f\n", average, dispersia)
	}
}

/*
func SummInList[T int | float64] (list []T) float64 {
	res := 0.0
	for i := range list {
		res += float64(list[i])
	}

	return res
}
*/

func SummInList(list []int) float64 {
	res := 0.0
	for i := range list {
		res += float64(list[i])
	}

	return res
}

func SummInListDouble(list []float64) float64 {
	res := 0.0
	for i := range list {
		res += float64(list[i])
	}

	return res
}
