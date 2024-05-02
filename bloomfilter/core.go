package bloomfilter

import (
	"hash/fnv"
	"math"
	"strconv"
)

type bFilter struct {
	predictEleCount int64
	errorPercent    float64

	eleCount  uint64
	hashCount int64

	data []int64
}

func New(predictEleCount int64, errorPercent float64) *bFilter {

	b := new(bFilter)
	b.predictEleCount = predictEleCount
	b.errorPercent = errorPercent

	b.eleCount = calcEleCount(predictEleCount, errorPercent)

	b.hashCount = int64(math.Max(calcHashCount(b.eleCount, predictEleCount), 1))

	dataCount := b.eleCount / 64

	if b.eleCount%64 > 0 {
		dataCount++
	}
	if dataCount < 1 {
		dataCount = 1
	}

	b.data = make([]int64, dataCount)

	return b
}

func (b *bFilter) Add(any string) bool {

	//hash
	hashs := make([]uint64, b.hashCount)
	new64 := fnv.New64()
	for i := int64(0); i < b.hashCount; i++ {
		new64.Reset()
		_, err := new64.Write([]byte(any + strconv.FormatInt(i, 10)))
		if err != nil {
			panic("add element error:" + any)
		}
		hashs[i] = new64.Sum64()
	}

	//insert
	exist := true

	for _, hash := range hashs {
		if !b.exist(hash) {
			b.set(hash)
			exist = false
		}
	}

	return exist
}

func (b *bFilter) Filter(any string) bool {
	hashs := make([]uint64, b.hashCount)
	new64 := fnv.New64()
	for i := int64(0); i < b.hashCount; i++ {
		new64.Reset()
		_, err := new64.Write([]byte(any + strconv.FormatInt(i, 10)))
		if err != nil {
			panic("add element error:" + any)
		}
		hashs[i] = new64.Sum64()
	}

	for _, hash := range hashs {
		if !b.exist(hash) {
			return false
		}
	}

	return true
}

func (b *bFilter) exist(hash uint64) bool {
	// [1-63] [64-127] [....]
	i := hash % uint64(len(b.data)*64) % uint64(len(b.data))
	return b.data[i]&int64(1<<i) != 0
}

func (b *bFilter) set(hash uint64) {
	i := hash % uint64(len(b.data)*64) % uint64(len(b.data))
	b.data[i] = b.data[i] | int64(1<<i)
}

func calcHashCount(count uint64, predictEleCount int64) float64 {
	f := float64(count) / float64(predictEleCount) * math.Ln2

	return f
}

func calcEleCount(count int64, percent float64) uint64 {
	f := math.Max(1, -float64(count)*math.Log(percent)/math.Ln2*math.Ln2)
	return uint64(f)

}
