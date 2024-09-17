package mock

import (
	"github.com/didi/tg-example/models"
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

func MockPreciseRank(items []*models.ItemType) []*models.ItemType {
	return mockRank(items)
}

func MockReRank(items []*models.ItemType) []*models.ItemType {
	return mockRank(items)
}

func mockRank(items []*models.ItemType) []*models.ItemType {
	source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

	for i,_ := range items {
		items[i].Score = r.Float64()
	}

	return items
}

func MockEsGeneralRecall(reqInfo *models.RequestContext, mockCostTime int) []*models.ItemType{
	return MockRecall(mockCostTime)
}

func MockEsAoiRecall(items []*models.ItemType, mockCostTime int) []*models.ItemType{
	if len(items)==0{
		items = MockRecall(mockCostTime)
	}

	return items
}

func MockEsPreciseRecall(items []*models.ItemType, mockCostTime int) []*models.ItemType{
	if len(items)==0{
		items = MockRecall(mockCostTime)
	}
	return items
}

func MockHistRecall(mockCostTime int) []*models.ItemType{
	return MockRecall(mockCostTime)
}

func MockHotRecall(mockCostTime int) []*models.ItemType{
	return MockRecall(mockCostTime)
}

func MockQacRecall(mockCostTime int) []*models.ItemType{
	return MockRecall(mockCostTime)
}

func MockRecall(mockCostTime int) []*models.ItemType{
	time.Sleep(time.Millisecond * time.Duration(mockCostTime))
	source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

	items := make([]*models.ItemType, 10, 10)
	for i:=0;i<10;i++{
		item :=&models.ItemType{
			GoodsId: int64(r.Intn(maxGoodsId)),
			Price: r.Float64(),
		}
		items[i]=item
	}

	return items
}

func MockRoughRank(items []*models.ItemType, items2 []*models.ItemType, items3 []*models.ItemType, items4 []*models.ItemType) []*models.ItemType {
	resultMap := make(map[int64]*models.ItemType)
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

	results := make([]*models.ItemType,len(resultMap), len(resultMap))
	i:=0
	for _, v := range resultMap{
		results[i] = v
		i++
	}

	return results
}