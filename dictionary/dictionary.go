package dictionary

import "errors"

// Dictionary : 自定义字典类型
type Dictionary map[string]string

// Search : 查找字典键值
func (dictionary Dictionary) Search(word string) (string, error) {
	value, exist := dictionary[word]

	if !exist {
		return "", errors.New("could not find the word")
	}
	return value, nil
}
