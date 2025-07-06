package main

import (
	"fmt"
	"tests/address"
)

func main() {
	addressType := address.AddressType("Street 123")
	fmt.Println(addressType) // Output: street
}
