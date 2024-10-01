package cmd

import (
	"context"
	"kompack-go-api/config"
	"kompack-go-api/internal/factory"
	cHttp "kompack-go-api/internal/http"
	"kompack-go-api/pkg/database"
	"kompack-go-api/pkg/tracer"
	"kompack-go-api/pkg/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "The rest command to handle RESTful operations",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv(".env", "env", ".")

		tracer.InitLogger()

		f := factory.NewFactory()
		g := gin.New()

		if config.C.GinMode == "release" {
			gin.SetMode(gin.ReleaseMode)
		}

		cHttp.NewHttp(g, f)

		srv := &http.Server{
			Addr:    ":" + config.C.Port,
			Handler: g,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("failed start server: %s", err.Error())
			}
		}()

		wait := util.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"http-server": func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
			"postgres": func(ctx context.Context) error {
				sql := database.GetConnection()

				return database.Close(sql)
			},
		})

		<-wait
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
