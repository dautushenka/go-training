package practice3

import (
	"fmt"
	"time"
	"strconv"
	"os"
	"strings"
)

var loc *time.Location;

func init() {
	loc, _ = time.LoadLocation("Europe/Minsk");
}

type Bot struct {
	name string
	copy func(string, ...string) string
}

func RunBot2() {
	var input string;
	bot := CreateBot(chooselang());
	bot.printOptionList();
	for {
		if _, err := fmt.Scanln(&input); err != nil {
			input = "";
		}
		input = strings.ToLower(input);
		if bot.checkExitCommand(input) {
			bot.sayBye();
			break;
		}
		bot.ExecuteCommand(input);
	}
	os.Exit(0);
}

func CreateBot(lang string) Bot {
	copy := GetCopy(lang);

	return Bot{copy("botName"), copy};
}

func (s *Bot) printOptionList() {
	fmt.Println(s.copy("introduction"));
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d. %s\n", i, s.copy("option" + strconv.Itoa(i)));
	}
}

func (s *Bot) checkExitCommand(command string) bool {
	switch command {
		default: return false;
		case "5": fallthrough;
		case s.copy("command5"): return true;
	}
}

func (s *Bot) ExecuteCommand(command string) {
	switch command {
		case "1": fallthrough;
		case s.copy("command1"): s.sayHello(); 
			break;
		case "2": fallthrough;
		case s.copy("command2"): s.sayTime();
			break;
		case "3": fallthrough;
		case s.copy("command3"): s.sayDate();
			break;
		case "4": fallthrough;
		case s.copy("command4"): s.sayDay();
			break;
		default: 
			fmt.Println(s.copy("unknownCmd"));
	}
}

func (s *Bot) sayHello() {
	fmt.Println(s.copy("hello", s.name))
}

func (s *Bot) sayBye() {
	fmt.Println(s.copy("bye"))
}

func (s *Bot) sayTime() {
	t := time.Now().In(loc)
	fmt.Println(s.copy("time", t.Format("15:04")));
}
  
func (s *Bot) sayDate() {
	t := time.Now().In(loc)

	fmt.Println(s.copy("date", t.Format("01-02-2006")));
}

func (s *Bot) sayDay() {
	t := time.Now().In(loc);

	fmt.Println(s.copy("day", s.copy(t.Format("Monday"))));
}