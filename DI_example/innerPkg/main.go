package InnerPkg

import "fmt"

type Config struct {
	Name string
	Port int
}

type DI struct {
	Config *Config
}

func init() {
	fmt.Println("InnerPkg init fired")
}

func NewDI(config *Config) DI {
	return DI{Config: config}
}
