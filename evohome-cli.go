package main

import (
    "evohome"
    "flag"
    "fmt"
    "os"
    "github.com/abiosoft/ishell"
    "github.com/c-bata/go-prompt"
)

var client *evohome.Evohome

func main() {
    // Command-line arguments
    helpArg := flag.Bool("help", false, "Display help text")
    usernameArg := flag.String("username", "", "Username for authenticating to Honeywell's Evohome service")
    passwordArg := flag.String("password", "", "password for authenticating to Honeywell's Evohome service")
    flag.Parse()

    if *helpArg {
        flag.PrintDefaults()
        os.Exit(1)
    }

    // Show authentication shell
    authShell := ishell.New()
    authShell.Println("Evohome CLI\n")

    var username string
    if *usernameArg == "" {
        authShell.Print("Username: ")
        username = authShell.ReadLine()
    } else {
        username = *usernameArg
    }

    if username == "" {
        fmt.Println("Exiting...")
        os.Exit(1)
    }

    var password string
    if *passwordArg == "" {
        authShell.Print("Password: ")
        password = authShell.ReadPassword()
    } else {
        password = *passwordArg
    }

    authShell.Close()

    // Connect to Evohome
    client = evohome.NewEvohome(username, password)
    if client == nil || !client.Initialized() {
        fmt.Println("\nConnection/authentication error")
        fmt.Println("Exiting...")
        os.Exit(1)
    }

    // Show main shell
    p := prompt.New(
        executor,
        completer,
        prompt.OptionPrefix(">>> "),
        prompt.OptionTitle("Evohome CLI"),
    )
    p.Run()
}

func clientInitialized() (bool) {
    return client != nil && client.Initialized()
}
