package ziface

/*
IMessage 将请求的消息封装到一个message中
*/
type IMessage interface {
	// GetMessageId 消息Id
	GetMessageId() uint32
	// GetMessageLength 消息的长度
	GetMessageLength() uint32
	// GetMessage 获取消息
	GetMessage() []byte
	// SetMessageId 设置消息Id
	SetMessageId(id uint32)
	// SetMessageLength 设置消息长度
	SetMessageLength(len uint32)
	// SetMessage 设置消息内容
	SetMessage(data []byte)
}
