# 简介
go语言的orderedmap结构

## 安装

go get -u github.com/jjiale/orderedmap

## 使用

支持 Set, Get, Delete,Add,Sub 等方法:

```go
m := orderedmap.NewOrderedMap()
m.Set("foo", "bar")
m.Get("foo")
m.Del("foo")
m.Set("foo",1)
m.Add("foo",1)  // foo = 2
```

### 迭代

```go
for _, key := range m.Keys() {
	value, _:= m.Get(key)
	fmt.Println(key, value)
}
```
