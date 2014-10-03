package leap

import (
    "fmt"
    "os"
    "strconv"
)

//Returns true if the input year is a leap year, false otherwise.
func IsLeapYear(year int) bool {
    if year%4 == 0 {
        if year%100 == 0 {
            if year%400 == 0 {
                return true
            }
            return false
        }
        return true
    }
    return false
}

//could also do it like this!
//func IsLeapYear(year int) bool {
//    return year%4 == 0 && (year%100 != 0 || year%400 == 0)
//}

func main() {
    y, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Printf("error: %v/n", err)
    }

    year := IsLeapYear(y)
    if year {
        fmt.Printf("%v is a leap year!", year)
    } else {
        fmt.Printf("%v is not a leap year.", year)
    }
}