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


def create_tables():
    # 创建玩家表
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS players (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            image_url VARCHAR(255)
        )
    """)
    # 创建战队信息表
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS team_info (
            id INT AUTO_INCREMENT PRIMARY KEY,
            squadron VARCHAR(255),
            win_rate VARCHAR(255),
            win_desc VARCHAR(255),
            kda VARCHAR(255),
            kda_desc VARCHAR(255),
            output_per_minute VARCHAR(255),
            economy_per_minute VARCHAR(255),
            cs_per_minute VARCHAR(255),
            dragon_control_rate VARCHAR(255),
            average_kills VARCHAR(255),
            average_assists VARCHAR(255),
            average_deaths VARCHAR(255),
            baron_control_rate VARCHAR(255)
        )
    """)


# create_tables()

html = '''
<div class="team-info" data-v-5bc87674=""><div class="team-players" data-v-5bc87674=""><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/6500519/p/245/3111561470029_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">ON</span></div><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/2373870/p/247/1013275138590_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">WEI</span></div><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/6500519/p/245/3111553020179_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">Bin</span></div><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/6500519/p/245/3111562640755_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">Xun</span></div><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/6500519/p/245/3111554128667_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">Elk</span></div><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/6500519/p/245/3111555228484_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">knight</span></div><div class="player" data-v-5bc87674=""><img src="//i1.hdslb.com/bfs/esport/m/6500519/p/245/3111560376992_300X300.png@48w_48h.avif" loading="lazy" class="player-img" data-v-5bc87674=""><span data-v-5bc87674="">LvMao</span></div><!--v-if--></div><div class="team-data" data-v-5bc87674=""><div class="rings" data-v-5bc87674=""><div class="win" data-v-5bc87674=""><div class="ring" data-v-180491d5="" data-v-5bc87674=""><svg width="100" height="100" viewBox="0 0 100 100" data-v-180491d5=""><circle cx="50" cy="50" r="45" stroke-width="10" stroke-linecap="round" stroke="rgba(17, 64, 240, 0.1)" fill="none" data-v-180491d5=""></circle><circle class="ring-item" cx="50" cy="50" r="45" stroke-width="10" stroke-dasharray="243.15927138785,39.5840674352314" stroke-linecap="round" stroke="#0094ff" fill="none" data-v-180491d5=""></circle></svg><div class="ring-data" data-v-5bc87674=""><span class="ring-data-number" data-v-5bc87674="">86<svg width="9" height="19" viewBox="0 0 8 13" fill="none" xmlns="http://www.w3.org/2000/svg" data-v-5bc87674=""><path d="M5.23438 10.3984C5.23438 10.9245 5.43229 11.1875 5.82812 11.1875C6.22396 11.1875 6.42188 10.9245 6.42188 10.3984V9.53906C6.42188 9.01302 6.22396 8.75 5.82812 8.75C5.43229 8.75 5.23438 9.01302 5.23438 9.53906V10.3984ZM4.32031 9.42188C4.33073 8.86979 4.49219 8.46615 4.80469 8.21094C5.09115 7.96615 5.43229 7.84375 5.82812 7.84375C6.23438 7.84375 6.58594 7.96615 6.88281 8.21094C7.16927 8.46615 7.31771 8.86979 7.32812 9.42188V10.5156C7.32812 10.7917 7.29167 11.026 7.21875 11.2188C7.14583 11.4219 7.03385 11.5859 6.88281 11.7109C6.58594 11.9661 6.23438 12.0938 5.82812 12.0938C5.43229 12.0938 5.09115 11.9661 4.80469 11.7109C4.49219 11.4661 4.33073 11.0677 4.32031 10.5156V9.42188ZM0.65625 2.09375C0.666667 1.54167 0.828125 1.13802 1.14062 0.882812C1.42708 0.638021 1.76823 0.515625 2.16406 0.515625C2.57031 0.515625 2.92188 0.638021 3.21875 0.882812C3.50521 1.13802 3.65365 1.54167 3.66406 2.09375V3.1875C3.66406 3.46354 3.6276 3.69792 3.55469 3.89062C3.48177 4.09375 3.36979 4.25781 3.21875 4.38281C2.92188 4.63802 2.57031 4.76562 2.16406 4.76562C1.76823 4.76562 1.42708 4.63802 1.14062 4.38281C0.828125 4.13802 0.666667 3.73958 0.65625 3.1875V2.09375ZM1.57031 3.07031C1.57031 3.59635 1.76823 3.85938 2.16406 3.85938C2.5599 3.85938 2.75781 3.59635 2.75781 3.07031V2.21094C2.75781 1.6849 2.5599 1.42188 2.16406 1.42188C1.76823 1.42188 1.57031 1.6849 1.57031 2.21094V3.07031ZM6.16406 0.609375H7.21875L1.82812 12H0.773438L6.16406 0.609375Z" fill="currentcolor"></path></svg></span><span class="ring-data-text" data-v-5bc87674="">胜率</span></div></div><div class="win-desc" data-v-5bc87674="">6胜 1败 </div></div><div class="kda" data-v-5bc87674=""><div class="ring" data-v-180491d5="" data-v-5bc87674=""><svg width="100" height="100" viewBox="0 0 100 100" data-v-180491d5=""><circle cx="50" cy="50" r="45" stroke-width="10" stroke-linecap="round" stroke="rgba(242, 32, 77, 0.1)" fill="none" data-v-180491d5=""></circle><circle class="ring-item" cx="50" cy="50" r="45" stroke-width="10" stroke-dasharray="70.68583470577035,212.05750411731105" stroke-linecap="round" stroke="#F2204D" fill="none" data-v-180491d5=""></circle><circle class="ring-item" cx="50" cy="50" r="45" stroke-width="10" stroke-dasharray="118.75220230569418,163.99113651738722" stroke-linecap="round" stroke="rgba(242, 32, 77, 0.6)" fill="none" data-v-180491d5=""></circle><circle class="ring-item" cx="50" cy="50" r="45" stroke-width="10" stroke-dasharray="282.7433388230814,0" stroke-linecap="round" stroke="rgba(242, 32, 77, 0.1)" fill="none" data-v-180491d5=""></circle></svg><div class="ring-data" data-v-5bc87674=""><span class="ring-data-number" data-v-5bc87674="">4.7</span><span class="ring-data-text" data-v-5bc87674="">KDA</span></div></div><div class="kda-desc" data-v-5bc87674="">18.5/12.8/42.2</div></div></div><ul class="data-detail" data-v-5bc87674=""><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">分均输出</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">2943.6</span><span data-v-5bc87674=""></span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">分均经济</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">2027.7</span><span data-v-5bc87674=""></span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">分均补刀</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">35.3</span><span data-v-5bc87674=""></span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">小龙控制率</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">58.2</span><span data-v-5bc87674="">%</span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">场均击杀</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">18.5</span><span data-v-5bc87674=""></span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">场均助攻</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">42.2</span><span data-v-5bc87674=""></span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">场均死亡</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">12.8</span><span data-v-5bc87674=""></span></span></li><li class="detail-item" data-v-5bc87674=""><span class="title" data-v-5bc87674="">大龙控制率</span><span class="number" data-v-5bc87674=""><span data-v-5bc87674="">71.9</span><span data-v-5bc87674="">%</span></span></li></ul></div></div>
'''
# 图片文件保存目录
save_dir = "../resource/player"


