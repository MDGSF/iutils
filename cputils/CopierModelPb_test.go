package cputils

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/MDGSF/iutils/toptr"
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
	err := CopyModelPb(&toValue, &fromValue)
	if err != nil {
		t.Errorf("CopyModelPb failed: %v", err)
	}

	// Test case for time.Time to int64 conversion
	t.Run("TimeField to IntField conversion", func(t *testing.T) {
		fromValue := testStruct1{TimeField: time.Now()}
		toValue := testStruct2{}
		err := CopyModelPb(&toValue, &fromValue)
		if err != nil {
			t.Errorf("CopyModelPb failed: %v", err)
		}

		if toValue.TimeField != fromValue.TimeField.Unix() {
			t.Errorf("Expected IntField to be %d, got %d", fromValue.TimeField.Unix(), toValue.TimeField)
		}
	})

	// Test case for sql.NullTime to int64 conversion
	t.Run("NullTimeField to IntField conversion", func(t *testing.T) {
		nullTime := sql.NullTime{Time: time.Now(), Valid: true}
		fromValue := testStruct3{TimeField: nullTime}
		toValue := testStruct2{}
		err := CopyModelPb(&toValue, &fromValue)
		if err != nil {
			t.Errorf("CopyModelPb failed: %v", err)
		}

		if toValue.TimeField != nullTime.Time.Unix() {
			t.Errorf("Expected IntField to be %d, got %d", nullTime.Time.Unix(), toValue.TimeField)
		}
	})

	// Add more test cases as needed
}

// TestCopyModelPb02 is a Go function that tests the CopyModelPb function by comparing the fields of two struct types.
//
// The function takes a testing.T parameter and does not return anything.
func TestCopyModelPb02(t *testing.T) {
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

	err := CopyModelPb(&to, &from)
	if err != nil {
		t.Errorf("CopyModelPb failed: %v", err)
	}

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

func TestCopyModelPb03(t *testing.T) {
	type Model struct {
		Time     time.Time
		NullTime sql.NullTime
		TimePtr  *time.Time
		String   string
		Int64    int64
		Int32    int32
	}

	current_time := time.Now()

	model := Model{
		Time: time.Now(),
		NullTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		TimePtr: &current_time,
		String:  "test",
		Int64:   123456789,
		Int32:   123,
	}

	expected := Model{
		Time:     model.Time,
		NullTime: model.NullTime,
		TimePtr:  model.TimePtr,
		String:   model.String,
		Int64:    model.Int64,
		Int32:    model.Int32,
	}

	var toValue Model
	err := CopyModelPb(&toValue, model)
	if err != nil {
		t.Errorf("CopyModelPb failed: %v", err)
	}

	if !reflect.DeepEqual(toValue, expected) {
		t.Errorf("CopyModelPb() failed, expected %+v, got %+v", expected, toValue)
	}
}

func TestCopyModelPb04(t *testing.T) {
	current_time := time.Now()
	current_time_ts := current_time.Unix()

	type FromModel struct {
		Time     time.Time
		NullTime sql.NullTime
		TimePtr  *time.Time
		String   *string
		Int64    *int64
		Int32    *int32
		Int      *int
	}

	fromModel := FromModel{
		Time: current_time,
		NullTime: sql.NullTime{
			Time:  current_time,
			Valid: true,
		},
		TimePtr: &current_time,
		String:  toptr.StringPtr("test"),
		Int64:   toptr.Int64Ptr(123456789),
		Int32:   toptr.Int32Ptr(123),
		Int:     toptr.IntPtr(456),
	}

	type ToModel struct {
		Time     int64
		NullTime int64
		TimePtr  int64
		String   string
		Int64    int64
		Int32    int32
		Int      int
	}

	expected := ToModel{
		Time:     current_time_ts,
		NullTime: current_time_ts,
		TimePtr:  current_time_ts,
		String:   "test",
		Int64:    123456789,
		Int32:    123,
		Int:      456,
	}

	var toModel ToModel
	err := CopyModelPb(&toModel, fromModel)
	if err != nil {
		t.Errorf("CopyModelPb() failed, expected no error, got %v", err)
	}

	if !reflect.DeepEqual(toModel, expected) {
		t.Errorf("CopyModelPb() failed, expected %+v, got %+v", expected, toModel)
	}
}

func TestCopyModelPb05(t *testing.T) {
	current_time := time.Now()

	type FromModel struct {
		NullTime sql.NullTime
	}

	fromModel := FromModel{
		NullTime: sql.NullTime{
			Time:  current_time,
			Valid: false,
		},
	}

	type ToModel struct {
		NullTime int64
	}

	expected := ToModel{
		NullTime: 0,
	}

	var toModel ToModel
	err := CopyModelPb(&toModel, fromModel)
	if err != nil {
		t.Errorf("CopyModelPb() failed, expected no error, got %v", err)
	}

	if !reflect.DeepEqual(toModel, expected) {
		t.Errorf("CopyModelPb() failed, expected %+v, got %+v", expected, toModel)
	}
}

func TestCopyModelPb06(t *testing.T) {
	type FromModel struct {
		TimePtr *time.Time
	}

	fromModel := FromModel{
		TimePtr: nil,
	}

	type ToModel struct {
		TimePtr int64
	}

	expected := ToModel{
		TimePtr: 0,
	}

	var toModel ToModel
	err := CopyModelPb(&toModel, fromModel)
	if err != nil {
		t.Errorf("CopyModelPb() failed, expected no error, got %v", err)
	}

	if !reflect.DeepEqual(toModel, expected) {
		t.Errorf("CopyModelPb() failed, expected %+v, got %+v", expected, toModel)
	}
}
