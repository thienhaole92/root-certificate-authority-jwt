package cmd

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"os"

	"github.com/thienhaole92/root-certificate-authority-jwt/context"
	"github.com/thienhaole92/root-certificate-authority-jwt/util"
)

type SignCSRCmd struct {
	RootCAKey     string `arg:"" required:"" name:"root-ca-key" help:"Path to root CA key." type:"path"`
	RootCACert    string `arg:"" required:"" name:"root-ca-cert" help:"Path to root CA cert." type:"path"`
	CSR           string `arg:"" required:"" name:"csr" help:"Path to CSR file." type:"path"`
	ValidityBound int    `arg:"" required:"" name:"validity-bound" help:"The length of time in day the certificate can be used." type:"int"`
	Name          string `arg:"" required:"" name:"name" help:"The certificate name."`
}

func (c *SignCSRCmd) Run(ctx *context.Context) error {
	// load the root CA private key from the file
	rcakf, err := os.ReadFile(c.RootCAKey)
	if err != nil {
		return err
	}

	rcapkp, _ := pem.Decode(rcakf)
	if rcapkp == nil {
		return errors.New(`fail to decode root CA key`)
	}

	priv, err := x509.ParsePKCS1PrivateKey(rcapkp.Bytes)
	if err != nil {
		return err
	}

	// load the root CA cert from the file
	rcacf, err := os.ReadFile(c.RootCACert)
	if err != nil {
		return err
	}

	rcacp, _ := pem.Decode(rcacf)
	if rcacp == nil {
		return errors.New(`fail to decode root CA cert`)
	}

	cert, err := x509.ParseCertificate(rcacp.Bytes)
	if err != nil {
		panic(err)
	}

	csrf, err := os.ReadFile(c.CSR)
	if err != nil {
		return err
	}

	crsp, _ := pem.Decode(csrf)
	if crsp == nil {
		return errors.New(`fail to decode CSR file`)
	}

	csr, err := x509.ParseCertificateRequest(crsp.Bytes)
	if err != nil {
		return err
	}

	if err = csr.CheckSignature(); err != nil {
		return err
	}

	// create client certificate template
	serial, err := util.GenerateSerial()
	if err != nil {
		return err
	}

	clientCRTTemplate := &x509.Certificate{
		Signature:          csr.Signature,
		SignatureAlgorithm: csr.SignatureAlgorithm,
		PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
		PublicKey:          csr.PublicKey,
		SerialNumber:       serial,
		Issuer:             cert.Subject,
		Subject:            csr.Subject,
		NotBefore:          time.Now(),
		NotAfter:           time.Now().AddDate(c.ValidityBound, 0, 0),
		KeyUsage:           x509.KeyUsageDigitalSignature,
	}

	// create client certificate from template and CA public key
	clientcrt, err := x509.CreateCertificate(rand.Reader, clientCRTTemplate, cert, csr.PublicKey, priv)
	if err != nil {
		panic(err)
	}

	// save the certificate
	clientcrtf, err := os.Create(fmt.Sprintf(`%s.pem`, c.Name))
	if err != nil {
		panic(err)
	}
	pem.Encode(clientcrtf, &pem.Block{Type: "CERTIFICATE", Bytes: clientcrt})
	return clientcrtf.Close()
}
