package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/godcong/chronos"
	"github.com/godcong/fate"
	"github.com/godcong/fate/config"
	"os"
)

func main() {

	//cfg := config.DefaultConfig() 参数如下
	//config.Config{
	//	HardFilter: false,
	//	//输出最大笔画数
	//	StrokeMax: 3,
	//	//输出最小笔画数
	//	StrokeMin: 18,
	//	//立春修正（出生日期为立春当日时间为已过立春八字需修正）
	//	FixBazi: true,
	//	//三才五格过滤
	//	SupplyFilter: true,
	//	//生肖过滤
	//	ZodiacFilter: true,
	//	//周易八卦过滤
	//	BaguaFilter: true,
	//	//连接DB：
	//	Database: config.Database{
	//		Host:         "localhost",
	//		Port:         "3306",
	//		User:         "root",
	//		Pwd:          "111111",
	//		Name:         "fate",
	//		MaxIdleCon:   0,
	//		MaxOpenCon:   0,
	//		Driver:       "mysql",
	//		File:         "",
	//		Dsn:          "",
	//		ShowSQL:      false,
	//		ShowExecTime: false,
	//	},
	//})
	//出生日期
	fmt.Println("Enter Date(ex: 2020/02/01 18:16): ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	born := chronos.New(text)

	//姓氏
	lastName := "刘"
	cfg := config.DefaultConfig()
	cfg.BaguaFilter = true
	cfg.ZodiacFilter = true
	cfg.SupplyFilter = true
	cfg.HardFilter = true
	cfg.StrokeMin = 3
	cfg.StrokeMax = 24
	cfg.Database = config.Database{
		Host:         "localhost",
		Port:         "3306",
		User:         "root",
		Pwd:          "***",
		Name:         "fate",
		MaxIdleCon:   0,
		MaxOpenCon:   0,
		Driver:       "mysql",
		File:         "",
		Dsn:          "",
		ShowSQL:      false,
		ShowExecTime: false,
	}
	cfg.FileOutput = config.FileOutput{
		OutputMode: config.OutputModeLog,
		Path:       "name.log",
	}

	f := fate.NewFate(lastName, born.Solar().Time(), fate.ConfigOption(cfg))

	e := f.MakeName(context.Background())
	if e != nil {
		panic(e)
	}
}
