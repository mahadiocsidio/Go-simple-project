package main

import (
	"fmt"
	"os/exec"
	"os"
	"runtime"
)

func main(){
	for{
		fmt.Println("Aplikasi Kalkulator Sederhana")
		fmt.Println("Pilih Operasi")
		fmt.Println("1. Penambahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Opsi: ")
		
		var choice int
		fmt.Scanln(&choice)
		
		if choice == 5{
			fmt.Printf("Keluar dari Program")
			break
		}

		var num1,num2 float32
		fmt.Printf("Masukkan Angka Pertama: \n")
		fmt.Scanln(&num1)
		fmt.Printf("Masukkan Angka Kedua: \n")
		fmt.Scanln(&num2)

		switch choice{
		case 1:
			fmt.Printf("Hasil %.2f + %.2f = %2f\n", num1,num2, num1+num2)
		case 2:
			fmt.Printf("Hasil %.2f - %.2f = %2f\n", num1,num2, num1-num2)
		case 3:
			fmt.Printf("Hasil %.2f * %.2f = %2f\n", num1,num2, num1*num2)
		case 4:
			if num2 !=0{
				fmt.Printf("Hasil %.2f / %.2f = %2f\n", num1,num2, num1/num2)
			}else{
				fmt.Println("Tidak bisa membagi dengan 0")
			}
		default:
			fmt.Println("Pilihan tidak valid")
		}

		fmt.Println()
		clearScreen()
	}
}
func clearScreen(){
	waitForEnter()
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