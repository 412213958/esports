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
save_directory = './resource/articles'

# 确保目录存在，如果不存在则创建
if not os.path.exists(save_directory):
    os.makedirs(save_directory)

# 下载图片并更新数据库字段
for link in image_links:
    image_url = link[0]
    # 使用 urlparse 解析 URL
    parsed_url = urllib.parse.urlparse(image_url)
    # 获取路径部分
    path = parsed_url.path
    # 分割路径并获取最后一个元素作为文件名
    file_name = path.split('/')[-1]

    # 从链接中提取文件名
    # file_name = os.path.split(image_url)[1]
    response = requests.get(image_url)
    if response.status_code == 200:
        with open(os.path.join(save_directory, file_name), 'wb') as file:
            file.write(response.content)
        # 更新数据库中的 image_link 字段为文件的本地路径
        local_path = "http://127.0.0.1:8081"+os.path.join(save_directory, file_name)
        update_query = f"UPDATE articles SET image_link = '{local_path}' WHERE image_link = '{image_url}'"
        cursor.execute(update_query)
        conn.commit()