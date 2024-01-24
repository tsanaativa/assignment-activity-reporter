package main

import (
	"bufio"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/tsanaativa-vinnera/assignment-activity-reporter/app"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	app.NewApp().RunApp(scanner)
}
