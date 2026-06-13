package cli

import (
	"github.com/spf13/cobra"
	"github.com/tamnd/threads-cli/pkg/thid"
)

func newIDCmd(a *App) *cobra.Command {
	return &cobra.Command{
		Use:   "id <input>",
		Short: "Classify any Threads handle, id, shortcode, or URL (offline)",
		Long: `Classify a Threads identifier without touching the network.

Accepts a handle (@name or bare), a numeric pk, a shortcode, or a post URL, and
reports the kind, the resolved handle/shortcode/pk, and the canonical URL.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func() { _ = a.Out.Flush() }()
			for _, in := range readArgsOrStdin(args) {
				if err := a.Out.Emit(identityRow(thid.Classify(in))); err != nil {
					return err
				}
			}
			return nil
		},
	}
}
