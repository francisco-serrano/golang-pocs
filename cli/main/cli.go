package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func exampleA() {
	viper.BindEnv("GOMAXPROCS")

	val := viper.Get("GOMAXPROCS")

	fmt.Println("GOMAXPROCS", val)

	viper.Set("GOMAXPROCS", 10)

	val = viper.Get("GOMAXPROCS")

	fmt.Println("GOMAXPROCS", val)

	viper.BindEnv("NEW_VARIABLE")

	val = viper.Get("NEW_VARIABLE")

	if val == nil {
		fmt.Println("NEW_VARIABLE not defined")
		return
	}

	fmt.Println(val)
}

func exampleB() {
	flag.Int("i", 100, "i parameter")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("i")
	fmt.Println(i)
}

func exampleC() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./myConfig.json")

	fmt.Println(os.Getwd())

	fmt.Printf("using config: %s\n", viper.ConfigFileUsed())

	viper.ReadInConfig()

	if viper.IsSet("item1.key1") {
		fmt.Println("item1.key1:", viper.Get("item1.key1"))
	} else {
		fmt.Println("item1.key1 not set!")
	}

	if viper.IsSet("item2.key3") {
		fmt.Println("item2.key3:", viper.Get("item2.key3"))
	} else {
		fmt.Println("item2.key3 is not set!")
	}

	if !viper.IsSet("item3.key1") {
		fmt.Println("item3.key1 is not set!")
	}
}

func main() {
	//exampleA()
	//exampleB()
	exampleC()
}