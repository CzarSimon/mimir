package domain

import (
	"testing"

	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func TestCreateArticleUpdate(t *testing.T) {
	a := news.Article{ID: "article-0"}
	oldSubj := []news.Subject{
		news.Subject{Symbol: "s-0"},
		news.Subject{Symbol: "s-1"},
	}
	newSubj := []news.Subject{
		news.Subject{Symbol: "s-2"},
	}
	repeatedSubjects := []news.Subject{
		news.Subject{Symbol: "s-1"},
	}
	mergedSubjects := []news.Subject{
		news.Subject{Symbol: "s-0"},
		news.Subject{Symbol: "s-1"},
		news.Subject{Symbol: "s-2"},
	}

	oldRefs := []news.Referer{
		news.Referer{ExternalID: "a-0"},
		news.Referer{ExternalID: "a-1"},
	}
	newRef := news.Referer{ExternalID: "a-2"}
	repeatedRef := news.Referer{ExternalID: "a-1"}
	mergedRefs := []news.Referer{
		news.Referer{ExternalID: "a-0"},
		news.Referer{ExternalID: "a-1"},
		news.Referer{ExternalID: "a-2"},
	}

	u1 := CreateArticleUpdate(a, oldSubj, repeatedSubjects, oldRefs, repeatedRef)
	assertArticleUpdate(t, u1, a, oldSubj, oldRefs)
	if u1.Type != NO_UPDATE {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NO_UPDATE, u1.Type)
	}

	u2 := CreateArticleUpdate(a, oldSubj, newSubj, oldRefs, repeatedRef)
	assertArticleUpdate(t, u2, a, mergedSubjects, oldRefs)
	if u2.Type != NEW_SUBJECTS {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NEW_SUBJECTS, u2.Type)
	}

	u3 := CreateArticleUpdate(a, oldSubj, repeatedSubjects, oldRefs, newRef)
	assertArticleUpdate(t, u3, a, oldSubj, mergedRefs)
	if u3.Type != NEW_REFERENCES {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NEW_REFERENCES, u3.Type)
	}

	u4 := CreateArticleUpdate(a, oldSubj, newSubj, oldRefs, newRef)
	assertArticleUpdate(t, u4, a, mergedSubjects, mergedRefs)
	if u4.Type != NEW_SUBJECTS_AND_REFERENCES {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NEW_SUBJECTS_AND_REFERENCES, u4.Type)
	}
}

func assertArticleUpdate(t *testing.T, u ArticleUpdate, eA news.Article, eS []news.Subject, eR []news.Referer) {
	if u.Article.ID != eA.ID {
		t.Errorf("Article.ID wrong. Expected: %s Got: %s", u.Article.ID, eA.ID)
	}

	if len(u.Subjects) != len(eS) {
		t.Fatalf("Subjects length missmatch. Expected: %d Got: %d", len(u.Subjects), len(eS))
	}
	for i, sub := range u.Subjects {
		if sub.Symbol != eS[i].Symbol {
			t.Errorf("%d - Subject.Symbol wrong. Expected: %s Got: %s", i, eS[i].Symbol, sub.Symbol)
		}
	}

	if len(u.Referers) != len(eR) {
		t.Fatalf("References length missmatch. Expected: %d Got: %d", len(u.Referers), len(eR))
	}
	for i, ref := range u.Referers {
		if ref.ExternalID != eR[i].ExternalID {
			t.Errorf("%d - Author.ExternalID wrong. Expected: %s Got: %s", i, eR[i].ExternalID, ref.ExternalID)
		}
	}
}
