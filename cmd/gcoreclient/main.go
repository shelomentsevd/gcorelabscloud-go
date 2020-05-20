package main

import (
	"fmt"
	"os"

	flavors2 "github.com/G-Core/gcorelabscloud-go/client/flavors/v1/flavors"

	images2 "github.com/G-Core/gcorelabscloud-go/client/images/v1/images"

	instances2 "github.com/G-Core/gcorelabscloud-go/client/instances/v1/instances"

	k8s2 "github.com/G-Core/gcorelabscloud-go/client/k8s/v1/k8s"

	keypairs2 "github.com/G-Core/gcorelabscloud-go/client/keypairs/v1/keypairs"

	keystones2 "github.com/G-Core/gcorelabscloud-go/client/keystones/v1/keystones"

	limits2 "github.com/G-Core/gcorelabscloud-go/client/limits/v1/limits"

	networks2 "github.com/G-Core/gcorelabscloud-go/client/networks/v1/networks"

	"github.com/G-Core/gcorelabscloud-go/client/projects/v1/projects"

	quotas2 "github.com/G-Core/gcorelabscloud-go/client/quotas/v1/quotas"

	regions2 "github.com/G-Core/gcorelabscloud-go/client/regions/v1/regions"

	snapshots2 "github.com/G-Core/gcorelabscloud-go/client/snapshots/v1/snapshots"

	subnets2 "github.com/G-Core/gcorelabscloud-go/client/subnets/v1/subnets"

	tasks2 "github.com/G-Core/gcorelabscloud-go/client/tasks/v1/tasks"

	volumes2 "github.com/G-Core/gcorelabscloud-go/client/volumes/v1/volumes"

	"github.com/G-Core/gcorelabscloud-go/client/flags"
	"github.com/G-Core/gcorelabscloud-go/client/floatingips/v1/floatingips"
	"github.com/G-Core/gcorelabscloud-go/client/heat"
	"github.com/G-Core/gcorelabscloud-go/client/loadbalancers/v1/loadbalancers"
	"github.com/G-Core/gcorelabscloud-go/client/securitygroups/v1/securitygroups"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var AppVersion = "v0.2.11"

var commands = []*cli.Command{
	&networks2.NetworkCommands,
	&tasks2.TaskCommands,
	&keypairs2.KeypairCommands,
	&volumes2.VolumeCommands,
	&subnets2.SubnetCommands,
	&flavors2.FlavorCommands,
	&loadbalancers.LoadBalancerCommands,
	&instances2.InstanceCommands,
	&heat.HeatsCommand,
	&securitygroups.SecurityGroupCommands,
	&floatingips.FloatingIPCommands,
	&snapshots2.SnapshotCommands,
	&images2.ImageCommands,
	&regions2.RegionCommands,
	&projects.ProjectCommands,
	&keystones2.KeystoneCommands,
	&quotas2.QuotaCommands,
	&limits2.LimitCommands,
	&k8s2.ClusterCommands,
	&k8s2.ClusterPoolCommands,
}

func buildClientCommands(commands []*cli.Command) ([]*cli.Command, []cli.Flag, string) {
	clientType := os.Getenv("GCLOUD_CLIENT_TYPE")
	tokenClientUsage := fmt.Sprintf("GCloud API client\n%s", flags.TokenClientHelpText)
	passwordClientUsage := fmt.Sprintf("GCloud API client\n%s", flags.PasswordClientHelpText)
	if clientType == "client" {
		return commands, flags.TokenClientFlags, tokenClientUsage
	} else if clientType == "password" {
		return commands, flags.PasswordClientFlags, passwordClientUsage
	}
	return []*cli.Command{
		{
			Name:        "token",
			Aliases:     nil,
			Usage:       tokenClientUsage,
			Subcommands: commands,
			Flags:       flags.TokenClientFlags,
		},
		{
			Name:  "password",
			Usage: passwordClientUsage,
			Flags: flags.PasswordClientFlags,
			Before: func(c *cli.Context) error {
				return c.Set("client-type", "password")
			},
			Subcommands: commands,
		},
	}, nil, ""
}

func main() {

	flags.AddOutputFlags(commands)
	commands, appFlags, usage := buildClientCommands(commands)
	app := cli.NewApp()
	app.Version = AppVersion
	app.EnableBashCompletion = true
	app.Commands = commands
	if appFlags != nil {
		app.Flags = appFlags
	}
	if len(usage) > 0 {
		app.Usage = usage
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Errorf("Cannot initialize application: %+v", err)
		os.Exit(1)
	}
}