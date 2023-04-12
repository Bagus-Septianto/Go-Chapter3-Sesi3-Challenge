# Go-Chapter3-Sesi3-Challenge
Repositori ini adalah hasil pengembangan dari  [Go-Chapter3-Sesi2-Challenge](https://github.com/Bagus-Septianto/Go-Chapter3-Sesi2-Challenge "Bagus-Septianto/Go-Chapter3-Sesi2-Challenge"). perubahan yang sangat terlihat adalah dari struktur direktorinya, sekarang file-file yang ada sudah sesuai dengan tugasnya (sebelumnya 1 file bisa sangat "gemuk").

## Penjelasan Singkat
- `app` digunakan untuk menyimpan `controllers`, `models`, `repository`, `services`
  - `controllers` digunakan untuk menyimpan file-file yang berfungsi untuk komunikasi dengan client (REST/JSON). Route API ada disini dan akan memanggil fungsi yang ada di `services`
  - `models` digunakan untuk menyimpan `struct` yang akan dipakai.
  - `repository` digunakan untuk menyimpan file-file yang berfungsi untuk komunikasi dengan database (GORM). fungsi yang ada di folder ini akan dipanggil oleh fungsi yang ada di `services`
  - `services` digunakan untuk menyimpan file-file yang berfungsi untuk mengolah *business logic* yang didapatkan dari `repository`
- `pkg` digunakan untuk menyimpan `common`, `database`, `helpers`
  - `common` digunakan untuk menyimpan konstanta yang digunakan pada projek ini
  - `database` digunakan untuk menyimpan fungsi untuk koneksi ke database
  - `helpers` digunakan untuk menyimpan fungsi-fungsi bantuan

## Struktur Direktori
```
challenge294/
├── app/
│   ├── controllers/
│   │   ├── middlewares/
│   │   │   ├── authentication.go
│   │   │   └── authorization.go
│   │   ├── product/
│   │   │   ├── create.go
│   │   │   ├── read.go
│   │   │   ├── readall.go
│   │   │   ├── update.go
│   │   │   ├── delete.go
│   │   │   └── product.go
│   │   └── user/
│   │       ├── login.go
│   │       ├── register.go
│   │       └── user.go
│   ├── models/
│   │   ├── gormModel.go
│   │   ├── product.go
│   │   └── user.go
│   ├── repository/
│   │   ├── product/
│   │   │   ├── create.go
│   │   │   ├── read.go
│   │   │   ├── readall.go
│   │   │   ├── update.go
│   │   │   ├── delete.go
│   │   │   ├── product.go
│   │   │   └── product_mock.go
│   │   ├── user/
│   │   │   ├── login.go
│   │   │   ├── register.go
│   │   │   └── user.go
│   │   └── repository.go
│   └── services/
│       ├── product/
│       │   ├── create.go
│       │   ├── read.go
│       │   ├── read_test.go
│       │   ├── readall.go
│       │   ├── readall_test.go
│       │   ├── update.go
│       │   ├── delete.go
│       │   └── product.go
│       ├── user/
│       │   ├── login.go
│       │   ├── register.go
│       │   └── user.go
│       └── services.go
└── pkg/
    ├── common/
    │   └── const.go
    ├── database/
    │   └── db.go
    └── helpers/
        ├── bcrypt.go
        ├── headerValue.go
        └── jwt.go
```

## Command yang mungkin berguna
```ps
# run/start the app
go run main.go

# run test
go test -v ./app/services/...
```