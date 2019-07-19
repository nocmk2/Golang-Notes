Golang-Notes
==============
 <sup> [The Go Blog](https://blog.golang.org/) 
  | [Go Documentation](https://golang.org/doc/) | [Go wiki](https://github.com/golang/go/wiki) | [The Go Programming Language (book)](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440)
</sup>

![SheetMusic](assets/1330-7981-piano.jpg)


Keywords
-------

### break
``` golang
func showHand() []string {
	gotcards := make([]string, 0)
	for _, i := range "♥♠♣♦" {
		for _, j := range "AKQJ098765432" {
			gotcards = append(gotcards, string(i)+string(j))
			if j == '8' { // I wanna play show-hand
				break // break the inner loop
			}
		}
	}
	// fmt.Println(gotcards)
	return gotcards // finaly i got 28 show-hand cards
}

func topCards(n int) []string {
	gotcards := make([]string, 0)
OutterLoop:
	for _, i := range "♥♠♣♦" {
		for cnt, j := range "AKQJ098765432" {
			gotcards = append(gotcards, string(i)+string(j))
			if cnt == n-1 { // I wanna play show-hand
				break OutterLoop // break the Outter loop using lable
			}
		}
	}
	// fmt.Println(gotcards)
	return gotcards // finaly i got 28 show-hand cards
}

func playPoker(t string, n int) []string {
	if t == "showHand" {
		return showHand()
	} else if t == "topCards" {
		return topCards(n)
	} else {
		panic("wrong poker game type:  showHand or topCards")
	}
}
```

``` golang
func gBoom(s rune) []string {
	boom := make([]string, 0)
	for _, i := range "♥♠♣♦" {
	second:
		for _, j := range "AKQJ098765432" {
			fmt.Println(string(i) + string(j))
			switch j {
			case s:
				boom = append(boom, string(i)+string(j))
				// single break without label only break the switch  not the second for loop
				// break only break out to the innermost "for" "switch" "select"
				break second
			}
		}
	}
	return boom
}
```

### case
``` golang
case 1, 3, 7, 7:  // compile error  two 7 here
```

``` golang
case 1, 3, seven(), 7: f1()


func seven() int {
	return 7
}
```

``` golang
case 1>2: f1()
```

- case constant string  compare to if
- https://github.com/golang/go/issues/10000

``` golang

```

### chan
```
```

### const
```
```

### continue
```
```

### default
```
```

### defer
```
```

### else
```
```

### fallthrough
```
```

### for
```
```

### func
```
```

### go
```
```

### goto
```
```

### if
```
```

### import
```
```

### interface
```
```

### map
```
```

### package
```
```

### range
```
```

### return
```
```

### select
```
```

### struct
```
```

### switch
```
```

### type
```
```

### var
```
```

Constants
---------

### true false
```
```

### iota
```
```

### nil
```
```

Types
-----

### int int8 int16 int32 int64
```
```

### uint uint8 uint16 uint32 uint64 uintptr
```
```

### float32 float64 complex128 complex64
```
```

### bool
```
```

### byte rune
```
```

### string
```
```

### error
```
```



Functions
---------

### make
```
```

### len
```
```

### cap
```
```

### new
```
```

### append
```
```

### copy
```
```

### close
```
```

### delete
```
```

### complex
```
```

### real
```
```

### imag
```
```

### panic
```
```

### recover
```
```

Others
------

### ... (ellipsis)
```
```

