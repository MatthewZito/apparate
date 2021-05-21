package testing

import (
	"os"
	"testing"

	"github.com/MatthewZito/apparate/internal"
)

const (
	KEY          = "key1"
	VAL          = "val1"
	TEST_DB_PATH = "test_db"
)

var RunTest = CreateForEach(setUp, tearDown)

func setUp() {
	os.Remove(TEST_DB_PATH)
}

func tearDown() {
	os.Remove(TEST_DB_PATH)
}

func TestMain(t *testing.T) {
	RunTest(func() {
		p := internal.Portal{
			Alias: KEY,
			Path:  VAL,
		}

		db, err := internal.Open(TEST_DB_PATH)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		/* PUT */
		if err = db.Put(&p); err != nil {
			t.Fatal(err)
		}

		/* GET */
		p2 := internal.Portal{
			Alias: "key1",
		}

		if err := db.Get(&p2); err != nil {
			t.Fatal(err)
		} else if p2.Path != VAL {
			t.Fatalf("expected \"%s\", got \"%s\"", VAL, p2.Path)
		}

		/* GET NOT EXTANT */
		p3 := internal.Portal{
			Alias: "E_NO_EXIST",
		}

		if err := db.Get(&p3); err != internal.ErrNotFound {
			t.Fatalf("got \"%s\", expected absence", p3.Path)
		}

		/* DELETE */
		if err := db.Delete(&p); err != nil {
			t.Fatal(err)
		}

		if err := db.Delete(&p); err != internal.ErrNotFound {
			t.Fatalf("delete returned %v, expected %s", err, internal.ErrNotFound.Error())
		}
	})
}

func CreateForEach(setUp func(), tearDown func()) func(func()) {
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown()
	}
}
