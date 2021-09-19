package rain

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vineetksingh/rainy/internal/config"

	"github.com/vineetksingh/rainy/internal/cmd"
	"github.com/vineetksingh/rainy/internal/cmd/build"
	"github.com/vineetksingh/rainy/internal/cmd/cat"
	consolecmd "github.com/vineetksingh/rainy/internal/cmd/console"
	"github.com/vineetksingh/rainy/internal/cmd/deploy"
	"github.com/vineetksingh/rainy/internal/cmd/diff"
	rainfmt "github.com/vineetksingh/rainy/internal/cmd/fmt"
	"github.com/vineetksingh/rainy/internal/cmd/info"
	"github.com/vineetksingh/rainy/internal/cmd/logs"
	"github.com/vineetksingh/rainy/internal/cmd/ls"
	"github.com/vineetksingh/rainy/internal/cmd/merge"
	"github.com/vineetksingh/rainy/internal/cmd/pkg"
	"github.com/vineetksingh/rainy/internal/cmd/rm"
	"github.com/vineetksingh/rainy/internal/cmd/tree"
	"github.com/vineetksingh/rainy/internal/cmd/watch"
	"github.com/vineetksingh/rainy/internal/console"
)

// Cmd is the rain command's entrypoint
var Cmd = &cobra.Command{
	Use:     "rain",
	Long:    "Rain is a command line tool for working with AWS CloudFormation templates and stacks",
	Version: config.VERSION,
}

const usageTemplate = `Usage:{{if .Runnable}}
  <cyan>{{.UseLine}}</>{{end}}{{if .HasAvailableSubCommands}}
  <cyan>{{.CommandPath}}</> [<gray>command</>]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{range $group := groups}}{{ $group }}:{{range $c := $.Commands}}{{if $c.IsAvailableCommand}}{{if eq $c.Annotations.Group $group}}
  <cyan>{{rpad $c.Name $c.NamePadding }}</> {{$c.Short}}{{end}}{{end}}{{end}}

{{end}}Other Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}{{if .Annotations.Group}}{{else}}
  <cyan>{{rpad .Name .NamePadding }}</> {{.Short}}{{end}}{{end}}{{end}}{{end}}{{if and .HasParent .HasAvailableFlags}}

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`

const stackGroup = "Stack commands"
const templateGroup = "Template commands"

func addCommand(label string, profileOptions bool, c *cobra.Command) {
	if label != "" {
		c.Annotations = map[string]string{"Group": label}
	}

	if profileOptions {
		c.Flags().StringVarP(&config.Profile, "profile", "p", "", "AWS profile name; read from the AWS CLI configuration file")
		c.Flags().StringVarP(&config.Region, "region", "r", "", "AWS region to use")
	}

	Cmd.AddCommand(c)
}

func init() {
	// Stack commands
	addCommand(stackGroup, true, cat.Cmd)
	addCommand(stackGroup, true, deploy.Cmd)
	addCommand(stackGroup, true, logs.Cmd)
	addCommand(stackGroup, true, ls.Cmd)
	addCommand(stackGroup, true, rm.Cmd)
	addCommand(stackGroup, true, watch.Cmd)

	// Template commands
	addCommand(templateGroup, false, build.Cmd)
	addCommand(templateGroup, false, diff.Cmd)
	addCommand(templateGroup, false, rainfmt.Cmd)
	addCommand(templateGroup, false, merge.Cmd)
	addCommand(templateGroup, true, pkg.Cmd)
	addCommand(templateGroup, false, tree.Cmd)

	// Other commands
	addCommand("", true, consolecmd.Cmd)
	addCommand("", true, info.Cmd)

	// Customise usage
	Cmd.Annotations = map[string]string{"Groups": fmt.Sprintf("%s|%s", stackGroup, templateGroup)}

	cobra.AddTemplateFunc("groups", func() []string {
		return []string{stackGroup, templateGroup}
	})

	oldUsageFunc := Cmd.UsageFunc()
	Cmd.SetUsageFunc(func(c *cobra.Command) error {
		Cmd.SetUsageTemplate(console.Sprint(usageTemplate))
		return oldUsageFunc(c)
	})

	Cmd.PersistentFlags().BoolVarP(&console.NoColour, "no-colour", "", false, "Disable colour output")

	cmd.AddDefaults(Cmd)
}
