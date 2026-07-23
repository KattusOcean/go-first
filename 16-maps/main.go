package main

import "fmt"

func main() {
	// Map declaration and inzialization
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["Spain"] = "Madrid"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["France"] = "Paris"

	// Print capitals using keys
	for country := range countryCapitalMap {
		fmt.Println("Country:", country, "| Capital:", countryCapitalMap[country])
	}

	// Searches a exact match without the need of a loop
	capital, ok := countryCapitalMap["India"]
	if ok {
		fmt.Println("Country: India | Capital:", capital)
	} else {
		fmt.Println("Does not exists!")
	}

	// Delete entry
	delete(countryCapitalMap, "France")
	fmt.Println("France was deleted successfully")
	for country := range countryCapitalMap {
		fmt.Println("Country:", country, "| Capital:", countryCapitalMap[country])
	}
}
