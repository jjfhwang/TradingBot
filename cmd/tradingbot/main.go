// cmd/tradingbot/main.go
package main

import (
"flag"
"log"
"os"

"tradingbot/internal/tradingbot"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := tradingbot.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
