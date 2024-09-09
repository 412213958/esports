from bs4 import BeautifulSoup
import mysql.connector
import requests
import os

# 创建数据库连接
conn = mysql.connector.connect(
    host="192.168.5.151",
    user="esports",
    password="123456",
    database="esports"
)
# 创建游标
cursor = conn.cursor()

# 执行查询语句获取 image_url 字段的值
cursor.execute("SELECT image_url FROM players")
results = cursor.fetchall()

# 处理并更新 image_url 字段的值
for row in results:
    original_url = row[0]
    new_url = original_url.replace('logo', 'player')
    # 执行更新语句将新的 URL 写回数据库
    cursor.execute("UPDATE players SET image_url = %s WHERE image_url = %s", (new_url, original_url))

# 提交更改
conn.commit()

# 关闭游标和连接
cursor.close()
conn.close()