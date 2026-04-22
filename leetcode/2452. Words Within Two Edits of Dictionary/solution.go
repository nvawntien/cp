func twoEditWords(queries []string, dictionary []string) []string {
    n := len(queries[0])
    ans := []string{}

    for _, s1 := range queries {
        for _, s2 := range dictionary {
            diff := 0
            for i := 0; i < n; i++ {
                if s1[i] != s2[i] {
                    diff++
                }
            }
            if diff <= 2 {
                ans = append(ans, s1)
                break
            }        
        }
    }

    return ans
}