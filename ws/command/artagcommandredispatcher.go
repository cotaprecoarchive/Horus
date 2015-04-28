package command

import (
	"regexp"
	"strings"

	cmmd "github.com/CotaPreco/Horus/command"
	"github.com/CotaPreco/Horus/tag"
	wst "github.com/CotaPreco/Horus/ws/tag"
)

var (
	ARTAG_REGEXP = regexp.MustCompile("(?i)^((?:A|R)TAG)\\s([a-zA-Z0-9_\\-:\\*\\s]+)$")
)

type ARTagCommandRedispatcher struct {
	bus cmmd.CommandBus
}

func NewARTagCommandRedispatcher(bus cmmd.CommandBus) *ARTagCommandRedispatcher {
	return &ARTagCommandRedispatcher{
		bus,
	}
}

func (h *ARTagCommandRedispatcher) CanHandle(cmd cmmd.Command) bool {
	switch cmd.(type) {
	case *SimpleTextCommand:
		var c = cmd.(*SimpleTextCommand)
		return ARTAG_REGEXP.MatchString(c.CommandStr)
	}

	return false
}

func (h *ARTagCommandRedispatcher) Handle(cmd cmmd.Command) {
	var c = cmd.(*SimpleTextCommand)

	var commandAndTags = strings.Split(c.CommandStr, string([]byte{32}))

	var command = strings.ToUpper(commandAndTags[0])
	var tags []tag.Tag

	if command == "ATAG" || command == "RTAG" {
		// ...collect and accumulate tags ignoring whose is `err != nil`
		for _, t := range commandAndTags[1:] {
			tag, err := tag.NewTag(t)

			if err == nil {
				tags = append(tags, tag)
			}
		}
	}

	switch strings.ToUpper(commandAndTags[0]) {
	case "ATAG":
		h.bus.Dispatch(&wst.ATAGCommand{
			Connection: c.Connection,
			Tags:       tags,
		})
		break

	case "RTAG":
		h.bus.Dispatch(&wst.RTAGCommand{
			Connection: c.Connection,
			Tags:       tags,
		})
		break
	}
}
