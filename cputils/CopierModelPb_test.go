package cputils

import (
	"database/sql"
	"testing"
	"time"
)

// TestCopyModelPb01 is a test function for CopyModelPb.
//
// It does not take any parameters and does not return anything.
func TestCopyModelPb01(t *testing.T) {
	type testStruct1 struct {
		TimeField time.Time
	}

	type testStruct2 struct {
		TimeField int64
	}

	type testStruct3 struct {
		TimeField sql.NullTime
	}

	fromValue := testStruct1{TimeField: time.Now()}
	toValue := testStruct2{}
	CopyModelPb(&toValue, &fromValue)

	// Test case for time.Time to int64 conversion
	t.Run("TimeField to IntField conversion", func(t *testing.T) {
		fromValue := testStruct1{TimeField: time.Now()}
		toValue := testStruct2{}
		CopyModelPb(&toValue, &fromValue)

		if toValue.TimeField != fromValue.TimeField.Unix() {
			t.Errorf("Expected IntField to be %d, got %d", fromValue.TimeField.Unix(), toValue.TimeField)
		}
	})

	// Test case for sql.NullTime to int64 conversion
	t.Run("NullTimeField to IntField conversion", func(t *testing.T) {
		nullTime := sql.NullTime{Time: time.Now(), Valid: true}
		fromValue := testStruct3{TimeField: nullTime}
		toValue := testStruct2{}
		CopyModelPb(&toValue, &fromValue)

		if toValue.TimeField != nullTime.Time.Unix() {
			t.Errorf("Expected IntField to be %d, got %d", nullTime.Time.Unix(), toValue.TimeField)
		}
	})

	// Add more test cases as needed
}

// TestCopierModelPb is a Go function that tests the CopyModelPb function by comparing the fields of two struct types.
//
// The function takes a testing.T parameter and does not return anything.
func TestCopierModelPb(t *testing.T) {
	type SFrom struct {
		Id                  int
		Id2                 int32
		Id3                 int64
		FileName            string
		CreateAtTimeStamp   int64
		CreateAt            time.Time
		UploadSize          sql.NullInt64
		DomainName          sql.NullString
		CollectionStartTime sql.NullTime
	}

	type TTo struct {
		Id                  int
		Id2                 int32
		Id3                 int64
		FileName            string
		CreateAt            int64
		UploadSize          int64
		DomainName          string
		CollectionStartTime int64
	}

	from := make([]SFrom, 0)
	{
		currentTimeStamp := time.Now()
		m1 := SFrom{
			Id:                  123,
			Id2:                 456,
			Id3:                 789,
			FileName:            "foo.txt",
			CreateAtTimeStamp:   currentTimeStamp.Unix(),
			CreateAt:            currentTimeStamp,
			UploadSize:          sql.NullInt64{Int64: 11111111111111, Valid: true},
			DomainName:          sql.NullString{String: "mydomainname", Valid: true},
			CollectionStartTime: sql.NullTime{Time: currentTimeStamp, Valid: true},
		}
		from = append(from, m1)
	}

	to := make([]TTo, 0)

	CopyModelPb(&to, &from)

	onefrom := from[0]
	oneto := to[0]

	if onefrom.Id != oneto.Id {
		t.Errorf("onefrom.Id = %v, oneto.Id = %v", onefrom.Id, oneto.Id)
	}
	if onefrom.Id2 != oneto.Id2 {
		t.Errorf("onefrom.Id2 = %v, oneto.Id2 = %v", onefrom.Id2, oneto.Id2)
	}
	if onefrom.Id3 != oneto.Id3 {
		t.Errorf("onefrom.Id3 = %v, oneto.Id3 = %v", onefrom.Id3, oneto.Id3)
	}
	if onefrom.FileName != oneto.FileName {
		t.Errorf("onefrom.FileName = %v, oneto.FileName = %v", onefrom.FileName, oneto.FileName)
	}
	if onefrom.CreateAtTimeStamp != oneto.CreateAt {
		t.Errorf("onefrom.CreateAtTimeStamp = %v, oneto.CreateAt = %v", onefrom.CreateAtTimeStamp, oneto.CreateAt)
	}
	if onefrom.UploadSize.Int64 != oneto.UploadSize {
		t.Errorf("onefrom.UploadSize.Int64 = %v, oneto.UploadSize = %v", onefrom.UploadSize.Int64, oneto.UploadSize)
	}
	if onefrom.DomainName.String != oneto.DomainName {
		t.Errorf("onefrom.DomainName.String = %v, oneto.DomainName = %v", onefrom.DomainName.String, oneto.DomainName)
	}
	if onefrom.CreateAtTimeStamp != oneto.CollectionStartTime {
		t.Errorf("onefrom.CreateAtTimeStamp = %v, oneto.CollectionStartTime = %v", onefrom.CreateAtTimeStamp, oneto.CollectionStartTime)
	}
}
