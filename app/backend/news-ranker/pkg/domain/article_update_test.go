package domain

import (
	"testing"

	"github.com/CzarSimon/mimir/app/backend/pkg/schema/news"
)

func TestCreateArticleUpdate(t *testing.T) {
	a := news.Article{ID: "article-0"}
	oneSubj := []news.Subject{
		news.Subject{Name: "s-0"},
	}
	twoSubj := []news.Subject{
		news.Subject{Name: "s-0"},
		news.Subject{Name: "s-1"},
	}
	oneRef := []news.Author{
		news.Author{ID: "a-0"},
	}
	twoRefs := []news.Author{
		news.Author{ID: "a-0"},
		news.Author{ID: "a-1"},
	}

	u1 := CreateArticleUpdate(a, oneSubj, oneSubj, oneRef, oneRef)
	assertArticleUpdate(t, u1, a, oneSubj, oneRef)
	if u1.Type != NO_UPDATE {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NO_UPDATE, u1.Type)
	}

	u2 := CreateArticleUpdate(a, oneSubj, twoSubj, twoRefs, twoRefs)
	assertArticleUpdate(t, u2, a, twoSubj, twoRefs)
	if u2.Type != NEW_SUBJECTS {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NEW_SUBJECTS, u2.Type)
	}

	u3 := CreateArticleUpdate(a, twoSubj, twoSubj, oneRef, twoRefs)
	assertArticleUpdate(t, u3, a, twoSubj, twoRefs)
	if u3.Type != NEW_REFERENCES {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NEW_REFERENCES, u3.Type)
	}

	u4 := CreateArticleUpdate(a, oneSubj, twoSubj, oneRef, twoRefs)
	assertArticleUpdate(t, u4, a, twoSubj, twoRefs)
	if u4.Type != NEW_SUBJECTS_AND_REFERENCES {
		t.Errorf("CreateArticleUpdate failed. Expected type: %d Got: %d",
			NEW_SUBJECTS_AND_REFERENCES, u4.Type)
	}
}

func assertArticleUpdate(t *testing.T, u ArticleUpdate, eA news.Article, eS []news.Subject, eR []news.Author) {
	if u.Article.ID != eA.ID {
		t.Errorf("Article.ID wrong. Expected: %s Got: %s", u.Article.ID, eA.ID)
	}

	if len(u.Subjects) != len(eS) {
		t.Errorf("Subjects length missmatch. Expected: %d Got: %d", len(u.Subjects), len(eS))
	}
	for i, sub := range u.Subjects {
		if sub.Name != eS[i].Name {
			t.Errorf("%d - Subject.Name wrong. Expected: %s Got: %s", i, eS[i].Name, sub.Name)
		}
	}

	if len(u.References) != len(eR) {
		t.Errorf("References length missmatch. Expected: %d Got: %d", len(u.References), len(eR))
	}
	for i, ref := range u.References {
		if ref.ID != eR[i].ID {
			t.Errorf("%d - Author.ID wrong. Expected: %s Got: %s", i, eR[i].ID, ref.ID)
		}
	}
}
