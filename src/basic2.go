package src

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Printf("sum = %d\n", sum)

	sum2 := 1
	// 初期化と後処理のステートメントは不要
	// while文にあたる
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Printf("sum = %d\n", sum2)

	// 無限ループ
	//for {}

	// switch: 条件が一致すればbreakする
	fmt.Print("Go runs on")
	os := runtime.GOOS
	switch os { //switch os := runtime.GOOS os この書き方もあり
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!", t.Hour())
	case t.Hour() < 17:
		fmt.Println("Good afternoon!", t.Hour())
	default:
		fmt.Println("Good evening!", t.Hour())
	}

	// defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる
	defer fmt.Println("world")
	fmt.Println("hello")

	// deferが複数ある場合はLIFOの順に実行される
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
