package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Jenn"},
			result: "Hello, Jenn!",
		},
		{
			items: []string{"Jenn", "Nyah"},
			result: "Hello, Jenn, Nyah!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}