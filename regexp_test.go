package main

import (
	"regexp"
	"testing"
)

var re = regexp.MustCompile(`\d+`)

func TestRe1(t *testing.T) {
	str := "http://example.com/1323024?key=value"

	ptn := `/\d+(\?.*)$`
	re, e := regexp.Compile(ptn)

	if e != nil {
		t.Fatalf("Failed to compile %s", ptn)
	}

	match := re.FindString(str)
	t.Logf("[DEBUG] matched: %s", match)

	if match != "/1323024?key=value" {
		t.Errorf("Failed to FindString('%s') by %s", str, re)
	}

	match2 := re.FindStringSubmatch(str)
	t.Logf("[DEBUG] matched: %s", match2)

	if len(match2) != 2 {
		t.Fatalf("Failed to FindStringSubmatch('%s') by %s", str, re)
	}

	exp1, exp2 := "/1323024?key=value", "?key=value"
	if match2[0] != exp1 || match2[1] != exp2 {
		t.Errorf("Expected $1 = '%s' $2 = '%s'", exp1, exp2)
	}
}

func TestRe2(t *testing.T) {
	str := "{a: 1, hoge: 20, fuga: 300, b: ''}"

	ptn := `\w+: \d+`
	re, e := regexp.Compile(ptn)

	if e != nil {
		t.Fatalf("Failed to compile %s", ptn)
	}

	match := re.FindAllString(str, -1)
	t.Logf("[DEBUG] matched: %s", match)
	if len(match) != 3 {
		t.Fatalf("Failed to FindAllStringSubmatch('%s') by ", str, re)
	}

	e1, e2, e3 := "a: 1", "hoge: 20", "fuga: 300"
	b := match[0] == "a: 1" && match[1] == "hoge: 20" && match[2] == "fuga: 300"
	if !b {
		t.Errorf("Expected $1 = %s, $2 = %s, $3 = %s but got %s %s %s", e1, e2, e3, match[0], match[1], match[2])
	}
}
