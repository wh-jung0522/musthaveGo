//ex22.4/ex22.4.go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) // ❶ 항목이 5개인 링 생성

	n := r.Len() // ❷ 링 개수 반환

	for i := 0; i < n; i++ {
		r.Value = 'A' + i // ❸ 순회하면 모든 항목에 값 대입
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value) // ❹ 순회하며 값 출력
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value) // ➎ 순하며 값 출력
		r = r.Prev()
	}
}