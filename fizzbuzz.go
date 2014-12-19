package fizzbuzz

func Say(arabic int) string {

    var answer string
    var rules []Rule

    rules = []Rule{&FizzBuzzRule{}, &BuzzRule{}, &FizzRule{}, &OtherRule{}}

    for _, rule := range rules {
        if rule.IsHandle(arabic){
            answer = rule.GetAnswer(arabic)
            break
        }
    }
    return answer
}
