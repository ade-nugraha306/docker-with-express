# Express With Docker

---

Ini contoh project Express dengan Docker Container.

---

**Cara Jalankan?**
<br>

Cara menjalankan project ini dibagi menjadi dua:

- Jika ingin menjalankan project dalam Environment **Production** maka jalankan perintah:

```bash
docker compose -f docker-compose.yml -f docker-compose.prod.yml up --build
```
<br>

- Jika ingin menjalankan project dalam Environment **Development** maka jalankan perintah:

```bash
docker compose up
```

---

Dengan adanya **Docker** maka kamu tidak perlu lagi menjalan kan perintah `npm install` karna **Docker** akan **menginstall dan menjalankan semuanya untuk kamu.**

