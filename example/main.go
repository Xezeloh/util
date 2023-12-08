package main

import (
	"fmt"
	"github.com/Xezeloh/util/pkg/function"
)

func main() {
	function.NewFirstNonZeroValueConsumeChain[error]().
		Consume(func() error {
			fmt.Println("1")
			return nil
		}).
		Consume(func() error {
			fmt.Println("2")
			return nil
		})
}
