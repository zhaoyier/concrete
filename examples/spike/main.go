package main

import (
	"git.ezbuy.me/ezbuy/concrete/common/web"
	"git.ezbuy.me/ezbuy/concrete/examples/spike/service"
)

func main() {
	serve := web.NewWebServe()
	serve.Register(new(service.Hello))
	serve.WebStart("8080")
}
