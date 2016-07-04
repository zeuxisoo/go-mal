package main

import (
    "os"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-mal/cmd"
)

const APP_VER = "0.1.0"

func main() {
    app := cli.NewApp()
    app.Name = "Mal"
    app.Usage = "Mal Fake Service"
    app.Version = APP_VER
    app.Commands = []cli.Command{
        cmd.CmdServe,
    }
    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
