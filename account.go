package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"text/tabwriter"
)

func addAccount(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username or email: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimSuffix(username, "\n") // reader adds \n which we have to remove

	if _, ok := config.Accounts[username]; ok {
		fmt.Println("You are already logged into this account!")
		return nil
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
	if auth.Error != nil && *auth.Error == "ForbiddenOperationException" {
		fmt.Println("Invalid credentials")
		return nil
	}
	fmt.Printf("Logged in as %s\n", auth.SelectedProfile.Name)

	config.ClientToken = auth.ClientToken
	config.Accounts[username] = Account{
		AuthToken: auth.AccessToken,
		Nick:      auth.SelectedProfile.Name,
	}
	config.SelectedAccount = username
	return updateConfig()
}

func listAccounts(ctx *cli.Context) error {
	if len(config.Accounts) == 0 {
		fmt.Println("You aren't logged in into any account")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	_, _ = fmt.Fprintln(w, "\tUSERNAME\tNICK")
	selectedAccount := config.Accounts[config.SelectedAccount]
	_, _ = fmt.Fprintf(w, "Selected\t%s\t%s\n", config.SelectedAccount, selectedAccount.Nick)
	for username, account := range config.Accounts {
		if username != config.SelectedAccount {
			_, _ = fmt.Fprintf(w, "\t%s\t%s\n", username, account.Nick)
		}
	}

	err := w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func selectAccount(ctx *cli.Context) error {
	if !ctx.Args().Present() {
		fmt.Println("You must specify account as an argument")
		return nil
	}

	name := ctx.Args().First()
	if _, ok := config.Accounts[name]; !ok {
		fmt.Printf("There is no account called %s\n", name)
	} else {
		if config.SelectedAccount == name {
			fmt.Printf("Account %s is already set as default\n", name)
		} else {
			config.SelectedAccount = name
			fmt.Printf("Selected %s as deafult account to use when launching Minecraft\n", name)
			return updateConfig()
		}
	}
	return nil
}

func removeAccount(ctx *cli.Context) error {
	return nil
}
