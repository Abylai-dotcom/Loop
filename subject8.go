package main

import "fmt"

func main() {
	// //sub 1
	// for i := 1; i <= 20; i++ {
	// 	fmt.Println(i)

	// }
	// // sub 2
	// sum := 0

	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// 	fmt.Println(sum)

	// }

	// sub 3

	// var number int

	// for i := 1; i <= 10; i++ {
	// 	for j := 1; j <= 10; j++ {
	// 		res := i * j
	// 		number = res
	// 		fmt.Print(res, "\t")
	// 	}
	// 	fmt.Println(number)
	// }

	// sub 4

	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		i += 2
		fmt.Println(i)
	}

	// sub 5

	number := 12345
	counter := 0
	for number > 0 {
		counter++
		number /= 10
	}
	fmt.Println(counter)

	// sub 6

	text := "Developer"

	for _, v := range text {
		fmt.Printf("%c\n", v)
	}

	// sub 7
	balance := 3000
	for {
		var choice int
		fmt.Print("Введите число от 0 до 3: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("Текущий баланс:", balance)
		case 2:
			balance += 500
			fmt.Println("Баланс после пополнения:", balance)
		case 3:
			balance -= 200
			fmt.Println("Баланс после снятия:", balance)
		case 0:
			fmt.Println("Выход из программы")
			break

		default:
			fmt.Println("Неверный выбор, попробуйте снова")
			continue
		}
		if choice == 0 {
			break
		}
	}

}
