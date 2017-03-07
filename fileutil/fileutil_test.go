package fileutil

import (
	"fmt"
	"testing"
)

func TestGetName(t *testing.T) {
	testCase := []struct {
		name string
		want string
	}{
		{"Test.txt", "Test"},
		{"test", "test"},
		{"Test A.txt", "Test A"},
		{"Test_A.txt", "Test_A"},
		{"Test.A.txt", "Test.A"},
		{"Test_A.B.txt", "Test_A.B"},
		{"Test_A.B..txt", "Test_A.B."},
		{"Test A..txt", "Test A."},
		{"Test A...txt", "Test A.."},
	}

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			n := GetName(tc.name)
			if n != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, n)
			}
		})
	}
}

func TestGetExtension(t *testing.T) {
	testCase := []struct {
		name string
		want string
	}{
		{"Test.txt", "txt"},
		{"test", ""},
		{"Test A.txt", "txt"},
		{"Test_A.txt", "txt"},
		{"Test.A.txt", "txt"},
		{"Test_A.B.txt", "txt"},
		{"Test_A.B..txt", "txt"},
		{"Test A..txt", "txt"},
		{"Test A...txt", "txt"},
	}

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			e := GetExtension(tc.name)
			if e != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, e)
			}
		})
	}
}

func TestDownload(t *testing.T) {
	TestCase := []struct {
		url  string
		want string
		err  bool
	}{
		{"https://golang.org/doc/gopher/frontpage.png", "frontpage.png", false},
		{"https://golang.org/doc/gopher/", "", true},
		{"", "", true},
		{"https://golang.org/", "", true},
		{"https://golang.org", "golang.org", true},
	}

	for _, tc := range TestCase {
		t.Run(fmt.Sprintf("%s", tc.url), func(t *testing.T) {
			f, err := Download(tc.url, ".")

			if err != nil {
				if tc.err == false {
					t.Errorf("Error : Expected %t, got %s", tc.err, err)
				}
			}

			if err == nil {
				if f.Name() != tc.want {
					t.Errorf("Expected %s, got %s", tc.want, f.Name())
				}
			}
		})
	}
}
