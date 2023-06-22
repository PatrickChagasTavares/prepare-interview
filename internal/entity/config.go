package entity

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Help    bool
	Version bool
	Lang    string
	Token   string
	Role    string
	Skills  string
	Name    string
}

const (
	OperationHelper uint8 = iota
	OperationVersion
	OperationCompany
	OperationPrepare
)

func LoadComands() *Config {
	var (
		helpFlag  = flag.Bool("help", false, "")
		tokenFlag = flag.String("token", "", "")
	)
	flag.Parse()

	flag.Usage = replaceFlagUsage

	return &Config{
		Help:  *helpFlag,
		Token: *tokenFlag,
	}
}

func (c *Config) Operation() uint8 {
	if c.Help || len(flag.Args()) < 1 {
		return OperationHelper
	}

	switch flag.Arg(0) {
	case "version":
		return OperationVersion
	case "company":
		c.ParseBy("company")
		return OperationCompany
	case "prepare":
		c.ParseBy("prepare")
		return OperationPrepare
	default:
		return OperationHelper
	}
}

func (c *Config) HasToken() bool {
	return c.Token != ""
}

func (c *Config) ParseBy(name string) {
	args := flag.Args()[1:]

	flagSet := flag.NewFlagSet(name, flag.ExitOnError)
	langFlag := flagSet.String("lang", "Portuguese", "")

	if name == "prepare" {
		roleFlag := flagSet.String("role", "", "")
		skillsFlag := flagSet.String("skills", "", "")

		if err := flagSet.Parse(args); err != nil {
			log.Fatalln(err)
		}

		c.Role = *roleFlag
		c.Skills = *skillsFlag
		c.Lang = *langFlag
		return
	}

	if name == "company" {
		nameFlag := flagSet.String("name", "", "")

		if err := flagSet.Parse(args); err != nil {
			log.Fatalln(err)
		}

		c.Name = *nameFlag
		c.Lang = *langFlag
		return
	}
}

func replaceFlagUsage() {
	fmt.Fprintln(os.Stderr, `Usage: interview OPTIONS COMMAND [arg...]
	interview [ -help ]
	
	Options:
	  -name            Ex.: "Twitter" || "Google"
	  -role            Ex.: "Software Engineer" || "Suport TI"
	  -skills          Ex.: Golang, NodeJS and AWS || Pacote office
	  -lang            Ex: english or portuguese. Default is Portuguese
	  -token           API key gerated by openAI
	  -help            Print usage
	
	Commands:
	  company      [-name N] search information about company.
	  prepare      [-role R] [-skills S] make text to help you to prepare interview, using Role R and skills S
	  version      Print current migration version
	`)
}
