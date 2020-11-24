package cmd

import (
	"context"
	"github.com/ermos/cli"
	"github.com/ermos/gomon/internal/builder"
	"github.com/ermos/gomon/internal/watcher"
	"regexp"
)

type DevHandler struct {}

func (DevHandler) Description(c cli.CLI) string {
	return "Start live server that allows to reload your golang application"
}

func (DevHandler) Run(ctx context.Context, c cli.CLI) error {
	var dir, ext []string
	if c.Options["dir"] != nil {
		dir = parseList(c.Options["dir"][0])
	}
	if c.Options["ext"] != nil {
		ext = parseList(c.Options["ext"][0])
	}
	ch := make(chan bool)
	watcher.Watch(ch, dir, ext)
	builder.Build(ch)
	return nil
}

func parseList(list string) (result []string) {
	var re = regexp.MustCompile(`(?m)(.+?)(?:,|$)`)
	for _, match := range re.FindAllString(list, -1) {
		result = append(result, match)
	}
	return
}