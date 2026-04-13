func isValid(s string) bool {
    arr := []byte {}

    for i := 0; i < len(s); i++ {
        if s[i] == ')' {
            if len(arr) == 0 || arr[len(arr)-1] != '(' {
                return false
            } else if arr[len(arr)-1] == '(' {
                arr = arr[:len(arr)-1]
            }
        } else if s[i] == ']' {
            if len(arr) == 0 || arr[len(arr)-1] != '[' {
                return false
            } else if arr[len(arr)-1] == '[' {
                arr = arr[:len(arr)-1]
            }
        } else if s[i] == '}' {
            if len(arr) == 0 || arr[len(arr)-1] != '{' {
                return false
            } else if arr[len(arr)-1] == '{' {
                arr = arr[:len(arr)-1]
            }
        } else {
            arr = append(arr, s[i])
        }
    }

    if len(arr) != 0 {
        return false
    }
    
    return true;
}