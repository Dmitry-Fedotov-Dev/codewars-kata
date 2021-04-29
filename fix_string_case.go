package kata

import(
    "strings"
)

func solve(str string) string {
    low_char := 0;
    high_char := 0;
    str_lower := strings.ToLower(str)
    var answer string

    for i := 0; i < len(str); i++ {
      if str[i] == str_lower[i] {
        low_char++
      } else {
        high_char++
      }
    } 
    if low_char >= high_char{
     answer = strings.ToLower(str)
         } else {
      answer = strings.ToUpper(str)
    } 
    return answer
 
}
