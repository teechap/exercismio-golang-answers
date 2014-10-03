package clock

import "strconv"

type Clock struct {
    hrs int
    mins int
}

func New(hours, minutes int) Clock {
    switch {
    case minutes >= 60:
        return New(hours + 1, minutes - 60)
    case minutes <= -60:
        return New(hours - 1, minutes + 60)
    case hours < 0:
        return New(hours + 24, minutes)
    case hours > 23:
        return New(hours - 24, minutes)
    case minutes < 0:
        return New(hours - 1, minutes + 60)
    }
    return Clock{hrs: hours, mins: minutes}
}

func (c Clock) Add(minutes int) Clock {
    return New(c.hrs, c.mins + minutes)
}

func (c Clock) String() string {
    var hours, minutes string
    if c.hrs >= 10 {
        hours = strconv.Itoa(c.hrs)
    } else {
        hours = "0" + strconv.Itoa(c.hrs)
    }
    if c.mins >= 10 {
        minutes = strconv.Itoa(c.mins)
    } else {
        minutes = "0" + strconv.Itoa(c.mins)
    }
    result := hours + ":" + minutes
    return result
}