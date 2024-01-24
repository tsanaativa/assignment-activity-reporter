package app

import (
	"bufio"
	"fmt"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/activityreporter"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/customerror"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) RunApp(scanner *bufio.Scanner) error {
	socialGraph := activityreporter.NewSocialGraph()

	var occuredErr error
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
			err := HandleSetup(scanner, socialGraph)
			occuredErr = err

		case "2":
			err := HandleAction(scanner, socialGraph)
			occuredErr = err

		case "3":
			err := HandleDisplay(scanner, socialGraph)
			occuredErr = err

		case "4":
			HandleTrending(socialGraph)

		case "5":
			fmt.Println("Good bye!")
			exit = true

		default:
			occuredErr = printAndReturnError(customerror.ErrInvalidMenu)
		}
	}
	return occuredErr
}

func HandleSetup(scanner *bufio.Scanner, socialGraph *activityreporter.SocialGraph) error {
	input := promptInput(scanner, "Setup social graph: ")
	inputSlice := strings.Fields(input)

	if len(inputSlice) == 3 {

		if inputSlice[1] == "follows" {

			username1, username2 := inputSlice[0], inputSlice[2]

			user1, ok := socialGraph.IsUserExist(username1)
			if !ok {
				user1 = activityreporter.NewUser(username1, socialGraph)
				socialGraph.AddNewUser(user1)
			}

			user2, ok := socialGraph.IsUserExist(username2)
			if !ok {
				user2 = activityreporter.NewUser(username2, socialGraph)
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

	return printAndReturnError(customerror.ErrInvalidKeyword)
}

func HandleAction(scanner *bufio.Scanner, socialGraph *activityreporter.SocialGraph) error {
	input := promptInput(scanner, "Enter user Actions: ")
	inputSlice := strings.Fields(input)

	switch len(inputSlice) {
	case 3:
		return HandleUpload(inputSlice, socialGraph)

	case 4:
		return HandleLike(inputSlice, socialGraph)

	default:
		return printAndReturnError(customerror.ErrInvalidKeyword)
	}
}

func HandleUpload(inputSlice []string, socialGraph *activityreporter.SocialGraph) error {
	if inputSlice[1] == "uploaded" && inputSlice[2] == "photo" {
		username := inputSlice[0]

		val, ok := socialGraph.IsUserExist(username)
		if ok {
			err := val.UploadPhoto()
			if err != nil {
				return printAndReturnError(err)
			}
			return nil
		}

		return printAndReturnError(customerror.ErrUnknownUser(username))
	}

	return printAndReturnError(customerror.ErrInvalidKeyword)
}

func HandleLike(inputSlice []string, socialGraph *activityreporter.SocialGraph) error {
	if inputSlice[1] == "likes" && inputSlice[3] == "photo" {
		username1, username2 := inputSlice[0], inputSlice[2]

		val, ok := socialGraph.IsUserExist(username1)
		if ok {
			val2, ok2 := socialGraph.IsUserExist(username2)
			if ok2 {
				err := val2.LikedPhotoBy(val)
				if err == nil {
					return nil
				}
				return printAndReturnError(err)
			}

			return printAndReturnError(customerror.ErrUnknownUser(username2))
		}

		return printAndReturnError(customerror.ErrUnknownUser(username1))
	}

	return printAndReturnError(customerror.ErrInvalidKeyword)
}

func HandleDisplay(scanner *bufio.Scanner, socialGraph *activityreporter.SocialGraph) error {
	input := promptInput(scanner, "Display activity for: ")
	val, ok := socialGraph.IsUserExist(input)
	if ok {
		fmt.Printf("%s activities:\n", val.Username)

		for _, v := range val.ActivityReport() {
			fmt.Println(v)
		}

		return nil
	}
	return printAndReturnError(customerror.ErrUnknownUser(input))
}

func HandleTrending(socialGraph *activityreporter.SocialGraph) []*activityreporter.User {
	fmt.Println("Trending photos:")

	for i, v := range socialGraph.Trending() {
		likesCount := v.LikesCount()
		fmt.Printf("%d. %s photo got %d like", i+1, v.Username, likesCount)

		if likesCount > 1 {
			fmt.Printf("s")
		}

		fmt.Println()
	}

	return socialGraph.Trending()
}

func printAndReturnError(err error) error {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
