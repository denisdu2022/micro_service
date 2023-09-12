package hanlder

//服务名前边增加了handler 前缀是为了解决名称冲突的问题(如果有多个service,在全局上只能有一个)

const HelloServiceName = "handler/HelloService"

//类

type NewHelloService struct {
}

//类的方法

func (s *NewHelloService) Hello(request string, reply *string) error {
	*reply = "您好: " + request
	return nil
}
