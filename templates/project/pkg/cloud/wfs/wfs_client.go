/**
 * Created by Wangwei on 2019-03-27 11:36.
 */

package wfs

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/donnie4w/go-logger/logger"
	"github.com/wangwei123/wfs/httpserver/protocol"
)

type WfsClient struct {
	ServerUrl string
}

func (this *WfsClient) PostFile(bs []byte, name, fileType string) (err error) {
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTHttpPostClient(this.ServerUrl)
	if err != nil {
		logger.Error("err:", err.Error())
	}
	client := protocol.NewIWfsClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		logger.Error("err:", err.Error())
	}
	defer transport.Close()
	wf := protocol.NewWfsFile()
	wf.FileBody = bs
	wf.Name = &name
	wf.FileType = &fileType
	_, er := client.WfsPost(wf)
	if er != nil {
		err = er
		logger.Debug("err:", err.Error())
	}
	return
}

func (this *WfsClient) GetFile(name string) (bs []byte, err error) {
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTHttpPostClient(this.ServerUrl)
	if err != nil {
		logger.Error("err:", err.Error())
	}
	client := protocol.NewIWfsClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		logger.Error("err:", err.Error())
	}
	defer transport.Close()
	wf, er := client.WfsRead(name)
	if er == nil {
		bs = wf.GetFileBody()
		logger.Debug("len(bs):", len(bs))
	}
	err = er
	return
}

func (this *WfsClient) DelFile(name string) (err error) {
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTHttpPostClient(this.ServerUrl)
	if err != nil {
		logger.Error("err:", err.Error())
	}
	client := protocol.NewIWfsClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		logger.Error("err:", err.Error())
	}
	defer transport.Close()
	ack, er := client.WfsDel(name)
	if er == nil {
		logger.Debug("ack:", ack)
	}
	err = er
	return
}
