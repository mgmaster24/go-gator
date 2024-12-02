package command

import (
	"fmt"

	"github.com/mgmaster24/gator/internal"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CmdMap map[string]func(*internal.State, Command) error
}

func (c *Commands) Register(name string, f func(*internal.State, Command) error) {
	c.CmdMap[name] = f
}

func (c *Commands) Run(s *internal.State, cmd Command) error {
	if val, ok := c.CmdMap[cmd.Name]; ok {
		err := val(s, cmd)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("Command %s doesn't exist!", cmd.Name)
}
