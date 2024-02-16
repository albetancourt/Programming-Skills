package main

import (
	"flag"
	"fmt"
	"os"
)

const bedroomCost int = 40
const bathroomCost int = 30
const squareMeterCost int = 90

type house struct {
	sizeSquareMeters int
	numBedrooms      int
	numBathrooms     int
	price            int
}

func (h *house) calculatePrice() {
	h.price = h.numBedrooms*bedroomCost + h.numBathrooms*bathroomCost + h.sizeSquareMeters*squareMeterCost
}

var houses = map[int]*house{
	1: {
		sizeSquareMeters: 200,
		numBedrooms:      3,
		numBathrooms:     2,
		price:            0,
	},
	2: {
		sizeSquareMeters: 150,
		numBedrooms:      2,
		numBathrooms:     2,
		price:            0,
	},
	3: {
		sizeSquareMeters: 100,
		numBedrooms:      2,
		numBathrooms:     1,
		price:            0,
	},
	4: {
		sizeSquareMeters: 100,
		numBedrooms:      1,
		numBathrooms:     2,
		price:            0,
	},
	5: {
		sizeSquareMeters: 80,
		numBedrooms:      1,
		numBathrooms:     1,
		price:            0,
	},
}

func main() {
	calculateHousePrices()

	housesCommand := flag.NewFlagSet("houses", flag.ExitOnError)
	listCommand := housesCommand.Bool("list", false, "List all houses")
	describeCommand := housesCommand.Int("describe", 0, "Describe a specific house by ID")

	if len(os.Args) < 2 || os.Args[1] != "houses" {
		showUsage()
	}

	housesCommand.Parse(os.Args[2:])
	if *listCommand {
		listHouses()
	} else if *describeCommand != 0 {
		describeHouse(*describeCommand)
	} else {
		showUsage()
	}
}

func showUsage() {
	fmt.Println("Usage: go run main.go houses [ -list | -describe <ID>]")
	os.Exit(1)
}

func calculateHousePrices() {
	for _, h := range houses {
		h.calculatePrice()
	}
}

func listHouses() {
	for id := range houses {
		describeHouse(id)
	}
}

func describeHouse(id int) {
	h := houses[id]

	fmt.Printf(
		"House: %d, %d square meters, %d bedrooms, %d bathrooms, Price: %d USD \n",
		id, h.sizeSquareMeters, h.numBedrooms, h.numBathrooms, h.price,
	)
}
