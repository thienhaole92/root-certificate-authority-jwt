package cmd

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/thienhaole92/root-certificate-authority-jwt/context"
	"github.com/thienhaole92/root-certificate-authority-jwt/util"
)

type GenerateCertCmd struct {
	Key           string `arg:"" required:"" name:"key" help:"Path to CA key file." type:"path"`
	ValidityBound int    `arg:"" required:"" name:"validity-bound" help:"The length of time in day the certificate can be used." type:"int"`
	Name          string `arg:"" required:"" name:"name" help:"The certificate name."`
	Organization  string `arg:"" optional:"" name:"organization" type:"string"`
	Country       string `arg:"" optional:"" name:"country" type:"string"`
	Province      string `arg:"" optional:"" name:"province" type:"string"`
	Locality      string `arg:"" optional:"" name:"locality" type:"string"`
	StreetAddress string `arg:"" optional:"" name:"street-address" type:"string"`
	PostalCode    string `arg:"" optional:"" name:"postal-code" type:"string"`
}

func (c *GenerateCertCmd) Run(ctx *context.Context) error {
	// load the private key from the file
	kf, err := os.ReadFile(c.Key)
	if err != nil {
		return err
	}

	p, _ := pem.Decode(kf)
	if p == nil {
		return errors.New(`fail to decode private key`)
	}

	priv, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return err
	}

	// create the CA
	serial, err := util.GenerateSerial()
	if err != nil {
		return err
	}

	ca := &x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			Organization:  []string{c.Organization},
			Country:       []string{c.Country},
			Province:      []string{c.Province},
			Locality:      []string{c.Locality},
			StreetAddress: []string{c.StreetAddress},
			PostalCode:    []string{c.PostalCode},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(c.ValidityBound, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	cert, err := x509.CreateCertificate(rand.Reader, ca, ca, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	cf, err := os.Create(fmt.Sprintf(`%s.pem`, c.Name))
	if err != nil {
		return err
	}

	pem.Encode(cf, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	return cf.Close()
}
