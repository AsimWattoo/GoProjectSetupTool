package commands

import (
	"fmt"
	"os/exec"
)

type CommandExecutor struct {
	commands []Command
}

func (ce *CommandExecutor) AddCommand(name string, args ...string) error {
	fn, exists := GlobalExecutions[name]

	if !exists {
		return fmt.Errorf("%s not implemented yet", name)
	}

	ce.commands = append(ce.commands, Command{name: name, args: args, execute: fn})
	return nil
}

func (ce *CommandExecutor) AddRawCommand(name string, args ...string) error {
	command := Command{
		name: name,
		args: args,
		execute: func(arg []string) error {
			cmd := exec.Command(name, args...)
			_, err := cmd.CombinedOutput()
			return err
		},
	}

	ce.commands = append(ce.commands, command)
	return nil
}

func (ce *CommandExecutor) ExecuteCommands() error {
	for index, command := range ce.commands {
		fmt.Printf("Executing %d/%d - %s ", index+1, len(ce.commands), command.name)

		for _, arg := range command.args {
			fmt.Printf("%s ", arg)
		}

		fmt.Printf("\n")

		err := command.execute(command.args)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewCommandExecutor() *CommandExecutor {
	commandExecutor := &CommandExecutor{
		commands: []Command{},
	}
	return commandExecutor
}
