package cli

import (
	"fmt"
	"time"

	"github.com/elpinal/coco3/cli"
)

type CLI cli.CLI

func (c *CLI) Run(args []string) int {
	if h := time.Now().Hour(); h >= 21 || h <= 5 {
		fmt.Println("Don't be awake during 09:00 p.m. ... 05:00 a.m.")
		return 3
	}
	return (*cli.CLI).Run((*cli.CLI)(c), args)
}