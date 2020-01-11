package goal_test

import (
	"fmt"

	"github.com/sweetycode/goal"
)

func ExampleStringIn() {
	gcLang := []string{"golang", "java", "python"}
	b1 := goal.StringIn("golang", gcLang)
	b2 := goal.StringIn("cpp", gcLang)
	fmt.Println(b1, b2)
	//Output:
	// true false
}

func ExamplePrefixed() {
	s1 := goal.Prefixed("/bin/bash", "/usr")
	s2 := goal.Prefixed("/usr/bin/bash", "/usr")
	fmt.Println(s1)
	fmt.Println(s2)
	//Output:
	// /usr/bin/bash
	// /usr/bin/bash
}

func ExampleSuffixed() {
	s1 := goal.Suffixed("/bin/bash", " -x")
	s2 := goal.Suffixed("/bin/bash -x", " -x")
	fmt.Println(s1)
	fmt.Println(s2)
	//Output:
	// /bin/bash -x
	// /bin/bash -x
}

func ExampleSuffixedNewline() {
	s1 := goal.SuffixedNewline("hello world")
	s2 := goal.SuffixedNewline("hello world\n")
	fmt.Print(s1)
	fmt.Print(s2)
	//Output:
	// hello world
	// hello world
}

func ExampleSplit2() {
	var s1, s2 string
	s1, s2 = goal.Split2("www.pystarter.com", ".")
	fmt.Printf("%q, %q\n", s1, s2)
	s1, s2 = goal.Split2("www.pystarter.com", "/")
	fmt.Printf("%q, %q\n", s1, s2)
	//Output:
	// "www", "pystarter.com"
	// "www.pystarter.com", ""
}

func ExampleRSplit2() {
	var s1, s2 string
	s1, s2 = goal.RSplit2("www.pystarter.com", ".")
	fmt.Printf("%q, %q\n", s1, s2)
	s1, s2 = goal.RSplit2("www.pystarter.com", "/")
	fmt.Printf("%q, %q\n", s1, s2)
	//Output:
	// "www.pystarter", "com"
	// "www.pystarter.com", ""
}

func ExampleAtoi() {
	n1 := goal.Atoi("99")
	n2 := goal.Atoi("hello")
	n3 := goal.Atoi("99 hello")
	n4 := goal.Atoi("99.9")
	fmt.Println(n1, n2, n3, n4)
	//Output:
	// 99 0 0 0
}

func ExampleItoa() {
	s := goal.Itoa(99)
	fmt.Println(s)
	//Output:
	// 99
}

func ExampleJSONEncode() {
	data := goal.JSONEncode(map[string]interface{}{
		"errcode": 0,
	})
	fmt.Println(string(data))
	//Output:
	//{"errcode":0}
}

func ExampleJSONEncodeString() {
	s := goal.JSONEncodeString(map[string]interface{}{
		"errcode": 0,
	})
	fmt.Println(s)
	//Output:
	//{"errcode":0}
}
