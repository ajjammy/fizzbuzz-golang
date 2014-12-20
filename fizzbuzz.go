package fizzbuzz

var rules = []Rule{&FizzBuzzRule{}, &BuzzRule{}, &FizzRule{}, &OtherRule{}}

func Say(arabic int) string {
    var answer string

    for _, rule := range rules {
        if rule.IsHandle(arabic){
            answer = rule.GetAnswer(arabic)
            break
        }
    }

    return answer
}
