
# Course API

REST API yang dibuat menggunakan Golang dengan GIN Framework dan database mysql , dan juga menggunakan JWT untuk keperluan autentikasi


## Installation

Install Course API 

1. Clone repository ini dengan menggunakan perintah.

```bash
https://github.com/alfaa19/gin-restAPI-course.git 
```
    
2. Pastikan anda sudah menginstall Go dan database MySQL , apabila belum anda dapat mendownload Go dan MySQL di link di bawah ini.
    ```bash
    https://go.dev/dl/  #Go 
    https://dev.mysql.com/downloads #MySQL
    ```

3. Pastikan anda sudah memilikin akun cloudinary, apabila belum anda dapat menuju link berikut untuk mendaftar.
```bash
https://console.cloudinary.com/
```
4. Setelah anda memiliki akun cloudinary, copy API Enviroment Variable yang ada di dashboard kemudian simpan kedalam file .env.example yang ada pada file project

    ![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)

5. Selanjutnya anda dapat membuka file project dan mengatur configurasi yang ada pada file .env.example setelah mengatur konfigurasi jangan lupa rubah nama file .env.example menjadi .env
![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)

## Environment Variables

Untuk menjalankan project ini , anda perlu menambahkan beberapa Environment Variables ke dalam file .env

`MYSQL_USER`
`MYSQL_PASSWORD`
`MYSQL_DBNAME`
`MYSQL_HOST`
`MYSQL_PORT`

`CLOUDINARY_URL`  


## Run Locally


pergi ke project directory

```bash
  cd my-project
```

Start  server

```bash
  go run main.go
```


## API Reference

#### Get all courses

```http
  GET /api/v1/courses
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. Your JWT Token |



#### Get item

```http
  GET /api/course/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of course to fetch |
| `token` | `string` | **Required**. Your JWT Token |


#### Get total courses

```http
  GET /api/v1/courses/total-courses
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. Your JWT Token |

#### Get total free courses

```http
  GET /api/v1/courses/total-free-courses
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. Your JWT Token |

#### Post new course

```http
  POST /api/v1/course
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token` | `string` | **Required**. Your JWT Token |
| `title` | `string` | **Required**. Course Title |
| `desc` | `string` | **Not Required**. Course Description |
| `price` | `float64` | **Required**. Course Price |
| `banner` | `image` | **Required**. Course banner |
| `category_id` | `int` | **Required**. Course category id |


#### Detail

Untuk Detail Dokumentasi dapat di lihat pada postman documentation berikut :
```bash
https://documenter.getpostman.com/view/19929865/2s93kz6QnF
```

