package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jrobic/my-cinema/movies-api/src/domain"
	"github.com/jrobic/my-cinema/movies-api/src/infra/db"
	"github.com/jrobic/my-cinema/movies-api/src/infra/repository"
	"github.com/jrobic/my-cinema/movies-api/src/usecase"
)

func TestIngestMoviesFileUsecase(t *testing.T) {
	t.Run("should ingest movies from csv file", func(t *testing.T) {
		want := []domain.Movie{
			{
				ID:                  uuid.Must(uuid.FromString("53aa35c8-e659-44b2-882f-f6056e443c99")),
				TmdbID:              "758323",
				Title:               "The Pope's Exorcist",
				Genres:              []string{"Horror", "Mystery", "Thriller"},
				OriginalLanguage:    "en",
				Overview:            "Father Gabriele Amorth Chief Exorcist of the Vatican investigates a young boy's terrifying possession and ends up uncovering a centuries-old conspiracy the Vatican has desperately tried to keep hidden.",
				Popularity:          5953.227,
				ProductionCompanies: []string{"Screen Gems", "2.0 Entertainment", "Jesus & Mary", "Worldwide Katz", "Loyola Productions", "FFILME.RO"},
				ReleaseDate:         "2023-04-05",
				Budget:              1800000000,
				Revenue:             6567581600,
				Runtime:             103,
				Status:              "Released",
				Tagline:             "Inspired by the actual files of Father Gabriele Amorth, Chief Exorcist of the Vatican.",
				VoteAverage:         7.433,
				VoteCount:           545,
				Credits: []string{
					"Russell Crowe", "Daniel Zovatto", "Alex Essoe", "Franco Nero", "Peter DeSouza", "Feighoney", "Laurel Marsden", "Cornell John", "Ryan O'Grady", "Bianca Bardoe",
					"Santi Bayón",
					"Paloma Bloyd",
					"Alessandro Gruttadauria",
					"River Hawkins",
					"Jordi Collet",
					"Carrie Munro",
					"Marc Velasco",
					"Edward Harper",
					"Jones",
					"Matthew Sim",
					"Victor Solé",
					"Tom Bonington",
					"Andrea Dugoni",
					"Ed White",
					"Laila Barwick",
					"Gennaro Diana",
					"Pablo Raybould",
					"Ralph Ineson",
					"Derek Carroll",
					"Ella Cannon",
				},
				Keywords: []string{
					"spain",
					"rome italy",
					"vatican",
					"pope",
					"pig",
					"possession",
					"conspiracy",
					"devil",
					"exorcist",
					"skepticism",
					"catholic priest",
					"1980s",
					"supernatural horror",
				},
				PosterPath:   "/9JBEPLTPSm0d1mbEcLxULjJq9Eh.jpg",
				BackdropPath: "/hiHGRbyTcbZoLsYYkO4QiCLYe34.jpg",
			},
			{
				ID:                  uuid.Must(uuid.FromString("53aa35c9-e659-44b2-882f-f6056e443c99")),
				TmdbID:              "640146",
				Title:               "Ant-Man and the Wasp: Quantumania",
				Genres:              []string{"Action", "Adventure", "Science Fiction"},
				OriginalLanguage:    "en",
				Overview:            "Super-Hero partners Scott Lang and Hope van Dyne along with with Hope's parents Janet van Dyne and Hank Pym and Scott's daughter Cassie Lang find themselves exploring the Quantum Realm interacting with strange new creatures and embarking on an adventure that will push them beyond the limits of what they thought possible.",
				ReleaseDate:         "2023-02-15",
				Popularity:          4425.387,
				ProductionCompanies: []string{"Marvel Studios", "Kevin Feige Productions"},
				Budget:              2000000000,
				Revenue:             47576622800,
				Runtime:             125,
				Status:              "Released",
				Tagline:             "Witness the beginning of a new dynasty.",
				VoteAverage:         6.507,
				VoteCount:           2811,
				Credits: []string{
					"Paul Rudd",
					"Evangeline Lilly",
					"Jonathan Majors",
					"Kathryn Newton",
					"Michelle Pfeiffer",
					"Michael Douglas",
					"Corey Stoll",
					"Bill Murray",
					"William Jackson Harper",
					"David Dastmalchian",
					"Jamie Andrew Cutler",
					"Katy O'Brian",
					"Mark Weinman",
					"Randall Park",
					"Ross Mullan",
					"Tom Clark",
					"Leon Cooke",
					"Nathan Blees",
					"Durassie Kiangangu",
					"Liran Nathan",
					"Sam Symons",
					"Grahame Fox",
					"Nicola Peluso",
					"Harrison Daniels",
					"Brahmdeo Shannon Ramana",
					"Russell Balogh",
					"Leonardo Taiwo",
					"Osian Roberts",
					"Lucas Gerstel",
					"Mia Gerstel",
					"Tracy Jeffrey",
					"Dinah Jeffrey",
					"Judy Jeffrey",
					"John Nayagam",
					"Greta Nayagam",
					"Cathy Chan",
					"Adam Sai",
					"Jamie Sai",
					"Jakari Fraser",
					"Patricia Belcher",
					"Mark Oliver Everett",
					"Ruben Rabasa",
					"Melanie Garcia",
					"Gregg Turkington",
					"Sierra Katow",
					"Ryan Bergara",
					"Marielle Scott",
					"Jake Millgard",
					"Dey Young",
					"Briza Covarrubias",
					"Tess Aubert",
					"David J. Castillo",
					"Sir Cornwell",
					"Alan Heitz",
					"Esther McAuley",
					"Aisling Maria Andreica",
					"Milton Lopes",
					"Roger Craig Smith",
					"Matthew Wood",
					"Loveday Smith",
					"John Townsend",
					"Tom Hiddleston",
					"Owen Wilson",
					"Abby Ryder Fortson",
				},
				Keywords: []string{
					"hero",
					"ant",
					"sequel",
					"superhero",
					"based on comic",
					"family",
					"superhero team",
					"aftercreditsstinger",
					"duringcreditsstinger",
					"marvel cinematic universe (mcu)",
				},
				PosterPath:   "/qnqGbB22YJ7dSs4o6M7exTpNxPz.jpg",
				BackdropPath: "/m8JTwHFwX7I7JY5fPe4SjqejWag.jpg",
			},
		}

		moviesRepo := repository.NewMoviesInMemoryRepository([]*domain.Movie{})
		usecases := usecase.NewAppUsecases(moviesRepo)

		data, _ := usecases.IngestMoviesFile("./__tests__/movies-small.csv")

		if data.Count != len(want) {
			t.Errorf("got %d movies, want %d", data.Count, len(want))
		}

		movies, _ := moviesRepo.FindAll()

		if len(movies) != 2 {
			t.Errorf("got %d, want %d", len(movies), 2)
		}
	})
}

