package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FranciscoMendes10866/gql-api/graph/generated"
	"github.com/FranciscoMendes10866/gql-api/graph/helpers"
	"github.com/FranciscoMendes10866/gql-api/graph/model"
	resty "github.com/go-resty/resty/v2"
)

func (r *queryResolver) GetAllMangas(ctx context.Context) ([]*model.GetAllMangasResponse, error) {
	val, notFound := helpers.MemcachedClient.Get("mangas-all")

	if notFound == nil {
		list := []*model.GetAllMangasResponse{}
		buffer := val.Value
		json.Unmarshal(buffer, &list)
		fmt.Println("[%a] Get from cache", "mangas-all")
		return list, nil
	}

	var result []*model.GetAllMangasResponse
	client := resty.New()
	_, err := client.R().SetHeader("x-dango-manga-key", "superSecretKey").SetResult(&result).Get("http://localhost:3333/api/manga/get-all")
	if err != nil {
		return nil, err
	}

	buffer, _ := json.Marshal(result)
	helpers.SaveToCache("mangas-all", buffer)

	return result, nil
}

func (r *queryResolver) GetMangaDetails(ctx context.Context, mangaID string) (*model.GetMangaDetailsResponse, error) {
	val, notFound := helpers.MemcachedClient.Get(mangaID)

	if notFound == nil {
		saved := &model.GetMangaDetailsResponse{}
		buffer := val.Value
		json.Unmarshal(buffer, &saved)
		fmt.Println("[%a] Get Manga from cache", mangaID)
		return saved, nil
	}

	var result *model.GetMangaDetailsResponse
	client := resty.New()
	_, err := client.R().SetHeader("x-dango-manga-key", "superSecretKey").SetPathParams(map[string]string{"mangaId": mangaID}).SetResult(&result).Get("http://localhost:3333/api/manga/get-details/{mangaId}")
	if err != nil {
		return nil, err
	}

	buffer, _ := json.Marshal(result)
	helpers.SaveToCache(mangaID, buffer)

	return result, nil
}

func (r *queryResolver) GetChapter(ctx context.Context, chapterID string) (*model.GetChapterResponse, error) {
	val, notFound := helpers.MemcachedClient.Get(chapterID)

	if notFound == nil {
		saved := &model.GetChapterResponse{}
		buffer := val.Value
		json.Unmarshal(buffer, &saved)
		fmt.Println("[%a] Get Chapter from cache", chapterID)
		return saved, nil
	}

	var result *model.GetChapterResponse
	client := resty.New()
	_, err := client.R().SetHeader("x-dango-manga-key", "superSecretKey").SetPathParams(map[string]string{"chapterId": chapterID}).SetResult(&result).Get("http://localhost:3333/api/chapter/{chapterId}")
	if err != nil {
		return nil, err
	}

	buffer, _ := json.Marshal(result)
	helpers.SaveToCache(chapterID, buffer)

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
