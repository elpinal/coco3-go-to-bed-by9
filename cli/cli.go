package cli

import (
	"fmt"
	"time"

	"github.com/elpinal/coco3/cli"
)

type CLI cli.CLI

func (c *CLI) Run(args []string) int {
	if h := time.Now().Hour(); h >= 21 || h < 6 {
		fmt.Println("Don't be awake during 09:00 p.m. .. 06:00 a.m.")
		// This return value is an arbitrary number to distinguish from
		// any other errors ((*cli.CLI).Run may not return it.)
		return 3
	}
	return (*cli.CLI)(c).Run(args)
}