func TestIngestMoviesFileUsecaseMongo(t *testing.T) {
	t.Run("should ingest movies from csv file", func(t *testing.T) {
		want := []domain.Movie{
			{
				ID:                  uuid.Must(uuid.FromString("53aa35c8-e659-44b2-882f-f6056e443c99")),
				TmdbID:              "758323",
				Title:               "The Pope's Exorcist",
				Genres:              []string{"Horror", "Mystery", "Thriller"},
				OriginalLanguage:    "en",
				Overview:            "Father Gabriele Amorth Chief Exorcist of the Vatican investigates a young boy's terrifying possession and ends up uncovering a centuries-old conspiracy the Vatican has desperately tried to keep hidden.",
				Popularity:          5953.227,
				ProductionCompanies: []string{"Screen Gems", "2.0 Entertainment", "Jesus & Mary", "Worldwide Katz", "Loyola Productions", "FFILME.RO"},
				ReleaseDate:         "2023-04-05",
				Budget:              1800000000,
				Revenue:             6567581600,
				Runtime:             103,
				Status:              "Released",
				Tagline:             "Inspired by the actual files of Father Gabriele Amorth, Chief Exorcist of the Vatican.",
				VoteAverage:         7.433,
				VoteCount:           545,
				Credits: []string{
					"Russell Crowe", "Daniel Zovatto", "Alex Essoe", "Franco Nero", "Peter DeSouza", "Feighoney", "Laurel Marsden", "Cornell John", "Ryan O'Grady", "Bianca Bardoe",
					"Santi Bayón",
					"Paloma Bloyd",
					"Alessandro Gruttadauria",
					"River Hawkins",
					"Jordi Collet",
					"Carrie Munro",
					"Marc Velasco",
					"Edward Harper",
					"Jones",
					"Matthew Sim",
					"Victor Solé",
					"Tom Bonington",
					"Andrea Dugoni",
					"Ed White",
					"Laila Barwick",
					"Gennaro Diana",
					"Pablo Raybould",
					"Ralph Ineson",
					"Derek Carroll",
					"Ella Cannon",
				},
				Keywords: []string{
					"spain",
					"rome italy",
					"vatican",
					"pope",
					"pig",
					"possession",
					"conspiracy",
					"devil",
					"exorcist",
					"skepticism",
					"catholic priest",
					"1980s",
					"supernatural horror",
				},
				PosterPath:   "/9JBEPLTPSm0d1mbEcLxULjJq9Eh.jpg",
				BackdropPath: "/hiHGRbyTcbZoLsYYkO4QiCLYe34.jpg",
			},
			{
				ID:                  uuid.Must(uuid.FromString("53aa35c9-e659-44b2-882f-f6056e443c99")),
				TmdbID:              "640146",
				Title:               "Ant-Man and the Wasp: Quantumania",
				Genres:              []string{"Action", "Adventure", "Science Fiction"},
				OriginalLanguage:    "en",
				Overview:            "Super-Hero partners Scott Lang and Hope van Dyne along with with Hope's parents Janet van Dyne and Hank Pym and Scott's daughter Cassie Lang find themselves exploring the Quantum Realm interacting with strange new creatures and embarking on an adventure that will push them beyond the limits of what they thought possible.",
				ReleaseDate:         "2023-02-15",
				Popularity:          4425.387,
				ProductionCompanies: []string{"Marvel Studios", "Kevin Feige Productions"},
				Budget:              2000000000,
				Revenue:             47576622800,
				Runtime:             125,
				Status:              "Released",
				Tagline:             "Witness the beginning of a new dynasty.",
				VoteAverage:         6.507,
				VoteCount:           2811,
				Credits: []string{
					"Paul Rudd",
					"Evangeline Lilly",
					"Jonathan Majors",
					"Kathryn Newton",
					"Michelle Pfeiffer",
					"Michael Douglas",
					"Corey Stoll",
					"Bill Murray",
					"William Jackson Harper",
					"David Dastmalchian",
					"Jamie Andrew Cutler",
					"Katy O'Brian",
					"Mark Weinman",
					"Randall Park",
					"Ross Mullan",
					"Tom Clark",
					"Leon Cooke",
					"Nathan Blees",
					"Durassie Kiangangu",
					"Liran Nathan",
					"Sam Symons",
					"Grahame Fox",
					"Nicola Peluso",
					"Harrison Daniels",
					"Brahmdeo Shannon Ramana",
					"Russell Balogh",
					"Leonardo Taiwo",
					"Osian Roberts",
					"Lucas Gerstel",
					"Mia Gerstel",
					"Tracy Jeffrey",
					"Dinah Jeffrey",
					"Judy Jeffrey",
					"John Nayagam",
					"Greta Nayagam",
					"Cathy Chan",
					"Adam Sai",
					"Jamie Sai",
					"Jakari Fraser",
					"Patricia Belcher",
					"Mark Oliver Everett",
					"Ruben Rabasa",
					"Melanie Garcia",
					"Gregg Turkington",
					"Sierra Katow",
					"Ryan Bergara",
					"Marielle Scott",
					"Jake Millgard",
					"Dey Young",
					"Briza Covarrubias",
					"Tess Aubert",
					"David J. Castillo",
					"Sir Cornwell",
					"Alan Heitz",
					"Esther McAuley",
					"Aisling Maria Andreica",
					"Milton Lopes",
					"Roger Craig Smith",
					"Matthew Wood",
					"Loveday Smith",
					"John Townsend",
					"Tom Hiddleston",
					"Owen Wilson",
					"Abby Ryder Fortson",
				},
				Keywords: []string{
					"hero",
					"ant",
					"sequel",
					"superhero",
					"based on comic",
					"family",
					"superhero team",
					"aftercreditsstinger",
					"duringcreditsstinger",
					"marvel cinematic universe (mcu)",
				},
				PosterPath:   "/qnqGbB22YJ7dSs4o6M7exTpNxPz.jpg",
				BackdropPath: "/m8JTwHFwX7I7JY5fPe4SjqejWag.jpg",
			},
		}

		wantLen := len(want)

		mongo, err := db.NewDBMongo("mongodb://root:root@localhost:27017", "mycinema_ingest_test", 10*time.Second)
		if err != nil {
			t.Errorf("could not connect to mongodb: %v", err)
		}

		defer func() {
			mongo.DB.Drop(context.TODO())
			mongo.Close()
		}()

		moviesRepo := repository.NewMoviesMongoRepository(mongo)
		usecases := usecase.NewAppUsecases(moviesRepo)

		data, err := usecases.IngestMoviesFile("./__tests__/movies-small.csv")

		if err != nil {
			t.Errorf("Error while ingesting movies: %v", err)
		}

		if data.Count != wantLen {
			t.Errorf("got %d movies, want %d", data.Count, wantLen)
		}

		movies, _ := moviesRepo.FindAll()

		if len(movies) != wantLen {
			t.Errorf("got %d, want %d", len(movies), wantLen)
		}
	})
}
