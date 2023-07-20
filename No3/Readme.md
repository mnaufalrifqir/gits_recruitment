## Penjelasan Kompleksitas

1. Iterasi untuk memeriksa validitas input:
- Kompleksitas: O(n)
- Penjelasan: Kode ini melakukan iterasi sebanyak n kali (panjang string brackets) untuk memeriksa setiap karakter dalam string apakah merupakan karakter bracket yang valid ('{', '}', '[', ']', '(', atau ')'). Jika ditemukan karakter yang tidak valid, maka program mencetak "Invalid input" dan menghentikan iterasi. Karena iterasi ini hanya dilakukan sekali, kompleksitasnya adalah O(n), di mana n adalah panjang string brackets.

2. Fungsi balancedBracket:
- Kompleksitas: O(n)
- Penjelasan: Fungsi balancedBracket melakukan iterasi sebanyak n kali (panjang string brackets) untuk memeriksa keseimbangan bracket. Selama iterasi, fungsi ini menggunakan slice checkArr untuk menyimpan karakter bracket yang belum ditutup. Setiap kali menemukan karakter buka bracket '{', '[', atau '(', karakter tersebut akan ditambahkan ke checkArr. Dan setiap kali menemukan karakter tutup bracket '}', ']', atau ')', fungsi akan memeriksa apakah checkArr tidak kosong dan apakah karakter tutup tersebut sesuai dengan karakter buka terakhir dalam checkArr. Jika tidak, fungsi akan mengembalikan "NO". Setelah selesai iterasi, jika checkArr kosong, berarti semua bracket telah ditutup dengan benar, dan fungsi akan mengembalikan "YES". Jika tidak kosong, berarti ada bracket yang belum ditutup, dan fungsi akan mengembalikan "NO". Karena iterasi ini hanya dilakukan sekali, kompleksitasnya adalah O(n), di mana n adalah panjang string brackets.

## Kesimpulan
Keseluruhan kompleksitas kode di atas adalah O(n), di mana n adalah panjang string brackets. Ini karena kedua bagian kode hanya melakukan iterasi sebanyak panjang string brackets satu kali.