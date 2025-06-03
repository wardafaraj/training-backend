package webserver

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
)

// Renderer fetches the template render
func Renderer() *echoview.ViewEngine {
	gvc := goview.Config{
		Root:      "webserver",
		Extension: ".html",
		Master:    "layouts/master",
		Funcs: template.FuncMap{

			"sub": func(a, b int) int {
				return a - b
			},

			"add": func(a, b int) int {
				return a + b
			},
			"subfloat": func(a, b float32) string {
				return fmt.Sprintf("%0.2f", a-b)

			},

			"addfloat": func(a, b float32) string {
				return fmt.Sprintf("%0.2f", a-b)
			},

			"amountInFloat": func(a float32) string {
				return fmt.Sprintf("%0.2f", a)
			},

			"inc": func(i int) int {
				return i + 1
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
			"date_formatter": func(d string) string {
				return d[0:10]
			},

			"title": func() string {
				return "CRM"
			},

			"timestamp": func(t time.Time) string {
				return t.Format("02/01/2006 15:04")
			},
			"menu": func(path string, data ...string) string {
				for _, v := range data {

					if strings.HasPrefix(path, v) {
						return "menu-open"
					}
				}
				return ""
			},
			"active": func(path string, data ...string) string {
				for _, v := range data {
					if path == v {
						return "active"
					}
				}
				return ""
			},
			"hasRole": func(roles []string, role string) bool {
				r := hasOne(roles, role)
				return r
			},

			"hasAnyRole": func(roles []string, role ...string) bool {
				r := hasAny(role, roles)
				return r
			},
			"hasPermission": func(permissions []string, permission string) bool {
				p := hasOne(permissions, permission)
				return p
			},
		},

		DisableCache: true,
	}
	return echoview.New(gvc)
}

func hasAny(s1 []string, s2 []string) bool {
	for _, a := range s1 {
		for _, b := range s2 {
			if a == b {
				return true
			}
		}
	}
	return false
}

func hasOne(s1 []string, s2 string) bool {
	for _, s := range s1 {
		if s == s2 {
			return true
		}
	}
	return false
}
