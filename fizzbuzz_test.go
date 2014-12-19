package fizzbuzz

import "testing"

var m = map[int]string{
    1: "1",
    2: "2",
    3: "fizz",
    4: "4",
    5: "buzz",
    6: "fizz",
    9: "fizz",
    10: "buzz",
    12: "fizz",
    15: "fizzbuzz",
}

func TestFizzBuzz(t *testing.T){

    for arabic, expected := range m {
        result := Say(arabic)

        if result != expected {
            t.Errorf("Expected %s but got %s", expected, result)
        }
    }

}
