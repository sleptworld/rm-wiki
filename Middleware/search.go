package Middleware

import (
	"fmt"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"github.com/sleptworld/test/DB"
	"github.com/sleptworld/test/tools"
	"gorm.io/gorm"
	"strconv"
)

var (
	Searcher = riot.Engine{}
	path     = "./riot-index"
	opt      = types.EngineOpts{
		Using:       0,
		PinYin:      true,
		UseStore:    true,
		StoreFolder: path,
	}
	
	SearchOpt = types.RankOpts{
		ScoringCriteria: nil,
		ReverseOrder:    false,
		OutputOffset:    0,
		MaxOutputs:      0,
	} 
)

func InitSearcher(db *gorm.DB) error {

	if tools.FileIsExist(path) {
		Searcher.Init(opt)
	} else {
		Searcher.Init(opt)
		if err := addAllEntries(db); err != nil {
			return err
		} else {
			return nil
		}
	}
	defer Searcher.Close()
	return nil
}

func AddDoc(e interface{}, condition string) {
	switch condition {
	case "Entry":
		index := string(rune(int((e).(DB.Entry).ID)))
		Searcher.Index(index, types.DocData{
			Content: e.(DB.Entry).Content,
		})
	default:
		fmt.Println("wrong")
	}
	Searcher.Flush()
}
func RemoveDoc(index string) {
	Searcher.RemoveDoc(index)
	Searcher.Flush()
}

func addAllEntries(db *gorm.DB) error {
	var rs []DB.Entry
	if result := db.Model(&DB.Entry{}).Find(&rs); result.Error == nil {
		for _, e := range rs {
			index := strconv.FormatUint(uint64(e.ID), 10)
			Searcher.Index(index, types.DocData{
				Content: e.Content,
			})
		}

		Searcher.Flush()
		return nil
	} else {
		return result.Error
	}
}

func Search(content string, opt *types.RankOpts) types.SearchResp{

	res := Searcher.Search(types.SearchReq{
		Text:     content,
		RankOpts: opt,
	})

	return res
}
