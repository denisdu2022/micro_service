package hanlder

//解决名称冲突的问题

const HelloServiceName = "handler/HelloService"

type NewHelloService struct {
}

func (s *NewHelloService) Hello(request string, reply *string) error {
	*reply = "Hello: " + request
	return nil
}
