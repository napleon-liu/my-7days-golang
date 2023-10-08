package geecache

/*
*	ByteView只有一个数据成员， b []byte, b将会存储真实的缓存值，选择byte类型是为了能够支持任意数据类型的存储，例如字符串、图片等。
*	实现 Len() int 方法，我们在 lru.Cache 的实现中，要求被缓存对象必须实现 Value 接口，即 Len() int 方法，返回其所占的内存大小。
*	b 是只读的，使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改。
 */

// A ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice
func (c ByteView) ByteSlice() []byte {
	return cloneBytes(c.b)
}

// cloneBytes make a new buffer and copy b into it
func cloneBytes(c []byte) []byte {
	b := make([]byte, len(c))
	copy(b, c)
	return b
}

// String returns the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.b)
}
