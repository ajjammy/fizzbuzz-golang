package fizzbuzz

import "strconv"

type OtherRule struct{}

func (r *OtherRule) IsHandle(arabic int) bool {
    return true
}

func (r *OtherRule) GetAnswer(arabic int) string {
    return strconv.Itoa(arabic)
}
