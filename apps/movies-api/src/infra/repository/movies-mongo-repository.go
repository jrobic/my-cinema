package repository

import (
	"context"
	"time"

	"github.com/jrobic/my-cinema/movies-api/src/domain"
	"github.com/jrobic/my-cinema/movies-api/src/infra/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MoviesMongoRepository struct {
	db *db.DBMongo
}

func NewMoviesMongoRepository(db *db.DBMongo) *MoviesMongoRepository {
	return &MoviesMongoRepository{
		db: db,
	}
}

func (r *MoviesMongoRepository) Insert(movie domain.Movie) error {
	collection := r.db.DB.Collection("movies")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, movie)

	if err != nil {
		return err
	}

	return nil
}

func (r *MoviesMongoRepository) FindAll() ([]*domain.Movie, error) {
	movies := []*domain.Movie{}

	collection := r.db.DB.Collection("movies")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{}, options.Find().SetLimit(100))
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MoviesMongoRepository) InsertMany(movies []domain.Movie) error {
	collection := r.db.DB.Collection("movies")

	var docs []interface{}

	for _, movie := range movies {
		docs = append(docs, movie)

		// BulkWrite is hyper slow compared to InsertMany
		// 	var docs []mongo.WriteModel

		// for _, movie := range movies {
		// 	op := mongo.NewUpdateOneModel()

		// 	op.SetFilter(bson.D{{Key: "tmdb_id", Value: movie.TmdbID}})
		// 	op.SetUpdate(bson.D{{Key: "$set", Value: movie}})
		// 	op.SetUpsert(true)

		// 	docs = append(docs, op)
	}

	// _, err := collection.BulkWrite(context.TODO(), docs)

	_, err := collection.InsertMany(context.TODO(), docs)

	if err != nil {
		return err
	}

	return nil
}
