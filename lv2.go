package main

import "fmt"

type Movie struct {
	Name   string
	Time   string
	Actors string
	Story  string
}

func main() {
	m := Movie{
		Name:   "信条 TENET",
		Time:   "2020年9月4日",
		Actors: "约翰·大卫·华盛顿、罗伯特·帕丁森、伊丽莎白·德比齐、迈克尔·凯恩、肯尼思·布拉纳",
		Story: "一群蒙面匪徒闯入乌克兰一个歌剧院劫持人质，真实目标是要抢夺一个装有神秘物质的手提箱。主角和同伴假扮成特警也要抢夺手提箱，但任务离奇失败，\n" +
			"主角为保护秘密“被假死”。主角在假死被救后，被赋予新的更大的任务，通过女科学家的帮助，主角从射出的子弹倒着回到枪膛的实验中了解到\n" +
			"逆转时间”的概念，并被告知整个世界可能因为逆时间技术消失。主角从此不再有身份，加入“信条”组织，成为无名氏，他的任务是保护全\n" +
			"世界不要被逆时间毁灭。为了阻止萨特企图毁灭人类的行为，以无名氏、尼尔和凯特为首的队伍们来到萨特的基地利用逆时间\n" +
			"完成了任务，毁掉了萨特手中的时间炸弹，最终阻止了人类世界在逆时间中灭亡\n",
	}
	fmt.Printf("请输入你的命令\n1.获得名字\n2.上映时间\n3.演员列表\n4.电影简介\n5.退出程序\n")
	var option int
	for {
		fmt.Scanf("%d\n", &option)
		if option == 1 {
			fmt.Println(m.Name)
		} else if option == 2 {
			fmt.Println(m.Time)
		} else if option == 3 {
			fmt.Println(m.Actors)
		} else if option == 4 {
			fmt.Println(m.Story)
		} else if option == 5 {
			println("信条确实好康,推荐去康")
			return
		}
	}
}
