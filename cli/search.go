package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

func newSearchCmd(a *App) *cobra.Command {
	var typ string
	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Keyword search across public posts",
		Long: `Replay the logged-out search query for a keyword.

This path depends on a rotating doc_id. When the current doc_id does not expose
search to anonymous callers, the stream ends honestly after what it found.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func() { _ = a.Out.Flush() }()
			ctx := cmd.Context()
			query := strings.Join(args, " ")
			a.progress("searching %q", query)
			for r, err := range a.Client.Search(ctx, query, a.Limit) {
				if err != nil {
					return err
				}
				if err := a.Out.Emit(searchRow(&r)); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&typ, "type", "top", "top|recent")
	return cmd
}
