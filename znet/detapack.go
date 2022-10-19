package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

// DataPack /**
type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头的长度的方法
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen是uint32,4字节， id也是uint32
	return 8
}

// Pack 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个要放回的数据流
	dataBuff := bytes.NewBuffer([]byte{})
	// 二进制写入长度
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageLength()); err != nil {
		return nil, err
	}
	// 二进制写入版本号
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageId()); err != nil {
		return nil, err
	}
	// 将data数据写入buffer中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessage()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// Unpack 将包的head信息读出来，再根据head的信息里的长度和版本去读取数据
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据读数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	// 解压head信息，得到dataLen和messageId
	msg := &Message{}
	// dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 如果dataLen超出了我们的size，不做处理
	if utils.GlobalConfig.MaxPackageSize > 0 && msg.DataLen > utils.GlobalConfig.MaxPackageSize {
		return nil, errors.New("too large msg data recv")
	}
	return msg, nil
}
