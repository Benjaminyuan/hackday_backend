package crawl

import (
	"encoding/json"
	_"fmt"
)

func ParseAllProvinceData(c []byte) *ProvinceAction {
	provinceAcion := &ProvinceAction{}
	json.Unmarshal(c, provinceAcion)
	return provinceAcion
}
func ParseAllChinaData(c []byte) *ChinaAction {
	chinaAction := &ChinaAction{}
	json.Unmarshal(c, chinaAction)
	return chinaAction
}
func ParseSelectChoseList(c []byte) *SelectChooseListAction {
	selectChooselist := &SelectChooseListAction{}
	json.Unmarshal(c, selectChooselist)
	return selectChooselist
}
func ParseTypeProp(c []byte) *TypeProp {
	typeProp := &TypeProp{}
	json.Unmarshal(c, typeProp)
	return typeProp
}
func ParseRankData(c []byte) *RankData {
	rankData := &RankData{}
	json.Unmarshal(c, rankData)
	return rankData
}
