package toptr

import "testing"

func TestIntPtr(t *testing.T) {
	testCases := []struct {
		input    int
		expected *int
	}{
		{10, new(int)},
		{-10, new(int)},
		{0, new(int)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := IntPtr(tc.input)
		if *result != *tc.expected {
			t.Errorf("IntPtr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestInt8Ptr(t *testing.T) {
	testCases := []struct {
		input    int8
		expected *int8
	}{
		{10, new(int8)},
		{-10, new(int8)},
		{0, new(int8)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := Int8Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("Int8Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestInt16Ptr(t *testing.T) {
	testCases := []struct {
		input    int16
		expected *int16
	}{
		{10, new(int16)},
		{-10, new(int16)},
		{0, new(int16)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := Int16Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("Int16Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestInt32Ptr(t *testing.T) {
	testCases := []struct {
		input    int32
		expected *int32
	}{
		{10, new(int32)},
		{-10, new(int32)},
		{0, new(int32)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := Int32Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("Int32Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestInt64Ptr(t *testing.T) {
	testCases := []struct {
		input    int64
		expected *int64
	}{
		{10, new(int64)},
		{-10, new(int64)},
		{0, new(int64)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := Int64Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("Int64Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestUIntPtr(t *testing.T) {
	testCases := []struct {
		input    uint
		expected *uint
	}{
		{10, new(uint)},
		{0, new(uint)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := UIntPtr(tc.input)
		if *result != *tc.expected {
			t.Errorf("UIntPtr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestUInt8Ptr(t *testing.T) {
	testCases := []struct {
		input    uint8
		expected *uint8
	}{
		{10, new(uint8)},
		{0, new(uint8)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := UInt8Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("UInt8Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestUInt16Ptr(t *testing.T) {
	testCases := []struct {
		input    uint16
		expected *uint16
	}{
		{10, new(uint16)},
		{0, new(uint16)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := UInt16Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("UInt16Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestUInt32Ptr(t *testing.T) {
	testCases := []struct {
		input    uint32
		expected *uint32
	}{
		{10, new(uint32)},
		{0, new(uint32)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := UInt32Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("UInt32Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestUInt64Ptr(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected *uint64
	}{
		{10, new(uint64)},
		{0, new(uint64)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := UInt64Ptr(tc.input)
		if *result != *tc.expected {
			t.Errorf("UInt64Ptr(%d) = %d; want %d", tc.input, *result, *tc.expected)
		}
	}
}

func TestStringPtr(t *testing.T) {
	testCases := []struct {
		input    string
		expected *string
	}{
		{"test", new(string)},
		{"", new(string)},
	}

	for _, tc := range testCases {
		*tc.expected = tc.input
		result := StringPtr(tc.input)
		if *result != *tc.expected {
			t.Errorf("StringPtr(%q) = %q; want %q", tc.input, *result, *tc.expected)
		}
	}
}
