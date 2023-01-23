package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	flags := []cli.Flag {
		cli.StringFlag {
			Name: "host",
			Value: "zapimoveis.com.br",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca Ips de endereços na internet",
			Flags: flags,
			Action: buscaIps,
		},
		{
			Name: "server",
			Usage: "Busca nome de servidores na internet",
			Flags: flags,
			Action: buscaNomeServidor,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaNomeServidor(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}