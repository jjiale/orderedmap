package orderedmap

type ListNode struct {
	before *ListNode
	after  *ListNode
	key    interface{}
	val    interface{}
}

type OrderedMap struct {
	dataMap     map[interface{}]interface{}
	head        *ListNode
	accessOrder bool // false按插入顺序排序，true按访问顺序
}

func NewOrderedMap() *OrderedMap {
	lhm := &OrderedMap{make(map[interface{}]interface{}), nil, false}

	return lhm
}

func (t *OrderedMap) Set(k, v interface{}) bool {
	// 往尾部插入
	t.addBefore(k, v)
	t.dataMap[k] = v
	return true
}

func (t *OrderedMap) addBefore(k, v interface{}) {
	// 往尾部插入
	node := &ListNode{
		key: k,
		val: v,
	}

	// 查找一波
	if _, ok := t.dataMap[k]; ok {
		if t.accessOrder == false {
			return
		}

		// 在链表里删除
		t.DelListNode(k)
	}
	// 插入头部
	t.head, node = node, t.head
	t.head.before = node

	if node != nil {
		node.after = t.head
	}
}

func (t *OrderedMap) DelListNode(k interface{}) {
	node := t.head
	for node != nil {
		if node.key != k {
			node = node.before
			continue
		}

		if node.before != nil {
			node.before.after = node.after
		}
		if node.after != nil {
			node.after.before = node.before
		}
		if node.before == nil && node.after == nil {
			t.head = nil
		}
		break
	}
	return
}

func (t *OrderedMap) Get(k interface{}) (interface{}, bool) {
	v, ok := t.dataMap[k]
	if t.accessOrder && ok {
		t.addBefore(k, v)
	}

	return v, ok
}

func (t OrderedMap) Keys() []interface{} {
	keys := []interface{}{}
	if t.head == nil {
		return keys
	}

	node := t.head
	for node != nil {
		if t.accessOrder {
			keys = append(keys, node.key)
		} else {
			keys = append([]interface{}{node.key}, keys...)
		}
		node = node.before
	}
	return keys
}

func (t OrderedMap) Del(k interface{}) {
	t.DelListNode(k)
	delete(t.dataMap, k)
}

// -------------- 快速计数 ----------
func (t *OrderedMap) Add(k interface{}, num int) bool {
	v, ok := t.Get(k)
	if !ok {
		return t.Set(k, num)
	}
	oriNum, ok := v.(int)
	if !ok {
		return false
	}

	return t.Set(k, oriNum+num)
}

func (t *OrderedMap) Sub(k interface{}, num int) bool {
	return t.Add(k, -num)
}
