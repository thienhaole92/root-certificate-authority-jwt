package cmd

import (
	"fmt"

	"github.com/thienhaole92/root-certificate-authority-jwt/context"
)

type VersionCmd struct {
}

func (c *VersionCmd) Run(ctx *context.Context) error {
	fmt.Println("v0.0.1")
	return nil
}
