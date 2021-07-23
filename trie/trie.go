package trie

import (
	"fmt"
	"testing"
)

const maxCharInt = 128

type Trie struct {
	isWord   bool
	elem     interface{}
	children [maxCharInt]*Trie
}

//NewTrieTee create a trietree.
func NewTrieTee() *Trie {
	return &Trie{}
}

// Inserts a word into the trietree.
func (t *Trie) Insert(word string, elem interface{}) error {
	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i]
		if c >= maxCharInt || c <= 0 {
			return fmt.Errorf("char [%c] is not support for trietree", c)
		}

		if cur.children[c] == nil {
			cur.children[c] = &Trie{}
		}
		cur = cur.children[c]
		if i == 0 {
			cur.isWord = true
			cur.elem = elem
		}
	}
	return nil
}

func deleteFromNode(t *Trie, c byte) error {
	if t == nil {
		return fmt.Errorf("can not find node to delete")
	}
	if c >= maxCharInt || c <= 0 {
		return fmt.Errorf("char [%c] is not support for trietree", c)
	}
	t.children[c] = nil
	return nil
}

func (t *Trie) setNode2noword(word string) error {
	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i]
		if c >= maxCharInt || c <= 0 {
			return fmt.Errorf("char [%c] is not support for trietree", c)
		}

		if cur.children[c] == nil {
			return fmt.Errorf("can not find node to delete")
		}
		cur = cur.children[c]
	}

	cur.isWord = false
	cur.elem = nil
	return nil
}

func (t *Trie) Delete(word string) error {
	node, c := t.getDeleteFromNode(word)
	if node == nil {
		return t.setNode2noword(word)
	} else {
		return deleteFromNode(node, c)
	}
}

// getDeleteFromNode 找到最浅的可删除节点
func (t *Trie) getDeleteFromNode(word string) (*Trie, byte) {
	wordlen := len(word)
	if wordlen == 0 {
		return nil, 0
	}

	cur := t
	nodelist := make([]*Trie, wordlen+1)
	childcountlist := make([]int, wordlen+1)

	for i := wordlen - 1; i >= 0; i-- {
		c := word[i]
		if c >= maxCharInt || c <= 0 {
			return nil, 0
		}

		if cur.children[c] == nil {
			return nil, 0
		}
		nodelist[i] = cur
		for j := 0; j < maxCharInt; j++ {
			if cur.children[j] != nil || cur.isWord {
				childcountlist[i+1]++
			}
		}
		cur = cur.children[c]
	}

	// 不是叶子，不需要删除
	if !cur.isWord {
		return nil, 0
	}

	// 还有子节点，不需要删除
	for j := 0; j < maxCharInt; j++ {
		if cur.children[j] != nil {
			return nil, 0
		}
	}

	// 查找最浅可删除节点
	for i := 0; i < wordlen; i++ {
		if childcountlist[i] > 1 {
			//fmt.Printf("Delete from [%d]/[%c] %d, delete[%c]\n", i, word[i], childcountlist[i], word[i-1])
			return nodelist[i-1], word[i-1]
		}
	}

	//fmt.Printf("Delete from [root], delete[%c]\n", word[wordlen-1])
	return t, word[wordlen-1]
}

// Returns if the word is in the trietree.
func (t *Trie) Search(word string) (bool, interface{}) {
	cur := t
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i]
		if c >= maxCharInt || c <= 0 {
			return false, nil
		}

		if cur.children[c] == nil {
			return false, nil
		}
		cur = cur.children[c]
	}
	return cur.isWord, cur.elem
}

// Returns if the word is in the trietree.
func (t *Trie) EndsSearch(word string) (bool, interface{}) {
	cur := t
	find := false
	var elem interface{}
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i]
		if c >= maxCharInt || c <= 0 {
			return false, nil
		}

		if cur.children[c] == nil {
			if find {
				return true, elem
			}
			return false, nil
		}
		cur = cur.children[c]
		if cur.isWord == true {
			if i > 0 && word[i-1] == '.' {
				find = true
				elem = cur.elem
			}
		}
	}
	return cur.isWord, cur.elem
}

