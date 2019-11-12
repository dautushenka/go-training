package practice3

import (
	"fmt"
)

var langs = map[string]map[string]string {
	"en": map[string]string {
		"lang": "English",
		"botName": "Bot",
		"hello": "Hello, I am %s",
		"bye": "Bye",
		"unknownCmd": "Unknown command, please try again",
		"introduction": "Please select command or enter command",
		"option1": "say hello",
		"command1": "hello",
		"option2": "say time",
		"command2": "time",
		"option3": "say date",
		"command3": "date",
		"option4": "say day of week",
		"command4": "day",
		"option5": "Exit",
		"command5": "exit",
		"time": "Now is %s",
		"date": "Today is %s",
		"day": "Today is %s",
	},
	"ru": map[string]string {
		"lang": "Русский",
		"botName": "Бот",
		"hello": "Привет, я %s",
		"bye": "Пока",
		"unknownCmd": "Неизвестная команда, попробуйте еще раз",
		"introduction": "Выберите вариант или введите команду",
		"option1": "Приветсвие",
		"command1": "привет",
		"option2": "Время",
		"command2": "время",
		"option3": "Дата",
		"command3": "дата",
		"option4": "День недели",
		"command4": "день",
		"option5": "Выход",
		"command5": "пока",
		"time": "Сейчас %s",
		"date": "Сегодня %s",
		"day": "Сегодня %s",
		"Moday": "Понедельник",
		"Tuesday": "Вторник",
		"Wednesday": "Среда",
		"Thursday": "Четверг",
		"Friday": "Пятница",
		"Saturday": "Суббота",
		"Sunday": "Воскресенье",
	},
};

func GetCopy(lang string) func(key string, args ...string) string {
	return func(key string, args ...string) string {
		inter := make([]interface{}, len(args));
		for i, v := range args {
			inter[i] = v;
		}
		if v, ok := langs[lang][key]; ok {
			return fmt.Sprintf(v, inter...);
		}
		return key;
		
	}
}

func getAvailabeLangs() (result map[string]string) {
	result = map[string]string{};
	for l, m := range langs {
		result[l] = m["lang"];
	}
	return;
}