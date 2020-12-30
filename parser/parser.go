package parser

import (
	"github.com/thatisuday/commando"
)

// Parse executes command line arguments
func Parse() {
	commando.
		SetExecutableName("semver").
		SetVersion("1.0.0").
		SetDescription("This tool manages semantic versioning.")

	commando.
		Register(nil).
		AddArgument("version,v", "Version", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			version := args["version"].Value
			println(version)
		})
	commando.Parse(nil)
}
