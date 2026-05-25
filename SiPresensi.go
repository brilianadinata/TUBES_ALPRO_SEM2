package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
    "sync"
)

type Mahasiswa struct {
    NIM   string
    Nama  string
    Kelas string
}

type Storage struct {
    mu        sync.Mutex
    mahasiswa map[string]Mahasiswa
}

func NewStorage() *Storage {
    return &Storage{
        mahasiswa: make(map[string]Mahasiswa),
    }
}

func (s *Storage) TambahMahasiswa(mhs Mahasiswa) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.mahasiswa[mhs.NIM]; exists {
        return errors.New("⚠️  Mahasiswa dengan NIM ini sudah ada")
    }

    s.mahasiswa[mhs.NIM] = mhs
    return nil
}

func (s *Storage) TampilkanMahasiswa() {
    s.mu.Lock()
    defer s.mu.Unlock()

    if len(s.mahasiswa) == 0 {
        fmt.Println("\n⚠️  Data mahasiswa masih kosong.")
        return
    }

    fmt.Println("\n╔════════════════════════════════════════════════════════╗")
    fmt.Printf("║ %-10s │ %-20s │ %-10s ║\n", "NIM", "NAMA", "KELAS")
    fmt.Println("╠══════════╪══════════════════════╪═══════════╣")

    for _, mhs := range s.mahasiswa {
        fmt.Printf("║ %-10s │ %-20s │ %-10s ║\n", mhs.NIM, mhs.Nama, mhs.Kelas)
    }

    fmt.Println("╚════════════════════════════════════════════════════════╝")
}

func (s *Storage) CariMahasiswa(nim string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    mhs, exists := s.mahasiswa[nim]

    if !exists {
        fmt.Println("\n⚠️  Data mahasiswa tidak ditemukan.")
        return
    }

    fmt.Println("\n╔════════════════════════════════╗")
    fmt.Println("║       DATA MAHASISWA           ║")
    fmt.Println("╠════════════════════════════════╣")
    fmt.Printf("║ NIM   : %-22s ║\n", mhs.NIM)
    fmt.Printf("║ Nama  : %-22s ║\n", mhs.Nama)
    fmt.Printf("║ Kelas : %-22s ║\n", mhs.Kelas)
    fmt.Println("╚════════════════════════════════╝")
}

func (s *Storage) HapusMahasiswa(nim string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.mahasiswa[nim]; !exists {
        return errors.New("⚠️  Data tidak ditemukan")
    }

    delete(s.mahasiswa, nim)
    return nil
}

func (s *Storage) UpdateMahasiswa(mhs Mahasiswa) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, exists := s.mahasiswa[mhs.NIM]; !exists {
        return errors.New("⚠️  Data tidak ditemukan")
    }

    s.mahasiswa[mhs.NIM] = mhs
    return nil
}

func garis() {
    fmt.Println("====================================")
}

func header() {
    fmt.Println("╔════════════════════════════════════════════════╗")
    fmt.Println("║           🎓 SISTEM DATA MAHASISWA 🎓          ║")
    fmt.Println("║                                                ║")
    fmt.Println("╚════════════════════════════════════════════════╝")
}

func menu() {
    fmt.Println("\n╔════════════════════════════════════════╗")
    fmt.Println("║               MENU UTAMA               ║")
    fmt.Println("╠════════════════════════════════════════╣")
    fmt.Println("║ [1] Tambah Mahasiswa                   ║")
    fmt.Println("║ [2] Tampilkan Semua Mahasiswa          ║")
    fmt.Println("║ [3] Cari Mahasiswa                     ║")
    fmt.Println("║ [4] Update Mahasiswa                   ║")
    fmt.Println("║ [5] Hapus Mahasiswa                    ║")
    fmt.Println("║ [0] Keluar                             ║")
    fmt.Println("╚════════════════════════════════════════╝")
    fmt.Print("Pilih Menu : ")
}

func input(reader *bufio.Reader, text string) string {
    fmt.Print(text)
    data, _ := reader.ReadString('\n')
    return strings.TrimSpace(data)
}

func main() {
    store := NewStorage()
    reader := bufio.NewReader(os.Stdin)

    for {
        header()
        menu()

        var pilih int
        fmt.Scanln(&pilih)

        switch pilih {

        case 1:
            fmt.Println("\n--- Tambah Mahasiswa ---")

            nim := input(reader, "Masukkan NIM   : ")
            nama := input(reader, "Masukkan Nama  : ")
            kelas := input(reader, "Masukkan Kelas : ")

            err := store.TambahMahasiswa(Mahasiswa{
                NIM:   nim,
                Nama:  nama,
                Kelas: kelas,
            })

            if err != nil {
                fmt.Println("Error :", err)
            } else {
                fmt.Println("✅ Data berhasil ditambahkan!")
            }

        case 2:
            store.TampilkanMahasiswa()

        case 3:
            fmt.Println("\n--- Cari Mahasiswa ---")

            nim := input(reader, "Masukkan NIM : ")
            store.CariMahasiswa(nim)

        case 4:
            fmt.Println("\n--- Update Mahasiswa ---")

            nim := input(reader, "Masukkan NIM   : ")
            nama := input(reader, "Masukkan Nama  : ")
            kelas := input(reader, "Masukkan Kelas : ")

            err := store.UpdateMahasiswa(Mahasiswa{
                NIM:   nim,
                Nama:  nama,
                Kelas: kelas,
            })

            if err != nil {
                fmt.Println("Error :", err)
            } else {
                fmt.Println("✅ Data berhasil diupdate!")
            }

        case 5:
            fmt.Println("\n--- Hapus Mahasiswa ---")

            nim := input(reader, "Masukkan NIM : ")

            err := store.HapusMahasiswa(nim)

            if err != nil {
                fmt.Println("Error :", err)
            } else {
                fmt.Println("✅ Data berhasil dihapus!")
            }

        case 0:
            fmt.Println("\nTerima kasih dan semoga harimu menyenangkan.")
            fmt.Println("Sampai jumpa di lain waktu!")
            return

        default:
            fmt.Println("\n⚠️  Menu tidak tersedia.")
        }

        fmt.Println("\nTekan ENTER untuk lanjut...")
        reader.ReadString('\n')
    }
}