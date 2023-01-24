package main

import (
	"fmt"
	"math"
)

type Play struct {
	Hamlet struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"hamlet"`
	AsLike struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"aslike"`
	Othello struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"othello"`
}

type Invoice struct {
	Customer     string `json:"customer"`
	Performances []struct {
		PlayID   string `json:"playID"`
		Audience int    `json:"audience"`
	} `json:"performances"`
}

func format() map[string]interface{
	return map[string]interface{
		"style":                 "currency",
	"currency":              "USD",
	"minimumFractionDigits": 2,
	}
}

func statement(invoice Invoice, plays Play) string {
	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %v\n", invoice.Customer)
	for _, perf := range invoice.Performances {
		play := plays[perf.(map[string]string)["playID"]]
		thisAmount := 0
		switch play["type"] {
		case "tragedy":
			thisAmount = 40000
			if perf.(map[string]interface{})["audience"].(int) > 30 {
				thisAmount += 1000 * (perf.(map[string]interface{})["audience"].(int) - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.(map[string]interface{})["audience"].(int) > 20 {
				thisAmount += 10000 + 500*(perf.(map[string]interface{})["audience"].(int)-20)
			}
			thisAmount += 300 * perf.(map[string]interface{})["audience"].(int)
		default:
			return fmt.Sprintf("unknown type: %v", play["type"])
		}
		volumeCredits += int(math.Max(float64(perf.(map[string]interface{})["audience"].(int))-30, 0))
		if play["type"] == "comedy" {
			volumeCredits += perf.(map[string]interface{})["audience"].(int) / 5
		}
		result += fmt.Sprintf(" %v: %v (%v seats)\n", play["name"], format(thisAmount/100), perf.(map[string]interface{})["audience"])
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %v\n", format(totalAmount/100))
	result += fmt.Sprintf("You earned %v credits\n", volumeCredits)
	return result
}

