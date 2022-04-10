package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	//rsc.io/quote=このパッケージは、くだらないことわざを収集します
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
