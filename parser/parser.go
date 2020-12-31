package parser

import (
	"fmt"
	"os"

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
		SetShortDescription("Parses the semantic version").
		SetDescription("This command shows the Major, Minor, Patch, Release and Build fields of the informed version.").
		AddArgument("semver", "The semantic version", "").
		AddFlag("verbose,v", "Shows field names", commando.Bool, false).
		AddFlag("linebreak,l", "Shows one field per line", commando.Bool, false).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			semver := args["semver"].Value
			verbose, _ := flags["verbose"].GetBool()
			linebreak, _ := flags["linebreak"].GetBool()
			v, err := app.ParseSemVer(semver)
			if err != nil {
				fmt.Printf("%q is not a valid semantic version\n", semver)
				return
			}
			app.Show(os.Stdout, v, verbose, linebreak)
		})

	commando.
		Register("up").
		SetShortDescription("Adds 1 to the version").
		SetDescription("This command increases a semantic version by one, defaulting to the patch.").
		AddArgument("semver", "The semantic version", "").
		AddFlag("major,M", "Increases the major version", commando.Bool, false).
		AddFlag("minor,m", "Increases the minor version", commando.Bool, false).
		AddFlag("patch,p", "Increases the patch version", commando.Bool, true).
		AddFlag("verbose,v", "Shows field names", commando.Bool, false).
		AddFlag("linebreak,l", "Shows one field per line", commando.Bool, false).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			semver := args["semver"].Value
			major, _ := flags["major"].GetBool()
			minor, _ := flags["minor"].GetBool()
			verbose, _ := flags["verbose"].GetBool()
			linebreak, _ := flags["linebreak"].GetBool()
			v, err := app.Up(semver, major, minor)
			if err != nil {
				fmt.Printf("%q is not a valid semantic version\n", semver)
				return
			}
			app.Show(os.Stdout, v, verbose, linebreak)
		})

	commando.
		Register("down").
		SetShortDescription("Subtracts 1 to the version").
		SetDescription("This command decreases a semantic version by one, defaulting to the patch.").
		AddArgument("semver", "The semantic version", "").
		AddFlag("major,M", "Decreases the major version", commando.Bool, false).
		AddFlag("minor,m", "Decreases the minor version", commando.Bool, false).
		AddFlag("patch,p", "Decreases the patch version", commando.Bool, true).
		AddFlag("verbose,v", "Shows field names", commando.Bool, false).
		AddFlag("linebreak,l", "Shows one field per line", commando.Bool, false).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			semver := args["semver"].Value
			major, _ := flags["major"].GetBool()
			minor, _ := flags["minor"].GetBool()
			verbose, _ := flags["verbose"].GetBool()
			linebreak, _ := flags["linebreak"].GetBool()
			v, err := app.Down(semver, major, minor)
			if err != nil {
				fmt.Printf("%q is not a valid semantic version\n", semver)
				return
			}
			app.Show(os.Stdout, v, verbose, linebreak)
		})

	commando.
		Register("set").
		SetShortDescription("Sets fields").
		SetDescription("This command sets every field of the version.").
		AddArgument("semver", "The semantic version", "").
		AddFlag("major,M", "Decreases the major version", commando.String, "").
		AddFlag("minor,m", "Decreases the minor version", commando.String, "").
		AddFlag("patch,p", "Decreases the patch version", commando.String, "").
		AddFlag("release,p", "Decreases the patch version", commando.String, "").
		AddFlag("build,p", "Decreases the patch version", commando.String, "").
		AddFlag("verbose,v", "Shows field names", commando.Bool, false).
		AddFlag("linebreak,l", "Shows one field per line", commando.Bool, false).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			semver := args["semver"].Value
			major, _ := flags["major"].GetString()
			minor, _ := flags["minor"].GetString()
			patch, _ := flags["patch"].GetString()
			release, _ := flags["release"].GetString()
			build, _ := flags["build"].GetString()
			verbose, _ := flags["verbose"].GetBool()
			linebreak, _ := flags["linebreak"].GetBool()
			v, err := app.Set(semver, major, minor, patch, release, build)
			if err != nil {
				fmt.Printf("%q is not a valid semantic version\n", semver)
				return
			}
			app.Show(os.Stdout, v, verbose, linebreak)
		})
	commando.Parse(nil)
}
