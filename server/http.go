package server

import (
	"fmt"
	"github.com/go-be/be"
	"net/http"
)

type Http struct {
	routers map[string]be.Router
}

// Start 启动服务
func (server *Http) Start(port int) {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		if server.routers != nil && len(server.routers) > 0 {

			path := r.URL.Path
			appName := ""
			i := 1
			l := len(path)
			for i < l {
				if path[i] == '/' {
					appName = path[1:i]
					break
				}
				i++
			}

			fmt.Println(appName)

			if appName != "" {
				if router, ok := server.routers[appName]; ok {

					req := new(be.Request)
					req.Init(r)

					res := new(be.Response)
					res.Init(rw)

					c := new(be.Context)
					c.Init(req, res)

					// 回收资源
					defer c.Gc()

					router.Execute(c)
				}
			}
		}
	})

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		return
	}
}

func (server *Http) AddRouter(appName string, router be.Router) *Http {
	if server.routers == nil {
		server.routers = make(map[string]be.Router)
	}

	server.routers[appName] = router
	return server
}
