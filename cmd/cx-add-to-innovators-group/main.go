package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/hokiegeek/keycloak-stuff"
)

func main() {
	configFilePtr := flag.String("config", "", "the yaml file with the Keycloak config")
	usersFilePtr := flag.String("users", "", "the CSV with the users to add")

	flag.Parse()

	// Load the config
	cfg, err := keycloak_stuff.LoadConfigFromFile(*configFilePtr)
	if err != nil {
		panic(err)
	}

	// Load the list of users to create
	usersFile, err := os.Open(*usersFilePtr)
	if err != nil {
		panic(err)
	}
	defer usersFile.Close()

	users, err := csv.NewReader(usersFile).ReadAll()
	if err != nil {
		panic(err)
	}

	groupName := "sonatype-innovators"

	fmt.Printf("Adding users to '%s' group:\n", groupName)

	userEmails := make([]string, len(users[1:]))
	for i, user := range users[1:] {
		userEmails[i] = user[1]
	}

	k, err := keycloak_stuff.New(cfg)
	if err != nil {
		panic(err)
	}

	if err := keycloak_stuff.AddUsersToGroup(k, userEmails, groupName); err != nil {
		fmt.Println(errors.Unwrap(err))
	}
}
