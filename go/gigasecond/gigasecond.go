package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

//GIGASECOND large number seconds
const GIGASECOND = time.Duration(1e9) * time.Second

//AddGigasecond adds largenumner today
func AddGigasecond(t time.Time) time.Time {
	return t.Add(GIGASECOND)
}
