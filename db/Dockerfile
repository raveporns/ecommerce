# ใช้ PostgreSQL image ล่าสุดเป็น base
FROM postgres:latest

# ตั้งค่าตัวแปรสิ่งแวดล้อมจาก .env
ARG POSTGRES_DB
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD

# ตั้งค่าตัวแปรสิ่งแวดล้อมสำหรับ PostgreSQL
ENV POSTGRES_DB=${POSTGRES_DB}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}

# คัดลอกสคริปต์ SQL ถ้ามี
# COPY ./init.sql /docker-entrypoint-initdb.d/
