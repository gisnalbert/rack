package main

import (
	"fmt"
	"strconv"

	"github.com/convox/rack/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/convox/rack/cmd/convox/stdcli"
)

func init() {
	stdcli.RegisterCommand(cli.Command{
		Name:        "instances",
		Description: "list your Convox rack's instances",
		Usage:       "",
		Action:      cmdInstancesList,
	})
}

func cmdInstancesList(c *cli.Context) {
	instances, err := rackClient(c).GetInstances()

	if err != nil {
		stdcli.Error(err)
		return
	}

	t := stdcli.NewTable("ID", "IP", "STATUS", "PROCESSES", "CPU", "MEM")

	for _, i := range instances {
		t.AddRow(i.Id, i.Ip, i.Status, strconv.Itoa(i.Processes),
			fmt.Sprintf("%0.2f%%", i.Cpu*100),
			fmt.Sprintf("%0.2f%%", i.Memory*100))
	}
	t.Print()
}
