package main

import (
	"encoding/base64"
	"fmt"
)

type Book struct {
	Name       string
	Author     string
	Publishing string
	Year       int
	Taken      bool
}

func generateID(name string) int {
	module := 91
	hash := 0
	powNum := 1
	for i := 0; i < len(name); i++ {
		hash += int(name[i]-'a'+1) * powNum
		powNum *= module
	}
	return hash
}

func generateIDforMap(name string) string {
	data := []byte(name)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

type Storage interface {
	PutToStorage(Book)
	SearchByID(string) interface{}
}

type StorageMap struct {
	Memory map[string]Book
}

func (storage *StorageMap) PutToStorage(item Book) {
	itemID := generateIDforMap(item.Name)
	storage.Memory[itemID] = item
	item.Taken = false
	fmt.Println("Ждем вас снова")
}

func (storage *StorageMap) SearchByID(name string) interface{} {
	itemID := generateIDforMap(name)
	if _, exists := storage.Memory[itemID]; exists {
		fmt.Println("Книга найдена!")
		return storage.Memory[itemID]
	}
	return "Такой книги нет("
}

type StorageSlice struct {
	Memory    []Book
	MemoryMap map[int]int
}

func (storage *StorageSlice) PutToStorage(item Book) {
	itemID := generateID(item.Name)
	index := len(storage.Memory)
	storage.Memory = append(storage.Memory, item)
	storage.MemoryMap[itemID] = index
	storage.Memory[index].Taken = false
	fmt.Println("Ждем вас снова")
}

func (storage *StorageSlice) SearchByID(name string) interface{} {
	itemID := generateID(name)
	if _, exists := storage.MemoryMap[itemID]; exists {
		fmt.Println("Книга найдена!")
		return storage.Memory[storage.MemoryMap[itemID]]
	}
	return "Такой книги нет("
}

type LibraryMain interface {
	Search(name string) interface{}
	Put(item Book)
}

type Library struct {
	MainStorage Storage
}

func (lib *Library) Search(name string) interface{} {
	return lib.MainStorage.SearchByID(name)
}

func (lib *Library) Put(item Book) {
	lib.MainStorage.PutToStorage(item)
}

func main() {
	HarryPotter := Book{"Harry Potter", "Rowling", "Escmo", 2005, false}
	Hobbit := Book{"Hobbit", "Tolkin", "Escmo", 2010, false}
	HungerGames := Book{"The Hunger Games", "Kollins", "Machaon", 2015, false}
	Hamlet := Book{"Hamlet", "Shakespeare", "Something", 2020, false}
	Makbet := Book{"Makbet", "Shakespeare", "SomePublishing", 2019, false}
	var storageForMap Storage = &StorageMap{make(map[string]Book)}
	MapLibrary := &Library{storageForMap}
	MapLibrary.Put(HarryPotter)
	MapLibrary.Put(Hobbit)
	MapLibrary.Put(HungerGames)
	MapLibrary.Put(Hamlet)
	MapLibrary.Put(Makbet)
	fmt.Println(MapLibrary.Search("Harry Potter"))
	fmt.Println(MapLibrary.Search("Hobbit"))
	var storageForSlice Storage = &StorageSlice{make([]Book, 3), make(map[int]int)}
	SliceLibrary := &Library{storageForSlice}
	SliceLibrary.Put(HarryPotter)
	SliceLibrary.Put(Hobbit)
	SliceLibrary.Put(HungerGames)
	SliceLibrary.Put(Hamlet)
	SliceLibrary.Put(Makbet)
	fmt.Println(SliceLibrary.Search("Makbet"))
	fmt.Println(SliceLibrary.Search("Lalala"))
	fmt.Println(SliceLibrary.Search("Hamlet"))
}
