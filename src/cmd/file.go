package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"

	//"google.golang.org/grpc/grpclog"
	"log"
)

// serveCmd represents the serve command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "start  file service",
	Long:  `start  file service`,
	Run: func(cmd *cobra.Command, args []string) {
		serveFileMain()
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}

func serveFileMain() {
	log.Println("file starting ...")

	http.Handle("/", http.FileServer(http.Dir(viper.GetString("basePath")+"/shot")))
	e := http.ListenAndServe(":8888", nil)
	fmt.Println(e)
}
