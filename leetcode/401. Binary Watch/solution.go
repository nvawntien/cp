func readBinaryWatch(turnedOn int) []string {
    res := []string{}

    for h := 0; h < 12; h++ {
        for m := 0; m < 60; m++ {
            if bits(h)+bits(m) == turnedOn {
                res = append(res, fmt.Sprintf("%d:%02d", h, m))
            }
        }
    }
    return res
}

func bits(x int) int {
    cnt := 0
    for x > 0 {
        cnt += x & 1
        x >>= 1
    }
    return cnt
}
