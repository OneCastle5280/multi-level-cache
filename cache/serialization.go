package cache

// Serialization
// @Description: 序列化接口
type Serialization interface {
	//
	// Marshal
	//  @Description: 序列化
	//  @param v
	//  @return []byte
	//  @return error
	//
	Marshal(v any) ([]byte, error)

	//
	// Unmarshal
	//  @Description: 反序列化
	//  @param data
	//  @param v
	//  @return error
	//
	Unmarshal(data []byte, v any) error

	//
	// unionKey
	//  @Description: 序列化方式的唯一标识
	//  @return string
	//
	unionKey() string
}