def download_and_save_image(image_url, save_dir):
    """
    下载图片并保存到指定目录

    参数：
    image_url (str): 图片的 URL
    save_dir (str): 保存图片的目录
    """
    try:
        response = requests.get(image_url, verify=False)
        # 提取文件名
        file_name = image_url.split("/")[-1]
        # 保存图片到指定目录
        save_path = os.path.join(save_dir, file_name)
        with open(save_path, "wb") as f:
            f.write(response.content)
        print("已下载图片文件:", file_name)
    except Exception as e:
        print(f"下载图片时出错: {e}")


def insert_player_data(name, image_url):
    query = "INSERT INTO players (name, image_url) VALUES (%s, %s)"
    cursor.execute(query, (name, image_url))


def insert_team_info(squadron, win_rate, win_desc, kda, kda_desc, output_per_minute, economy_per_minute, cs_per_minute,
                     dragon_control_rate, average_kills, average_assists, average_deaths, baron_control_rate):
    query = "INSERT INTO team_info (squadron,win_rate, win_desc, kda, kda_desc, output_per_minute, economy_per_minute, cs_per_minute, dragon_control_rate, average_kills, average_assists, average_deaths, baron_control_rate) VALUES (%s,%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
    cursor.execute(query, (squadron,
                           win_rate, win_desc, kda, kda_desc, output_per_minute, economy_per_minute, cs_per_minute,
                           dragon_control_rate,
                           average_kills, average_assists, average_deaths, baron_control_rate))


# 使用BeautifulSoup解析HTML
soup = BeautifulSoup(html, 'html.parser')

# 提取赛事信息
team_info = soup.find('div', class_='team-info')

# 打印所有球员信息
players = team_info.find_all('div', class_='player')
for player in players:
    img = player.find('img')
    span = player.find('span')
    name = span.text
    image_url = "http://121.36.26.12:8081/resource/logo/" + img['src'].split("/")[-1]
    # 下载并保存图片到本地
    download_and_save_image("http:" + img['src'], save_dir)
    insert_player_data(name, image_url)
    print(f"Player: {span.text}, Image: {image_url}")

# 提取并打印胜率信息
win_data = team_info.find('div', class_='win')
ring_data_number = win_data.find('span', class_='ring-data-number')
win_desc = win_data.find('div', class_='win-desc')
win_rate = ring_data_number.text
win_rate_desc = win_desc.text

# 提取并打印KDA信息
kda_data = team_info.find('div', class_='kda')
kda_ring_data_number = kda_data.find('span', class_='ring-data-number')
kda_desc = kda_data.find('div', class_='kda-desc')
kda_value = kda_ring_data_number.text
kda_desc_value = kda_desc.text

# 提取并打印详细数据
data_details = team_info.find('ul', class_='data-detail')
output_per_minute = None
economy_per_minute = None
cs_per_minute = None
dragon_control_rate = None
average_kills = None
average_assists = None
average_deaths = None
baron_control_rate = None

for detail in data_details.find_all('li', class_='detail-item'):
    title = detail.find('span', class_='title').text
    number = detail.find('span', class_='number').text
    if title == "分均输出":
        output_per_minute = number
    elif title == "分均经济":
        economy_per_minute = number
    elif title == "分均补刀":
        cs_per_minute = number
    elif title == "小龙控制率":
        dragon_control_rate = number
    elif title == "场均击杀":
        average_kills = number
    elif title == "场均助攻":
        average_assists = number
    elif title == "场均死亡":
        average_deaths = number
    elif title == "大龙控制率":
        baron_control_rate = number

# 插入战队信息到表中
insert_team_info("BLG", win_rate, win_rate_desc, kda_value, kda_desc_value, output_per_minute, economy_per_minute,
                 cs_per_minute, dragon_control_rate, average_kills, average_assists, average_deaths, baron_control_rate)

# 提交更改并关闭连接
conn.commit()
cursor.close()
conn.close()
