package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/template"
)

var bob = &struct {
	Name string
	Age  int
}{
	Name: "Bob",
	Age:  23,
}

type NullWriter struct{}

func (w *NullWriter) Write(b []byte) (int, error) { return len(b), nil }

func BenchmarkAppendStr(b *testing.B) {
	w := &NullWriter{}
	buff := []byte{}
	for i := 0; i < b.N; i++ {
		buff = buff[0:0]
		buff = append(buff, "Hi, my name is "...)
		buff = append(buff, bob.Name...)
		buff = append(buff, " and I'm "...)
		buff = append(buff, string(bob.Age)...)
		buff = append(buff, " years old."...)
		w.Write(buff)
	}
}

func BenchmarkBufferString(b *testing.B) {
	w := &NullWriter{}
	buff := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		buff.Reset()
		buff.Write([]byte("Hi, my name is "))
		buff.Write([]byte(bob.Name))
		buff.Write([]byte(" and I'm "))
		buff.Write([]byte(string(bob.Age)))
		buff.Write([]byte(" years old."))
		w.Write(buff.Bytes())
	}
}

func BenchmarkBufferString2(b *testing.B) {
	w := &NullWriter{}
	buff := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		buff.Reset()
		buff.WriteString("Hi, my name is ")
		buff.WriteString(bob.Name)
		buff.WriteString(" and I'm ")
		buff.WriteString(string(bob.Age))
		buff.WriteString(" years old.")
		w.Write(buff.Bytes())
	}
}

func BenchmarkFmtFormat(b *testing.B) {
	w := &NullWriter{}
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(w, "Hi, my name is %s and I'm %d years old.", bob.Name, bob.Age)
	}
}

func BenchmarkConcat1(b *testing.B) {
	w := &NullWriter{}
	for i := 0; i < b.N; i++ {
		fmt.Fprint(w, "Hi, my name is "+bob.Name+" and I'm "+string(bob.Age)+" years old.")
	}
}

func BenchmarkConcat2(b *testing.B) {
	w := &NullWriter{}
	for i := 0; i < b.N; i++ {
		fmt.Fprint(w, strings.Join([]string{"Hi, my name is ", bob.Name, " and I'm ", string(bob.Age), " years old."}, ""))
	}
}

func BenchmarkTemplate(b *testing.B) {
	t := template.Must(template.New("").Parse(
		`Hi, my name is {{.Name}} and I'm {{.Age}} years old.`))
	w := &NullWriter{}
	for i := 0; i < b.N; i++ {
		t.Execute(w, bob)
	}
}
