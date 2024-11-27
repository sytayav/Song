package song

import "context"

type Repository interface {
	Create(ctx context.Context, s *Song) (string, error)
	FindAll(ctx context.Context) (s []Song, err error)
	FindOne(ctx context.Context, id string) (Song, error)
	Update(ctx context.Context, s Song) error
	Delete(ctx context.Context, id string) error
}
