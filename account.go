package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

func addAccount(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username or email: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimSuffix(username, "\n") // reader adds \n which we have to remove

	for _, account := range config.Accounts {
		if account.Username == username {
			fmt.Println("You are already logged into this account!")
			return nil
		}
	}

	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Print("\n")
	if err != nil {
		return err
	}

	auth, err := auth(username, string(password))
	if err != nil {
		return err
	}
	fmt.Printf("Logged in as %s\n", auth.SelectedProfile.Name)

	config.ClientToken = auth.ClientToken
	config.Accounts = append(config.Accounts, Account{
		AuthToken: auth.AccessToken,
		Nick:      auth.SelectedProfile.Name,
		Username:  username,
	})
	return updateConfig()
}

func listAccounts(ctx *cli.Context) error {

	return nil
}
