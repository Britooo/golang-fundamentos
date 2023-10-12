package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Sempre que for utilizar pacotes de terceiros, é necessário fazer o download do mesmo
// Para isso, é necessário utilizar o comando go get
// Exemplo: go get github.com/urfave/cli
// Para usar o pacote, é necessário importar o mesmo
// e usar o nome após a última barra

// Gerar retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Primeira CLI feita em Go"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := configurarFlags()

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na Internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores na Internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func configurarFlags() []cli.Flag {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "sptech.school",
		},
	}

	return flags
}

func buscarServidores(c *cli.Context) {
	host := c.String("host") // Busca o valor da flag host

	fmt.Printf("Buscando servidores do host: %s\n", host)

	// Pacote net
	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal("Não foi possível encontrar o servidor do host informado\n", erro)
	} else {
		for index, servidor := range servidores {
			fmt.Printf("Servidor[%d]: %s\n", (index + 1), servidor.Host)
		}
	}
}

func buscarIps(c *cli.Context) {
	host := c.String("host") // Busca o valor da flag host

	fmt.Printf("Buscando IPs do host: %s\n", host)

	// Pacote net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal("Não foi possível encontrar o IP do host informado\n", erro)
	} else {
		for index, ip := range ips {
			fmt.Printf("IP[%d]: %s\n", (index + 1), ip)
		}
	}
}
