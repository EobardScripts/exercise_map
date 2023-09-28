package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func commonElements(arr1, arr2 []string) string {
	mp := make(map[string]int)
	checkDublicate := make(map[string]int)
	result := "Общие элементы: "
	mark := len(result)

	for i, el := range arr1 {
		mp[el] = i
	}

	for i, el := range arr2 {
		_, ok := mp[arr2[i]]
		if ok {
			_, ch := checkDublicate[arr2[i]]
			if !ch {
				result += el + " "
				checkDublicate[arr2[i]] = i
			} else {
				continue
			}
		}
	}

	if mark == len(result) {
		result = "Общие элементы отсутствуют"
	}
	return result
}

func main() {
	sizeArr1 := 0
	sizeArr2 := 0

	fmt.Println("Введите размер первого массива: ")
	_, err := fmt.Scanln(&sizeArr1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Введите размер второго массива: ")
	_, err = fmt.Scanln(&sizeArr2)
	if err != nil {
		log.Fatalln(err)
	}

	if sizeArr1 == 0 || sizeArr2 == 0 {
		log.Fatalln("Массив не может быть нулевого размера!")
	}

	arr1 := make([]string, sizeArr1)
	arr2 := make([]string, sizeArr2)
	_ = arr2

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Заполните первый массив элементами: ")
	for i := 0; i < sizeArr1; i++ {
		arr1[i], _ = reader.ReadString('\n')
	}

	fmt.Println("Заполните второй массив элементами: ")
	for i := 0; i < sizeArr2; i++ {
		arr2[i], _ = reader.ReadString('\n')
	}

	fmt.Println(commonElements(arr1, arr2))
}
