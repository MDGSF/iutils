package toptr

func IntPtr(i int) *int          { return &i }
func Int8Ptr(i int8) *int8       { return &i }
func Int16Ptr(i int16) *int16    { return &i }
func Int32Ptr(i int32) *int32    { return &i }
func Int64Ptr(i int64) *int64    { return &i }
func UIntPtr(i uint) *uint       { return &i }
func UInt8Ptr(i uint8) *uint8    { return &i }
func UInt16Ptr(i uint16) *uint16 { return &i }
func UInt32Ptr(i uint32) *uint32 { return &i }
func UInt64Ptr(i uint64) *uint64 { return &i }
func StringPtr(s string) *string { return &s }
