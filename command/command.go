package command

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

var app=cli.NewApp()

var pizza=[]string{"enjoy your pizza with some delicious"}

func Commands(){
	app.Commands=[]cli.Command{
		{
			Name: "peppers",
			Aliases: []string{"p"},
			Usage: "Add peppers to your pizza",
			Action: func(c *cli.Context) {
				pe:="peppers"
				peppers:=append(pizza,pe)
				m:=strings.Join(peppers,"")
				log.Print(m)
			},
		},
		{
			Name: "pineapple",
			Aliases: []string{"pa"},
			Usage: "add pineapple to you pizza",
			Action: func(c *cli.Context) {
				pa:="pineapple"
				pineapple:=append(pizza,pa)
				m:=strings.Join(pineapple,"")
				log.Printf(m)
			},
		},
		{
			Name: "cheese",
			Aliases: []string{"c"},
			Usage: "add cheese to you pizza",
			Action: func() {
				che:="cheese"
				cheese:=append(pizza,che)
				m:=strings.Join(cheese,"")
				log.Println(m)
			},
		},
	}
}

func InitCommand(){
	Commands()
	err:=app.Run(os.Args)
	if  err!= nil{
		log.Fatal(err)
	}
}