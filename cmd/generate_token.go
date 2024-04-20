package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thienhaole92/root-certificate-authority-jwt/context"
)

type GenerateTokenCmd struct {
	Key  string `arg:"" required:"" name:"key" help:"Path to key file." type:"path"`
	Data string `arg:"" required:"" name:"data" help:"Custom data." type:"path"`
}

func (c *GenerateTokenCmd) Run(ctx *context.Context) error {
	kf, err := os.ReadFile(c.Key)
	if err != nil {
		return err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(kf)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = c.Data                    // Our custom data.
	claims["exp"] = now.Add(time.Hour).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()                // The time at which the token was issued.
	claims["nbf"] = now.Unix()                // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(key)
	if err != nil {
		return err
	}

	fmt.Println("jwt:", token)

	return nil
}
