package main

import (
	"context"
	"database/sql"
	"fmt"
)

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first: ", val)
	}(a)

	a = 20
	defer func(val int) {
		fmt.Println("second: ", val)
	}(a)

	a = 30

	fmt.Println("exiting: ", a)
	return a

}

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// use tx to do more database inserts here
	return nil
}
func main() {
	// fmt.Println(addTo(3))
	// fmt.Println(addTo(3, 2))
	// fmt.Println(addTo(3, 2, 4, 6, 8))

	// f := func(j int) {
	// 	fmt.Println("printing", j, "from inside of an anonymous function")
	// }

	// for i := 0; i < 5; i++ {
	// 	f(i)
	// }

	// twoBase := makeMult(2)
	// threeBase := makeMult(3)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(twoBase(i), threeBase(i))
	// }

	// defer
	// if len(os.Args) < 2 {
	// 	log.Fatal("no file specified")
	// }

	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	deferExample()
}
