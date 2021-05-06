package ds

import "errors"

var (
	KeyAlreadyExist = errors.New("key already exist")
)

const (
	MAX_KEY_NUM = 2
)

// B+树 节点实例
type Node struct {
	Pointers []interface{}
	Keys     []int
	Parent   *Node
	IsLeaf   bool
	NumKeys  int
	Next     *Node
}

type Record struct {
	Value []byte
}

// B+树 主体
type BPTree struct {
	Root *Node
}

func NewBPlusTree() *BPTree {
	return new(BPTree)
}

// B+树 插入
func (bp *BPTree) Insert(key, value []byte) error {
	//var pointer *Record
	//var leaf *Node
	// 先判断key是否已存在
	if _, err := bp.Find(key); err == nil {
		return KeyAlreadyExist
	}
	// 生成record结构体
	//record, err := makeRecord(value)
	//if err != nil {
	//	return err
	//}

	// 判断根节点
	if bp.Root == nil {
		//todo new bp-tree
	}

	return nil
}

// B+树 查找
func (bp *BPTree) Find(key []byte) (*Record, error) {
	return nil, nil
}

// B+树 删除
func (bp BPTree) Delete(key []byte) error {
	return nil
}

// B+树 树高
func (bp *BPTree) height() int {
	return 0
}

func makeRecord(value []byte) (*Record, error) {
	nr := new(Record)
	nr.Value = value
	return nr, nil
}
