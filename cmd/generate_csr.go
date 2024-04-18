package cmd

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"github.com/thienhaole92/root-certificate-authority-jwt/context"
)

type GenerateCSRCmd struct {
	Key                string `arg:"" required:"" name:"key" help:"Path to key file." type:"path"`
	Name               string `arg:"" required:"" name:"name" help:"The certificate name."`
	Email              string `arg:"" required:"" name:"email" type:"string"`
	CommonName         string `arg:"" required:"" name:"common-name" type:"string"`
	Country            string `arg:"" optional:"" name:"country" type:"string"`
	Province           string `arg:"" optional:"" name:"province" type:"string"`
	Locality           string `arg:"" optional:"" name:"locality" type:"string"`
	Organization       string `arg:"" optional:"" name:"organization" type:"string"`
	OrganizationalUnit string `arg:"" optional:"" name:"organizational-unit" type:"string"`
}

func (c *GenerateCSRCmd) Run(ctx *context.Context) error {
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

	subj := pkix.Name{
		CommonName:         c.CommonName,
		Country:            []string{c.Country},
		Province:           []string{c.Province},
		Locality:           []string{c.Locality},
		Organization:       []string{c.Organization},
		OrganizationalUnit: []string{c.OrganizationalUnit},
		ExtraNames: []pkix.AttributeTypeAndValue{
			{
				Type: asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1},
				Value: asn1.RawValue{
					Tag:   asn1.TagIA5String,
					Bytes: []byte(c.Email),
				},
			},
		},
	}

	template := x509.CertificateRequest{
		Subject:            subj,
		SignatureAlgorithm: x509.SHA512WithRSA,
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, &template, priv)
	if err != nil {
		return err
	}

	// pem encode
	csrf, err := os.Create(fmt.Sprintf(`%s.csr`, c.Name))
	if err != nil {
		return err
	}

	pem.Encode(csrf, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csr})
	return csrf.Close()
}
