package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-rod/rod"
	"log"
	"os"
	"strings"
	"time"
)

type OrderInfo struct {
	Email   string
	OrderID string
}

func OrderUpdater(orderInfo *OrderInfo) {
	page := rod.New().MustConnect().MustPage("https://www.walmart.com/account/trackorder").MustWindowFullscreen()
	page.MustElement("#email").MustInput(orderInfo.Email)
	page.MustElement("#fullOrderId").MustInput(orderInfo.OrderID)
	page.MustElement("#main-content > form > div.text-center > button > span > span > span").MustClick()
	time.Sleep(2 * time.Second)

	arrival := page.MustElement("#main-content > div > div.recent-orders-new > span > div.hide-content-max-m > ul > li > div > ul > li:nth-child(1) > div > div:nth-child(1) > span.shipping-status-text")
	orderPlaced := page.MustElementX("//*[@id=\"main-content\"]/div/div[2]/span/div[1]/ul/li/div/div[1]/div[1]/span[1]")
	name := page.MustElementX("//*[@id=\"show-details-recent-orders\"]/div/div/div/div[3]/div/div/div/div[1]")
	address := page.MustElementX("//*[@id=\"show-details-recent-orders\"]/div/div/div/div[3]/div/div/div/div[2]")
	product := page.MustElementX("//*[@id=\"main-content\"]/div/div[2]/span/div[2]/div/ul/li/ul/li/ul/li/div[2]/div[2]/div[1]/a/div/div")
	price := page.MustElementX("//*[@id=\"main-content\"]/div/div[2]/span/div[1]/ul/li/div/div[1]/div[1]/span[2]")
	card := page.MustElementX("//*[@id=\"show-details-recent-orders\"]/div/div/div/div[2]/div[1]/div/div/ul/li/div")
	creation := time.Now()
	reader, err := readData("WalmartOrders.csv")
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range reader {
		order := Orders{
			Email:   record[0],
			OrderID: record[1],
		}
		if !fileExists("WalmartOrders2.csv") {
			Upload, err := os.Create("WalmartOrders2.csv")
			if err != nil {
				log.Fatal("Cannot create Upload", err)
			}
			defer Upload.Close()
			writer := csv.NewWriter(Upload)
			defer writer.Flush()

			var data = [][]string{
				{"E-Mail", "OrderID", "Product", "Price", "Name", "Address", "Card Used", "Placed", "Arriving", "Updated"},
				{orderInfo.Email, orderInfo.OrderID, product.MustText(), price.MustText(), name.MustText(), address.MustText(), card.MustText(), orderPlaced.MustText(), arrival.MustText(), creation.Format("01-02-2006 15:04:05 Monday")},
			}
			for _, value := range data {
				err := writer.Write(value)
				if err != nil {
					log.Fatal("Cannot create Upload", err)
				}
			}

			fmt.Println("WalmartOrders2.csv created BOSS !!")
		} else if fileExists("WalmartOrders2.csv") {
			if orderInfo.Email == order.Email && orderInfo.OrderID == order.OrderID {
				path := "WalmartOrders2.csv"
				file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				if err != nil {
					log.Fatal(err)
				}

				defer file.Close()

				var data = [][]string{
					{orderInfo.Email, orderInfo.OrderID, product.MustText(), price.MustText(), name.MustText(), address.MustText(), card.MustText(), orderPlaced.MustText(), arrival.MustText(), creation.Format("01-02-2006 15:04:05 Monday")},
				}
				data = append(data)
				w := csv.NewWriter(file)
				w.WriteAll(data)
				fmt.Println("WalmartOrders2.csv updated BOSS !!")

			}
		}
	}
}

