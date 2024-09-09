import urllib

from bs4 import BeautifulSoup
import os
import requests
import mysql.connector

# 创建数据库连接
conn = mysql.connector.connect(
    host="192.168.174.151",
    user="esports",
    password="123456",
    database="esports"
)
cursor = conn.cursor()

# 执行查询获取 image_link 字段的值
cursor.execute("SELECT image_link FROM articles")
image_links = cursor.fetchall()

# 定义保存图片的目录
save_directory = '/resource/articles'

# 确保目录存在，如果不存在则创建
if not os.path.exists(save_directory):
    os.makedirs(save_directory)

# 处理并更新链接
for link in image_links:
    original_link = link[0]
    # 去除链接中的 "./"
    corrected_link = original_link.replace("articles", "articles/")
    # 更新数据库中的 image_link 字段
    update_query = f"UPDATE articles SET image_link = '{corrected_link}' WHERE image_link = '{original_link}'"
    cursor.execute(update_query)
    conn.commit()

cursor.close()
conn.close()