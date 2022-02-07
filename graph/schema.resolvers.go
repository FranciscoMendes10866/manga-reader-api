package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FranciscoMendes10866/gql-api/graph/generated"
	"github.com/FranciscoMendes10866/gql-api/graph/model"
	resty "github.com/go-resty/resty/v2"
)

func (r *queryResolver) GetAllMangas(ctx context.Context) ([]*model.GetAllMangasResponse, error) {
	var result []*model.GetAllMangasResponse
	client := resty.New()
	_, err := client.R().SetHeader("x-dango-manga-key", "superSecretKey").SetResult(&result).Get("http://localhost:3333/api/manga/get-all")
	if err != nil {
		return nil, err
	}

	fmt.Println(result)
	return result, nil
}

func (r *queryResolver) GetMangaDetails(ctx context.Context, mangaID string) (*model.GetMangaDetailsResponse, error) {
	var result *model.GetMangaDetailsResponse
	client := resty.New()
	_, err := client.R().SetHeader("x-dango-manga-key", "superSecretKey").SetPathParams(map[string]string{"mangaId": mangaID}).SetResult(&result).Get("http://localhost:3333/api/manga/get-details/{mangaId}")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
