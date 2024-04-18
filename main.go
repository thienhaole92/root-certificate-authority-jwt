package main

import (
	"github.com/alecthomas/kong"
	"github.com/thienhaole92/root-certificate-authority-jwt/cmd"
	"github.com/thienhaole92/root-certificate-authority-jwt/context"
)

var CLI struct {
	Debug bool `help:"Enable debug mode."`

	Version        cmd.VersionCmd        `cmd:"" help:"Print version."`
	GenerateRSAKey cmd.GenerateRSAKeyCmd `cmd:"" help:"Generate a random RSA private key."`
	GenerateCert   cmd.GenerateCertCmd   `cmd:"" help:"Generate a certificate and sign it with the CA key."`
	GenerateCSR    cmd.GenerateCSRCmd    `cmd:"" help:"Generate a certificate signing request."`
	SignCSR        cmd.SignCSRCmd        `cmd:"" help:"Sign the CSR with the Root CA."`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&context.Context{Debug: CLI.Debug})
	ctx.FatalIfErrorf(err)
}
