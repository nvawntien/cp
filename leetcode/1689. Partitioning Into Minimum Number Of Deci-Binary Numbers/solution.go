func minPartitions(n string) int {
    ans := 0
    for i := range n {
        if int(n[i] - '0') > ans {
            ans = int(n[i]-'0')
        }
    }   
    return ans
}