package domain

import (
	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

const (
	_ = iota
	NO_UPDATE
	NEW_SUBJECTS
	NEW_REFERENCES
	NEW_SUBJECTS_AND_REFERENCES
)

// UpdateType describes distinct type of update.
type UpdateType int

// ArticleUpdate bundles an update instruction with the data needed to perform it.
type ArticleUpdate struct {
	Type       UpdateType
	Article    news.Article
	Subjects   []news.Subject
	References []news.Author
}

// CreateArticleUpdate dicerns how an article has been updated
// and assembles the data needed to rank it again.
func CreateArticleUpdate(article news.Article, oldSub, newSub []news.Subject, oldRefs, newRefs []news.Author) ArticleUpdate {
	hasNewSubjects := len(newSub) > len(oldSub)
	hasNewReferences := len(newRefs) > len(oldRefs)

	return ArticleUpdate{
		Type:       dicernUpdateType(hasNewSubjects, hasNewReferences),
		Article:    article,
		Subjects:   newSub,
		References: newRefs,
	}
}

func dicernUpdateType(hasNewSubjects, hasNewReferences bool) UpdateType {
	if hasNewSubjects && hasNewReferences {
		return NEW_SUBJECTS_AND_REFERENCES
	} else if hasNewSubjects {
		return NEW_SUBJECTS
	} else if hasNewReferences {
		return NEW_REFERENCES
	}
	return NO_UPDATE
}
