package handler

//HelloServiceName 有前缀,是为了解决名称冲突的问题

const HelloServiceName = "handler/HelloService"

//只关心方法,不用关心名称

type NewHelloService struct {
}

func (s *NewHelloService) Hello(request string, reply *string) error {
	*reply = "您好: " + request
	return nil
}
