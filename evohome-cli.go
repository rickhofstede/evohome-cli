package main

import (
    "evohome"
    "fmt"
    "os"
    "github.com/abiosoft/ishell"
    "github.com/c-bata/go-prompt"
)

var client *evohome.Evohome

func main() {
    // Show authentication shell
    authShell := ishell.New()
    authShell.Println("Evohome CLI\n")

    authShell.Print("Username: ")
    username := authShell.ReadLine()

    authShell.Print("Password: ")
    password := authShell.ReadPassword()
    authShell.Close()

    // Connect to Evohome
    client = evohome.NewEvohome(username, password)
    if client == nil || !client.Initialized() {
        fmt.Println("\nConnection/authentication error")
        fmt.Println("Exiting...")
        os.Exit(0)
        return
    }

    // Show main shell
    p := prompt.New(
        mainExecutor,
        mainCompleter,
        prompt.OptionPrefix(">>> "),
        prompt.OptionTitle("Evohome CLI"),
    )
    p.Run()
}

func clientInitialized() (bool) {
    return client != nil && client.Initialized()
}