func OrderTrack(orderInfo *OrderInfo) {
	page := rod.New().MustConnect().MustPage("https://www.walmart.com/account/trackorder").MustWindowFullscreen()
	page.MustElement("#email").MustInput(orderInfo.Email)
	page.MustElement("#fullOrderId").MustInput(orderInfo.OrderID)
	page.MustElement("#main-content > form > div.text-center > button > span > span > span").MustClick()
	time.Sleep(2 * time.Second)
	if !(page.MustHasR("h2", "Order summary")) {
		fmt.Println("Walmart Does Not Have Your Order...(ㄒoㄒ)...")
	} else {
		arrival := page.MustElement("#main-content > div > div.recent-orders-new > span > div.hide-content-max-m > ul > li > div > ul > li:nth-child(1) > div > div:nth-child(1) > span.shipping-status-text")
		orderPlaced := page.MustElementX("//*[@id=\"main-content\"]/div/div[2]/span/div[1]/ul/li/div/div[1]/div[1]/span[1]")
		name := page.MustElementX("//*[@id=\"show-details-recent-orders\"]/div/div/div/div[3]/div/div/div/div[1]")
		address := page.MustElementX("//*[@id=\"show-details-recent-orders\"]/div/div/div/div[3]/div/div/div/div[2]")
		product := page.MustElementX("//*[@id=\"main-content\"]/div/div[2]/span/div[2]/div/ul/li/ul/li/ul/li/div[2]/div[2]/div[1]/a/div/div")
		price := page.MustElementX("//*[@id=\"main-content\"]/div/div[2]/span/div[1]/ul/li/div/div[1]/div[1]/span[2]")
		card := page.MustElementX("//*[@id=\"show-details-recent-orders\"]/div/div/div/div[2]/div[1]/div/div/ul/li/div")
		creation := time.Now()

		if fileExists("WalmartOrders.csv") {
			reader, err := readData("WalmartOrders.csv")
			if err != nil {
				log.Fatal(err)
			}
			for _, record := range reader {

				order := Orders{
					Email:   record[0],
					OrderID: record[1],
				}
				if order.Email == orderInfo.Email && order.OrderID == orderInfo.OrderID {
					fmt.Println("Already Exists !!")
					break
				}
			}
			path := "WalmartOrders.csv"
			file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()

			var data = [][]string{
				{orderInfo.Email, orderInfo.OrderID, product.MustText(), price.MustText(), name.MustText(), address.MustText(), card.MustText(), orderPlaced.MustText(), arrival.MustText(), creation.Format("01-02-2006 15:04:05 Monday")},
			}
			data = append(data)
			w := csv.NewWriter(file)
			w.WriteAll(data)
			fmt.Println("WalmartOrders.csv updated BOSS !!")
			//todo discord webhook

		} else {
			Upload, err := os.Create("WalmartOrders.csv")
			if err != nil {
				log.Fatal("Cannot create Upload", err)
			}
			defer Upload.Close()
			writer := csv.NewWriter(Upload)
			defer writer.Flush()

			var data = [][]string{
				{"E-Mail", "OrderID", "Product", "Price", "Name", "Address", "Card Used", "Placed", "Arriving", "Updated"},
				{orderInfo.Email, orderInfo.OrderID, product.MustText(), price.MustText(), name.MustText(), address.MustText(), card.MustText(), orderPlaced.MustText(), arrival.MustText(), creation.Format("01-02-2006 15:04:05 Monday")},
			}
			for _, value := range data {
				err := writer.Write(value)
				if err != nil {
					log.Fatal("Cannot create Upload", err)
				}
			}

			fmt.Println("WalmartOrders.csv created BOSS !!")
		}
	}
}

func continueAdding(answer string) {
	answer = strings.ToLower(answer)
	if answer == "no" || answer == "n" {
		return
	} else if answer == "" {
		return
	} else if !(answer == "no" || answer == "yes" || answer == "n" || answer == "y") {
		return
	} else if answer == "yes" || answer == "y" {
		email, err := ReadLine("Input E-Mail Below (xxxxx@xxxxxx.com)")
		orderID, err := ReadLine("Input OrderID Below (xxxxxxx-xxxxxx)")
		if err != nil {
			log.Fatalln(err)
		}
		switch true {

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
			cont, _ := ReadLine("Continue? (yes/no) (y/n) ")
			continueAdding(cont)
		}
	}
}
