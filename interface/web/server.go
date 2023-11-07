package web

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	middleware "github.com/yzx9/motion/interface/middleWare"
	"sync"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 18:42
*@Version: V1.0
 */

type Server struct {
	Error  error
	config *viper.Viper
	engine *gin.Engine
	addr   string
}

func (s *Server) GetGinEngine() *gin.Engine {
	return s.engine
}

var server *Server
var mutex sync.Mutex

func NewServer() {
	if server == nil {
		mutex.Lock()
		if server == nil {
			InitServer()
		}
		mutex.Unlock()
	}
}

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORSMiddleware())
	r.MaxMultipartMemory = 1 << 30
	server = &Server{
		engine: r,
		addr:   "192.168.2.115:8080",
	}
	server.SetRouter()
}
