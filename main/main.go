package main


import (
	"fmt"

	"github.com/tralexa/clui/demos/okro"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	o := &okro.Okro{}
	o.Main()
}
