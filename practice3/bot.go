package practice3

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"strings"
)

var lang string = "en";

var commandList = map[string]func() {
	"1": func() { sayHello(lang)},
	"2": func() { sayTime(lang)},
	"3": func() { sayDate(lang)},
	"4": func() { sayDay(lang)},
	"hello": func() { sayHello("en")},
	"time": func() { sayTime("en")},
	"date": func() { sayDate("en")},
	"day": func() { sayDay("en")},
	"привет": func() { sayHello("ru")},
	"время": func() { sayTime("ru")},
	"дата": func() { sayDate("ru")},
	"день": func() { sayDay("ru")},
}

var exitCommandList = []string{
	"5", "пока", "bye", "exit", "halt", "shoutdown",
}

func RunBot() {
	lang = chooselang();
	runCommunication()
}

func runCommunication() {
	var input string;
	copy := GetCopy(lang);
	fmt.Println(copy("introduction"));
	printOptionList();
	for {
		if _, err := fmt.Scanln(&input); err != nil {
			input = "";
		}
		input = strings.ToLower(input);
		if checkExitCommand(input) {
			fmt.Println(copy("bye"));
			break;
		}
		if cmd, ok := commandList[input]; ok {
			cmd();
			continue;
		}
		fmt.Println(copy("unknownCmd"));
	}
	os.Exit(0);
}

func checkExitCommand(input string) bool {
	for _, v := range exitCommandList {
		if (v == input) {
			return true;
		}
	}
	return false;
}

func printOptionList() {
	copy := GetCopy(lang);
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d. %s\n", i, copy("option" + strconv.Itoa(i)));
	}
}

func chooselang() string {
	fmt.Println("Hi, please choose default language:");
	langs := getAvailabeLangs();
	i := 0; langMap := make([]string, len(langs));
	for key, label := range langs {
		fmt.Printf("%d. %s\n", i + 1, label);
		langMap[i] = key;
		i++;
	}
	for {
		var input int;
		if _, err := fmt.Scanln(&input); err != nil {
			panic(err);
		}
		if (input < 1 || input > len(langMap)) {
			fmt.Println("Wrong value, please select right one");
			continue;
		}
		return langMap[input - 1];

		break;
	}
	return lang;
}

func sayHello(lang string) {
	copy := GetCopy(lang);
	fmt.Println(copy("hello", copy("botName")))
}

func sayTime(lang string) {
  name := "Europe/Minsk"
  t := time.Now()
  loc, _ := time.LoadLocation(name)
  t = t.In(loc)

  fmt.Println(GetCopy(lang)("time", t.Format("15:04")));
  
}

func sayDate(lang string) {
  name := "Europe/Minsk"
  t := time.Now()
  loc, _ := time.LoadLocation(name)
  t = t.In(loc)

  fmt.Println(GetCopy(lang)("date", t.Format("01-02-2006")));
}

func sayDay(lang string) {
  name := "Europe/Minsk"
  t := time.Now()
  loc, _ := time.LoadLocation(name)
  t = t.In(loc)

  copy := GetCopy(lang);
  fmt.Println(copy("day", copy(t.Format("Monday"))));
}