package main

import (
	"fmt"
	"sort"
	"strings"
)

type Item struct {
	ID int
}

func uniqueItems(items []Item) []Item {
	itemMap := make(map[int]bool)
	var uniqueItems []Item

	for _, item := range items {
		if !itemMap[item.ID] {
			itemMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}

	sortItems(uniqueItems)

	return uniqueItems
}

func sortItems(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
}

func parseInput(input string) []Item {
	var items []Item
	parts := strings.Split(input, ",")
	for _, part := range parts {
		var id int
		fmt.Sscanf(part, "%d", &id)
		items = append(items, Item{ID: id})
	}
	return items
}

func main() {
	var input string
	fmt.Println("Введіть значення для слайсу (через кому):")
	fmt.Scanln(&input)

	items := parseInput(input)

	unique := uniqueItems(items)

	fmt.Println("Унікальні значення:")
	for _, item := range unique {
		fmt.Printf("{%d} ", item.ID)
	}
	fmt.Println()
}

