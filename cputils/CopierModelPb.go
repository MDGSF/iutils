package cputils

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jinzhu/copier"
)

// CopyModelPb copies values from one struct to another with conversions for specific types.
//
// toValue: destination struct to copy values into.
// fromValue: source struct to copy values from.
func CopyModelPb(toValue interface{}, fromValue interface{}) {
	const Int64 int64 = 0
	var TimePtr *time.Time = nil
	var StringPtr *string = nil
	var Int64Ptr *int64 = nil
	var Int32Ptr *int32 = nil
	copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty:   true,
		CaseSensitive: false,
		DeepCopy:      true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: Int64,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					return s.Unix(), nil
				},
			},
			{
				SrcType: sql.NullTime{},
				DstType: Int64,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(sql.NullTime)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					return s.Time.Unix(), nil
				},
			},
			{
				SrcType: TimePtr,
				DstType: Int64,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(*time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					if s == nil {
						return 0, nil
					}

					return s.Unix(), nil
				},
			},
			{
				SrcType: StringPtr,
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(*string)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					if s == nil {
						return "", nil
					}

					return *s, nil
				},
			},
			{
				SrcType: Int64Ptr,
				DstType: Int64,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(*int64)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					if s == nil {
						return int64(0), nil
					}

					return *s, nil
				},
			},
			{
				SrcType: Int32Ptr,
				DstType: copier.Int,
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(*int32)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					if s == nil {
						return int(0), nil
					}

					return int(*s), nil
				},
			},
		},
	})
}
