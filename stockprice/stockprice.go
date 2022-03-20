// package that takes an input from stdin and retrieves the current
// stock price for that symbol

package stockprice

import (
    "fmt"
    "context"
    "bufio"
    "os"
    "strings"
    "encoding/json"
    finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type PriceInformation struct {
    C *float32 `json:"currentPrice, omitempty"`
    H *float32 `json:"highOfDay, omitempty"`
    L *float32 `json:"lowOfDay, omitempty"`
    O *float32 `json:"open, omitempty"`
    Pc *float32 `json:"previousClose, omitempty"`
    D *float32 `json:"change, omitempty"`
    Dp *float32 `json:"percentChange, omitempty"`
}

func GetStockPrice() {

    cfg := finnhub.NewConfiguration()
    cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FH_TOKEN"))
    finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter ticker: ")
    symbol, err := reader.ReadString('\n')

    if err != nil {
        fmt.Printf("Call failed with error: %s. \n ", err)
    } else {
       fmt.Printf("Retrieving current price for %s\n", symbol)
       quote, _, apierr := finnhubClient.Quote(context.Background()).Symbol(strings.ToUpper(strings.TrimSpace(symbol))).Execute()
       if apierr != nil {
           fmt.Printf("API call failed with error: %s. \n", apierr)
       } else {
            currentQuote := PriceInformation {
                C: quote.C,
                H: quote.H,
                L: quote.L,
                O: quote.O,
                Pc: quote.Pc,
                D: quote.D,
                Dp: quote.Dp,
            }
            res, err := json.Marshal(currentQuote)
            if err != nil {
                fmt.Printf("error converting json: %s. \n", err)
            }
            fmt.Println(string(res))
       }
    }

}

