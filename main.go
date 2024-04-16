package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	orderValuesNum = 8
)

// we dont need to store id and dealId as ints to perfom the algo.
type order struct {
	id         string
	dealId     string
	email      string
	address    string
	city       string
	state      string
	code       string
	creditCard string
}

func main() {
	var input string

	fmt.Scanln(&input)

	numOrders, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	validOrders := make(map[string]*order)
	fradulentOrders := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < numOrders; i++ {
		scanner.Scan()
		line := scanner.Text()
		orderValues := strings.Split(line, ",")
		if len(orderValues) != orderValuesNum {
			fmt.Println("invalid order number")
			os.Exit(1)
		}

		order := &order{
			id:         orderValues[0],
			dealId:     orderValues[1],
			email:      orderValues[2],
			address:    orderValues[3],
			city:       orderValues[4],
			state:      orderValues[5],
			code:       orderValues[6],
			creditCard: orderValues[7],
		}

		isInvalidOrder := false
		for _, v := range validOrders {
			normalizedInputOrderEmail := normalizeEmail(order.email)
			normalizedStoredOrderEmail := normalizeEmail(v.email)

			normalizedInputOrderAddress := normalizeAddress(order.address)
			normalizedStoredOrderAddress := normalizeAddress(v.address)

			normalizedInputOrderState := normalizeState(order.state)
			normalizedStoredOrderState := normalizeState(v.state)

			if normalizedInputOrderEmail == normalizedStoredOrderEmail && order.dealId == v.dealId && order.creditCard != v.creditCard {
				fradulentOrders = append(fradulentOrders, v.id)
				fradulentOrders = append(fradulentOrders, order.id)
				isInvalidOrder = true
				break
			}

			if (normalizedInputOrderAddress == normalizedStoredOrderAddress || order.city == v.city || normalizedInputOrderState == normalizedStoredOrderState || order.code == v.code) && order.dealId == v.dealId && order.creditCard != v.creditCard {
				fradulentOrders = append(fradulentOrders, v.id)
				fradulentOrders = append(fradulentOrders, order.id)
				isInvalidOrder = true
				break
			}

		}

		if !isInvalidOrder {
			validOrders[order.id] = order
		}
	}
	fmt.Println(strings.Join(fradulentOrders, ","))
}

func stringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func normalizeEmail(s string) string {
	s = strings.ToLower(s)
	sLen := strings.Split(s, "@")
	if len(sLen) == 2 {
		normalizedUsername := strings.ReplaceAll(sLen[0], ".", "")
		normalizedUsernameLen := strings.Split(normalizedUsername, "+")
		s = strings.Join([]string{normalizedUsernameLen[0], sLen[1]}, "@")
	}

	return s
}

func normalizeAddress(s string) string {

	s = strings.ToLower(s)

	addresesAbbreviated := map[string]string{
		"street": "st",
		"road":   "rd",
	}
	addressWords := strings.Split(s, " ")

	abbreviatedAddress := []string{}
	for _, streetWord := range addressWords {
		if v, ok := addresesAbbreviated[streetWord]; ok {
			abbreviatedAddress = append(abbreviatedAddress, v)
		} else {
			abbreviatedAddress = append(abbreviatedAddress, streetWord)
		}
	}

	return strings.Join(abbreviatedAddress, " ")

}

func normalizeState(s string) string {
	s = strings.ToLower(s)
	statesAbbreviated := []string{"il", "ca", "ny"}
	for _, state := range statesAbbreviated {
		if strings.HasPrefix(s, state) {
			s = state
		}
	}

	return s
}
