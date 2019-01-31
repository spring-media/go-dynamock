package dynamock

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
)

type AttributeValueExpect struct {
	AttributeValue *dynamodb.AttributeValue
	Compare        Compare
}

type Compare string

const (
	CompareKeyExists Compare = "KEY"
	CompareKeyValue  Compare = "KEY_VALUE"
	CompareNotEmpty  Compare = "NOT_EMPTY"
)

func (AttributeValueExpect) CreateAttributeValueExpect(attribute *dynamodb.AttributeValue, compare Compare) AttributeValueExpect {
	return AttributeValueExpect{AttributeValue: attribute, Compare: compare}
}

func Equals(x map[string]*AttributeValueExpect, y map[string]*dynamodb.AttributeValue) (bool, error) {
	if len(x) != len(y) {
		return false, fmt.Errorf("length of array not equal")
	}

	for k, v := range x {
		if val, ok := y[k]; ok {
			switch v.Compare {
			case CompareKeyExists:
				break
			case CompareKeyValue:
				if !reflect.DeepEqual(v.AttributeValue, val) {
					return false, fmt.Errorf("value %v do not equals %v", v.AttributeValue, val)
				}
				break
			case CompareNotEmpty:
				if len(val.String()) == 0 {
					return false, fmt.Errorf("value for %v cannot be empty", k)
				}
				break
			}
		} else {
			return false, fmt.Errorf("key %s not found", k)
		}
	}
	return true, nil
}
