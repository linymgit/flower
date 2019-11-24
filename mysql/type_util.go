package mysql

func Change2int64(peek interface{}) (id int64, ok bool) {
	if peek == nil {
		return
	}
	switch peek.(type) {
	case float64:
		id = int64(peek.(float64))
	case float32:
		id = int64(peek.(float32))
	case int:
		id = int64(peek.(int))
	case int64:
		id = peek.(int64)
	case int32:
		id = int64(peek.(int32))
	case uint:
		id = int64(peek.(uint))
	case uint64:
		id = int64(peek.(uint64))
	case uint32:
		id = int64(peek.(uint32))
	case uint16:
		id = int64(peek.(uint16))
	case uint8:
		id = int64(peek.(uint8))
	default:
		// TODO
		return
	}
	ok = true
	return
}