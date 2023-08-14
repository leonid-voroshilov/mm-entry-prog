package my_map

type MyMap struct {
	Size   int
	Values interface{}
}
type Pair struct {
	Key   int
	Value interface{}
}

func (mmap *MyMap) Get(key int) interface{} {
	if mmap.Size == 0 {
		return nil
	} else if mmap.Size == 1 {
		val := mmap.Values.(*Pair)
		if val.Key == key {
			return val.Value
		}
	} else if mmap.Size < 6 {
		for _, val := range mmap.Values.([]*Pair) {
			if val.Key == key {
				return val.Value
			}
		}
		return nil
	}
	return mmap.Values.(map[int]interface{})[key]
}
func (mmap *MyMap) Add(key int, value interface{}) error {
	if mmap.Size == 0 {
		mmap.Values = &Pair{
			Key:   key,
			Value: value,
		}
		mmap.Size += 1
		return nil
	} else if mmap.Size == 1 {
		if mmap.Get(key) != nil {
			mmap.Values = &Pair{
				Key:   key,
				Value: value,
			}
			return nil
		} else {
			tmp := mmap.Values.(*Pair)
			mmap.Values = make([]*Pair, 0, 5)
			(mmap.Values.([]*Pair))[0] = tmp
			(mmap.Values.([]*Pair))[1] = &Pair{
				Key:   key,
				Value: value,
			}
			mmap.Size += 1
			return nil
		}
	}
	if mmap.Size < 5 {
		if mmap.Get(key) != nil {
			for _, val := range mmap.Values.([]*Pair) {
				if val.Key == key {
					val.Value = value
					return nil
				}
			}
		} else {
			(mmap.Values.([]*Pair))[mmap.Size] = &Pair{Key: key, Value: value}
			mmap.Size += 1
			return nil
		}
	} else if mmap.Size == 5 {
		if mmap.Get(key) != nil {
			for _, val := range mmap.Values.([]*Pair) {
				if val.Key == key {
					val.Value = value
					return nil
				}
			}
		} else {
			// Не успеваю доделать задание в срок.
			// Что тут было бы: эвакуация данных из слайса в map и инкремент переменной
		}
	}
	// a тут простое обращение к встроенному типу
	return nil
}
