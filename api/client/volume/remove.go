package volume

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand(dockerCli *client.DockerCli) *cobra.Command {
	return &cobra.Command{
		Use:     "rm VOLUME [VOLUME]...",
		Aliases: []string{"remove"},
		Short:   "Remove a volume",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cli.MinRequiredArgs(args, 1, cmd); err != nil {
				return err
			}
			return runRemove(dockerCli, args)
		},
	}
}

func runRemove(dockerCli *client.DockerCli, volumes []string) error {
	client := dockerCli.Client()
	var status = 0

	for _, name := range volumes {
		if err := client.VolumeRemove(context.Background(), name); err != nil {
			fmt.Fprintf(dockerCli.Err(), "%s\n", err)
			status = 1
			continue
		}
		fmt.Fprintf(dockerCli.Err(), "%s\n", name)
	}

	if status != 0 {
		return cli.StatusError{StatusCode: status}
	}
	return nil
}
