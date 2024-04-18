package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/thienhaole92/root-certificate-authority-jwt/context"
)

type GenerateRSAKeyCmd struct {
	Size int    `arg:"" required:"" name:"size" help:"Bit size."`
	Name string `arg:"" required:"" name:"name" help:"Key name."`
}

func (c *GenerateRSAKeyCmd) Run(ctx *context.Context) error {
	priv, err := rsa.GenerateKey(rand.Reader, c.Size)
	if err != nil {
		return err
	}

	// encode the private key to the PEM format
	pbl := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}

	pkf, err := os.Create(fmt.Sprintf(`%s.pem`, c.Name))
	if err != nil {
		return err
	}

	if err := pem.Encode(pkf, pbl); err != nil {
		return err
	}

	return pkf.Close()
}
