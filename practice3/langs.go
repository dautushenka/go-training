package practice3

import (
	"fmt"
)

var langs = map[string]map[string]string {
	"en": map[string]string {
		"lang": "English",
		"hello": "Hello, I am bot",
		"bye": "Bye",
		"unknownCmd": "Unknown command, please try again",
		"introduction": "Please select command or enter command",
		"option1": "say hello",
		"option2": "say time",
		"option3": "say date",
		"option4": "say day of week",
		"option5": "Exit",
		"time": "Now is %s",
		"date": "Today is %s",
		"day": "Today is %s",
	},
	"ru": map[string]string {
		"lang": "Русский",
		"hello": "Привет, я Бот",
		"bye": "Пока",
		"unknownCmd": "Неизвестная команда, попробуйте еще раз",
		"introduction": "Выберите вариант или введите команду",
		"option1": "Приветсвие",
		"option2": "Время",
		"option3": "Дата",
		"option4": "День недели",
		"option5": "Выход",
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