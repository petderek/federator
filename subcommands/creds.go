package subcommands

import (
	"flag"
	"io"
)

// CredsSubcommandParsedArgs holds all the bits of data that are
// needed for the link subcommand to work properly.
type CredsSubcommandParsedArgs struct {
	RoleArn, ExternalID, Region string
	OutputJSON                  bool
}

// CredsSubcommand holds the parsed args, when populated as well as internal
// state that is needed to make this work.
type CredsSubcommand struct {
	Parsed     CredsSubcommandParsedArgs
	subcommand *flag.FlagSet
}

func newCredsSubcommandParsedArgs() CredsSubcommandParsedArgs {
	return CredsSubcommandParsedArgs{}
}

// NewCredsSubcommand creates an empty container for all the
// data that will be set up by calling .Setup and wil be populated by
// calling .Parse.
func NewCredsSubcommand() CredsSubcommand {
	return CredsSubcommand{
		Parsed: newCredsSubcommandParsedArgs(),
	}
}

// Setup will setup the subcommand with flags and descriptions.
func (cmd *CredsSubcommand) Setup() {
	cmd.subcommand = flag.NewFlagSet("creds", flag.ExitOnError)

	cmd.subcommand.StringVar(
		&cmd.Parsed.RoleArn,
		"role-arn",
		"",
		"the role arn to assume for federating with AWS",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.ExternalID,
		"external-id",
		"",
		"the external ID that can optionally be provided if the assume role requires it",
	)
	cmd.subcommand.StringVar(
		&cmd.Parsed.Region,
		"region",
		"",
		"the region to make the call against, will be read from the CLI config if omitted",
	)
	cmd.subcommand.BoolVar(
		&cmd.Parsed.OutputJSON,
		"json",
		false,
		"output results as JSON rather than plain text",
	)
}

// Parse will parse the flags, according to the arguments setup in .Setup
func (cmd CredsSubcommand) Parse(args []string) error {
	return cmd.subcommand.Parse(args)
}

// SetOutput is a mirror of flag.FlagSet.SetOutput
func (cmd CredsSubcommand) SetOutput(output io.Writer) {
	cmd.subcommand.SetOutput(output)
}

// PrintDefaults is a mirror of flag.FlagSet.PrintDefaults
func (cmd CredsSubcommand) PrintDefaults() {
	cmd.subcommand.PrintDefaults()
}
