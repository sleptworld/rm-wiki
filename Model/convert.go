package Model
//
//import (
//	"github.com/sleptworld/test/DB"
//	"strconv"
//	"time"
//)
//
//
//func Entry2Entry(from *DB.Entry, to *SingleEntry) *SingleEntry{
//
//	var user Author
//	var cat Cat
//	var tags []Tag
//
//	for _, tag := range from.Tags {
//		t := Tag{
//			ID:   tag.ID,
//			Name: tag.Name,
//		}
//		tags = append(tags, t)
//	}
//
//	var histories []History
//	for _, history := range from.History {
//		t := History{
//			ID:        history.ID,
//			CreatedAt: history.CreatedAt,
//			Content:   history.Content,
//			Info:      history.Info,
//		}
//		histories = append(histories, t)
//	}
//
//	DB.FindUser(DB.Db, "id = ?", from.UserID, 1, &user)
//	DB.SearchCat(DB.Db, "id = ?", from.CatID, &cat)
//
//	return &SingleEntry{
//		ID:        from.ID,
//		CreatedAt: from.CreatedAt,
//		UpdatedAt: from.UpdatedAt,
//		Title:     from.Title,
//		Author:    user,
//		Tags:      tags,
//		Category:  cat,
//		Content:   from.Content,
//		History:   histories,
//		Info:      from.Info,
//	}
//}
//
//func Entries2Entries(from *[]DB.Entry) (to *[]AllEntry){
//
//	var temp []AllEntry
//	for _,value := range *from{
//		var tags []Tag
//		for _,v := range value.Tags{
//			tags = append(tags, Tag{
//				ID:   v.ID,
//				Name: v.Name,
//			})
//		}
//		temp = append(temp,AllEntry{
//			ID:        value.ID,
//			CreatedAt: value.CreatedAt,
//			UpdatedAt: value.UpdatedAt,
//			Title:     value.Title,
//			Content:   value.Content,
//			Tags:      tags,
//		})
//	}
//
//	return &temp
//}
//
//func User2User(from *DB.User){
//
//	e := Entries2Entries(&(*from).Entries)
//	History2AllHistory(&(*from).EditEntries,&edit)
//	Draft2AllDraft(&(*from).Drafts,&drafts)
//
//	DB.Db.Model(&DB.UserGroup{}).Where("id = ?",(*from).UserGroupID).First(&group)
//
//	var result = SingleUser{
//		ID:          from.ID,
//		CreateAt:    from.CreatedAt,
//		Name:        from.Name,
//		Email:       from.Email,
//		Birthday:    from.Birthday,
//		UserGroup:   userGroup{},
//		Avatar:      from.Avatar,
//		Description: from.Description,
//		Site:        from.Site,
//		Country:     from.Country,
//		Language:    from.Language,
//		Entries:     nil,
//		EditEntries: nil,
//		Drafts:      nil,
//		Mechanism:   from.Mechanism,
//		Sex:         from.Sex,
//		Profession:  from.Profession,
//	}
//
//
//}
//func Users2Users(from *[]DB.User) (to *[]AllUser) {
//
//	var result []AllUser
//
//
//	for _,value := range *from{
//
//		var group userGroup
//		DB.Db.Where(&DB.UserGroup{}).Where("id = ?",value.UserGroupID).First(&group)
//		temp := AllUser{
//			ID:        value.ID,
//			CreatedAt: value.CreatedAt,
//			Name:      value.Name,
//			Email:     value.Email,
//			Birthday:  value.Birthday,
//			UserGroup: group,
//		}
//		result = append(result, temp)
//	}
//	return &result
//}
//
//func History2AllHistory(from *[]DB.History,to *[]History){
//	for _,value := range *from{
//		*to = append((*to),History{
//			ID:        value.ID,
//			CreatedAt: value.CreatedAt,
//			Content:   value.Content,
//			Info:      value.Info,
//		})
//	}
//}
//
//func Draft2AllDraft(from *[]DB.Draft, to *[]Draft) {
//	for _,value := range *from{
//		*to = append((*to),Draft{
//			ID:      value.ID,
//			Title:   value.Title,
//			Content: value.Content,
//		})
//	}
//}
//func UserUpdate2User(from *UserUpdate) DB.User {
//	var result DB.User
//	result.ID = (*from).ID
//	result.Name = (*from).Name
//	result.Site = (*from).Site
//	result.Profession = (*from).Profession
//	result.Mechanism = (*from).Mechanism
//	result.Description = (*from).Description
//	result.Avatar = (*from).Avatar
//
//	DB.UserPretreatment(&result,(*from).Pwd)
//
//	return result
//}
//
//func EntryUpdate2Entry(from *UpdateEntry) DB.Entry{
//	var result DB.Entry
//	tags := DB.Tags2Entry((*from).Tags)
//	result.ID = (*from).ID
//	result.Content = (*from).Content
//	result.Tags = tags
//
//	return result
//}