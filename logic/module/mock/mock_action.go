package mock

import (
	"context"
	"github.com/didi/tg-flow/model"
	"github.com/didi/tg-flow/wfengine"
)

type MockAction struct {
	wfengine.ModelBase
}

var timecostMap map[string]int

func (m MockAction) DoAction(ctx context.Context, sc *model.StrategyContext) interface{} {
	/*timeCost := 10
	if tm,ok := timecostMap[m.Name]; ok {
		timeCost = tm
	}

	fmt.Println(time.Now(),"start", m.GetName(),"执行耗时为：", timeCost)
	time.Sleep(time.Duration(timeCost) * time.Millisecond)
	fmt.Println(time.Now(),"end", m.GetName())*/
	return nil
}

func (m MockAction) OnTimeout(context.Context, *model.StrategyContext) {
	//this is a sample, you can do sth when timeout happen
	//fmt.Println(fmt.Sprintf("时间:%v, 节点:%v超时,触发回调:%v",time.Now().Format("2006-01-02 15:04:05.000"), "", m.GetName()))
}

func init() {
	timecostMap = make(map[string]int)
	timecostMap["data.DataPrepare"] = 10
	timecostMap["recall.HistRecall"] = 40
	timecostMap["recall.HotRecall"] = 30
	timecostMap["recall.QacRecall"] = 60
	timecostMap["recall.EsGeneral"] = 30
	timecostMap["recall.EsAoi"] = 40
	timecostMap["recall.EsPrecise"] = 20
	timecostMap["rank.RoughRank"] = 15
	timecostMap["rank.PreciseRank"] = 10
	timecostMap["rank.ReRank"] = 5
	timecostMap["recall.Timeout"] = 0
	//希望名字带Timeout
	timecostMap["recall.RecallSuggestTimeoutZhPyMix"] = 0
	timecostMap["recall.RecallSuggestTimeoutZh"] = 0
}
