package cmd

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thienhaole92/root-certificate-authority-jwt/context"
)

type VerifyTokenCmd struct {
	RootCert    string `arg:"" required:"" name:"root-cert" help:"Path to root cert file." type:"path"`
	ServiceCert string `arg:"" required:"" name:"service-cert" help:"Path to service cert file." type:"path"`
	Token       string `arg:"" optional:"" name:"token" help:"JWT token to verify"`
}

func (c *VerifyTokenCmd) Run(ctx *context.Context) error {
	cf, err := os.ReadFile(c.RootCert)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(cf)
	if block == nil {
		return fmt.Errorf("failed in parsing root cert string")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	certRoot := x509.NewCertPool()
	certRoot.AddCert(cert)

	pkey, err := os.ReadFile(c.ServiceCert)
	if err != nil {
		return err
	}

	pkeyblock, _ := pem.Decode(pkey)
	if pkeyblock == nil {
		return fmt.Errorf("failed in parsing service cert string")
	}
	pkeycert, err := x509.ParseCertificate(pkeyblock.Bytes)
	if err != nil {
		return err
	}

	opts := x509.VerifyOptions{
		Roots: certRoot,
	}
	if _, err := pkeycert.Verify(opts); err != nil {
		return err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pkey)
	if err != nil {
		return err
	}

	tok, err := jwt.Parse(c.Token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return err
	}

	_, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return fmt.Errorf("validate: invalid")
	}

	fmt.Println("validate: valid")

	return nil
}
