package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/utils"
	chatDal "github.com/Wishrem/wuso/server/chat/dal"
	chat "github.com/Wishrem/wuso/server/chat/service"
	friendDal "github.com/Wishrem/wuso/server/friend/dal"
	"github.com/Wishrem/wuso/server/routes"
	userDal "github.com/Wishrem/wuso/server/user/dal"
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	path := flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path)

	chatDal.Init()
	friendDal.Init()
	userDal.Init()

	chat.Init(10)

	opts := idgen.NewIdGeneratorOptions(0)
	idgen.SetIdGenerator(opts)
}

func main() {
	r := routes.NewRouter()

	server := &http.Server{
		Addr:           config.Server.Addr,
		Handler:        r,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: config.Server.MaxHeaderBytes,
	}

	go utils.ListenSignal(func() {
		server.Shutdown(context.TODO())
	})

	log.Printf("server listening at %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
