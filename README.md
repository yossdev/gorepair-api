# gorepair-rest-api

Final backend project alterra academy kampus merdeka using clean architecture.

## MVP

- [x] Membuat akun pengguna

- [x] Menambah atau mengubah informasi pengguna

- [ ] Membuat akun pengusaha bengkel

- [ ] Menambah atau mengubah informasi bengkel (lokasi bengkel, jenis layanan, biaya layanan, jam operasional)

- [ ] Mendapatkan list bengkel (berdasarkan IP geolocation) dan informasi bengkel yang dipilih pengguna

- [ ] Membuat transaksi order service on-site atau penjemputan kendaraan

- [ ] Melihat status order (placed, on process, ready to take)

- [ ] Membatalkan atau menghapus transaksi

- [ ] Rating ke pengguna atau ke bengkel

## API Server technology stack is

- Server code: **go1.17.1**
- REST Server: [**fiber v2**](https://docs.gofiber.io/)
- Database: **MySQL**, **MongoDB**
- ORM: [**gorm v2**](https://gorm.io/docs/)

## Public API used

- [Free IP Geolocation API](https://freegeoip.app/)

## Other technology

- jwt
- bcrypt
- validator v10
- [logrus](https://pkg.go.dev/github.com/sirupsen/logrus@v1.8.1#section-readme) for logger
- [rotateFilehook](https://pkg.go.dev/github.com/snowzach/rotatefilehook@v0.0.0-20180327172521-2f64f265f58c#section-readme)
- [scribbleDB](https://github.com/nanobox-io/golang-scribble) for local_db (cache jwt)
- [viper](https://github.com/spf13/viper)
- [docker](https://www.docker.com/)
- [mockery](https://github.com/vektra/mockery) - Mocking framework
- [testify](https://github.com/stretchr/testify) - Testing toolkit