// Display the trietree.
func (t *Trie) Display(pre string) {
	cur := t
	for i := 0; i < maxCharInt-1; i++ {
		if cur.children[i] != nil {
			fmt.Printf("%s %c %v %v\n", pre, i, cur.children[i].isWord, cur.children[i].elem)
			cur.children[i].Display("-" + pre)
		}
	}
}

func (t *Trie) Destory() {
	for i := 0; i < maxCharInt; i++ {
		if t.children[i] != nil {
			t.children[i].Destory()
			t.children[i] = nil
		}
	}
}

func Run() {
	var t *testing.T
	TestTrieTree(t)
}

func TestTrieTree(t *testing.T) {
	trieTree := NewTrieTee()
	if trieTree == nil {
		t.Error("NewTrieTee Error")
	}

	err := trieTree.Insert("a.b.c.m", 11)
	err = trieTree.Insert("d.e.f.n", 12)
	err = trieTree.Insert("c.f.d.k", 13)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}
	trieTree.Display("www")
	trieTree.Delete("a.b.c.m")
	trieTree.Display("www")

	/*err = trieTree.Insert("a.b.c.d", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("a.a.b.c.d", 1001)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("a.a.a.b.c.d", 1002)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("1.2.c.d", 4)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("g.f.e.d", 2)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("*.p.e.d", 5)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	//trieTree.Display("-")

	f, d := trieTree.Search("a.b.c.d")
	if !f || d != 1 {
		t.Error("TrieTee Search Error", f, d)
	}

	f, d = trieTree.EndsSearch("aa.a.b.c.d")
	if !f || d != 1 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("aaa.b.c.d")
	if f {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("aa.a.a.b.c.d")
	if !f || d != 1001 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("a.a.b.c.d")
	if !f || d != 1001 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}
	f, d = trieTree.EndsSearch("aa.a.b.c.d")
	if !f || d != 1 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("aaa.a.b.c.d")
	if !f || d != 1 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("aaa.a.a.a.b.c.d")
	if !f || d != 1002 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}
	f, d = trieTree.EndsSearch("aaaa.a.a.b.c.d")
	if !f || d != 1001 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("e.a.b.c.d")
	if !f || d != 1 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.EndsSearch("e.a1.b.c.d")
	if f {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	f, d = trieTree.Search("1.2.c.d")
	if !f || d != 4 {
		t.Error("TrieTee Search Error", f, d)
	}

	f, d = trieTree.Search("3.2.c.d")
	if f {
		t.Error("TrieTee Search Error", f, d)
	}

	err = trieTree.Insert("0zonetest.com", 51)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("test.com", 41)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	f, d = trieTree.EndsSearch("0zonetest.com")
	if !f || d != 51 {
		t.Error("TrieTee EndsSearch Error", f, d)
	}

	trieTree.Display("-")
	trieTree.Destory()
	trieTree.Display("+")

	// if f {
	// 	t.Error("TrieTee Search Error", f, d)
	// }

	// f, d = trieTree.WildcardSearch(".f.e.d")
	// if f {
	// 	t.Error("TrieTee Search Error", f, d)
	// }*/

}

