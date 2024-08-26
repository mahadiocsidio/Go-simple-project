package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var todolist [] string
const filename ="todo_list.txt"

func main() {
	loadTodos()
	for{
		clearScreen()

		fmt.Println("\n Todo List")
		fmt.Println("1. Lihat daftar tugas")
		fmt.Println("2. Tambah tugas")
		fmt.Println("3. Hapus tugas")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			viewTodos()
			waitForEnter()
		case 2:
			addTodo()
			waitForEnter()
		case 3:
			deleteTodo()
			waitForEnter()
		case 4:
			fmt.Println("Keluar dari aplikasi.")
			saveTodos()
			os.Exit(0)
		default:
		fmt.Println("Pilihan tidak valid.")
		}
	}
}

func viewTodos(){
	if len(todolist) == 0{
		fmt.Println("Tidak ada tugas yang harus dikerjakan.")
		return
	}
	fmt.Println("\n Daftar Tugas:")
	for i, todo := range todolist{
		fmt.Printf("%d. %s\n",i+1, todo)
	}
}

func addTodo(){
	var newTodo string
	fmt.Print("Masukkan tugas baru: ")
	fmt.Scanln(&newTodo)
	todolist = append(todolist, strings.TrimSpace(newTodo))
	fmt.Println("Tugas berhasil ditambahkan.")
	saveTodos()	
}

func deleteTodo(){
	var number int
	viewTodos()
	fmt.Print("Masukkan nomor tugas yang akan dihapus: ")
	fmt.Scanln(&number)
	if number >= 1 && number <= len(todolist){
		todolist = append(todolist[:number-1], todolist[number:]...)
		fmt.Println("Tugas berhasil dihapus.")
		return
	}
	fmt.Println("Nomor tugas tidak valid")
	saveTodos()
}

func saveTodos(){
	file, err := os.Create(filename)
	if err != nil{
		fmt.Println("gagal Menyimpan daftar tugas", err)
		return
	}
	defer file.Close()

	for _, todo := range todolist{
		file.WriteString(todo + "\n")
	}
}

func loadTodos(){
	file, err := os.Open(filename)
	if err != nil{
		if os.IsNotExist(err){
			return
		}
		fmt.Println("gagal membuka file daftar tugas", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		todolist = append(todolist, scanner.Text())
	}
	if err := scanner.Err(); err !=nil{
		fmt.Println("gagal membaca file daftar tugas", err)
	}
}

func clearScreen(){
	var cmd *exec.Cmd
	if runtime.GOOS == "windows"{
		cmd = exec.Command("cmd","/c", "cls")
	} else{
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func waitForEnter() {
	fmt.Println("\nTekan Enter untuk kembali ke menu")
	fmt.Scanln() // Menunggu hingga pengguna menekan Enter
}