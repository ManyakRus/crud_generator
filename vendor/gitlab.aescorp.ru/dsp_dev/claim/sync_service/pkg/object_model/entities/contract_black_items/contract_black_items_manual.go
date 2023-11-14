package contract_black_items

import "github.com/vmihailenco/msgpack/v5"

// NewBlackListItem -- Новая запись чёрного списка
func NewBlackListItem() ContractBlackItem {
	return ContractBlackItem{}
}

func AsBlackListItem(b []byte) (ContractBlackItem, error) {
	c := NewBlackListItem()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewBlackListItem(), err
	}
	return c, nil
}

func BlackListItemAsBytes(c *ContractBlackItem) ([]byte, error) {
	b, err := msgpack.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}
