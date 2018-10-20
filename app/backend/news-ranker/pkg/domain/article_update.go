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
	Referers   []news.Referer
	NewReferer news.Referer
}

func (u ArticleUpdate) ToScapeTarget() news.ScrapeTarget {
	article := u.Article
	return news.ScrapeTarget{
		URL:            article.URL,
		Subjects:       u.Subjects,
		Referer:        u.NewReferer,
		ReferenceScore: article.ReferenceScore,
		Title:          article.Title,
		Body:           article.Body,
		ArticleID:      article.ID,
	}
}

// CreateArticleUpdate dicerns how an article has been updated
// and assembles the data needed to rank it again.
func CreateArticleUpdate(article news.Article, oldSub, newSub []news.Subject, referers []news.Referer, newReferer news.Referer) ArticleUpdate {
	mergedSubjects := mergeSubjects(oldSub, newSub)
	mergedReferers := mergeReferers(referers, newReferer)

	hasNewSubjects := len(mergedSubjects) > len(oldSub)
	hasNewReferers := len(mergedReferers) > len(referers)

	return ArticleUpdate{
		Type:       dicernUpdateType(hasNewSubjects, hasNewReferers),
		Article:    article,
		Subjects:   mergedSubjects,
		Referers:   mergedReferers,
		NewReferer: newReferer,
	}
}

func mergeSubjects(old, newSubjects []news.Subject) []news.Subject {
	subjectSet := createSubjectSet(old)
	merged := make([]news.Subject, len(old))
	copy(merged, old)

	for _, newSub := range newSubjects {
		if _, ok := subjectSet[newSub.Symbol]; !ok {
			merged = append(merged, newSub)
		}
	}
	return merged
}

func createSubjectSet(subjects []news.Subject) map[string]bool {
	subjectSet := make(map[string]bool)
	for _, subject := range subjects {
		subjectSet[subject.Symbol] = true
	}
	return subjectSet
}

func mergeReferers(referers []news.Referer, newReferer news.Referer) []news.Referer {
	merged := make([]news.Referer, len(referers))
	copy(merged, referers)

	for _, referer := range referers {
		if referer.ExternalID == newReferer.ExternalID {
			return merged
		}
	}
	merged = append(merged, newReferer)
	return merged
}

func dicernUpdateType(hasNewSubjects, hasNewReferers bool) UpdateType {
	if hasNewSubjects && hasNewReferers {
		return NEW_SUBJECTS_AND_REFERENCES
	} else if hasNewSubjects {
		return NEW_SUBJECTS
	} else if hasNewReferers {
		return NEW_REFERENCES
	}
	return NO_UPDATE
}
