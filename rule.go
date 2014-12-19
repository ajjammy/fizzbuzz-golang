package fizzbuzz

type Rule interface{
    IsHandle(arabic int) bool
    GetAnswer(arabic int) string
}
