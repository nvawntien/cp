func angleClock(hour int, minutes int) float64 {
    m := float64(minutes) * 6.0
    h := float64(hour%12)*30.0 + float64(minutes)*0.5

    diff := math.Abs(h - m)

    if diff > 180.0 {
        return 360.0 - diff
    }

    return diff
}