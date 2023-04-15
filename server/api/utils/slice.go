package utils

import "reflect"

func SliceArray[T any](arr []T, limit int32, offset int32) []T {
	start := int(limit * offset)
	end := int(limit*offset + limit)

	if start < 0 {
		start = 0
	}

	if end > len(arr) {
		end = len(arr)
	}

	if start > end {
		return nil
	}

	return arr[start:end]
}

func FilterByField(items interface{}, field string, value interface{}) interface{} {
	filteredItems := reflect.MakeSlice(reflect.TypeOf(items), 0, 0)

	itemsValue := reflect.ValueOf(items)
	for i := 0; i < itemsValue.Len(); i++ {
		item := itemsValue.Index(i)
		fieldValue := item.FieldByName(field).Interface()

		if fieldValue == value {
			filteredItems = reflect.Append(filteredItems, item)
		}
	}

	return filteredItems.Interface()
}
