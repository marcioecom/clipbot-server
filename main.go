package main

import (
	"github.com/marcioecom/clipbot-server/api"
	"github.com/marcioecom/clipbot-server/helper"
)

func main() {
	helper.InitLogger()
	helper.LoadEnvs()

	api.Start()
}
