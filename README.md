## Register Login System With Go MongoDB

---

This for Issues "docker with mongo #3" which where the client ask for made project that using:

- [x] Docker as Container
- [x] Using MongoDB
- [x] Using RESTFUL APIs

But currently, i didn't use **REDIS** because i didn't have time to learn it yet. And also i didn't made the Frontend for better user experience

---

## How to use the App?

firstly you run the docker with:

```bash
docker compose up -d
```

then you go to **Postman**/**Thunder Client**/**Insomnia** depends what you like.

Then how to register?

- **Register**

Type the following example request body in the **Postman :**

```json
{
    "email": "your@email.com"
    "username": "yourusername"
    "password": "yourpassword"
}
```

And if it **Success** the response should be:

```json
{
  "status": "success"
}
```

Then how to Login?

- **Login**

Type the following example request body in the **Postman :**

```json
{
    "email": "your@email.com"
    "password": "yourpassword"
}
```

And if it **Success** the response should be:

```json
{
    "message": "Login successful!"
}
```

<br>

---

## Sistem Registrasi Dengan Go dan MongoDB

---

Ini untuk Isu "docker with mongo #3" dimana klien meminta untuk membuat project menggunakan:

- [x] Docker sebagai Container
- [x] Menggunakan MongoDB
- [x] Menggunakan API RESTFUL

Tapi sekarang, saya tidak menggunakan **REDIS** karena tidak mempunyai waktu untuk mempelajarinya. Dan juga saya tidak membuat Frontend untuk pengalaman pengguna yang lebih baik

---

## Cara Menggunakan Aplikasi

pertama kamu jalankan docker dengan perintah:

```bash
docker compose up -d
```

lalu kamu pergi ke **Postman**/**Thunder Client**/**Insomnia** sesuai yang kamu suka.

Bagaimana cara Registrasi?

- **Registrasi**

Ketik contoh request body berikut ke Postman:

```json
{
    "email": "your@email.com"
    "username: "yourusername"
    "password": "yourpassword"
}
```

Jika berhasil maka respon nya akan seperti ini:

```json
{
  "status": "success"
}
```

Bagaimana cara Login?

- **Login**

Ketik contoh request body berikut ke Postman:

```json
{
    "email": "your@email.com"
    "password": "yourpassword"
}
```

Jika berhasil maka respon nya akan seperti ini:

```json
{
    "message": "Login successful!"
}
```
