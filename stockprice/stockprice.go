// package that takes an input from stdin and retrieves the current// stock price for that symbol

package stockprice

import (
    "fmt"
    "bufio"
    "os"
)

func GetStockPrice() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter ticker: ")
    text, err := reader.ReadString('\n')

    if err != nil {
        fmt.Printf("Call failed with error: %s. \n ", err)
    } else {
       fmt.Printf("Retrieving current price for %s\n", text)
    }

}

