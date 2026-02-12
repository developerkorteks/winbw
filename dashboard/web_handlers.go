package dashboard

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/winbu.tv/config"
)

// WebHandler handles web dashboard pages
type WebHandler struct {
	dynamicConfig *config.DynamicConfig
	templates     *template.Template
}

// NewWebHandler creates a new web dashboard handler
func NewWebHandler(dc *config.DynamicConfig) (*WebHandler, error) {
	// Parse templates
	templates, err := template.ParseGlob(filepath.Join("dashboard", "templates", "*.html"))
	if err != nil {
		return nil, err
	}

	return &WebHandler{
		dynamicConfig: dc,
		templates:     templates,
	}, nil
}

// ShowDashboard renders the main dashboard page
func (h *WebHandler) ShowDashboard(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Parse dashboard template
	dashboardTemplate := template.Must(template.ParseFiles(
		filepath.Join("dashboard", "templates", "layout.html"),
		filepath.Join("dashboard", "templates", "dashboard.html"),
	))
	
	data := gin.H{
		"Title": "Dashboard",
	}

	err := dashboardTemplate.Execute(c.Writer, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: "+err.Error())
		return
	}
}

// ShowConfig renders the configuration page
func (h *WebHandler) ShowConfig(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Parse config template
	configTemplate := template.Must(template.ParseFiles(
		filepath.Join("dashboard", "templates", "layout.html"),
		filepath.Join("dashboard", "templates", "config.html"),
	))
	
	data := gin.H{
		"Title": "Configuration",
	}

	err := configTemplate.Execute(c.Writer, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: "+err.Error())
		return
	}
}

// ShowHealth renders the health check page
func (h *WebHandler) ShowHealth(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	data := gin.H{
		"Title": "Health Check",
	}

	err := h.templates.ExecuteTemplate(c.Writer, "layout.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: "+err.Error())
		return
	}
}
