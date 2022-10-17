package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"spider/spider"
	"strconv"
)

// serveCmd represents the serve command
var weiboCmd = &cobra.Command{
	Use:   "weibo",
	Short: "start  spider job",
	Long:  `start  spider job`,
	Run: func(cmd *cobra.Command, args []string) {
		weiboMain()
	},
}

var uid string

func init() {
	rootCmd.AddCommand(weiboCmd)
	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "id", "", "id")
	//fmt.Println(cfgFile)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&uid, "uid", "0", "uid")

}

func startWeibo() {

	id, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	url := fmt.Sprintf("https://weibo.com/ajax/statuses/mymblog?uid=%d&page=1&feature=0", id)
	weibo := spider.Weibo{
		Url:     url,
		Id:      id,
		ShotUrl: fmt.Sprintf("https://weibo.com/%d", id),
		Path:    viper.GetString("basePath"),
	}
	weibo.Start()

}

func weiboMain() {
	log.Println("question cron starting ...")
	log.Println("question cron start init db ...")
	log.Println("question cron init db finish.")
	log.Println("question cron start")
	startWeibo()
}