func TestDeleteTree(t *testing.T) {
	trieTree := NewTrieTee()
	if trieTree == nil {
		t.Error("NewTrieTee Error")
	}

	err := trieTree.Insert("www.baidu.com", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("cache.baidu.com", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("db.baidu.com", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	//trieTree.GetDeleteFromNode("www.baidu.com")

	err = trieTree.Insert("api.baidu.com", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("gtm.api.baidu.com", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("dns.api.baidu.com", 3)
	err = trieTree.Insert("s.api.baidu.com", 4)
	err = trieTree.Insert("mns.api.baidu.com", 12)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	f, d := trieTree.Search("dns.api.baidu.com")
	if !f || d != 3 {
		t.Error("TrieTee Search Error", f, d)
	}

	//trieTree.Display("+")
	trieTree.Delete("dns.api.baidu.com")
	//trieTree.Display("+")
	f, d = trieTree.Search("mns.api.baidu.com")
	if !f || d != 12 {
		t.Error("TrieTee Search Error", f, d)
	}
	f, d = trieTree.Search("s.api.baidu.com")
	if !f || d != 4 {
		t.Error("TrieTee Search Error", f, d)
	}

	f, _ = trieTree.Search("dns.api.baidu.com")
	if f {
		t.Error("TrieTee Search Error", f)
	}

	err = trieTree.Insert("x.y.z", 198)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	f, d = trieTree.Search("x.y.z")
	if !f || d != 198 {
		t.Error("TrieTee Search Error", f, d)
	}

	trieTree.Delete("x.y.z")

	f, d = trieTree.Search("x.y.z")
	if f {
		t.Error("TrieTee Search Error", f)
	}
}

func TestDeleteTree2(t *testing.T) {
	trieTree := NewTrieTee()
	if trieTree == nil {
		t.Error("NewTrieTee Error")
	}

	err := trieTree.Insert("sub1.baidu.com", 1)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("baidu.com", 2)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	err = trieTree.Insert("sub2.baidu.com", 3)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	f, x := trieTree.EndsSearch("www.sub1.baidu.com")
	if !f || x != 1 {
		t.Error("TrieTee Search Error", f)
	}
	f, x = trieTree.EndsSearch("www.baidu.com")
	if !f || x != 2 {
		t.Error("TrieTee Search Error", f)
	}

	f, x = trieTree.EndsSearch("wbaidu.com")
	if f {
		t.Error("TrieTee Search Error", f)
	}

	f, x = trieTree.EndsSearch("wwwsub1.baidu.com")
	if !f || x != 2 {
		t.Error("TrieTee Search Error", f)
	}

	f, x = trieTree.EndsSearch("www.sub2.baidu.com")
	if !f || x != 3 {
		t.Error("TrieTee Search Error", f)
	}

	err = trieTree.Delete("baidu.com")
	if err != nil {
		t.Error("TrieTee Delete Error")
	}

	f, x = trieTree.EndsSearch("www.baidu.com")
	if f {
		t.Error("TrieTee Search Error", f)
	}

	err = trieTree.Insert("baidu.com", 2)
	if err != nil {
		t.Error("TrieTee Insert Error")
	}

	f, x = trieTree.EndsSearch("www.baidu.com")
	if !f || x != 2 {
		t.Error("TrieTee Search Error", f)
	}

	f, x = trieTree.EndsSearch("baidu.com")
	if !f || x != 2 {
		t.Error("TrieTee Search Error", f)
	}

	f, x = trieTree.EndsSearch("www.sub1.baidu.com")
	if !f || x != 1 {
		t.Error("TrieTee Search Error", f)
	}
	err = trieTree.Delete("sub1.baidu.com")
	if err != nil {
		t.Error("TrieTee Delete Error")
	}
	f, x = trieTree.EndsSearch("www.sub1.baidu.com")
	if !f || x != 2 {
		t.Error("TrieTee Search Error", f)
	}

}

func TestDeleteTree3(t *testing.T) {
	trieTree := NewTrieTee()
	trieTree.Insert("sub1.c.com", 1)
	trieTree.Insert("sub1.a.com", 1)
	trieTree.Insert("sux1.a.com", 1)
	trieTree.Insert("sub2.a.com", 1)
	trieTree.Display("-")
	trieTree.Delete("sub1.a.com")
	trieTree.Display("-")
	trieTree.Delete("sub1.c.com")
	trieTree.Display("-")
}
