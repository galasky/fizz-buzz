package controllers

import (
	"github.com/revel/revel"
	"strconv"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// return the correct value for each number
func addResult(string1, string2 string, int1, int2, number int) string {
	if number % int1 == 0 {
		if number % int2 == 0 {
			return string1 + string2
		}
		return string1
	}
	if number % int2 == 0 {
		return string2
	}
	return strconv.Itoa(number)
}

// Endpoint
func (c App) Generate(string1, string2 string, int1, int2, limit int) revel.Result {
	var results []string

	// build the result for the limit given
	for i := 1; i <= limit; i++ {
		results = append(results, addResult(string1, string2, int1, int2, i))
	}

	return c.RenderJSON(results)
}