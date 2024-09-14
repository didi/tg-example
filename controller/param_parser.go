package controller

import (
		"git.xiaojukeji.com/map-arch/tg-example/global/constants"
	"git.xiaojukeji.com/map-arch/tg-example/models"
	"github.com/gin-gonic/gin"
		"strconv"
	"strings"
)

func getRequestContext(c *gin.Context) (*models.RequestContext, error) {
	cityId := parseIntMust(c.Query("city_id"), 0)
	category := c.Query("category")
	keyword := c.Query("keyword")
	uid := parseIntMust(c.Query("uid"), 0)
	sceneId := parseIntMust(c.Query("scene_id"), 0)
	flowId := parseIntMust(c.Query("flow_id"), 0)
	isDebug := parseBoolMust(c.Query("is_debug"), false)
	rc := &models.RequestContext{
		CityId: cityId,
		Category: category,
		Keyword: keyword,
		Uid: uid,
		SceneId: sceneId,
		FlowId:  flowId,
		IsDebug: isDebug,
	}

	//TODO ZYF 注意，我们建议上游传scene_id过来，但很多上游系统比算法平台更早，没有scene_id，所以只能先兼容处理了。
	if rc.SceneId == 0 {
		rc.SceneId = constants.SceneIdTgExample
	}
	return rc, nil
}

func parseIntMust(str string, defaultValue int64) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}

	return val
}

func parseBoolMust(str string, defaultValue bool) bool {
	strLower := strings.ToLower(str)
	if strLower != "true" && strLower != "false" {
		return defaultValue
	}

	return strLower == "true"
}