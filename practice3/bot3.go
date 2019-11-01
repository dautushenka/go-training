package practice3

import (
	"fmt"
	"os"
	"strings"
)

type IBot interface {
	sayHello()
	sayBye()
	sayTime()
	sayDate()
	sayDay()
	t(key string, args ...string) string
}

func RunBot3() {
	botEng := CreateBot("en")
	botRu := CreateBot("ru")

	var bot IBot
	selectedLang := chooselang()
	if selectedLang == "en" {
		bot = &botEng
	} else {
		bot = &botRu
	}
	runCommunication3(bot)
}

func runCommunication3(bot IBot) {
	var input string
	fmt.Println(bot.t("introduction"))
	printOptionList()
	for {
		if _, err := fmt.Scanln(&input); err != nil {
			input = ""
		}
		input = strings.ToLower(input)

		switch input {
		case "1":
			fallthrough
		case bot.t("command1"):
			bot.sayHello()
			break
		case "2":
			fallthrough
		case bot.t("command2"):
			bot.sayTime()
			break
		case "3":
			fallthrough
		case bot.t("command3"):
			bot.sayDate()
			break
		case "4":
			fallthrough
		case bot.t("command4"):
			bot.sayDay()
			break
		case "5":
			fallthrough
		case bot.t("command5"):
			fmt.Println(bot.t("bye"))
			os.Exit(0)
		default:
			fmt.Println(bot.t("unknownCmd"))
		}
	}
}
