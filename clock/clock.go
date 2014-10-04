package clock

import "fmt"

type Clock struct {
    hrs, mins int
}

func New(hours, minutes int) Clock {
    var hr, min int
    time := hours * 60 + minutes
    if time > 0 {

        for time >= 1440 {
            time = time - 1440
        }
        hr = time / 60
        min = time % 60

        return Clock{hrs: hr, mins: min}

    } else if time < 0 {

        for time < 0 {
            time = time + 1440
        }
        hr = time / 60
        min = time % 60

        return Clock{hrs: hr, mins: min}

    } else {
        return Clock{hrs: 0, mins: 0}
    }
}

func (c Clock) Add(minutes int) Clock {
    return New(c.hrs, c.mins + minutes)
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c.hrs, c.mins)
}