//go:generate mockgen -destination=./mock/service.go -package=mock . Repository,Service

package crud

import "github.com/selmison/code-micro-videos/models"

type Repository interface {
	Service
}

type Service interface {
	GetCategories(limit int) (models.CategorySlice, error)
	FetchCategory(name string) (models.Category, error)
	AddCategory(c CategoryDTO) error
	RemoveCategory(name string) error
	UpdateCategory(name string, c CategoryDTO) error

	GetCastMembers(limit int) (models.CastMemberSlice, error)
	FetchCastMember(name string) (models.CastMember, error)
	AddCastMember(c CastMemberDTO) error
	RemoveCastMember(name string) error
	UpdateCastMember(name string, c CastMemberDTO) error

	GetGenres(limit int) (models.GenreSlice, error)
	FetchGenre(name string) (models.Genre, error)
	AddGenre(c GenreDTO) error
	RemoveGenre(name string) error
	UpdateGenre(name string, c GenreDTO) error

	GetVideos(limit int) (models.VideoSlice, error)
	FetchVideo(name string) (models.Video, error)
	AddVideo(c VideoDTO) error
	RemoveVideo(name string) error
	UpdateVideo(name string, c VideoDTO) error
}

// NewService creates a crud service with the necessary dependencies
func NewService(r Repository) *service {
	return &service{r}
}
