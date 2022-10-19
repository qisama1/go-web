package znet

type Message struct {
	Id      uint32 // 消息id
	DataLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func (m Message) GetMessageId() uint32 {
	return m.Id
}

func (m Message) GetMessageLength() uint32 {
	return m.DataLen
}

func (m Message) GetMessage() []byte {
	return m.Data
}

func (m Message) SetMessageId(id uint32) {
	m.Id = id
}

func (m Message) SetMessageLength(len uint32) {
	m.DataLen = len
}

func (m Message) SetMessage(data []byte) {
	m.Data = data
}
