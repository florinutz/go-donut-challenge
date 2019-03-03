package cmds

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/fatih/color"

	"github.com/florinutz/go-donut-challenge/pkg/app"
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/spf13/cobra"
)

// todo there are a lot of optimisations left to do here
// e.g. grouping params, validation
func BuildOrderCmd(app *app.App) *cobra.Command {
	var tpl *template.Template

	var newOrder coinbasepro.Order

	requiredEnvVars := []string{"COINBASE_PRO_PASSPHRASE", "COINBASE_PRO_KEY", "COINBASE_PRO_SECRET"}

	errIfEnvVarMissing := func(name string) error {
		if _, ok := os.LookupEnv(name); !ok {
			return fmt.Errorf("Env var %s is required\n\nAll required env vars: %s\n", name,
				strings.Join(requiredEnvVars, ", "))
		}

		return nil
	}

	cmd := &cobra.Command{
		Use:   "order",
		Short: "creates an order",
		Long: fmt.Sprintf("Simple wrapper for %s\nRequired env vars: %s\nMore details: %s",
			color.New(color.FgGreen).Sprint("https://docs.pro.coinbase.com/#orders"),
			strings.Join(requiredEnvVars, ", "),
			color.New(color.FgGreen).Sprint("https://github.com/preichenberger/go-coinbasepro#setup"),
		),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			for _, envVar := range requiredEnvVars {
				if err = errIfEnvVarMissing(envVar); err != nil {
					return
				}
			}

			tickerDisplay := `{{.}}` // todo elaborate display of the new order

			tpl, err = template.New("ticker").Parse(tickerDisplay)

			return
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			order, err := app.PlaceOrder(&newOrder)
			if err != nil {
				return err
			}

			_, _ = fmt.Fprintf(app.Out, "Order placed\n\n")

			err = tpl.Execute(app.Out, order)
			if err != nil {
				return err
			}

			return nil
		},
	}

	// common :

	cmd.Flags().StringVar(&newOrder.ClientOID, "client-oid", "",
		"[optional] Order ID selected by you to identify your order")

	cmd.Flags().StringVar(&newOrder.Type, "type", "limit", "order type (limit or market)")

	cmd.Flags().StringVar(&newOrder.Side, "side", "buy", "order side (buy or sell)")

	cmd.Flags().StringVarP(&newOrder.ProductID, "product-id", "p", "",
		"a valid product id (see the products subcommand)")

	cmd.Flags().StringVar(&newOrder.Stp, "stp", "",
		"[optional] Self-trade prevention flag")

	cmd.Flags().StringVar(&newOrder.Stop, "stop", "",
		"[optional] Either loss or entry. Requires stop_price to be defined.")

	cmd.Flags().StringVar(&newOrder.StopPrice, "stop-price", "",
		"[optional] Only if stop is defined. Sets trigger price for stop order.")

	// limit :

	cmd.Flags().StringVar(&newOrder.Price, "price", "",
		"[optional] Only if stop is defined. Sets trigger price for stop order.")

	cmd.Flags().StringVar(&newOrder.Size, "size", "",
		"Amount of BTC to buy or sell")

	cmd.Flags().StringVarP(&newOrder.TimeInForce, "time-in-force", "f", "GTC",
		"[optional] GTC, GTT, IOC, or FOK (default is GTC)")

	cmd.Flags().StringVar(&newOrder.CancelAfter, "cancel-after", "",
		"[optional] min, hour, day. Requires time_in_force to be GTT")

	cmd.Flags().BoolVar(&newOrder.PostOnly, "post-only", false,
		"[optional] Post only flag. Invalid when time_in_force is IOC or FOK")

	// market :

	// * One of size or funds is required.

	// size 	[optional]* Desired amount in BTC

	cmd.Flags().StringVar(&newOrder.Funds, "funds", "",
		"[optional] Post only flag. Invalid when time_in_force is IOC or FOK")

	return cmd
}
