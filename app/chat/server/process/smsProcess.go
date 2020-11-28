package process

import (
	"learn-go/app/chat/conf"
	"learn-go/app/chat/utils"
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
func (smsProcess *SmsProcess) SendMsg(errorCode error, data string) (err error) {

	// 获取结果数据
	msgDataBS, err := utils.GetSendData(errorCode, data)
	if err != nil {
		return err
	}

	// 发送长度再发送数据
	intByte := make([]byte, 4)
	binary.BigEndian.PutUint32(intByte, uint32(len(msgDataBS)))
	n, err := smsProcess.Conn.Write(intByte)
	if n != 4 || err != nil {
		return err
	}

	n, err = smsProcess.Conn.Write(msgDataBS)
	if err != nil {
		return err
	}
	return nil
}

// 发送消息
func (smsProcess *SmsProcess) SendUser(data string) (err error) {

	// 获取结果数据
	msg := conf.Msg{
		Type: conf.SendUserType,
		Data: data,
	}
	msgDataBS, err := json.Marshal(msg)
	if err != nil {
		err = conf.ErrorMarshal
		return
	}

	// 发送长度再发送数据
	intByte := make([]byte, 4)
	binary.BigEndian.PutUint32(intByte, uint32(len(msgDataBS)))
	n, err := smsProcess.Conn.Write(intByte)
	if n != 4 || err != nil {
		return err
	}

	n, err = smsProcess.Conn.Write(msgDataBS)
	if err != nil {
		return err
	}
	return nil
}

// 服务器推送
// 方法调用频率高时，需要考虑操作优化
// TODO: 用户消息相同时每个重复的序列化、求长度 .......
func (smsProcess *SmsProcess) ServerPush(msgDataBS []byte) (err error) {

	// 发送长度再发送数据
	intByte := make([]byte, 4)
	binary.BigEndian.PutUint32(intByte, uint32(len(msgDataBS)))
	n, err := smsProcess.Conn.Write(intByte)
	if n != 4 || err != nil {
		return err
	}

	n, err = smsProcess.Conn.Write(msgDataBS)
	if err != nil {
		return err
	}
	return nil
}

// 接收消息
func (smsProcess *SmsProcess) ReadMsg() (msg conf.Msg, err error) {

	// 读取头部
	_, err = smsProcess.Conn.Read(smsProcess.buf[:4])
	if err != nil {
		return
	}

	// 读取消息体
	msgLen := int(binary.BigEndian.Uint32(smsProcess.buf[0:4]))
	n, err := smsProcess.Conn.Read(smsProcess.buf[:msgLen])
	if err != nil {
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
