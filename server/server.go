package server

import (
	"flower/config"
	"flower/log"
	"flower/router"
	"fmt"
	"github.com/valyala/fasthttp"
)

type Server struct {
	Port   int
	server *fasthttp.Server
}

func (s *Server) Start() {
	s.server = &fasthttp.Server{
		Handler:                            router.Handler,
		ErrorHandler:                       nil,
		HeaderReceived:                     nil,
		Name:                               "FLOWER SERVER",
		Concurrency:                        0,
		DisableKeepalive:                   false,
		ReadBufferSize:                     0,
		WriteBufferSize:                    0,
		ReadTimeout:                        0,
		WriteTimeout:                       0,
		IdleTimeout:                        0,
		MaxConnsPerIP:                      0,
		MaxRequestsPerConn:                 0,
		MaxKeepaliveDuration:               0,
		TCPKeepalive:                       false,
		TCPKeepalivePeriod:                 0,
		MaxRequestBodySize:                 0,
		ReduceMemoryUsage:                  false,
		GetOnly:                            false,
		LogAllErrors:                       false,
		DisableHeaderNamesNormalizing:      false,
		SleepWhenConcurrencyLimitsExceeded: 0,
		NoDefaultServerHeader:              false,
		NoDefaultContentType:               false,
		ConnState:                          nil,
		Logger:                             nil,
		KeepHijackedConns:                  false,
	}

	fmt.Printf("starting flower server on %d...", config.Conf.ServerPort)
	log.InfoF("starting flower server on %d...", config.Conf.ServerPort)
	err := s.server.ListenAndServe(fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.ErrorF("start server error[%v]", err)
	}

}

func (s *Server) Shutdown() error {
	return s.server.Shutdown()
}
