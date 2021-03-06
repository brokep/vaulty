package routing

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Router interface {
	LookupRoute(req *http.Request) *Route
}

type router struct {
	routes []*Route
}

func NewRouter() *router {
	return &router{}
}

func (r *router) LookupRoute(req *http.Request) *Route {
	log.Debugf("LookupRoute for %s request to %s", req.Method, req.URL.String())

	for _, route := range r.routes {
		if route.Match(req) {
			return route
		}
	}
	return nil
}

func (r *router) SetRoutes(routes []*Route) {
	r.routes = routes
}
