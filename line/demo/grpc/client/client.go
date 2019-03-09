package main

import (
	"fmt"
	"github.com/scryinfo/dot/dots/grpc/client"
	pb "github.com/scryinfo/dot/line/demo/pb"
	"github.com/scryinfo/dot/line/lineimp"
	"log"
)

func main()  {
	l := lineimp.New()
	l.ToLifer().Create(nil)

	gclient.Add(l)

	err := l.Rely()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = l.CreateDots()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = l.ToLifer().Start(false)
	if err != nil {
		fmt.Println(err)
		return
	}

	f := &gclient.GrpcClient{}
	l.ToInjecter().Inject(f)

	conn := f.Grpcs.GetConn()

	ctx := f.Grpcs.GetCtx()


	//用户实现 start

	c := pb.NewTestClient(conn)

	d := pb.NewGreeterClient(conn)

	t :=pb.NewRechargeClient(conn)

	c1, err := c.SayHello(ctx, &pb.TestRequest{Name: "shrimpliao"})
	d1,err := d.SayHello1(ctx,&pb.HelloRequest{Name:"111111"})
	t1,err :=t.GetRechargeAddr(ctx,&pb.GetRechargeAddrRequest{DddAddr:"456"})

	fmt.Println("err",err)

	log.Printf("@@@c1: %s", c1.Message)
	log.Printf("@@@d1: %s", d1.Message)
	log.Printf("@@@t1: %s",t1)

	//用户实现 end

	f.Grpcs.Stop(false)
	f.Grpcs.Destroy(false)
}