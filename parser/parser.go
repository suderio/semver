package parser

import (
	"fmt"

	"github.com/suderio/semver/app"
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
		AddArgument("semver", "Semantic Version", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			semver := args["semver"].Value
			v, err := app.ParseSemVer(semver)
			if err != nil {
				fmt.Printf("%q is not a valid semantic version\n", semver)
				return
			}
			fmt.Println(v.Major, v.Minor, v.Patch, v.Release, v.Build)
		})
	commando.Parse(nil)
}
