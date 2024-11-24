# ใช้ภาพ PostgreSQL เป็น base image
FROM postgres:latest

# ถ้าต้องการเพิ่มการตั้งค่าหรือไฟล์เพิ่มเติมสามารถคัดลอกไปยัง /docker-entrypoint-initdb.d/
COPY ./init.sql /docker-entrypoint-initdb.d/

# เปิดพอร์ตที่ใช้สำหรับ PostgreSQL
EXPOSE 5432
