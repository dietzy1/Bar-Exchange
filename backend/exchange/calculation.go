package exchange

import (
	"strconv"

	"github.com/dietzy1/Bar-Exchange/service"
)

//We have X amounts of events of people buying drinks within Y amount of time

//I could make it so that prices are are created based on the amount of purchases within a specific category
//So we need to find out how many drinks of the unique name were purchased within the timeframe and specific category
//This way we can ensure that the drinks will alternate in price within their own category
//The multiplier should scale so its bigger for the first drink and smaller for the last drink
//The multiplier should also be configurable so its possible to change how big price swings are
//The timeframe should be configurable so its possible to change how many events are counted within the timeframe - this will also affect how big the price swings are

// Function which takes in the aggregated data and calculates the price of each drink based on the configuration
func (e *exchange) calculatePrice(input []service.Beverage) []service.Beverage {

	//Count how many drinks of a unique name there are
	drinkCount := make(map[string]int)

	for _, drink := range input {
		drinkCount[drink.Name]++
	}

	// Calculate the price for each drink based on the count and configuration
	for i, beverage := range input {
		uniqueCount := drinkCount[beverage.Name]
		priceMultiplier := e.conf.priceMultiplier * float64(uniqueCount)

		// Calculate the new price for the drink
		newPrice := beverage.BasePrice + (priceMultiplier * beverage.BasePrice)

		// Update the drink's price in the input data
		input[i].Price = strconv.FormatFloat(newPrice, 'f', 2, 64)
		//convert float64 to string

	}

	return input

}
