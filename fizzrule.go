package fizzbuzz

type FizzRule struct{}

func (r *FizzRule) IsHandle(arabic int) bool {
    return arabic % 3 == 0
}

func (r *FizzRule) GetAnswer(arabic int) string {
    return "fizz"
}
