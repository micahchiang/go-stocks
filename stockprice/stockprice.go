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
            res, err := json.Marshal(quote)
            if err != nil {
                fmt.Printf("error converting json: %s. \n", err)
            }
            fmt.Println(string(res))
       }
    }

}

