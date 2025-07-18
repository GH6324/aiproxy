package fastjsonserializer

import (
	"context"
	"fmt"
	"reflect"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/common/conv"
	"gorm.io/gorm/schema"
)

type JSONSerializer struct{}

func (*JSONSerializer) Scan(
	ctx context.Context,
	field *schema.Field,
	dst reflect.Value,
	dbValue any,
) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = conv.StringToBytes(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}

		if len(bytes) == 0 {
			field.ReflectValueOf(ctx, dst).Set(reflect.Zero(field.FieldType))
			return nil
		}

		err = sonic.Unmarshal(bytes, fieldValue.Interface())
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())

	return
}

func (*JSONSerializer) Value(
	_ context.Context,
	_ *schema.Field,
	_ reflect.Value,
	fieldValue any,
) (any, error) {
	return sonic.Marshal(fieldValue)
}

func init() {
	schema.RegisterSerializer("fastjson", new(JSONSerializer))
}
