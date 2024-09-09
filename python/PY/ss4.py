from bs4 import BeautifulSoup
import mysql.connector
import requests
import os
import urllib3

html_content = """
<ul id="accordion" class="swiper-wrapper" data-v-3713f460=""><li class="swiper-slide" data-v-3713f460="" style="width: 40px; opacity: 0.5; cursor: pointer;"><a href="/p/1425675" data-v-3713f460=""><img src="https://img.scoregg.com/z/2373870/p/247/2215342652674.jpg?imageMogr2/crop/520x260/gravity/center" data-v-3713f460=""></a> <div data-v-3713f460="" style="opacity: 1; display: none;"><a href="/p/1425675" data-v-3713f460=""><h2 data-v-3713f460=""><span data-v-3713f460="">精选</span>第三周最佳：JKL斩获MVP选手，milkyway获最佳新秀</h2> <p class="gray_3" data-v-3713f460="">LPL官方公布了第三周MVP选手及最佳新秀，TES下路选手JackeyLove获本周MVP选手，FPX打野选手milkway获得最佳新秀。第三周最佳新秀——FPX.milkyway第三周MVP选手——TES.Jackeylove</p></a></div></li><li class="swiper-slide" data-v-3713f460="" style="cursor: pointer; width: 40px; opacity: 0.5;"><a href="/p/1425674" data-v-3713f460=""><img src="https://img.scoregg.com/z/2373870/p/247/2215331642891.jpg?imageMogr2/crop/520x260/gravity/center" data-v-3713f460=""></a> <div data-v-3713f460="" style="opacity: 1; display: none;"><a href="/p/1425674" data-v-3713f460=""><h2 data-v-3713f460=""><span data-v-3713f460="">精选</span>组内赛第三周最佳阵容：晴天 银河 Angel JackeyLove Meiko</h2> <p class="gray_3" data-v-3713f460="">英雄联盟赛事官方公布了LPL夏季赛组内赛第三周的最佳阵容，选手分别是：上单——UP.Qingtian打野——FPX.milkyway中单——OMG.AngelADC——TES.JackeyLove辅助——TES.Meiko</p></a></div></li><li class="swiper-slide" data-v-3713f460="" style="cursor: pointer; width: 40px; opacity: 0.5;"><a href="https://www.scoregg.com/match/44341" data-v-3713f460=""><img src="https://img.scoregg.com/z/953482/p/247/2012333473648.jpg?imageMogr2/crop/520x260/gravity/center" data-v-3713f460=""></a> <div data-v-3713f460="" style="opacity: 1; display: none;"><a href="https://www.scoregg.com/match/44341" data-v-3713f460=""><h2 data-v-3713f460=""><span data-v-3713f460="">精选</span>[图文直播] 2024LPL夏季赛 JDG vs BLG</h2> <p class="gray_3" data-v-3713f460="">[图文直播] 2024LPL夏季赛 JDG vs BLG</p></a></div></li><li class="swiper-slide" data-v-3713f460="" style="cursor: default; width: 480px; opacity: 1;"><a href="/p/1424455" data-v-3713f460=""><img src="https://img.scoregg.com/z/2373870/p/247/1515471477916.jpg?imageMogr2/crop/520x260/gravity/center" data-v-3713f460=""></a> <div data-v-3713f460="" style="opacity: 1; display: block;"><a href="/p/1424455" data-v-3713f460=""><h2 data-v-3713f460=""><span data-v-3713f460="">精选</span>LPL组内赛前两周：Tarzan斩获MVP Yanxiang获最佳新秀</h2> <p class="gray_3" data-v-3713f460="">LPL官方发布了夏季赛组内赛前两周最佳选手及新秀，WBG打野Tarzan斩获MVP，WE打野Yanxiang获最佳新秀。第一&amp;第二周MVP选手——WBG.Tarzan第一&amp;第二周最佳新秀——WE.Yanxiang</p></a></div></li><li class="swiper-slide" data-v-3713f460="" style="cursor: pointer; width: 40px; opacity: 0.5;"><a href="/p/1424454" data-v-3713f460=""><img src="https://img.scoregg.com/z/2373870/p/247/1515460758465.jpg?imageMogr2/crop/520x260/gravity/center" data-v-3713f460=""></a> <div data-v-3713f460="" style="opacity: 1; display: none;"><a href="/p/1424454" data-v-3713f460=""><h2 data-v-3713f460=""><span data-v-3713f460="">精选</span>LPL组内赛第一/二周最佳阵容：Ale+Tarzan+左手+GALA+Kael</h2> <p class="gray_3" data-v-3713f460="">LPL夏季赛组内赛第一&amp;第二周最佳阵容情况公布，AL两人入选，WBGBLGLNG各一人。上单——AL.Ale打野——WBG.Tarzan（三次MVP）中单——BLG.knightADC——LNG.GALA辅助——AL.Kael</p></a></div></li></ul>
"""

# 禁用警告
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# 或者添加证书验证（需要根据实际情况进行配置）
http = urllib3.PoolManager(cert_reqs='CERT_REQUIRED')

# 创建数据库连接
conn = mysql.connector.connect(
    host="192.168.174.151",
    user="esports",
    password="123456",
    database="esports"
)
cursor = conn.cursor()

soup = BeautifulSoup(html_content, 'html.parser')

# 创建表
create_table_query = """
CREATE TABLE IF NOT EXISTS articles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    image_link VARCHAR(255),
    lp VARCHAR(255)
)
"""
cursor.execute(create_table_query)

for li in soup.find_all('li', class_='swiper-slide'):
    title = li.find('h2').text.strip()
    description = li.find('p', class_='gray_3').text.strip()
    image = li.find('img')['src']

    insert_query = """
    INSERT INTO articles (title, description, image_link)
    VALUES (%s, %s, %s)
    """
    cursor.execute(insert_query, (title, description, image))

conn.commit()
conn.close()