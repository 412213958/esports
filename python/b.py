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

# 创建游标对象
cursor = conn.cursor()

# 查询数据库表获取team1_logo和team2_logo字段
query = "SELECT id, team1_logo, team2_logo FROM race_calendar"
cursor.execute(query)
results = cursor.fetchall()

# 图片文件保存目录
save_dir = "../resource/logo"
os.makedirs(save_dir, exist_ok=True)

def download_and_save_image(image_url, save_path):
    """
    下载图片并保存到指定路径

    参数：
    image_url (str): 图片的 URL
    save_path (str): 保存图片的路径
    """
    try:
        response = requests.get(image_url, verify=False)
        with open(save_path, "wb") as f:
            f.write(response.content)
        print("已下载图片文件:", os.path.basename(save_path))
    except Exception as e:
        print(f"下载图片时出错: {e}")

# 遍历结果下载图片并保存
for result in results:
    id_value, field1, field2 = result
    updated_urls = []

    for field in [field1, field2]:
        if field:
            image_url = "https:" + field
            # 提取文件名
            file_name = image_url.split("/")[-1]
            # 保存图片到指定目录
            save_path = os.path.join(save_dir, file_name)
            download_and_save_image(image_url, save_path)
            updated_urls.append("http://121.36.26.12:8081/resource/logo/" + file_name)

    # 更新数据库对应字段的值
    try:
        update_query = "UPDATE race_calendar SET team1_logo = %s, team2_logo = %s WHERE id = %s"
        cursor.execute(update_query, (updated_urls[0], updated_urls[1], id_value))
    except Exception as e:
        print(f"更新数据库时出错: {e}")

# 提交事务并关闭游标和连接
try:
    conn.commit()
    cursor.close()
    conn.close()
except Exception as e:
    print(f"关闭连接时出错: {e}")