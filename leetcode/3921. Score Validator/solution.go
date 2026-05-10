func scoreValidator(events []string) []int {
    score := 0
    counter := 0
    freq := make(map[string]bool)
    
    for _, s := range events {
        if counter == 10 {
            break
        }
        
        if s == "0" {
            score += 0
            freq[s] = true
        }

        if s == "1" {
            score += 1
            freq[s] = true
        }

        if s == "2" {
            score += 2
            freq[s] = true
        }

        if s == "3" {
            score += 3
            freq[s] = true
        }

        if s == "4" {
            score += 4
            freq[s] = true
        }

        if s == "6" {
            score += 6
            freq[s] = true
        }

        if s == "W" {
            counter += 1
            freq[s] = true
        }

         if s == "WD" || s == "NB"{
            score += 1
            freq[s] = true
        }
    }

    return []int{score, counter}
}