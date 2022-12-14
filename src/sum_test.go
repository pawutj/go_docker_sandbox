package go_docker_sandbox

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 2)
	want := 3
	if result != want {
		t.Errorf("Want '%d' result '%d'", want, result)
	}
}
