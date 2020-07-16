package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/lobby"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"go.uber.org/zap"
)

var (
	ws     websocket.Upgrader
	lby    *lobby.Lobby
	logger *zap.Logger

	App = gin.New()
	_   = App.Use(cors.Default(), gin.Logger(), gin.Recovery())
)

func Start() {
	http.HandleFunc("/ws", handleClient)
	go func() {
		if err := http.ListenAndServe(fmt.Sprint(":", conf.WSPort), nil); err != nil {
			log.Println(err)
		}
	}()

	Regiter()
	s := &http.Server{
		Addr:           fmt.Sprint(":", conf.HttpPort),
		Handler:        App,
		ReadTimeout:    time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func Regiter() {
	App.GET("/ping", ping)
	if conf.Dev {
		// App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	App.GET("/", func(c *gin.Context) {
		serveHome(c.Writer, c.Request)
	})

	// order.Register(App)
	// member.Register(App)

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func ApplyLobby(l *lobby.Lobby) {
	lby = l
}

func ApplyLogger(z *zap.Logger) {
	logger = z
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "view/home.html")
}

func handleClient(w http.ResponseWriter, r *http.Request) {
	r.Header.Del("Origin")
	// start websocket
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		// handle error
		return
	}
	defer conn.Close()

	// create lobby.Player with conn
	c, err := checklogin(conn)
	if err != nil {
		log.Println(err)
		return
	}

	if p, r := lby.Player(c.Account); p != nil && r != nil {
		c.write(NewS2CGameNotFinished())
		time.Sleep(2 * time.Second)
		return
	}

	p := player.New(c.Account, c.Name, c)
	lby.Wellcome(p)
	p.Play()
}
