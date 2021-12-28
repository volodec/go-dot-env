package go_dot_env

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func SearchKey(key string) string {
	file, err := os.Open("./.env")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
		return ""
	}

	fileScanner := bufio.NewScanner(file)
	var str string
	// объект regexp, с помощью которого будет осуществлятся поиск строки с именем искомой переменной
	re := regexp.MustCompile(`^` + regexp.QuoteMeta(key) + `=`)

	// построчный поиск первого совпадения ключа в файле
	for fileScanner.Scan() {
		str = fileScanner.Text()
		match := re.MatchString(str)
		if match {
			break
		}
	}
	// объект regexp, с помощью которого будет взято слово перед знаком равно
	return re.ReplaceAllString(str, "")
}
