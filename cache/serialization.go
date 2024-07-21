package cache

import "encoding/json"

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

// JsonSerialization JSON 序列化
type JsonSerialization struct{}

func (JsonSerialization) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (JsonSerialization) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (JsonSerialization) unionKey() string {
	return "JSON"
}

func NewJsonSerialization() Serialization {
	return JsonSerialization{}
}
