package domain_test

import (
	"encoding/csv"
	"reflect"
	"strings"
	"testing"

	"github.com/jrobic/my-cinema/movies-api/src/domain"
)

func TestNewMovieFromCSVRecord(t *testing.T) {
	t.Run("should return a movie", func(t *testing.T) {
		line := "758323,The Pope's Exorcist,Horror-Mystery-Thriller,en,Father Gabriele Amorth Chief Exorcist of the Vatican investigates a young boy's terrifying possession and ends up uncovering a centuries-old conspiracy the Vatican has desperately tried to keep hidden.,5953.227,Screen Gems-2.0 Entertainment-Jesus & Mary-Worldwide Katz-Loyola Productions-FFILME.RO,2023-04-05,18000000.0,65675816.0,103.0,Released,\"Inspired by the actual files of Father Gabriele Amorth, Chief Exorcist of the Vatican.\",7.433,545.0,Russell Crowe-Daniel Zovatto-Alex Essoe-Franco Nero-Peter DeSouza-Feighoney-Laurel Marsden-Cornell John-Ryan O'Grady-Bianca Bardoe-Santi Bayón-Paloma Bloyd-Alessandro Gruttadauria-River Hawkins-Jordi Collet-Carrie Munro-Marc Velasco-Edward Harper-Jones-Matthew Sim-Victor Solé-Tom Bonington-Andrea Dugoni-Ed White-Laila Barwick-Gennaro Diana-Pablo Raybould-Ralph Ineson-Derek Carroll-Ella Cannon,spain-rome italy-vatican-pope-pig-possession-conspiracy-devil-exorcist-skepticism-catholic priest-1980s-supernatural horror,/9JBEPLTPSm0d1mbEcLxULjJq9Eh.jpg,/hiHGRbyTcbZoLsYYkO4QiCLYe34.jpg"
		want := domain.Movie{
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
			VoteAverage:         7.4,
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
		}

		csvReader := csv.NewReader(strings.NewReader(line))
		rec, _ := csvReader.Read()

		got := domain.NewMovieFromCSVRecord(rec)

		if got.TmdbID != want.TmdbID {
			t.Errorf("TmdbID got %q, want %q", got.TmdbID, want.TmdbID)
		}

		if got.Title != want.Title {
			t.Errorf("Title got %q, want %q", got.Title, want.Title)
		}

		if len(got.Genres) != len(want.Genres) {
			t.Errorf("Genres got %q, want %q", got.Genres, want.Genres)
		}

		if got.OriginalLanguage != want.OriginalLanguage {
			t.Errorf("OriginalLanguage got %q, want %q", got.OriginalLanguage, want.OriginalLanguage)
		}

		if got.Overview != want.Overview {
			t.Errorf("Overview got %q, want %q", got.Overview, want.Overview)
		}

		if got.Popularity != want.Popularity {
			t.Errorf("Popularity got %v, want %v", got.Popularity, want.Popularity)
		}

		if len(got.ProductionCompanies) != len(want.ProductionCompanies) {
			t.Errorf("ProductionCompanies got %q, want %q", got.ProductionCompanies, want.ProductionCompanies)
		}

		if got.ReleaseDate != want.ReleaseDate {
			t.Errorf("ReleaseDate got %q, want %q", got.ReleaseDate, want.ReleaseDate)
		}

		if got.Budget != want.Budget {
			t.Errorf("Budget got %v, want %v", got.Budget, want.Budget)
		}

		if got.Revenue != want.Revenue {
			t.Errorf("Revenue got %v, want %v", got.Revenue, want.Revenue)
		}

		if got.Runtime != want.Runtime {
			t.Errorf("Runtime got %v, want %v", got.Runtime, want.Runtime)
		}

		if got.Status != want.Status {
			t.Errorf("Status got %q, want %q", got.Status, want.Status)
		}

		if !reflect.DeepEqual(got.Tagline, want.Tagline) {
			t.Errorf("Tagline got %v, want %v", got.Tagline, want.Tagline)
		}

		if got.VoteAverage != want.VoteAverage {
			t.Errorf("VoteAverage got %v, want %v", got.VoteAverage, want.VoteAverage)
		}

		if got.VoteCount != want.VoteCount {
			t.Errorf("VoteCount got %v, want %v", got.VoteCount, want.VoteCount)
		}

		if len(got.Credits) != len(want.Credits) {
			t.Errorf("Credits got %q, want %q", got.Credits, want.Credits)
		}

		if len(got.Keywords) != len(want.Keywords) {
			t.Errorf("Keywords got %q, want %q", got.Keywords, want.Keywords)
		}

		if got.PosterPath != want.PosterPath {
			t.Errorf("PosterPath got %q, want %q", got.PosterPath, want.PosterPath)
		}

		if got.BackdropPath != want.BackdropPath {
			t.Errorf("BackdropPath got %q, want %q", got.BackdropPath, want.BackdropPath)
		}
	})
}
