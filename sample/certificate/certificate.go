// Scry Info.  All rights reserved.
// license that can be found in the license file.

package main

import (
	"fmt"
	"os"

	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/certificate"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/scryg/sutils/ssignal"
)

func main() {
	l, err := line.BuildAndStart(add) //first step create line and dots
	if err != nil {
		fmt.Println(err)
		return
	}
	defer line.StopAndDestroy(l, true) //fourth step stop and destroy dots

	dot.Logger().Infoln("dot ok")
	//second step ....

	dd, _ := l.ToInjecter().GetByLiveId(dot.LiveId(certificate.EcdsaTypeId))
	if d, ok := dd.(*certificate.Ecdsa); ok {
		err := makeSample(d)
		if err != nil {
			dot.Logger().Errorln(err.Error())
		}
	}

	ssignal.WatiCtrlC(func(s os.Signal) bool { //third wait for exit
		return false
	})
}

func add(l dot.Line) error {
	err := l.PreAdd(certificate.TypeLiveEcdsa())
	return err
}

//生成 ca证书，生成ca证书下的 server与client证书
func makeSample(cs *certificate.Ecdsa) error {

	caPri, err := certificate.MakePriKey()
	if err != nil {
		return err
	}

	ca, err := cs.GenerateCaCertKey(caPri, "ca.key", "ca.pem", []string{"scry"}, []string{"scry"})
	if err != nil {
		return err
	}

	err = cs.GenerateCertKey(ca, caPri, "server.key", "server.pem", []string{"scry"}, []string{"scry"})
	if err != nil {
		return err
	}

	err = cs.GenerateCertKey(ca, caPri, "client.key", "client.pem", []string{"scry"}, []string{"scry"})
	if err != nil {
		return err
	}

	return nil

}
