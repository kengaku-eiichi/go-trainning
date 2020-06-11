package animals

import "testing"

func TestElephantFeed(t *testing.T) {
	expect := "Grass"

	if result := ElephantFeed(); expect != result {
		t.Errorf("\nexpect is  : %v\nbut result : %v\n", expect, result)
	}
}
func TestMonkeyFeed(t *testing.T) {
	expect := "Banana"

	if result := MonkeyFeed(); expect != result {
		t.Errorf("\nexpect is  : %v\nbut result : %v\n", expect, result)
	}
}
func TestRabbitFeed(t *testing.T) {
	expect := "Carrot"

	if result := RabbitFeed(); expect != result {
		t.Errorf("\nexpect is  : %v\nbut result : %v\n", expect, result)
	}
}
