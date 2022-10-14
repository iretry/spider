package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"spider/spider"
	"time"
)

// serveCmd represents the serve command
var childCmd = &cobra.Command{
	Use:   "child",
	Short: "start  child job",
	Long:  `start  child job`,
	Run: func(cmd *cobra.Command, args []string) {
		childMain()
	},
}

var day string

func init() {
	rootCmd.AddCommand(childCmd)

	/**


	 */

	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "id", "", "id")
	//fmt.Println(cfgFile)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	childCmd.PersistentFlags().StringVar(&day, "day", "2022-04-01", "day")

}

type rowTips struct {
	Child string
	Mom   string
}

func startChild() {
	tips := map[int64]rowTips{
		29: {
			Child: "29周的胎儿现在坐高约26-27厘米，体重约1300克左右，这时的宝宝，如果是男孩，他的睾丸已经从腹中降下来;如果是女孩，您可以看到宝宝突起的小阴唇。",
			Mom:   "准妈妈的体重每周可以增加500克左右，子宫底持续增高，上升到肚脐上面的1-2横指处，子宫高度为24-26厘米。",
		},
		30: {Child: "30周的胎儿现在身高约44厘米，体重约1500克左右。胎儿在头部在继续增大，大脑发育也非常迅速。大脑和神经系统已经发达到一定的程度，皮下脂肪继续增长。", Mom: "准妈妈的体重每周可以增加500克左右，子宫底持续增高，上升到肚脐上面的1-2横指处，子宫高度为24-26厘米。"},
		31: {Child: "31周的胎儿身体和四肢继续长大，直到比例相当。胎儿现在的体重约为2000克左右。宝宝的皮下脂肪更加丰富了，皱纹减少。", Mom: "准妈妈的体重每周可以增加500克左右，子宫底持续增高，上升到肚脐上面的1-2横指处，子宫高度为24-26厘米"},
		32: {Child: "32周的胎儿身长约40厘米，顶臀长（坐高）约28厘米，体重约1.5~1.6千克。现在的胎儿与出生时的婴儿相似，但身体仍需要长胖些。他的手指甲和脚指甲已经完全长出来了。有些宝宝已经长了满头的头发，有些只长出了淡淡的绒毛。他的眼睛能区分光亮与黑暗。如果宝宝是男孩，他的睾丸可能已经从腹腔进入阴囊， 但是有的宝宝可能会在出生后当天才进入阴囊。", Mom: "准妈妈的体重已经比上个月增加了1800克左右，身体越来越沉重，由于子宫底已经升到了横膈膜处，增大的子宫，像四周“扩充底盘”，就会挤压着心脏、肺脏和小腹部，可能会引起食欲不佳、呼吸困难、尿频加重等症状。"},
		33: {Child: "33周的胎儿身长约48厘米，体重约2200克。胎儿的呼吸系统和消化系统发育已经接近成熟，33周的胎儿应当注意头的位置。", Mom: "准妈妈的体重已经比上个月增加了1800克左右，身体越来越沉重，由于子宫底已经升到了横膈膜处，增大的子宫，像四周“扩充底盘”，就会挤压着心脏、肺脏和小腹部，可能会引起食欲不佳、呼吸困难、尿频加重等症状。"},
		34: {Child: "34周的胎儿坐高约30厘米，体重2300克左右。胎儿现在圆圆的开始变胖。胎儿的皮下脂肪形成后将会在宝宝出生后调节体温。同时宝宝也在为分娩做准备了，宝宝的头转向下方， 头部进入骨盆。", Mom: "准妈妈的体重已经比上个月增加了1800克左右，身体越来越沉重，由于子宫底已经升到了横膈膜处，增大的子宫，像四周“扩充底盘”，就会挤压着心脏、肺脏和小腹部，可能会引起食欲不佳、呼吸困难、尿频加重等症状。"},
		35: {Child: "这时胎儿身长约50厘米左右，体重约2500克。胎儿现在圆圆的开始变胖。胎儿的皮下脂肪形成后将会在宝宝出生后调节体温。", Mom: "准妈妈的体重已经比上个月增加了1800克左右，身体越来越沉重，由于子宫底已经升到了横膈膜处，增大的子宫，像四周“扩充底盘”，就会挤压着心脏、肺脏和小腹部，可能会引起食欲不佳、呼吸困难、尿频加重等症状。"},
		36: {Child: "36周的胎儿大约重2.7千克，身长约45~50厘米。覆盖宝宝全身的绒毛和在羊水中保护宝宝皮肤的胎脂正在开始脱落。宝宝现在会吞咽这些脱落的物质和其他分泌物了。", Mom: "准妈妈的体重仍以每星期500克的速度在增长，子宫继续增长，子宫底高28-30厘米，已经升到心口窝，身体越发沉重。由于子宫壁和腹壁已经变得很薄，胎宝宝在肚子里活动时，准妈妈有时候能看到一个很明显的“鼓包”，那可能是胎宝宝的手脚或肘部。"},
		37: {Child: "到了第10个月，孕妇在最后的这个月会感觉很紧张心情烦躁焦急等，同时孕妇在这几周中身体会越来越感到沉重，要注意小心活动，避免长期站立，洗澡的时候避免滑倒等。总之，好好休息，密切注意自己身体的变化，随时{Child:做好临产的准备。", Mom: "准妈妈的体重仍以每星期500克的速度在增长，子宫继续增长，子宫底高28-30厘米，已经升到心口窝，身体越发沉重。由于子宫壁和腹壁已经变得很薄，胎宝宝在肚子里活动时，准妈妈有时候能看到一个很明显的“鼓包”，那可能是胎宝宝的手脚或肘部。"},
		38: {Child: "怀孕十个月的第二周，现在你的胎儿可能已经有3200克重了，身长也得有50厘米左右了。胎儿的头在你的骨盆腔内摇摆，周围有骨盆的骨架保护，很安全。这样也腾出了更多的地方长他的小胳膊、小腿、小屁股。很多胎儿这时{Child:头发已长得较长较多，大约有1—3厘米长，如果父母中某一方头发是自来卷的话，你的胎儿也很可能是个小卷毛头。", Mom: "准妈妈的体重仍以每星期500克的速度在增长，子宫继续增长，子宫底高28-30厘米，已经升到心口窝，身体越发沉重。由于子宫壁和腹壁已经变得很薄，胎宝宝在肚子里活动时，准妈妈有时候能看到一个很明显的“鼓包”，那可能是胎宝宝的手脚或肘部。"},
		39: {Child: "怀孕十个月的第三周，一般情况下男孩比女孩的平均体重略重一些。胎儿现在还在继续长肉，这些脂肪储备将会有助于宝宝出生后的体温调节。这个小家伙的身体各部份器官已发育完成，其中肺部是最后一个成熟的器官，在{Child:宝宝出生后几个小时内他才能建立起正常的呼吸模式", Mom: "准妈妈的体重仍以每星期500克的速度在增长，子宫继续增长，子宫底高28-30厘米，已经升到心口窝，身体越发沉重。由于子宫壁和腹壁已经变得很薄，胎宝宝在肚子里活动时，准妈妈有时候能看到一个很明显的“鼓包”，那可能是胎宝宝的手脚或肘部。"},
		40: {Child: "如果是选择顺产的妈咪朋友。快到临产期时不要再一个人外出散步了，因为在这个特殊的阶段，您很有可能会碰到意外之喜——宝宝提前降生。正常情况下只有当宫缩开始，宫颈不断扩张，包裹在胎儿和羊水外面的卵膜会在不断增加的压力下破裂，流出大量羊水，胎儿也将随之降生。", Mom: "到了怀孕40周，也就是到了分娩的日子，准妈妈的体重达到最高值，子宫底高30-35厘米，胎宝宝的位置有所下降，对胃和心脏的压迫减轻，准妈妈呼吸变得舒畅，同时身体为分娩做的准备已经成熟，不规则的宫缩在增加，开始出现临产征兆。"},
	}
	now := time.Now().Unix()

	pre, _ := time.ParseInLocation("2006-01-02", day, time.Local)

	week := (now-pre.Unix())/86400/7 + 1
	oDay := (now-pre.Unix())/86400%7 + 1
	if oDay != 1 {
		return
	}
	if row, ok := tips[week]; ok {
		postInfo := spider.JsonPostSample{
			Msgtype: "markdown",
			Markdown: spider.Markdown{
				Title: fmt.Sprintf("孕周:今天是%d周%d天", week, oDay),
				Text:  "#### " + fmt.Sprintf("今天是%d周%d天", week, oDay) + "\n > 宝宝的变化：" + row.Child + "\n\n > 妈妈的变化：" + row.Mom,
			},
		}
		postInfo.SamplePost()
	}

}

func childMain() {
	log.Println("child cron starting ...")
	log.Println("child cron start init db ...")
	log.Println("child cron init db finish.")
	log.Println("child cron start")
	startChild()
}
