package cmds

import (
	"fmt"
	"text/tabwriter"

	"github.com/florinutz/go-donut-challenge/pkg/app"
	"github.com/spf13/cobra"
)

func BuildProductsCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "products",
		Short: "show coinbase products",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			products, err := app.GetProducts()
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(app.Out, 0, 0, 3, ' ', 0)

			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", "ID", "Base Max Size", "Base Min Size",
				"Quote Increment")
			_, _ = fmt.Fprint(w, "\t\t\t\t\n")

			for _, product := range products {
				_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", product.ID, product.BaseMaxSize,
					product.BaseMinSize, product.QuoteIncrement)
			}

			_ = w.Flush()

			return nil
		},
	}
}
