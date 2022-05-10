package main

import "github.com/andersfylling/disgord"

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
