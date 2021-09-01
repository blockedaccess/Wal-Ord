package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func ReadLine(path string) (string, error) {

	fmt.Println("(`\\ .-') /`   ('-.                 _   .-')       ('-.      _  .-')    .-')" )
	fmt.Println("`.( OO ),'  ( OO ).-.            ( '.( OO )_    ( OO ).-. ( \\( -O )  (  OO) ")
	fmt.Println(",--./  .--.    / . --. /  ,--.       ,--.   ,--.)  / . --. /  ,------.  /     '._")
	fmt.Println("|      |  |    | \\-.  \\   |  |.-')   |   `.'   |   | \\-.  \\   |   /`. ' |'--...__)")
	fmt.Println("|  |   |  |, .-'-'  |  |  |  | OO )  |         | .-'-'  |  |  |  /  | | '--.  .--'")
	fmt.Println("|  |.'.|  |_) \\| |_.'  |  |  |`-' |  |  |'.'|  |  \\| |_.'  |  |  |_.' |    |  |")
	fmt.Println("|         |    |  .-.  | (|  '---.'  |  |   |  |   |  .-.  |  |  .  '.'    |  |")
	fmt.Println("|   ,'.   |    |  | |  |  |      |   |  |   |  |   |  | |  |  |  |\\  \\     |  |")
	fmt.Println("'--'   '--'    `--' `--'  `------'   `--'   `--'   `--' `--'  `--' '--'    `--'")
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("[%s]", path))
	fmt.Println("---------------------------------------------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println("---------------------------------------------------------------------------------")
		return scanner.Text(), nil
	} else {
		return "", errors.New("could not read scanner")
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
