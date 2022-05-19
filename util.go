package cmdlr3

import (
	"github.com/andersfylling/disgord"
)

func GetComponentsFromCommand(cmd *Command) []*disgord.MessageComponent {
	var ret []*disgord.MessageComponent

	for _, c := range cmd.Components {
		ret = append(ret, &disgord.MessageComponent{
			Type:        c.Type,
			Style:       c.Style,
			Label:       c.Label,
			CustomID:    c.CustomID,
			Url:         c.Url,
			Disabled:    c.Disabled,
			Options:     c.Options,
			Placeholder: c.Placeholder,
			MinValues:   c.MinValues,
			MaxValues:   c.MaxValues,
		})
	}

	return ret
}

func GetComponentsFromSubCommand(cmd *SubCommand) []*disgord.MessageComponent {
	var ret []*disgord.MessageComponent

	for _, c := range cmd.Components {
		ret = append(ret, &disgord.MessageComponent{
			Type:        c.Type,
			Style:       c.Style,
			Label:       c.Label,
			CustomID:    c.CustomID,
			Url:         c.Url,
			Disabled:    c.Disabled,
			Options:     c.Options,
			Placeholder: c.Placeholder,
			MinValues:   c.MinValues,
			MaxValues:   c.MaxValues,
		})
	}

	return ret
}

func checkHasSubCommands(cmd *SubCommand) bool {
	if cmd.SubCommands != nil {
		return true
	}
	return false
}

func (c *Command) ConvertSubcommandArray() []*disgord.ApplicationCommandOption {
	var opts []*disgord.ApplicationCommandOption

	if c.SubCommands == nil {
		return opts
	}

	for _, command := range c.SubCommands {
		if checkHasSubCommands(command) {
			opts = append(opts, &disgord.ApplicationCommandOption{
				Name:        command.Name,
				Description: command.Description,
				Type:        disgord.OptionTypeSubCommandGroup,
				Options:     command.ConvertSubcommandArray(),
			})

			continue
		}

		opts = append(opts, &disgord.ApplicationCommandOption{
			Name:        command.Name,
			Description: command.Description,
			Type:        disgord.OptionTypeSubCommand,
			Options:     nil,
		})
	}

	return opts
}

func (c *SubCommand) ConvertSubcommandArray() []*disgord.ApplicationCommandOption {
	var opts []*disgord.ApplicationCommandOption

	if c.SubCommands == nil {
		return opts
	}

	for _, command := range c.SubCommands {
		if checkHasSubCommands(command) {
			opts = append(opts, &disgord.ApplicationCommandOption{
				Name:        command.Name,
				Description: command.Description,
				Type:        disgord.OptionTypeSubCommandGroup,
				Options:     command.ConvertSubcommandArray(),
			})

			continue
		}

		opts = append(opts, &disgord.ApplicationCommandOption{
			Name:        command.Name,
			Description: command.Description,
			Type:        disgord.OptionTypeSubCommand,
			Options:     nil,
		})
	}

	return opts
}

func checkForSubCommand(cmd *Command, opts []*disgord.ApplicationCommandDataOption) *SubCommand {
	for _, subCmd := range cmd.SubCommands {
		if opts[0].Name == subCmd.Name {
			if subCmd.ConvertSubcommandArray() != nil {
				return checkForSubSubCommand(subCmd, opts[0].Options)
			}

			return subCmd
		}
	}

	return nil
}

func checkForSubSubCommand(cmd *SubCommand, opts []*disgord.ApplicationCommandDataOption) *SubCommand {
	for _, subCmd := range cmd.SubCommands {
		if opts[0].Name == subCmd.Name {
			if subCmd.ConvertSubcommandArray() != nil {
				return checkForSubSubCommand(cmd, opts[0].Options)
			}

			return subCmd
		}
	}

	return nil
}
