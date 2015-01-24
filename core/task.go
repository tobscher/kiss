package core

import (
	"os"

	"github.com/gophergala/go_ne/plugins/core"
	"errors"
)

type Task interface {
	Name() string
	Args() []string
}

func RunAll(runner Runner, config *Config) (error) {
	defer StopAllPlugins()

	for _, t := range config.Tasks {
		for _, s := range t.Steps {
			err := RunStep(runner, &s); if err != nil {
				return err;
			}
		}
	}

	return nil
}


func RunTask(runner Runner, config *Config, taskName string) (error) {
	defer StopAllPlugins()

	task, ok := config.Tasks[taskName]; if !ok {
		return errors.New("No task exists with that name")
	}
	
	for _, s := range task.Steps {
		err := RunStep(runner, &s); if err != nil {
			return err
		}
	}

	return nil
}


func RunStep(runner Runner, s *ConfigStep) (error) {
	var command *Command
	var err error
	
	if s.Plugin != nil {
		// Load plugin
		p, err := GetPlugin(*s.Plugin); if err != nil {
			return err
		}

		pluginArgs := plugin.Args{
			Environment: os.Environ(),
			Options:     s.Args,
		}

		command, err = p.GetCommand(pluginArgs); if err != nil {
			return err
		}
	} else {
		// Run arbitrary command
		command, err = NewCommand(*s.Command, s.Args); if err != nil {
			return err
		}
	}

	err = runner.Run(command); if err != nil {
		return err
	}
	
	return nil
}
