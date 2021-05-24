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
