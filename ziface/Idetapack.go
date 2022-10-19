package ziface

// IDataPack /*
type IDataPack interface {
	// GetHeadLen 获取头长度方法
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg IMessage) ([]byte, error)
	// Unpack 拆包方法
	Unpack([]byte) (IMessage, error)
}
