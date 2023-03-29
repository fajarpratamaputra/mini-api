package common

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"math"
	"strconv"
)

func ValidateRequestPayload(c echo.Context, rules govalidator.MapData, data interface{}) (i interface{}) {
	opts := govalidator.Options{
		Request: c.Request(),
		Data:    data,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	mappedError := v.ValidateJSON()
	if len(mappedError) > 0 {
		i = mappedError
	}

	return i
}

func TotalFormatting(value int) string {
	var tempValue float64
	var newValue string
	if value >= 1000 && value < 1000000 {
		tempValue = math.Floor(float64(value)) / 1000
		tempVal := fmt.Sprintf("%.3f", tempValue)

		if tempVal[2] == '9' || tempVal[3] == '9' {
			if tempVal[2] != '0' {
				newValue = fmt.Sprintf("%c.%cK", tempVal[0], tempVal[2])
			} else {
				val := int(tempValue)
				newValue = fmt.Sprintf("%dK", val)
			}
		} else {
			tempVal = fmt.Sprintf("%.2f", tempValue)
			if last := len(tempVal) - 1; last >= 0 {
				tempVal = tempVal[:last]
			}

			if last := len(tempVal) - 1; tempVal[last] != '0' {
				newValue = fmt.Sprintf("%sK", tempVal)
			} else {
				val := int(tempValue)
				newValue = fmt.Sprintf("%dK", val)
			}
		}

	} else if value >= 1000000 && value < 1000000000 {
		tempValue = math.Floor(float64(value)) / 1000000
		tempVal := fmt.Sprintf("%.6f", tempValue)

		if tempVal[2] == '9' || tempVal[3] == '9' {
			if tempVal[2] != '0' {
				newValue = fmt.Sprintf("%c.%cM", tempVal[0], tempVal[2])
			} else {
				val := int(tempValue)
				newValue = fmt.Sprintf("%dM", val)
			}
		} else {
			tempVal = fmt.Sprintf("%.2f", tempValue)
			if last := len(tempVal) - 1; last >= 0 {
				tempVal = tempVal[:last]
			}

			if last := len(tempVal) - 1; tempVal[last] != '0' {
				newValue = fmt.Sprintf("%sM", tempVal)
			} else {
				val := int(tempValue)
				newValue = fmt.Sprintf("%dM", val)
			}
		}

	} else if value >= 1000000000 {
		tempValue = math.Floor(float64(value)) / 1000000000
		tempVal := fmt.Sprintf("%.9f", tempValue)

		if tempVal[2] == '9' || tempVal[3] == '9' {
			if tempVal[2] != '0' {
				newValue = fmt.Sprintf("%c.%cB", tempVal[0], tempVal[2])
			} else {
				val := int(tempValue)
				newValue = fmt.Sprintf("%dB", val)
			}
		} else {
			tempVal = fmt.Sprintf("%.2f", tempValue)
			if last := len(tempVal) - 1; last >= 0 {
				tempVal = tempVal[:last]
			}

			if last := len(tempVal) - 1; tempVal[last] != '0' {
				newValue = fmt.Sprintf("%sB", tempVal)
			} else {
				val := int(tempValue)
				newValue = fmt.Sprintf("%dB", val)
			}
		}

	} else {
		newValue = strconv.Itoa(value)
	}

	return newValue
}
