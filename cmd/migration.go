package cmd

import (
	"fmt"
	"kompack-go-api/config"
	"kompack-go-api/pkg/database"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	i   bool
	mmf string
	f   string
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Provides several flags to handle database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if mmf != "" {
			err := database.CreateMigrationFile(mmf)
			if err != nil {
				fmt.Println(cases.Title(language.Indonesian).String(err.Error()))
			}
			return
		}

		config.LoadEnv(".env", "env", ".")

		database.CreateConnection()

		if f != "" {
			if f == "all" {
				err := database.MigrateAll()
				if err != nil {
					fmt.Println(cases.Title(language.Indonesian).String(err.Error()))
				}
				return
			}

			err := database.Migrate(f)
			if err != nil {
				fmt.Println(cases.Title(language.Indonesian).String(err.Error()))
			}
			return
		}

		if i {
			database.FirstMigrate()
			return
		}
	},
}

func init() {
	migrateCmd.Flags().StringVar(&mmf, "mmf", "", "this flag is used for creating migration file name")
	migrateCmd.Flags().StringVarP(&f, "file", "f", "", "this flag is used for migration by file")
	migrateCmd.Flags().BoolVarP(&i, "init-migration", "i", false, "this flag is used for first migration")
	migrateCmd.MarkFlagsMutuallyExclusive("mmf", "file", "init-migration")
	rootCmd.AddCommand(migrateCmd)
}
