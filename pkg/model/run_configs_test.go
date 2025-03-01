package model

import "testing"

func TestNewFlutterRunConfig(t *testing.T) {
	test := "test"
	mode := "debug"
	flavor := "flavor"
	target := "main.dart"
	define := "test.json"

	config := FlutterRunConfig{
		Name:               test,
		Mode:               mode,
		Flavor:             flavor,
		Target:             target,
		DartDefineFromFile: define,
	}

	if config.Name != test || config.Mode != mode || config.Flavor != flavor || config.Target != target || config.DartDefineFromFile != define {
		t.Fail()
	}

	err := config.AssertConfig()
	if err != nil {
		t.Fail()
	}

	config.Mode = "bleh"

	if err := config.AssertConfig(); err == nil {
		t.Fail()
	}
}
