package main

import (
    "fmt"
    "strings"
)

func main() {
    s1 := "rubyab"
    s2 := "railsb"
    
    answer := compareString(s1, s2)
    fmt.Println(answer)
}

func compareString(str1 string, str2 string) string {
    subjectChars1 := strings.Split(str1, "")
    subjectChars2 := strings.Split(str2, "")
    repetitionChars := make([]string, 0)
    
    for _, str := range subjectChars1 {
        if contains(subjectChars2, str) {
            if !contains(repetitionChars, str) {
                repetitionChars = append(repetitionChars, str)
            }
        }
    }
    
    return strings.Join(repetitionChars, "")
}

func contains(slice []string, str string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }
    return false
}
