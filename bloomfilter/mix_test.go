package bloomfilter

import (
	"fmt"
	"hash/fnv"
	"testing"
)

func TestHash(t *testing.T) {

	new64 := fnv.New64()
	new64.Write([]byte("hlllo"))
	fmt.Println(new64.Sum64())

	new64.Reset()
	new64.Write([]byte("hlllo"))
	fmt.Println(new64.Sum64())

	new64.Reset()
	new64.Write([]byte("hlllo1"))
	fmt.Println(new64.Sum64())
}

func TestBloomFilter(t *testing.T) {

	bf := New(100, 0.01)

	bf.Add("hello")
	bf.Add("hello1")
	bf.Add("hello2")

	//if !bf.Add("hello2") {
	//	t.Errorf("Add(hello2),false")
	//}

	if !bf.Filter("hello2") {
		t.Errorf("Filter(hello2),false")
	}
	if !bf.Filter("hello") {
		t.Errorf("Filter(hello),false")
	}
	if bf.Filter("hello111") {
		t.Errorf("Filter(hello111),true")
	}

}
