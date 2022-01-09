package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/FranciscoMendes10866/gql-api/graph/config"
	"github.com/FranciscoMendes10866/gql-api/graph/generated"
	"github.com/FranciscoMendes10866/gql-api/graph/model"
)

func (r *queryResolver) GetMangaList(ctx context.Context) ([]*model.Manga, error) {
	var mangas []*model.Manga
	config.Database.Table("manga_entities").Find(&mangas)
	return mangas, nil
}

func (r *queryResolver) GetMangaDetails(ctx context.Context, mangaID string) (*model.Manga, error) {
	var manga *model.Manga
	var chapters []*model.Chapter
	config.Database.Table("manga_entities").Where("id = ?", mangaID).First(&manga)
	config.Database.Table("chapter_entities").Where("manga_id = ?", mangaID).Find(&chapters)
	manga.Chapters = chapters
	return manga, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
