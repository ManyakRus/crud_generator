package contract_white_items

import "github.com/vmihailenco/msgpack/v5"

// NewWhiteListItem -- Новая запись белого списка
func NewWhiteListItem() ContractWhiteItem {
	return ContractWhiteItem{}
}

func AsWhiteListItem(b []byte) (ContractWhiteItem, error) {
	c := NewWhiteListItem()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewWhiteListItem(), err
	}
	return c, nil
}

func WhiteListItemAsBytes(c *ContractWhiteItem) ([]byte, error) {
	b, err := msgpack.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}
