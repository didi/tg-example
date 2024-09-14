package mock

import (
	"git.xiaojukeji.com/map-arch/tg-example/idl"
	"git.xiaojukeji.com/map-arch/tg-example/models"
"math/rand"
	"time"
)

var maxGoodsId = 100

type MockUtils struct {}

func MockUserProfileFromRedis(userId int64) map[string]string {
	userProfile := make(map[string]string)
	userProfile["age"] = "12"
	return userProfile
}

func MockPreciseRank(items []*idl.ItemType) []*idl.ItemType {
	return mockRank(items)
}

func MockReRank(items []*idl.ItemType) []*idl.ItemType {
	return mockRank(items)
}

func mockRank(items []*idl.ItemType) []*idl.ItemType {
	source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

	for i,_ := range items {
		items[i].Score = r.Float64()
	}

	return items
}

func MockEsGeneralRecall(reqInfo *models.RequestContext, mockCostTime int) []*idl.ItemType{
	return MockRecall(mockCostTime)
}

func MockEsAoiRecall(items []*idl.ItemType, mockCostTime int) []*idl.ItemType{
	if len(items)==0{
		items = MockRecall(mockCostTime)
	}

	return items
}

func MockEsPreciseRecall(items []*idl.ItemType, mockCostTime int) []*idl.ItemType{
	if len(items)==0{
		items = MockRecall(mockCostTime)
	}
	return items
}

func MockHistRecall(mockCostTime int) []*idl.ItemType{
	return MockRecall(mockCostTime)
}

func MockHotRecall(mockCostTime int) []*idl.ItemType{
	return MockRecall(mockCostTime)
}

func MockQacRecall(mockCostTime int) []*idl.ItemType{
	return MockRecall(mockCostTime)
}

func MockRecall(mockCostTime int) []*idl.ItemType{
	time.Sleep(time.Millisecond * time.Duration(mockCostTime))
	source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

	items := make([]*idl.ItemType, 10, 10)
	for i:=0;i<10;i++{
		item :=&idl.ItemType{
			GoodsId: int64(r.Intn(maxGoodsId)),
			Price: r.Float64(),
		}
		items[i]=item
	}

	return items
}

func MockRoughRank(items []*idl.ItemType, items2 []*idl.ItemType, items3 []*idl.ItemType, items4 []*idl.ItemType) []*idl.ItemType {
	resultMap := make(map[int64]*idl.ItemType)
	for _, item:= range items{
		if _, ok := resultMap[item.GoodsId]; !ok{
			resultMap[item.GoodsId] = item
		}
	}

	for _, item:= range items2{
		if _, ok := resultMap[item.GoodsId]; !ok{
			resultMap[item.GoodsId] = item
		}
	}

	for _, item:= range items3{
		if _, ok := resultMap[item.GoodsId]; !ok{
			resultMap[item.GoodsId] = item
		}
	}

	for _, item:= range items4{
		if _, ok := resultMap[item.GoodsId]; !ok{
			resultMap[item.GoodsId] = item
		}
	}

	results := make([]*idl.ItemType,len(resultMap), len(resultMap))
	i:=0
	for _, v := range resultMap{
		results[i] = v
		i++
	}

	return results
}