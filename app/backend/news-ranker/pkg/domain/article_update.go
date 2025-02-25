package domain

import (
	"github.com/CzarSimon/mimir/app/backend/pkg/id"
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
	mergedSubjects := mergeSubjects(oldSub, newSub, article.ID)
	mergedReferers := mergeReferers(referers, newReferer, article.ID)

	hasNewSubjects := len(mergedSubjects) > len(oldSub)
	hasNewReferers := len(mergedReferers) > len(referers)

	return ArticleUpdate{
		Type:       dicernUpdateType(hasNewSubjects, hasNewReferers),
		Article:    article,
		Subjects:   mergedSubjects,
		Referers:   mergedReferers,
		NewReferer: copyRefererWithIDs(newReferer, article.ID),
	}
}

func mergeSubjects(old, newSubjects []news.Subject, articleID string) []news.Subject {
	subjectSet := createSubjectSet(old)
	merged := make([]news.Subject, len(old))
	copy(merged, old)

	for _, newSub := range newSubjects {
		if _, ok := subjectSet[newSub.Symbol]; !ok {
			merged = append(merged, copySubjectWithIDs(newSub, articleID))
		}
	}
	return merged
}

func copySubjectWithIDs(subject news.Subject, articleID string) news.Subject {
	subject.ArticleID = articleID
	if subject.ID == "" {
		subject.ID = id.New()
	}
	return subject
}

func createSubjectSet(subjects []news.Subject) map[string]bool {
	subjectSet := make(map[string]bool)
	for _, subject := range subjects {
		subjectSet[subject.Symbol] = true
	}
	return subjectSet
}

func mergeReferers(referers []news.Referer, newReferer news.Referer, articleID string) []news.Referer {
	merged := make([]news.Referer, len(referers))
	copy(merged, referers)

	for _, referer := range referers {
		if referer.ExternalID == newReferer.ExternalID {
			return merged
		}
	}
	merged = append(merged, copyRefererWithIDs(newReferer, articleID))
	return merged
}

func copyRefererWithIDs(referer news.Referer, articleID string) news.Referer {
	referer.ArticleID = articleID
	if referer.ID == "" {
		referer.ID = id.New()
	}
	return referer
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
