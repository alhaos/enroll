package webServer

import (
	"embed"
	"fmt"
	"github.com/alhaos/enroll/internal/handlers"
	"github.com/gin-gonic/gin"
	"html/template"
)

//go:embed templates/*.gohtml
var templatesFS embed.FS

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
func New(config Config, handlers *handlers.Handler) (*Server, error) {

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	tmpl, err := template.ParseFS(templatesFS, "templates/*.gohtml")
	if err != nil {
		return nil, fmt.Errorf("unable to load templates: %w", err)
	}

	router.SetHTMLTemplate(tmpl)

	router.GET("/", handlers.IndexGetHandler)

	return &Server{
		config: config,
		Router: router,
	}, nil
}

// Run web server
func (s *Server) Run() error {
	err := s.Router.Run(s.config.Address)
	if err != nil {
		return err
	}
	return nil
}
