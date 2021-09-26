package main

import "fmt"

type Y interface{}
type FINT func(Y) Y
type X func(X) FINT

func yComb(f func(FINT) FINT) FINT {
	return func(x X) FINT {
		return f(func(y Y) Y {
			return x(x)(y)
		})
	}(func(x X) FINT {
		return f(func(y Y) Y {
			return x(x)(y)
		})
	})
}

func factGen(fact FINT) FINT {
	return func(n Y) Y {
		num := n.(int)

		if num == 0 {
			return 1
		} else {
			return Y(num*(fact(Y(num-1))).(int))
		}
	}
}

func yMem(f func(FINT) FINT, cache map[Y]Y) FINT {
	if cache == nil {
		cache = make(map[Y]Y)
	}

	return func(arg Y) Y {
		val, exists := cache[arg]
		if exists {
			return val
		}

		answer := f(func(n Y) Y {
			return yMem(f, cache)(n)
		})(arg)

		cache[arg] = answer

		return answer
	}
}

func main() {
	fmt.Println(yComb(factGen)(18))

	fib := yMem(func(g FINT) FINT {
		return func(n Y) Y {
			num := n.(int)

			if num == 0 {
				return 0
			}

			if num == 1 {
				return 1
			}

			return g(Y(num-1)).(int) + g(Y(num-2)).(int)
		}
	}, nil)

	fmt.Println(fib(100))
}