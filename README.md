In this project, users can have the front-end, they should enter the day of month and month name with uppercase, like `Monday`, `March`, in the end the result will be displayed in the left bottom corner of the screen



Uy vazifa
GET buyrug'ini yuborganingizda joriy vaqtni RFC 3339 formatida qaytaradigan kichik web-server yozing (masalan /time/RFC3339 endpoint)
Vaqtni JSON sifatida qaytarish imkoniyatini qo'shing. JSON yoki matn qaytarilishini nazorat qilish uchun Qabul qilish sarlavhasidan foydalaning (standart matn uchun). JSON quyidagi tarzda tuzilishi kerak:
{
  "day_of_week": "Monday",
  "day_of_month": 10,
  "month": "April",
  "year": 2023,
  "hour": 20,
  "minute": 15,
  "second": 20
}

# 2024-05-13T15:30:00-07:00
