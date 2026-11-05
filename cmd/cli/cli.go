package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library-app/library"
	"library-app/storage"
)

func ExecuteUI(centralCatalog *library.Catalog, storagePath string) {
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n Выберите действие:")
		fmt.Println("1. Регистрация новой книги")
		fmt.Println("2. Регистрация нового читателя")
		fmt.Println("3. Оформить выдачу на руки")
		fmt.Println("4. Принять возврат книги")
		fmt.Println("5. Вывести весь книжный фонд")
		fmt.Println("6. Вывести базу пользователей")
		fmt.Println("0. Завершить сессию")
		fmt.Print("Введите номер операции: ")

		if !in.Scan() {
			break
		}
		
		input := strings.TrimSpace(in.Text())
		action, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите корректное число.")
			continue
		}

		if action == 0 {
			fmt.Println("Синхронизация локальной базы данных...")
			if err := storage.WriteData(storagePath, centralCatalog); err != nil {
				fmt.Printf("Критический сбой сохранения: %v\n", err)
			}
			return
		}

		switch action {
		case 1:
			fmt.Print("Укажите название: ")
			in.Scan()
			t := strings.TrimSpace(in.Text())
			
			fmt.Print("Укажите автора: ")
			in.Scan()
			a := strings.TrimSpace(in.Text())
			
			fmt.Print("Укажите год издания: ")
			in.Scan()
			yStr := strings.TrimSpace(in.Text())
			
			y, err := strconv.Atoi(yStr)
			if err != nil {
				fmt.Println("Ошибка формата года.")
				continue
			}
			
			if _, err := centralCatalog.RegisterBook(t, a, y); err != nil {
				fmt.Println("Действие отклонено:", err)
			} else {
				fmt.Println("Успешно: карточка книги создана.")
			}

		case 2:
			fmt.Print("Имя гражданина Российской Федирации: ")
			in.Scan()
			f := strings.TrimSpace(in.Text())
			
			fmt.Print("Фамилия гражданина Российской Федирации: ")
			in.Scan()
			l := strings.TrimSpace(in.Text())
			
			if _, err := centralCatalog.RegisterClient(f, l); err != nil {
				fmt.Println("Действие отказано:", err)
			} else {
				fmt.Println("Пользователь зарегистрирован.")
			}

		case 3:
			fmt.Print("Идентификатор (ID) издания: ")
			in.Scan()
			bID, _ := strconv.Atoi(strings.TrimSpace(in.Text()))
			
			fmt.Print("Идентификатор (ID) читателя: ")
			in.Scan()
			rID, _ := strconv.Atoi(strings.TrimSpace(in.Text()))
			
			if err := centralCatalog.CheckoutItem(bID, rID); err != nil {
				fmt.Println("Отказ в выдаче:", err)
			} else {
				fmt.Println("Статус изменен: книга выдана.")
			}

		case 4:
			fmt.Print("Идентификатор (ID) возвращаемой книги: ")
			in.Scan()
			bID, _ := strconv.Atoi(strings.TrimSpace(in.Text()))
			
			if err := centralCatalog.ReturnItem(bID); err != nil {
				fmt.Println("Отказано:", err)
			} else {
				fmt.Println("Статус изменен: книга снова на полке.")
			}

		case 5:
			fmt.Println("\nДоступне книги")
			allBooks := centralCatalog.ListItems()
			if len(allBooks) == 0 {
				fmt.Println("В каталоге пока пусто.")
			}
			for _, item := range allBooks {
				fmt.Println(item)
			}

		case 6:
			fmt.Println("\nДобавлен новый пользователи")
			allUsers := centralCatalog.ListMembers()
			if len(allUsers) == 0 {
				fmt.Println("Нет зарегистрированных лиц.")
			}
			for _, user := range allUsers {
				fmt.Println(user)
			}

		default:
			fmt.Println("Внимание: операция не поддерживается.")
		}
	}
}
