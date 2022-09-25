package main

import (
	"fmt"
)

type Barang struct{
    Nama string
    Harga int
}

func main() {
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
            Nama: "Pakan Ayam",
            Harga: 25000,
        },
    }

    sortData(data)
    // fmt.Println(data)
    findMax(100000, data)
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

// find maximum total value of items that can be put in a knapsack of capacity W
func findMax(totalHarga int, data []Barang){
    var selectedItems []Barang
    var total int
    for i:=0; i<len(data); i++{
        if totalHarga - data[i].Harga >= 0 {
            total += data[i].Harga
            totalHarga -= data[i].Harga
            selectedItems = append(selectedItems, data[i])
        }
    }
    fmt.Println(total)
    fmt.Println(selectedItems)
}