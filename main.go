package main

import (
	"math/rand"
	"time"

	"github.com/tobscher/kiss/commands"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootCmd := commands.NewRootCommand()
	rootCmd.AddCommand(commands.NewVersionCommand(name, version))
	rootCmd.AddCommand(commands.NewRunCommand())
	rootCmd.AddCommand(commands.NewRunRoleCommand())
	rootCmd.AddCommand(commands.NewTasksCommand())
	rootCmd.AddCommand(commands.NewHostsCommand())
	rootCmd.AddCommand(commands.NewRolesCommand())

	rootCmd.Execute()
}
