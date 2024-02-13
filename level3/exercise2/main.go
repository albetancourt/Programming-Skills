package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const distanceFilePath string = "distances_in_km.csv"
const minimumPackages int = 100
const maximumPackages int = 500
const averageTruckVelocityKMH float64 = 80.0

var sourceState string
var destinationState string
var numberPackages int

type distanceMap map[string]map[string]float64

var distances distanceMap

func main() {
	parseArgs()

	err := validateInput()
	if err != nil {
		log.Fatal(err)
	}

	loadDistances()
	displayShippingInfo()
}

func displayShippingInfo() {
	fmt.Printf("Origin: %s\n", sourceState)
	fmt.Printf("Destination: %s\n", destinationState)
	fmt.Printf("Distance: %.2f km\n", distances[sourceState][destinationState])
	fmt.Printf("Shipping price: $%.2f USD\n", calculateShippingPrice())
	fmt.Printf("Estimated delivery time: %.2f hours\n", calculateDeliveryTimeinHours())
}

func calculateShippingPrice() float64 {
	pricePerKilometer := getPricePerKilometer()
	distanceInKilometers := distances[sourceState][destinationState]

	return float64(pricePerKilometer) * distanceInKilometers
}

func calculateDeliveryTimeinHours() float64 {
	distanceInKilometers := distances[sourceState][destinationState]

	return distanceInKilometers / averageTruckVelocityKMH
}

func getPricePerKilometer() int {
	if numberPackages > 200 {
		return 60
	}

	return 50
}

func parseArgs() {
	src := flag.String("src", "", "Source state")
	dest := flag.String("dest", "", "Destination state")
	numPackages := flag.Int("num-packages", 0, "Number of packages")
	flag.Parse()

	sourceState = *src
	destinationState = *dest
	numberPackages = *numPackages
}

func validateInput() error {
	if !(numberPackages >= minimumPackages) {
		return errors.New(fmt.Sprintf("The number of packages must be greater than %d", minimumPackages))
	}

	if !(numberPackages <= maximumPackages) {
		return errors.New(fmt.Sprintf("The number of packages must be less than %d", maximumPackages))
	}

	return nil
}

func loadDistances() {
	file, err := os.Open(distanceFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	distances = make(distanceMap)

	for _, record := range records[1:] {
		state1 := record[0]
		state2 := record[1]
		distance, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println("Error parsing distance:", err)
			continue
		}

		if _, ok := distances[state1]; !ok {
			distances[state1] = make(map[string]float64)
		}

		distances[state1][state2] = distance
	}
}
