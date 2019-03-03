package cmds

import (
	"text/template"

	"github.com/fatih/color"

	"github.com/florinutz/go-donut-challenge/pkg/app"
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/spf13/cobra"
)

func BuildTickerCmd(app *app.App) *cobra.Command {
	var tpl *template.Template

	return &cobra.Command{
		Use:   "ticker <product>",
		Short: "show a coinbase ticker",
		Long: `Current products: BCH-USD BCH-BTC BTC-GBP BTC-EUR BCH-GBP MKR-USDC BCH-EUR BTC-USD ZEC-USDC DNT-USDC 
LOOM-USDC DAI-USDC GNT-USDC ZIL-USDC MANA-USDC CVC-USDC ETH-USDC ZRX-EUR BAT-USDC ETC-EUR XRP-USD 
XRP-EUR XRP-BTC BTC-USDC ZRX-USD ETH-BTC ETH-EUR ETH-USD LTC-BTC LTC-EUR LTC-USD ETC-USD ETC-BTC 
ZRX-BTC ETC-GBP ETH-GBP LTC-GBP `,
		Args: cobra.ExactArgs(1),

		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			tplFuncMap := map[string]interface{}{
				"time": func(t coinbasepro.Time) string {
					return t.Time().Format("Mon Jan 2 15:04:05 -0700 MST 2006")
				},
				"highlight": color.New(color.FgGreen).Sprint,
			}
			tickerDisplay := `Id: {{.TradeID | highlight}}
Price: {{.Price | highlight}}
Size: {{.Size | highlight}}
Time: {{time .Time | highlight}}
Bid: {{.Bid | highlight}}
Ask: {{.Ask | highlight}}
Volume: {{.Volume | highlight}}
`
			tpl, err = template.New("ticker").Funcs(tplFuncMap).Parse(tickerDisplay)

			return
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			ticker, err := app.GetTicker(args[0])
			if err != nil {
				return err
			}

			err = tpl.Execute(app.Out, ticker)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
