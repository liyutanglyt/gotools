/**
 * Created by guanzhongkai on 2019/3/26 10:30.
 */
package socket

import (
	"io/ioutil"
	"net"
	"time"

	"github.com/gogf/gf/g/os/glog"
)

const (
	Network  = "tcp"
	Network4 = "tcp4"
)

func Client(params string, addr string, random string) (string, error) {
	// go 里面通过ResolveTCPAddr获取连接地址
	// net参数是"tcp4"、"tcp6"、"tcp"中的任意一个 , addr表示域名或者IP地址
	log := "[" + random + "][socket]"
	glog.Debug(log, "start -=-=-=-=-=-=->")
	start := time.Now()
	server, err := net.ResolveTCPAddr(Network4, addr)
	if err != nil {
		glog.Error(log, err)
		return "", err
	}

	conn, err := net.DialTCP(Network, nil, server)
	defer conn.Close()

	if err != nil {
		glog.Error(log, "连接失败", err)
		return "", err

	} else {
		// 往服务端发送数据
		_, err := conn.Write([]byte(params))
		if err != nil {
			glog.Error(log, "发送失败", err)
			return "", err

		}
		resp, err := ioutil.ReadAll(conn)
		if err != nil {
			glog.Error(log, "读取数据错误", err)
			return "", err

		}
		end := time.Now()
		glog.Debug(log, "end -=-=-=-=-=-=->")
		glog.Debug(log, "cos", end.Sub(start).Seconds())

		return string(resp), nil
	}
}
