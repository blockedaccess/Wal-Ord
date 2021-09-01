package main

import (
	"fmt"
	"log"
)

type Orders struct {
	Email string
	OrderID string
}

func main() {

	tasks, err := ReadLine("What do you want to do boss?] \n[1. Update WalmartOrders.csv] \n [2. Add New WalmartOrders")
	if err != nil {
		log.Fatalln(err)
	}

	switch{
	case tasks == "1":
		records, err := readData("WalmartOrders.csv")
		if err != nil {
			log.Fatal(err)
		}

		for _, record := range records {

			order := Orders{
				Email:   record[0],
				OrderID: record[1],
			}
			OrderUpdater(&OrderInfo{
					Email:   order.Email,
					OrderID: order.OrderID,
				})
				fmt.Printf("%s + %s\n", order.Email, order.OrderID)
			}
		break

	case tasks == "2":
		email, err := ReadLine("Input E-Mail Below (xxxxx@xxxxxx.com)")
		if err != nil {
			log.Fatalln(err)
		}
		orderID, err := ReadLine("Input OrderID Below (xxxxxxx-xxxxxx)")
		if err != nil {
			log.Fatalln(err)
		}

		switch {
		case validEmail(email) == false:
			fmt.Println("ε-(ｰдｰ)")
			fmt.Println("Not an E-mail bud...")
			break
		case validOrderID(orderID) == false:
			fmt.Println("(ｰдｰ)-ε")
			fmt.Println("Not an OrderID bud...")
			break
		case validEmail(email) == true && validOrderID(orderID) == true:
			OrderTrack(&OrderInfo{
				Email:   email,
				OrderID: orderID,
			})
			cont, _ := ReadLine("Continue? (yes/no) (y/n)")
			continueAdding(cont)
			break
		}

		break
	}

}
