package song

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/sytayav/Song/internal/song"
	"github.com/sytayav/Song/pkg/client/postgresql"
	"github.com/sytayav/Song/pkg/logging"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, song *song.Song) error {
	q := `
		INSERT INTO song 
		    (name, age) 
		VALUES 
		       ($1, $2) 
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, song.Song_name, song.Group_name, 123).Scan(&song.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []song.Song, err error) {
	q := `
		SELECT id, name FROM public.song;
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	songs := make([]song.Song, 0)

	for rows.Next() {
		var sng song.Song

		err = rows.Scan(&sng.ID, &sng.Song_name, &sng.Group_name)
		if err != nil {
			return nil, err
		}

		songs = append(songs, sng)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (song.Song, error) {
	q := `
		SELECT id, name FROM public.song WHERE id = $1
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var sng song.Song
	err := r.client.QueryRow(ctx, q, id).Scan(&sng.ID, &sng.Song_name, &sng.Group_name)
	if err != nil {
		return song.Song{}, err
	}

	return sng, nil
}

func (r *repository) Update(ctx context.Context, user song.Song) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) song.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
