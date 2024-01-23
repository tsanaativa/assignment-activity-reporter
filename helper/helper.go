package helper

import (
	"bufio"
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func RunActivityReporter() {
	exit := false
	menu := "Activity Reporter\n\n" +
		"1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit\n"

	for !exit {
		fmt.Println(menu)

		input := promptInput(scanner, "Enter menu: ")
		switch input {

		case "1":

		case "2":

		case "3":

		case "4":

		case "5":
			fmt.Println("Good bye!")
			exit = true

		default:
			printInvalidMenu()
		}
	}
}

func HandleSetup(input string) {

}

func HandleAction(input string) {

}

func HandleUpload(inputSlice []string) {

}

func HandleLike(inputSlice []string) {

}

func HandleDisplay(input string) {

}

func HandleTrending() {

}

func printInvalidMenu() {
	fmt.Println(customerror.ErrInvalidMenu)
}
