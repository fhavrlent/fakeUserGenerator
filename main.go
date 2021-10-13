package main

import (
	"strings"

	"github.com/SuperPaintman/nice/cli"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	app := cli.App{
		Name:  "fakeUserGenerator",
		Usage: cli.Usage("Generate random name, username and password."),
		Action: cli.ActionFunc(func(cmd *cli.Command) cli.ActionRunner {
			password := gofakeit.Password(true, true, true, true, false, 15)
			noun := gofakeit.Noun()
			adjective := gofakeit.Adjective()
			firstName := gofakeit.FirstName()
			lastName := gofakeit.LastName()

			return func(cmd *cli.Command) error {
				cmd.Printf("Name: %s %s\n", firstName, lastName)
				cmd.Printf("Username: %s%s%d\n", adjective, strings.Title(strings.ToLower(noun)), gofakeit.Number(1, 20))
				cmd.Printf("Password: %s\n", password)

				return nil
			}
		}),
		CommandFlags: []cli.CommandFlag{
			cli.HelpCommandFlag(),
			cli.VersionCommandFlag("1.0.0"),
		},
	}

	app.HandleError(app.Run())
}
