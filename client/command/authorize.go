package command

import (
	"context"
	"os"

	"github.com/spolu/warp"
	"github.com/spolu/warp/client"
	"github.com/spolu/warp/lib/errors"
	"github.com/spolu/warp/lib/out"
)

const (
	// CmdNmAuthorize is the command name.
	CmdNmAuthorize cli.CmdName = "authorize"
)

func init() {
	cli.Registrar[CmdNmAuthorize] = NewAuthorize
}

// Authorize authorizes write access to a warp client.
type Authorize struct {
	username string
	token    string
}

// NewAuthorize constructs and initializes the command.
func NewAuthorize() cli.Command {
	return &Authorize{}
}

// Name returns the command name.
func (c *Authorize) Name() cli.CmdName {
	return CmdNmAuthorize
}

// Help prints out the help message for the command.
func (c *Authorize) Help(
	ctx context.Context,
) {
	out.Normf("\nUsage: ")
	out.Boldf("warp authorize <username_or_token>\n")
	out.Normf("\n")
	out.Normf("  Grants write access to a client of the current warp.\n")
	out.Normf("\n")
	out.Errof("  Be extra careful!")
	out.Normf(" Please make sure that the user you are granting write\n")
	out.Normf("  access to is who you think they are. An attacker could take over your machine\n")
	out.Normf("  in a split second with write access to one of your warps.\n")
	out.Normf("\n")
	out.Normf("  If the username of a user is ambiguous (multiple users connnected with the\n")
	out.Normf("  same username), you must use the associated user token, as returned by the\n")
	out.Boldf("  state")
	out.Normf(" command.\n")
	out.Normf("\n")
	out.Normf("Arguments:\n")
	out.Boldf("  username_or_token\n")
	out.Normf("    The username or token of a connected user.\n")
	out.Valuf("    guest_JpJP50EIas9cOfwo goofy\n")
	out.Normf("\n")
	out.Normf("Examples:\n")
	out.Valuf("  warp authorize goofy\n")
	out.Valuf("  warp authorize guest_JpJP50EIas9cOfwo\n")
	out.Normf("\n")
}

// Parse parses the arguments passed to the command.
func (c *Authorize) Parse(
	ctx context.Context,
	args []string,
) error {
	if os.Getenv(warp.EnvWarpUnixSocket) == "" {
		return errors.Trace(
			errors.Newf("This command is only available from inside a warp."),
		)
	}

	return nil
}

// Execute the command or return a human-friendly error.
func (c *Authorize) Execute(
	ctx context.Context,
) error {

	result, err := cli.RunLocalCommand(ctx, warp.Command{
		Type: warp.CmdTpAuthorize,
		Args: []string{c.token},
	})
	if err != nil {
		return errors.Trace(err)
	}

	OutState(ctx, result.State)

	return nil
}
