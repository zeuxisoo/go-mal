package cmd

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/mal/modules/setting"
)

var CmdServe = cli.Command{
    Name       : "serve",
    Usage      : "Start mal fake server",
    Description: `Mal fake server is the only thing you need to run`,
    Action     : runServe,
    Flags      : []cli.Flag{
        stringFlag("address, a", "127.0.0.1", "Custom address for fake server"),
        intFlag("port, p", 25, "Custom port for fake server"),
    },
}

func runServe(ctx *cli.Context) error {
    setting.NewSetting()

    if ctx.IsSet("address") {
        setting.Address = ctx.String("address")
    }

    if ctx.IsSet("port") {
        setting.Port = ctx.Int("port")
    }

    return nil
}
