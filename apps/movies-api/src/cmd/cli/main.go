package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jrobic/my-cinema/movies-api/src/config"
	"github.com/jrobic/my-cinema/movies-api/src/infra/db"
	"github.com/jrobic/my-cinema/movies-api/src/infra/repository"
	"github.com/jrobic/my-cinema/movies-api/src/usecase"
	"github.com/spf13/cobra"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	versionCmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("movies-api version: %s\n", "0.1")
		},
	}

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Print the config",
		Run: func(cmd *cobra.Command, args []string) {
			cfg.Debug()
		},
	}

	rootCmd := &cobra.Command{
		Use: "movies-api-cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	ingestCmd := &cobra.Command{
		Use:     "import",
		Short:   "Import movies from a csv file",
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"i"},
		Example: `movies-api-cli import -f movies.csv`,
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]

			if len(file) < 4 || file[len(file)-4:] != ".csv" {
				fmt.Printf("File %s is not a .csv file\n", file)
				os.Exit(1)
			}

			if _, err := os.Stat(file); os.IsNotExist(err) {
				fmt.Printf("File %s does not exists\n", file)
				os.Exit(1)
			}

			fmt.Printf("Importing movies from %s\n", file)

			// Init Databases
			mongoDB, err := db.NewDBMongo(cfg.DBMongo.URI, cfg.DBMongo.DBName, 10*time.Second)
			defer mongoDB.Close()
			if err != nil {
				fmt.Printf("could not connect to mongodb: %v", err)
				os.Exit(1)
				return
			}

			moviesRepo := repository.NewMoviesMongoRepository(mongoDB)

			usecases := usecase.NewAppUsecases(moviesRepo)

			data, err := usecases.IngestMoviesFile(file)

			if err != nil {
				fmt.Printf("Error importing movies: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Imported %d movies\n", data.Count)
		},
	}

	rootCmd.AddCommand(versionCmd, configCmd, ingestCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
