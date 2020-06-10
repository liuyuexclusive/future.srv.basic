package main

import (
	"net/http"

	"github.com/liuyuexclusive/future.srv.basic/handler/messageHandler"
	"github.com/liuyuexclusive/future.srv.basic/handler/roleHandler"
	"github.com/liuyuexclusive/future.srv.basic/handler/userHandler"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/liuyuexclusive/utils/srv"

	micro "github.com/micro/go-micro/v2"

	message "github.com/liuyuexclusive/future.srv.basic/proto/message"
	role "github.com/liuyuexclusive/future.srv.basic/proto/role"
	user "github.com/liuyuexclusive/future.srv.basic/proto/user"
)

type start struct {
}

func (s *start) Start(service micro.Service) {
	service.Options()
	// Register Handler
	user.RegisterUserHandler(service.Server(), new(userHandler.Handler))
	role.RegisterRoleHandler(service.Server(), new(roleHandler.Handler))
	message.RegisterMessageHandler(service.Server(), new(messageHandler.Handler))

	// micro.RegisterSubscriber("go.micro.srv.basic1", service.Server(), func(ctx context.Context, msg *basic.TestMessage) error {
	// 	log.Log("Function Received message: ", msg.Name)
	// 	return nil
	// })
}

// main
func main() {

	go StartMonitor()

	// http.HandleFunc("/metrics", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("prometheus request %s\n", time.Now())
	// 	promhttp.Handler().ServeHTTP(rw, r)
	// })
	// go func() {
	// 	http.ListenAndServe(":9999", nil)
	// }()
	srv.Startup("go.micro.srv.basic", new(start), func(option *srv.Options) {
		option.IsLogToES = false
		option.IsTrace = false
		option.IsMonitor = false
	})

}

func StartMonitor() error {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8888", nil)
	return err
}
