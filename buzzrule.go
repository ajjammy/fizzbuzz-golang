package fizzbuzz

type BuzzRule struct{}

    func (r *BuzzRule) IsHandle(arabic int) bool {
        return arabic % 5 == 0
    }

    func (r *BuzzRule) GetAnswer(arabic int) string {
        return "buzz"
    }
