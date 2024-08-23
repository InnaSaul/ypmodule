package syscount

import (
	"fmt"
	"time"
)

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}
