package util

//"gin-blog/models"

// func Cron() {
// 	log.Println("Starting...")

// 	//根据本地时间创建一个新（空白）的 Cron job runner
// 	c := cron.New()

// 	//AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
// 	c.AddFunc("* * * * * *", func() {
// 		log.Println("Run models.CleanAllTag...")
// 		models.CleanAllTag()
// 	})
// 	c.AddFunc("* * * * * *", func() {
// 		log.Println("Run models.CleanAllArticle...")
// 		models.CleanAllArticle()
// 	})
// 	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制哦
// 	c.Start()
// 	//会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息
// 	t1 := time.NewTimer(time.Second * 10)
// 	//阻塞 select 等待 channel
// 	for {
// 		select {
// 		case <-t1.C:
// 			//会重置定时器，让它重新开始计时 （注意，本文适用于 “t.C已经取走，可直接使用 Reset”）
// 			t1.Reset(time.Second * 10)
// 		}
// 	}
// }
