package module

import (
	"git.xiaojukeji.com/map-arch/tg-example/logic/module/data"
		"git.xiaojukeji.com/map-arch/tg-example/logic/module/mock"
		"git.xiaojukeji.com/map-arch/tg-example/logic/module/rank"
	"git.xiaojukeji.com/map-arch/tg-example/logic/module/recall"
	"github.com/didi/tg-flow/wfengine"
)

const (
	//modules常量定义
	DataPrepareSamp1Const = "data.DataPrepare"

	RoughRankConst	= "rank.RoughRank"
	PreciseRankConst= "rank.PreciseRank"
	ReRankConst		= "rank.ReRank"
	HistRecallConst	= "recall.HistRecall"
	HotRecallConst	= "recall.HotRecall"
	QacRecallConst	= "recall.QacRecall"
	EsGeneralConst	= "recall.EsGeneral"
	EsAoiConst		= "recall.EsAoi"
	EsPreciseConst	= "recall.EsPrecise"
	TimeoutConst	= "recall.Timeout"
	isMock			= true
)

type ModuleObject struct {
	wfengine.ModuleObjBase
}

func (moduleObj ModuleObject) NewObj(moduleName string) wfengine.IModelBase {
	switch moduleName {
		case DataPrepareSamp1Const:
			return &data.DataPrepare{}
		case RoughRankConst:
			return &rank.RoughRank{}
		case PreciseRankConst:
			return &rank.PreciseRank{}
		case ReRankConst:
			return &rank.ReRank{}
		case HistRecallConst:
			return &recall.HistRecall{}
		case HotRecallConst:
			return &recall.HotRecall{}
		case QacRecallConst:
			return &recall.QacRecall{}
		case EsGeneralConst:
			return &recall.EsGeneral{}
		case EsAoiConst:
			return &recall.EsAoi{}
		case EsPreciseConst:
			return &recall.EsPrecise{}
		case TimeoutConst:
				return &recall.TimeoutAction{}
		default:
			if isMock {
				return &mock.MockAction{}
			}
	}
	return nil
}
