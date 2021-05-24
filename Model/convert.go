package Model

import (
	"github.com/sleptworld/test/DB"
	"strconv"
)

func Entry2Entry(from *DB.Entry, to *SingleEntry) {
	var user Author
	var cat Cat

	var tags []Tag
	for _, tag := range from.Tags {
		t := Tag{
			ID:   tag.ID,
			Name: tag.Name,
		}
		tags = append(tags, t)
	}

	var histories []History
	for _, history := range from.History {
		t := History{
			ID:        history.ID,
			CreatedAt: history.CreatedAt,
			Content:   history.Content,
			Info:      history.Info,
		}
		histories = append(histories, t)
	}

	id := strconv.Itoa(int(from.UserID))
	catid := strconv.Itoa(int(from.CatID))
	DB.FindUser(DB.Db, "id = ?", id, 1, &user)
	DB.SearchCat(DB.Db, "id = ?", catid, &cat)

	to.ID = from.ID
	to.CreatedAt = from.CreatedAt
	to.UpdatedAt = from.UpdatedAt
	to.Title = from.Title
	to.Author = user
	to.Tags = tags
	to.Category = cat
	to.Content = from.Content
	to.History = histories
	to.Info = from.Info
}

func Entry2AllEntry(from *[]DB.Entry,to *[]AllEntry){
	for _,value := range *from{
		var tags []Tag
		for _,v := range value.Tags{
			tags = append(tags, Tag{
				ID:   v.ID,
				Name: v.Name,
			})
		}
		*to = append((*to),AllEntry{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			Title:     value.Title,
			Content:   value.Content,
			Tags:      tags,
		})
	}
}

func User2AllUser(from *[]DB.User,to *[]AllUser)  {
	for _,value := range *from{
		var group userGroup
		DB.Db.Where(&DB.UserGroup{}).Where("id = ?",value.UserGroupID).First(&group)

		temp := AllUser{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			Name:      value.Name,
			Email:     value.Email,
			Birthday:  value.Birthday,
			UserGroup: group,
		}
		*to = append((*to), temp)
	}
}

func History2AllHistory(from *[]DB.History,to *[]History){
	for _,value := range *from{
		*to = append((*to),History{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			Content:   value.Content,
			Info:      value.Info,
		})
	}
}

func Draft2AllDraft(from *[]DB.Draft, to *[]Draft) {
	for _,value := range *from{
		*to = append((*to),Draft{
			ID:      value.ID,
			Title:   value.Title,
			Content: value.Content,
		})
	}
}

func User2User(from *DB.User,to *SingleUser)  {

	var group userGroup
	var entries []AllEntry
	var edit []History
	var drafts []Draft

	Entry2AllEntry(&(*from).Entries,&entries)
	History2AllHistory(&(*from).EditEntries,&edit)
	Draft2AllDraft(&(*from).Drafts,&drafts)

	DB.Db.Model(&DB.UserGroup{}).Where("id = ?",(*from).UserGroupID).First(&group)

	*to = SingleUser{
		ID:          (*from).ID,
		CreateAt:    (*from).CreatedAt,
		Name:        (*from).Name,
		Email:       (*from).Email,
		Birthday:    (*from).Birthday,
		UserGroup:   group,
		Avatar:      (*from).Avatar,
		Description: (*from).Description,
		Site:        (*from).Site,
		Country:     (*from).Country,
		Language:    (*from).Language,
		Entries:     entries,
		EditEntries: edit,
		Drafts:      drafts,
		Mechanism:   (*from).Mechanism,
		Sex:         (*from).Sex,
		Profession:  (*from).Profession,
	}
}

func UserUpdate2User(from *UserUpdate) DB.User {
	var result DB.User
	result.ID = (*from).ID
	result.Name = (*from).Name
	result.Site = (*from).Site
	result.Profession = (*from).Profession
	result.Mechanism = (*from).Mechanism
	result.Description = (*from).Description
	result.Avatar = (*from).Avatar

	DB.UserPretreatment(&result,(*from).Pwd)

	return result
}

func EntryUpdate2Entry(from *UpdateEntry) DB.Entry{
	var result DB.Entry
	tags := DB.Tags2Entry((*from).Tags)
	result.Content = (*from).Content
	result.Tags = tags

	return result
}