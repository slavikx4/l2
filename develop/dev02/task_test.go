package main

import "testing"

func TestUnpackString(t *testing.T) {

	var tests = []struct {
		numTest int
		input   string
		want    string
		textErr string
	}{
		{1, "a4bc2d5e", "aaaabccddddde", ""},
		{2, "abcd", "abcd", ""},
		{3, "45", "", "некорректная строка"},
		{4, "", "", ""},
	}

	for _, test := range tests {

		res, err := UnpackString(test.input)

		if err != nil {
			if err.Error() != test.textErr {
				t.Errorf("тест: %d\n ожидалось: %s\n результат: %s", test.numTest, test.textErr, err.Error())
			}
		}

		if res != test.want {
			t.Errorf("ожидалось: %s\n результат: %s", test.want, res)
		}
	}
}
