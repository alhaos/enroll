package webServer

import (
	"embed"
	"fmt"
	"github.com/alhaos/enroll/internal/handlers"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates/*.gohtml
var templatesFS embed.FS

//go:embed static/*
var staticFS embed.FS

// Config web server configurations struct
type Config struct {
	Address string `yml:"address"`
}

// Server general web serve struct
type Server struct {
	config Config
	Router *gin.Engine
}

// New is constructor from web service
func New(config Config, handler *handlers.Handler) (*Server, error) {

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	err := loadTemplates(router)
	if err != nil {
		return nil, err
	}

	registerRoutes(router, handler)

	registerStaticRoute(router)

	return &Server{
		config: config,
		Router: router,
	}, nil
}

// registerStaticRoute register static route
func registerStaticRoute(router *gin.Engine) {
	staticFsSub, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	router.StaticFS("static", http.FS(staticFsSub))
}

// registerRoutes register routes and handlers
func registerRoutes(router *gin.Engine, handler *handlers.Handler) {

	// register index route handler
	router.GET("/", handler.IndexGetHandler)
}

// loadTemplates load templates from embed fs
func loadTemplates(router *gin.Engine) error {

	tmpl, err := template.ParseFS(templatesFS, "templates/*.gohtml")
	if err != nil {
		return fmt.Errorf("unable to load templates: %w", err)
	}

	router.SetHTMLTemplate(tmpl)

	return nil
}

// Run web server
func (s *Server) Run() error {
	err := s.Router.Run(s.config.Address)
	if err != nil {
		return err
	}
	return nil
}
