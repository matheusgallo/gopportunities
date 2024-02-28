package main

import (
	"github.com/matheusgallo/gopportunities/config"
	"github.com/matheusgallo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	//init configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//init router
	router.Initialize()

}
