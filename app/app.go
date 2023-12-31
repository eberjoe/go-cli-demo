package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "eber.xyz",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de enderecos na internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca os nomes dos servidores na internet",
			Flags:  flags,
			Action: findNameServer,
		},
	}
	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findNameServer(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
