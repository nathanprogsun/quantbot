package client

import (
    "fmt"

    "github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "quantbot",
        Short: "Quantbot - Go trading bot",
    }

    cmd.AddCommand(newVersionCmd())
    cmd.AddCommand(newBotCmd())
    cmd.AddCommand(newAccountCmd())

    return cmd
}

func newVersionCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "version",
        Short: "Show version",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("quantbot v0.1.0")
        },
    }
}

func newBotCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "bot",
        Short: "Bot management",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("bot: placeholder (start/stop/status will be added)")
        },
    }
}

func newAccountCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "account",
        Short: "Account management",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("account: placeholder (add/list/balance will be added)")
        },
    }
}
