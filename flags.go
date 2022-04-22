package myip

import "flag"

var (
	ConfigPath string
)

func Flags() {
	flag.StringVar(&ConfigPath, "f", "", "Path to configuration file")

	flag.Parse()
}
