package process

import (
	"app/chat/conf"
	"encoding/binary"
	"encoding/json"
	"net"
)

type SmsProcess struct {
	Conn net.Conn
	buf  [4096]byte
}

func NewSmsProcess(conn net.Conn) *SmsProcess {
	return &SmsProcess{
		Conn: conn,
		buf:  [4096]byte{},
	}
}

// 发送消息
func (smsProcess *SmsProcess) SendMsg(msgType string, data string) (err error) {

	// 打包数据
	var msg conf.Msg
	msg.Type = msgType
	msg.Data = data
	msgDataBS, _ := json.Marshal(msg)

	// 发送长度再发送数据
	intByte := make([]byte, 4)
	binary.BigEndian.PutUint32(intByte, uint32(len(msgDataBS)))
	n, err := smsProcess.Conn.Write(intByte)
	if n != 4 || err != nil {
		return conf.ErrorConnect
	}

	n, err = smsProcess.Conn.Write(msgDataBS)
	if err != nil {
		return conf.ErrorConnect
	}
	return nil
}

// 接收消息
func (smsProcess *SmsProcess) ReadMsg() (msg conf.Msg, err error) {

	// 读取头部
	_, err = smsProcess.Conn.Read(smsProcess.buf[:4])
	if err != nil {
		err = conf.ErrorConnect
		return
	}

	// 读取消息体
	msgLen := int(binary.BigEndian.Uint32(smsProcess.buf[0:4]))
	n, err := smsProcess.Conn.Read(smsProcess.buf[:msgLen])
	if err != nil {
		err = conf.ErrorConnect
		return
	} else if n != msgLen {
		err = conf.ErrorMsgLength
		return
	}

	// 反序列化数据
	err = json.Unmarshal(smsProcess.buf[:msgLen], &msg)
	if err != nil {
		err = conf.ErrorUnmarshal
		return
	}

	return
}
