package fizzbuzz

type FizzBuzzRule struct{}

    func (r *FizzBuzzRule) IsHandle(arabic int) bool {
        return arabic % 3 == 0 && arabic % 5 == 0
    }

    func (r *FizzBuzzRule) GetAnswer(arabic int) string {
        return "fizzbuzz"
    }
