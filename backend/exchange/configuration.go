package exchange

import "time"

type configuration struct {
	timeframe        int           // 5 minutes
	priceMultiplier  float64       // 1.5
	intervalDuration time.Duration // 1 minute

}

func newBaseConfiguration() *configuration {
	return &configuration{
		timeframe:        5,
		priceMultiplier:  1.5,
		intervalDuration: time.Minute,
	}
}
