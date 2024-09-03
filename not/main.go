package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var notes []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Not Alma Uygulaması")
		fmt.Println("1. Not ekle")
		fmt.Println("2. Notları listele")
		fmt.Println("3. Notları dosyaya kaydet")
		fmt.Println("4. Çıkış")
		fmt.Print("Seçiminizi yapın: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addNote()
		case "2":
			listNotes()
		case "3":
			saveNotesToFile()
		case "4":
			fmt.Println("Çıkılıyor...")
			return
		default:
			fmt.Println("Geçersiz seçenek. Lütfen tekrar deneyin.")
		}
	}
}

func addNote() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Notunuzu girin: ")
	scanner.Scan()
	note := scanner.Text()
	notes = append(notes, note)
	fmt.Println("Not eklendi.")
}

func listNotes() {
	fmt.Print("Dosyaların bulunduğu dizini girin (örnek: /home/kullanici/notlar): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	directory := scanner.Text()

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Dizin okunurken bir hata oluştu:", err)
		return
	}

	fmt.Println("Dizindeki dosyalar:")
	for i, file := range files {
		if !file.IsDir() {
			fmt.Printf("%d: %s\n", i+1, file.Name())
		}
	}

	fmt.Print("Görmek istediğiniz dosyanın numarasını girin: ")
	scanner.Scan()
	choice := scanner.Text()

	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(files) {
		fmt.Println("Geçersiz seçim.")
		return
	}

	selectedFile := files[index-1].Name()
	filePath := filepath.Join(directory, selectedFile)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Dosya okunurken bir hata oluştu:", err)
		return
	}

	fmt.Printf("Dosya içeriği (%s):\n%s\n", selectedFile, string(data))
}

func saveNotesToFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Kaydedilecek dizini girin (örnek: /home/kullanici/notlar): ")
	scanner.Scan()
	directory := scanner.Text()

	fmt.Print("Dosya adını girin (örnek: notlar.txt): ")
	scanner.Scan()
	fileName := scanner.Text()

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Println("Dizin oluşturulurken bir hata oluştu:", err)
		return
	}

	filePath := filepath.Join(directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Dosya oluşturulurken bir hata oluştu:", err)
		return
	}
	defer file.Close()

	for _, note := range notes {
		_, err := file.WriteString(note + "\n")
		if err != nil {
			fmt.Println("Notlar dosyaya yazılırken bir hata oluştu:", err)
			return
		}
	}
	fmt.Println("Notlar dosyaya kaydedildi:", filePath)
}
