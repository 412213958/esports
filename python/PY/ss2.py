from bs4 import BeautifulSoup
import os
import requests
from bs4 import BeautifulSoup
import mysql.connector

# 创建数据库连接
conn = mysql.connector.connect(
    host="192.168.174.151",
    user="esports",
    password="123456",
    database="esports"
)
cursor = conn.cursor()

html_content = """<div title="热门圈子" data-v-28793cf8="" data-v-71580d08=""><div class="index_left_wrap" data-v-28793cf8=""><div class="index_slide_title" data-v-28793cf8=""><svg aria-hidden="true" class="icon" data-v-28793cf8=""><use xlink:href="#icon-guanfanghao" data-v-28793cf8=""></use></svg>
      官方俱乐部
      <a href="/star_list" class="index_slide_more gray_3 flr" data-v-28793cf8="">更多</a></div> <div class="index_left_main" data-v-28793cf8=""><div class="item" data-v-28793cf8=""><div class="index_left_border" data-v-28793cf8=""><a href="star?star_id=9" title="iG电子竞技俱乐部" data-v-28793cf8=""><img src="https://img.scoregg.com/z/0/p/174/1218022089753.png?imageMogr2/crop/24x24/gravity/center" class="fll" data-v-28793cf8="">iG电子竞技俱乐部</a></div></div><div class="item" data-v-28793cf8=""><div class="index_left_border" data-v-28793cf8=""><a href="star?star_id=26" title="天霸电子竞技俱乐部" data-v-28793cf8=""><img src="https://img.scoregg.com/z/6299199/p/201/1517401147685.png?imageMogr2/crop/24x24/gravity/center" class="fll" data-v-28793cf8="">天霸电子竞技俱乐部</a></div></div><div class="item" data-v-28793cf8=""><div class="index_left_border" data-v-28793cf8=""><a href="star?star_id=7" title="FPX电子竞技俱乐部" data-v-28793cf8=""><img src="https://img.scoregg.com/z/526460/p/181/0416423676610.png?imageMogr2/crop/24x24/gravity/center" class="fll" data-v-28793cf8="">FPX电子竞技俱乐部</a></div></div><div class="item" data-v-28793cf8=""><div class="index_left_border" data-v-28793cf8=""><a href="star?star_id=21" title="滔搏电子竞技俱乐部" data-v-28793cf8=""><img src="https://img.scoregg.com/z/2720502/p/195/1900313669547.png?imageMogr2/crop/24x24/gravity/center" class="fll" data-v-28793cf8="">滔搏电子竞技俱乐部</a></div></div><div class="item" data-v-28793cf8=""><div class="index_left_border" data-v-28793cf8=""><a href="star?star_id=1" title="WE电子竞技俱乐部" data-v-28793cf8=""><img src="https://img.scoregg.com/z/2962980/p/185/2116474678312.png?imageMogr2/crop/24x24/gravity/center" class="fll" data-v-28793cf8="">WE电子竞技俱乐部</a></div></div></div></div></div>
        """  # 这里替换为实际的 HTML 内容
# 创建表
# cursor.execute("""
#     CREATE TABLE IF NOT EXISTS tournaments (
#         id INT AUTO_INCREMENT PRIMARY KEY,
#         title VARCHAR(255),
#         image_src VARCHAR(255),
#         type_s VARCHAR(255)
#     )
# """)



soup = BeautifulSoup(html_content, 'html.parser')

# 提取所有赛事的信息
tournaments = soup.find_all('div', class_='item')
for tournament in tournaments:
    link_and_title = tournament.find('a', title=True)
    image = tournament.find('img')
    if link_and_title and image:
        title = link_and_title['title']
        src = image.get('src')
        print(f"标题: {title}, 图片链接: {src}")
        # 保存数据到数据库
        cursor.execute("INSERT INTO tournaments (title, image_src, type_s) VALUES (%s, %s, %s)", (title, src, "官方俱乐部"))

conn.commit()
cursor.close()
conn.close()