package main

import "testing"

func TestGreeting(t *testing.T) {

	emptyResult := greeting("") // should return "empty :("

	if emptyResult != "empty :("{
		t.Errorf("greeting(\"\") failed, expect %v, got %v", "empty :(", emptyResult)
	}else {
		t.Logf("greeting(\"\") success, expect %v, got %v", "empty :(", emptyResult)
	}

	result := greeting("Code.education Rocks!")

	if result != "Code.education Rocks!" {
		t.Errorf("greeting(\"Code.education Rocks!\") failed, expect %v, got %v", "Code.education Rocks!", result)
	}else {
		t.Logf("greeting(\"Code.education Rocks!\") success, expect %v, got %v", "empty :(", emptyResult)
	}
}
