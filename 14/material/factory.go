package material

import (
	"math"
)

type MaterialTable map[string]*MaterialTransform

func NewMaterialTable() MaterialTable {
	return MaterialTable{}
}

type MaterialMap map[string]int

func (r MaterialMap) Add(material string, count int) {
	r[material] += count
}
func (r MaterialMap) AddMaterialCount(mc MaterialCounter) {
	r.Add(mc.Material, mc.Count)
}
func (r MaterialMap) Sub(material string, count int) {
	r[material] -= count
}
func (r MaterialMap) SubMaterialCounter(mc MaterialCounter) {
	r.Sub(mc.Material, mc.Count)
}
func (r MaterialMap) HasMaterial(mc MaterialCounter) bool {
	return r[mc.Material] >= mc.Count
}

type MaterialQueue struct {
	Queue []MaterialCounter
	Head  int
}

func NewMaterialQueue() *MaterialQueue {
	return &MaterialQueue{
		make([]MaterialCounter, 0),
		-0,
	}
}
func (q *MaterialQueue) Finished() bool {
	return q.Head >= len(q.Queue)
}
func (q *MaterialQueue) Produce(t MaterialTransform, amountWish int) MaterialCounter {
	produced := t.Produces
	multiplier := int(math.Ceil(float64(amountWish) / float64(produced.Count)))

	for _, consumed := range t.Consumes {
		consumed.Count *= multiplier
		q.Queue = append(q.Queue, consumed)
	}

	produced.Count *= multiplier
	return produced
}
func (q *MaterialQueue) Consume() MaterialCounter {
	consumed := q.Queue[q.Head]
	q.Head++
	return consumed
}

type MaterialFactory struct {
	Table MaterialTable
	Belt  *MaterialQueue
	Stock MaterialMap
}

func NewMaterialFactory(mt MaterialTable) MaterialFactory {
	return MaterialFactory{
		Table: mt,
		Belt:  NewMaterialQueue(),
		Stock: make(MaterialMap),
	}
}

func (f MaterialFactory) initBelt() {
	f.Belt = NewMaterialQueue()
}

func (f MaterialFactory) ProduceRecursive(needed MaterialCounter) MaterialCounter {
	mTrans := f.Table[needed.Material]

	if mTrans != nil {
		for !f.Stock.HasMaterial(needed) {
			amountWish := needed.Count - f.Stock[needed.Material]
			produced := f.Belt.Produce(*mTrans, amountWish)
			f.Stock.AddMaterialCount(produced)
		}
	}
	if !f.Stock.HasMaterial(needed) {
		return MaterialCounter{needed.Material, f.Stock[needed.Material]}
	}
	f.Stock.SubMaterialCounter(needed)

	for !f.Belt.Finished() {
		stillNeedsToBeProduced := f.Belt.Consume()
		f.ProduceRecursive(stillNeedsToBeProduced)
	}
	return needed
}

func (f MaterialFactory) Produce(needed MaterialCounter) MaterialCounter {
	f.initBelt()
	mTrans := f.Table[needed.Material]
	f.Belt.Produce(*mTrans, needed.Count)

	for !f.Belt.Finished() {
		stillNeedsToBeProduced := f.Belt.Consume()

		mTrans := f.Table[stillNeedsToBeProduced.Material]

		if mTrans != nil {
			for !f.Stock.HasMaterial(stillNeedsToBeProduced) {
				amountWish := stillNeedsToBeProduced.Count - f.Stock[stillNeedsToBeProduced.Material]
				produced := f.Belt.Produce(*mTrans, amountWish)
				f.Stock.AddMaterialCount(produced)
			}
		}
		if !f.Stock.HasMaterial(stillNeedsToBeProduced) {
			return MaterialCounter{needed.Material, f.Stock[needed.Material]}
		}
		f.Stock.SubMaterialCounter(stillNeedsToBeProduced)

	}

	return needed
}

func (f MaterialFactory) Usage(material string) int {
	count := 0
	for _, mc := range f.Belt.Queue {
		if mc.Material == material {
			count += mc.Count
		}
	}
	return count
}

func (f MaterialFactory) ProduceWhileStock(neededMaterial string) int {
	need := MaterialCounter{
		Material: neededMaterial,
		Count:    1,
	}
	f.Produce(need)
	unitCost := f.Usage("ORE")
	producedCnt := 1

	//t := timer.New("Produce Round")
	//t.Start()
	for {
		need.Count = int(math.Max(float64(f.Stock["ORE"]/(unitCost*2)), float64(1)))
		//oreBf := f.Stock["ORE"]
		if f.Produce(need) != need {
			break
		}
		producedCnt += need.Count
		//fmt.Println(t.Elapsed(time.Now()),"ORE:",oreBf,"=>",f.Stock["ORE"])
	}
	return producedCnt
}
