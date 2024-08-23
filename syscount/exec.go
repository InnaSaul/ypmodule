package syscount

import (
	"fmt"
	"time"
)

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}

/*
func VeryLongTimeFunction () {
	// передаём в функцию metricTime значение текущего времени и откладываем её вызов до возврата
        defer     metricTime(time.Now())
		// Какие-то долгие вычисления
    }
*/
