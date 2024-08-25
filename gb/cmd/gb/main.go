package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type MainFile struct {
	Data map[string]Object
}

type Object struct {
	name string
	value int
}

var (
	data = MainFile{}
	expect1 = CURLY_OPEN
	expect2 = SQUARE_OPEN
)

const (
	ERROR = iota - 1
	CURLY_OPEN
	CURLY_CLOSE
	QUOT
	COMMA
	SQUARE_OPEN
	SQUARE_CLOSE
	COLON
)

type Token = int

func print_token(token Token) {
	switch token {
		case ERROR: fmt.Println("error")
    	case CURLY_OPEN: fmt.Println("curly_open")
		case CURLY_CLOSE: fmt.Println("curly_close")
		case QUOT: fmt.Println("quot")
		case COMMA: fmt.Println("comma")
		case SQUARE_OPEN: fmt.Println("square_open")
		case SQUARE_CLOSE: fmt.Println("square_close")
	}
}

func char_to_token(c byte) Token {
	switch c {
		case '{': return CURLY_OPEN
		case '}': return CURLY_CLOSE
		case '[': return SQUARE_OPEN
		case ']': return SQUARE_CLOSE
		case '"': return QUOT
		case ',': return COMMA
		case ':': return COLON
		default: return ERROR
	}
}

func next_token(scan []byte, expect1 Token, expect2 Token) (Token, int) {
	for idx, s := range scan {
		t := char_to_token(s)

	  	if t == expect1 || t == expect2 {
			return t, idx
	  }
	}
	return -1, -1
}

func get_next_valid_tokens(expect Token, quot bool) (Token, Token) {
	switch expect {
		case SQUARE_OPEN: return CURLY_OPEN, SQUARE_CLOSE
		case CURLY_OPEN: return CURLY_CLOSE, QUOT
		case CURLY_CLOSE: return COMMA, SQUARE_CLOSE
		case QUOT: 
			if quot {
				return COLON, CURLY_CLOSE
			} else {
				return QUOT, QUOT
			}
	}
	return -1, -1
}

func main() {
	//cmd := ""

	//if len(os.Args) == 2 {	
		//cmd = os.Args[1]
	//}

	//fmt.Printf("%s", cmd)

	s := bufio.NewScanner(os.Stdin)

	expect1 = CURLY_OPEN
    expect2 = SQUARE_OPEN

	for s.Scan() {
		idx := 0
		b := s.Bytes()
		quot := false
		//fmt.Printf("%s\n", string(b[idx:]))

		for { 
			tok, cur := next_token(b[idx:], expect1, expect2)
			if tok == -1 {
				break
			}

			print_token(tok)
			idx += cur + 1

			valid := expect2 != -1 && tok == expect2

			if !valid {
				valid = tok == expect1
			} else {
				expect1 = expect2
			}
			
			if expect1 == QUOT {
				quot = !quot
			}

			if valid {
//				fmt.Printf("2 ")
//				print_token(expect1)
				expect1, expect2 = get_next_valid_tokens(expect1, quot)
//				fmt.Printf("3 ")
//				print_token(expect1)
			}
		}
	}
}
