package main

import (
	"go.uber.org/zap"
	"koala/src/log"
	"koala/src/router"
	"koala/src/util/configUtil"
)

func main() {
	runServer()
}

func runServer()  {
	initServices()

	//other services

}

func initServices()  {
	//init config from koala.yaml
	_, err := configUtil.Conf.InitConfig()
	if err != nil {
		return
	}

	//init zap logger
	log.InitZapLogger()

	// init go-gin web router
	err = router.InitRouter(configUtil.GetListen())
	if err != nil {
		log.Error("Failed to init go-gin web router. Error message:", zap.Error(err))
		return
	}

	log.Info("Init server successfully.")
}
