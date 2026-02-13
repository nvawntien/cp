func longestBalanced(s string) int {
    res := []int {
        case1(s, 'a'),
        case1(s, 'b'),
        case1(s, 'c'),
        case2(s, 'a', 'b'),
        case2(s, 'a', 'c'),
        case2(s, 'b', 'c'),
        case3(s),
    }

    ans := 0
    for _, val := range res {
        if val > ans {
            ans = val
        }
    } 
    return ans
}

func case1(s string, c byte) int {
    curr, ans := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == c {
            curr++
            if curr > ans {
                ans = curr
            }
        } else {
            curr = 0
        }
    }
    return ans    
} 

func case2(s string, a, b byte) int {
    cnt1, cnt2 := 0, 0
    ans := 0
    pre := make(map[int]int)

    for i := 0; i < len(s); i++ {
        if s[i] == a {
            cnt1++
        } else if s[i] == b {
            cnt2++
        } else {
            cnt1 = 0
            cnt2 = 0
            pre = map[int]int{}
            continue
        }

        if cnt1 == cnt2 {
            if cnt1 * 2 > ans {
                ans = cnt1 << 1
            }
        } else {
            key := (cnt1 - cnt2) * 100001
            if pos, ok := pre[key]; ok {
                ans = max(ans, i-pos)
            } else {
                pre[key] = i
            }
        }
    }

    return ans 
}

func case3(s string) int {
    a, b, c := 0, 0, 0  
    pre := make(map[int]int)
    ans := 0

    for i := 0; i < len(s); i++ {
        switch s[i] {
            case 'a':
                a++
                case 'b':
                    b++
                    default:
                        c++
        }

        if a == b && b == c {
            ans = i+1
        } else {
            key := (a - b) * 100001 + (b-c)
            if pos, ok := pre[key]; ok {
                ans = max(ans, i - pos)
            } else {
                pre[key] = i
            }
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}