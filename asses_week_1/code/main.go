package main

import (
	"fmt"
)

type Barang struct{
	Nama string
	Harga int
}

func main() {
	var barang map[string]int
	barang = map[string]int{}

	barang["Benih Lele"] = 50000
	barang["Pakan lele cap menara"] = 25000
	barang["Probiotik A"] = 75000
	barang["Probiotik Nila B"] = 10000
	barang["Pakan Nila"] = 20000
	barang["Benih Nil"] = 20000
	barang["Cupang"] = 5000
	barang["Benih Nila"] = 30000
	barang["Benih Cupang"] = 10000
	barang["Probiotik B"] = 10000

	data := []Barang{
		Barang{
			Nama: "Benih lele",
			Harga: 50000,
		},
		Barang{
			Nama: "Pakan lele cap menara",
			Harga: 25000,
		},
		Barang{
			Nama: "Probiotik A",
			Harga: 75000,
		},
		Barang{
			Nama: "Probiotik Nila B",
			Harga: 10000,
		},
		Barang{
			Nama: "Pakan Nila",
			Harga: 20000,
		},
		Barang{
			Nama: "Benih Nil",
			Harga: 20000,
		},
		Barang{
			Nama: "Cupang",
			Harga: 5000,
		},
	}

	
	// fmt.Println(data)
	fmt.Println("Barang dengan Harga Rp10.000 : ")
	sepuluribu(barang)
	fmt.Println("------------------------------------------------")
	fmt.Println("Barang dengan harga termurah dan tertinggi : ")
	maxAndMinItem(barang)
	fmt.Println("------------------------------------------------")
	// sortData(data)
	findMaxItem(100000, data)
}

func sortData(data []Barang){
	for i:=1; i< len(data); i++{
		j := i
		for j > 0 {
            if data[j-1].Harga> data[j].Harga {
                data[j-1], data[j] = data[j], data[j-1]
            }
            j = j - 1
        } 
	}
}

func findMaxItem(totalHarga int, data []Barang){
    var selectedItems []Barang
    var total int
    for i:=0; i<len(data); i++{
        if totalHarga - data[i].Harga >= 0 {
            total += data[i].Harga
            totalHarga -= data[i].Harga
            selectedItems = append(selectedItems, data[i])
        }
    }
    fmt.Printf("Total harga : %d\n", total)
	fmt.Printf("Daftar Barang : %v", selectedItems)
    
}


// no 1 c
func sepuluribu(barang map[string]int) {
	for k, v := range barang {
		if v == 10000 {
			fmt.Println(k, v)
		}
	}
	
}


func maxAndMinItem(barang map[string]int) {
	var max int
	var nama_barang string
	var min int
	var barang_murah string
	for k,v := range barang{
		if v > max {
			max = v
			nama_barang = k
		}
		
	}
	for _, v := range barang{
		min = v
		break
	}
	for k, v := range barang{
		if v <= min {
			min = v
			barang_murah = k
		}
	}
	fmt.Printf("Barang termurah : %s Rp %d\n", barang_murah, min)
	fmt.Printf("barang termahal :  %s %d\n", nama_barang, max)
}

