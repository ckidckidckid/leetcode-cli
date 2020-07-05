package cmd

import (
	"github.com/kavimaluskam/leetcode-cli/pkg/api"
	"github.com/kavimaluskam/leetcode-cli/pkg/arg"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(showCmd)
	showCmd.Flags().IntP("id", "i", 0, "ID of Problem to be shown")
	showCmd.Flags().StringP("title", "t", "", "Title Slug of Problem to be shown")
	showCmd.Flags().BoolP("random", "r", false, "Random choice of Problem to be shown")
}

var showCmd = &cobra.Command{
	Use:     `show`,
	Aliases: []string{`dl`, `pick`, `show`},
	Short:   `Show individual problem`,
	Args:    arg.Show,
	RunE:    show,
}

func show(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetInt("id")
	title, _ := cmd.Flags().GetString("title")
	random, _ := cmd.Flags().GetBool("random")

	client, err := api.GetAuthClient()
	if err != nil {
		return err
	}

	problemDetail, err := client.GetProblemDetail(id, title, random)
	if err != nil {
		return err
	}

	problemDetail.ExportStdoutDetail()
	return nil
}