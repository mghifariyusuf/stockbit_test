package movie

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mghifariyusuf/stockbit_test.git/no_2/entity"
)

// Movie ...
type Movie struct {
	Title      string          `db:"title"`
	Year       string          `db:"year"`
	Rated      string          `db:"rated"`
	Released   string          `db:"released"`
	Runtime    string          `db:"runtime"`
	Genre      string          `db:"genre"`
	Director   string          `db:"director"`
	Writer     string          `db:"writer"`
	Actors     string          `db:"actors"`
	Plot       string          `db:"plot"`
	Language   string          `db:"language"`
	Country    string          `db:"country"`
	Awards     string          `db:"awards"`
	Poster     string          `db:"poster"`
	Ratings    json.RawMessage `db:"ratings"`
	Metascore  string          `db:"metascore"`
	ImdbRating string          `db:"imdb_rating"`
	ImdbVotes  string          `db:"imdb_votes"`
	ImdbID     string          `db:"imdb_id"`
	Type       string          `db:"type"`
	DVD        string          `db:"dvd"`
	BoxOffice  string          `db:"box_office"`
	Production string          `db:"production"`
	Website    string          `db:"website"`
}

// MySQL ...
type MySQL struct {
	db *sqlx.DB
}

// New ...
func New(db *sqlx.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}

// Upsert function to insert to db and update when imdb_id already
func (m *MySQL) Upsert(ctx context.Context, e entity.Movie) (err error) {
	ratings, err := json.Marshal(e.Ratings)
	if err != nil {
		log.Println(err)
		return err
	}

	db := Movie{
		Title:      e.Title,
		Year:       e.Year,
		Rated:      e.Rated,
		Released:   e.Released,
		Runtime:    e.Runtime,
		Genre:      e.Genre,
		Director:   e.Director,
		Writer:     e.Writer,
		Actors:     e.Actors,
		Plot:       e.Plot,
		Language:   e.Language,
		Country:    e.Country,
		Awards:     e.Awards,
		Poster:     e.Poster,
		Ratings:    ratings,
		Metascore:  e.Metascore,
		ImdbRating: e.ImdbRating,
		ImdbVotes:  e.ImdbVotes,
		ImdbID:     e.ImdbID,
		Type:       e.Type,
		DVD:        e.DVD,
		BoxOffice:  e.BoxOffice,
		Production: e.Production,
		Website:    e.Website,
	}

	_, err = m.db.NamedExecContext(ctx, `
		INSERT INTO movies (
			title,
			year,
			rated,
			released,
			runtime,
			genre,
			director,
			writer,
			actors,
			plot,
			language,
			country,
			awards,
			poster,
			ratings,
			metascore,
			imdb_rating,
			imdb_votes,
			imdb_id,
			type,
			dvd,
			box_office,
			production,
			website
		) VALUES (
			:title,
			:year,
			:rated,
			:released,
			:runtime,
			:genre,
			:director,
			:writer,
			:actors,
			:plot,
			:language,
			:country,
			:awards,
			:poster,
			:ratings,
			:metascore,
			:imdb_rating,
			:imdb_votes,
			:imdb_id,
			:type,
			:dvd,
			:box_office,
			:production,
			:website
		) ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			year = VALUES(year),
			rated = VALUES(rated),
			released = VALUES(released),
			runtime = VALUES(runtime),
			genre = VALUES(genre),
			director = VALUES(director),
			writer = VALUES(writer),
			actors = VALUES(actors),
			plot = VALUES(plot),
			language = VALUES(language),
			country = VALUES(country),
			awards = VALUES(awards),
			poster = VALUES(poster),
			ratings = VALUES(ratings),
			metascore = VALUES(metascore),
			imdb_rating = VALUES(imdb_rating),
			imdb_votes = VALUES(imdb_votes),
			type = VALUES(type),
			dvd = VALUES(dvd),
			box_office = VALUES(box_office),
			production = VALUES(production),
			website = VALUES(website)
	`, db)

	return nil
}
