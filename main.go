package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	calculate := 0
	var inputValue string
	fmt.Scanln(&inputValue)
	input := castStringToIntArray(inputValue)
	if len(input) == 4 {
		calculate = calculatePrice(adapterArrayToAction(input))
	}
	fmt.Printf(strconv.Itoa(calculate))
}

func adapterArrayToAction(parameters []int) CalculateAction {

	return CalculateAction{
		ActionValueInDay:   parameters[2],
		DaysToCalculate:    parameters[3],
		InitialActionValue: parameters[1],
		InitialDay:         parameters[0],
	}
}

func castStringToIntArray(parameters string) []int {
	value := strings.Split(parameters, " ")
	var castArray = []int{}
	for _, parameter := range value {
		castedValue, _ := strconv.Atoi(parameter)
		castArray = append(castArray, castedValue)
	}
	return castArray
}

// CalculateAction ...
type CalculateAction struct {
	InitialDay         int
	InitialActionValue int
	ActionValueInDay   int
	DaysToCalculate    int
}

func calculatePrice(action CalculateAction) int {

	valueOfAction := action.InitialActionValue
	day := action.InitialDay
	for i := 0; i < action.DaysToCalculate; i++ {
		day++
		if isPair(day) {
			valueOfAction += action.ActionValueInDay
		} else {
			valueOfAction = (valueOfAction - action.ActionValueInDay)
		}
	}
	return valueOfAction

}

func isPair(number int) bool {
	return number%2 == 0
}
