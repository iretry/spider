package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"spider/spider"
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
}


func startWeibo() {
	weibo := spider.Weibo{
		Url: "https://weibo.com/2970452952/profile?is_hot=1",
		Id: 2970452952,
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
