package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/activityreporter"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

var (
	scanner     = bufio.NewScanner(os.Stdin)
	socialGraph = activityreporter.NewSocialGraph()
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
			setupInput := promptInput(scanner, "Setup social graph: ")
			HandleSetup(setupInput)

		case "2":
			actionInput := promptInput(scanner, "Enter user Actions: ")
			HandleAction(actionInput)

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

func HandleSetup(input string) error {
	inputSlice := strings.Split(input, " ")

	if len(inputSlice) == 3 {

		if inputSlice[1] == "follows" {

			username1, username2 := inputSlice[0], inputSlice[2]

			user1, ok := socialGraph.IsUserExist(username1)
			if !ok {
				user1 = activityreporter.NewUser(username1)
				socialGraph.AddNewUser(user1)
			}

			user2, ok := socialGraph.IsUserExist(username2)
			if !ok {
				user2 = activityreporter.NewUser(username2)
				socialGraph.AddNewUser(user2)
			}

			err := user2.FollowedBy(user1)
			if err != nil {
				return printAndReturnError(err)
			}

			return nil
		}

		return printAndReturnError(customerror.ErrInvalidKeyword)
	}

	return printAndReturnError(customerror.ErrInvalidInput)
}

func HandleAction(input string) error {
	inputSlice := strings.Split(input, " ")

	switch len(inputSlice) {
	case 3:
		HandleUpload(inputSlice)
		return nil

	case 4:
		HandleLike(inputSlice)
		return nil

	default:
		return printAndReturnError(customerror.ErrInvalidInput)
	}
}

func HandleUpload(inputSlice []string) error {
	if inputSlice[1] == "uploaded" && inputSlice[2] == "photo" {
		username := inputSlice[0]

		val, ok := socialGraph.IsUserExist(username)
		if ok {
			val.UploadPhoto()
			return nil
		}

		return printAndReturnError(customerror.ErrUnknownUser(username))
	}

	return printAndReturnError(customerror.ErrInvalidKeyword)
}

func HandleLike(inputSlice []string) {

}

func HandleDisplay(input string) {

}

func HandleTrending() {

}

func printAndReturnError(err error) error {
	fmt.Println(err.Error())
	return err
}

func printInvalidMenu() {
	fmt.Println(customerror.ErrInvalidMenu)
}
