package main

import "testing"

func TestGreeting(t *testing.T) {

	emptyResult := test_greeting("") // should return "empty :("

	if emptyResult != "empty :("{
		t.Errorf("test_greeting(\"\") failed, expect %v, got %v", "empty :(", emptyResult)
	}else {
		t.Logf("test_greeting(\"\") success, expect %v, got %v", "empty :(", emptyResult)
	}

	result := test_greeting("Code.education Rocks!")

	if result != "Code.education Rocks!" {
		t.Errorf("test_greeting(\"Code.education Rocks!\") failed, expect %v, got %v", "Code.education Rocks!", result)
	}else {
		t.Logf("test_greeting(\"Code.education Rocks!\") success, expect %v, got %v", "empty :(", emptyResult)
	}

}

func test_greeting(text string) string {
	if (len(text) == 0 ) {
		return "empty :("
	} else {
		return text
	}
}
