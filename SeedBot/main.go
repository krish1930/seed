package main

import (
	"SeedBot/core"
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func main() {
	fmt.Println(`
      ___      ____     __     _  _      _  _   
   F __".   F __ ]    FJ    F L L]    FJ / ;  
  J |--\ L J |--| L  J  L  J   \| L  J |/ (|  
  | |  J | | |  | |  |  |  | |\   |  |     L  
  F L__J | F L__J J  F  J  F L\\  J  F L:\  L 
 J______/FJ\______/FJ____LJ__L \\__LJ__L \\__L
 |______F  J______F |____||__L  J__||__L  \L_|
                                                
`)
	fmt.Println(`made bу : krish`)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yml")
	if err != nil {
		panic(err)
	}

	core.ProcessBot(config.Default())
}
