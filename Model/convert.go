package Model

import (
	"github.com/sleptworld/test/DB"
)


func Entry2Entry(from *DB.Entry) *SingleEntry{

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

	DB.Init(from).Query(&user,1,"",nil)
	DB.Init(&DB.Cat{
		ID : from.CatID,
	}).Query(&cat,2,"",nil)

	return &SingleEntry{
		ID:        from.ID,
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
		Title:     from.Title,
		Author:    user,
		Tags:      tags,
		Category:  cat,
		Content:   from.Content,
		History:   histories,
		Info:      from.Info,
	}
}

func Entries2Entries(from *[]DB.Entry) (to *[]AllEntry){

	var temp []AllEntry
	for _,value := range *from{
		var tags []Tag
		for _,v := range value.Tags{
			tags = append(tags, Tag{
				ID:   v.ID,
				Name: v.Name,
			})
		}
		temp = append(temp,AllEntry{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			Title:     value.Title,
			Content:   value.Content,
			Tags:      tags,
		})
	}

	return &temp
}

func User2User(from *DB.User) *SingleUser{

	e := Entries2Entries(&(from.Entries))
	h := History2AllHistory(&(from.EditEntries))
	d := Draft2AllDraft(&(from.Drafts))

	var group userGroup

	group.GroupName = DB.GetRolesByUser(from.ID)[0]

	var result = SingleUser{
		ID:          from.ID,
		CreateAt:    from.CreatedAt,
		Name:        from.Name,
		Email:       from.Email,
		Birthday:    from.Birthday,
		UserGroup:   group,
		Avatar:      from.Avatar,
		Description: from.Description,
		Site:        from.Site,
		Country:     from.Country,
		Language:    from.Language,
		Entries:     *e,
		EditEntries: *h,
		Drafts:      *d,
		Mechanism:   from.Mechanism,
		Sex:         from.Sex,
		Profession:  from.Profession,
	}

	return &result
}
func Users2Users(from *[]DB.User) (to *[]AllUser) {

	var result []AllUser
	for _,value := range *from{

		var group userGroup
		group.GroupName = DB.GetRolesByUser(value.ID)[0]
		temp := AllUser{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			Name:      value.Name,
			Email:     value.Email,
			Birthday:  value.Birthday,
			UserGroup: group,
		}
		result = append(result, temp)
	}
	return &result
}

func History2AllHistory(from *[]DB.History) *[]History{
	var to []History
	for _,value := range *from{
		to = append(to,History{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			Content:   value.Content,
			Info:      value.Info,
		})
	}
	return &to
}

func Draft2AllDraft(from *[]DB.Draft) *[]Draft {
	var to []Draft
	for _,value := range *from{
		to = append(to,Draft{
			ID:      value.ID,
			Title:   value.Title,
			Content: value.Content,
		})
	}
	return &to
}