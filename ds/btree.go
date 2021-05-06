package ds

import (
	"errors"
	"sort"
)

// 阶数m：最所有多少孩子节点
// 关键字数n：每个节点最多有多少key<=m-1

var (
	NeedPageSplit = errors.New("page has more than limit keys, need page split")
)

const (
	BTREE_LEVEL_NUM = 5
)

type BNode struct {
	Parent   *BNode
	Children []*BNode
	Records  []*BRecord
	IsLeaf   bool
}

type BRecord struct {
	prev  *BNode
	next  *BNode
	key   int
	value []byte
}

type BTree struct {
	Root *BNode
}

func NewBTree() *BTree {
	bt := new(BTree)
	bt.Root = NewBNode()
	bt.Root.IsLeaf = true
	return bt
}

func (b *BTree) Insert(key int, value []byte) error {
	// cursor seek
	c := b.Cursor()
	c.seek()

	leaf := b.FindLeaf()
	// 找到叶子节点
	record := NewBTreeRecord(key, value)
	// 判断当前叶子节点个数是否<=m-1
	err := leaf.insertLeafElement(record)
	if err != nil {
		return err
	}
	// 满足则插入结束，不满足进行分裂

	return nil
}

func (b *BTree) Find(key int) ([]byte, error) {
	return nil, nil
}

func (b *BTree) FindLeaf() *BNode {
	var leaf *BNode
	leaf = b.Root
	return leaf
}

// 初始化节点
func NewBNode() *BNode {
	bn := new(BNode)
	bn.Records = make([]*BRecord, 0)
	return bn
}

// 新生成叶子节点
func (bn *BNode) NewLeafNodeWithRecords(records []*BRecord) *BNode {
	if bn.IsLeaf {
		bn.IsLeaf = false
	}
	newBn := new(BNode)
	newBn.IsLeaf = true
	newBn.Records = append(newBn.Records, records...)
	bn.addChild(newBn)
	return newBn
}

// 插入children中，建立父子关系
func (bn *BNode) addChild(b *BNode) {
	bn.Children = append(bn.Children, b)
	b.Parent = bn
	return
}

func (bn *BNode) pageSplit() {
	// 获取中间key
	midRecord, pos := bn.midKey()
	// 分裂出左右两子节点
	left := bn.NewLeafNodeWithRecords(bn.getSubRecords(0, pos))
	right := bn.NewLeafNodeWithRecords(bn.getSubRecords(pos+1, len(bn.Records)))
	// 建立指针关系
	midRecord.prev = left
	midRecord.next = right
	// 父节点只含有中间节点
	bn.Records = []*BRecord{midRecord}
}

// 找到页分裂时的中间key
func (bn *BNode) midKey() (*BRecord, int) {
	size := len(bn.Records)
	var pos int
	if size%2 == 0 {
		pos = size/2 - 1
		return bn.Records[pos], pos
	}
	pos = size / 2
	return bn.Records[pos], pos
}

// 获取部分key
func (bn *BNode) getSubRecords(start, end int) []*BRecord {
	return bn.Records[start:end]
}

// 向叶子节点插入数据
func (bn *BNode) insertLeafElement(r *BRecord) error {
	bn.Records = append(bn.Records, r)
	sort.Slice(bn.Records, func(i, j int) bool {
		return bn.Records[i].key < bn.Records[j].key
	})
	if len(bn.Records) > BTREE_LEVEL_NUM-1 {
		//todo page-split
		bn.pageSplit()
	}
	return nil
}

func NewBTreeRecord(key int, value []byte) *BRecord {
	br := new(BRecord)
	br.key = key
	br.value = value
	return br
}

type Records []*BRecord

func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i, j int) bool {
	return r[i].key < r[j].key
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// 用于查找的游标
type Cursor struct {
	bTree *BTree
	stack []*BNode
}

func (c *Cursor) seek() {

}

func (b *BTree) Cursor() *Cursor {
	return &Cursor{
		bTree: b,
		stack: make([]*BNode, 0),
	}
}
