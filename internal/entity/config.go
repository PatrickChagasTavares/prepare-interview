package entity

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Help    bool
	Version bool
	Lang    string
	Token   string
	Role    string
	Skills  string
}

const (
	OperationHelper uint8 = iota
	OperationVersion
	OperationDefault
)

func LoadComands() *Config {
	var (
		helpFlag    = flag.Bool("help", false, "")
		versionFlag = flag.Bool("version", false, "")
		langFlag    = flag.String("lang", "Portuguese", "")
		tokenFlag   = flag.String("token", "", "")
		roleFlag    = flag.String("role", "", "")
		skillsFlag  = flag.String("skills", "", "")
	)
	flag.Parse()

	flag.Usage = replaceFlagUsage

	return &Config{
		Help:    *helpFlag,
		Version: *versionFlag,
		Lang:    *langFlag,
		Token:   *tokenFlag,
		Role:    *roleFlag,
		Skills:  *skillsFlag,
	}
}

func (c *Config) Operation() uint8 {
	if c.Help {
		return OperationHelper
	}

	if c.Version {
		return OperationVersion
	}

	return OperationDefault
}

func (c *Config) HasToken() bool {
	return c.Token != ""
}

func replaceFlagUsage() {
	fmt.Fprintln(os.Stderr, `Usage: interview OPTIONS COMMAND [arg...]
	interview [ -version | -help ]
	
	Options:
	  -role            Ex.: "Software Engineer" || "Suport TI"
	  -skills          Ex.: Golang, NodeJS and AWS || Pacote office
	  -lang            Ex: english or portuguese. Default is Portuguese
	  -token           API key gerated by openAI
	  -version         Print version
	  -help            Print usage
	`)
}
