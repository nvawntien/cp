func minGenerations(points [][]int, target []int) int {
    visited := make(map[[3]int]bool)
    k := 0

    all := [][3]int{}
    t := [3]int{target[0], target[1], target[2]}
    
    for _, p := range points {
        if p[0] == target[0] && p[1] == target[1] && p[2] == target[2] {
            return 0
        }  
        pt := [3]int{p[0], p[1], p[2]}
        if !visited[pt] {
            all = append(all, pt)
            visited[pt] = true            
        }
    }
    
    for {
        k++
        n := len(all)
        newp := [][3]int{}

        for i := 0; i < n-1; i++ {
            for j := i+1; j < n; j++ {
                a := all[i]
                b := all[j]
                c := [3]int {
                    (a[0] + b[0]) >> 1,
                    (a[1] + b[1]) >> 1,
                    (a[2] + b[2]) >> 1,
                }

                if c == t {
                    return k
                }

                if !visited[c] {
                    visited[c] = true
                    newp = append(newp, c)
                }
            }
        }

        if (len(newp) == 0) {
            return -1
        }

        all = append(all, newp...)
    }

    return k
}