package models

type RequestContext struct {
	CityId		int64  `thrift:"cityid,1,required" json:"city_id"`
	Category	string `thrift:"category,2,required" json:"category"`
	Keyword		string `thrift:"keyword,3,required" json:"keyword"`
	Uid			int64  `thrift:"uid,4,required" json:"uid"`
	SceneId		int64  `thrift:"scene_id,5,required" json:"scene_id"`
	FlowId		int64  `thrift:"flow_id,6,required" json:"flow_id"`
	IsDebug		bool   `thrift:"is_debug,7,required" json:"is_debug"`
	UserProfile	map[string]string  `thrift:"userProfile,8,required" json:"user_profile"`
}
