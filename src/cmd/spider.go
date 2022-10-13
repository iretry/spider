package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"spider/spider"
	"strconv"
)

// serveCmd represents the serve command
var cronCmd = &cobra.Command{
	Use:   "spider",
	Short: "start  spider job",
	Long:  `start  spider job`,
	Run: func(cmd *cobra.Command, args []string) {
		cronMain()
	},
}

var tieId string
var page string
var endPage string

func init() {
	rootCmd.AddCommand(cronCmd)
	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "id", "", "id")
	//fmt.Println(cfgFile)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	cronCmd.PersistentFlags().StringVar(&tieId, "id", "0","帖子id")
	cronCmd.PersistentFlags().StringVar(&page, "spage", "1","开始页数")
	cronCmd.PersistentFlags().StringVar(&endPage, "epage", "1","结束页数")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func startCron() {
	sPage,_ := strconv.Atoi(page)
	ePage,_ := strconv.Atoi(endPage)
	var spiders spider.Spider
	var jokersInfo spider.Joker
	jokersInfo.Url = "http://tieba.baidu.com/p/"+tieId+"?pn=%d"
	jokersInfo.Page = uint32(sPage)
	jokersInfo.EndPage = uint32(ePage)
	spiders = jokersInfo
	spiders.Start()
}

func cronMain() {
	log.Println("question cron starting ...")
	log.Println("question cron start init db ...")
	log.Println("question cron init db finish.")
	log.Println("question cron start")
	startCron()
}
