package MapLearning

import "testing"

func TestMap(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

	t.Run("class Search",func(t *testing.T)  {
		dic := Dictionary{"test": "this is just a test"}
		got,err := dic.dicSearch("test")
		want := "this is just a test"

		if err == nil {
            t.Fatal("expected to get an error.")
        }

		if got != want {
			t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
		}
	})

		t.Run("error class Search",func(t *testing.T)  {
		dic := Dictionary{"test": "this is just a test"}
		_,err := dic.dicSearch("unknown")
		want := "this is just a test"

		if err == nil {
            t.Fatal("expected to get an error.")
        }

		if got != want {
			t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
		}
	})
}
