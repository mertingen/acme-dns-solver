package app

import (
	"acmesolver/app/models"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/go-acme/lego/v4/registration"
	"log"
	"os"
)

func Start(domain string) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Email: os.Getenv("EMAIL"),
		Key:   privateKey,
	}

	conf := lego.NewConfig(&user)
	conf.CADirURL = os.Getenv("CA_DIR_URL")
	conf.Certificate.KeyType = certcrypto.RSA2048

	client, err := lego.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	dnsProvier, err := route53.NewDNSProvider()
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Challenge.SetDNS01Provider(dnsProvier)
	if err != nil {
		log.Fatal(err)
	}

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Fatal(err)
	}
	user.Registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{domain},
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(certificates.Domain)
	fmt.Println(string(certificates.Certificate))
	fmt.Println(string(certificates.PrivateKey))
}
