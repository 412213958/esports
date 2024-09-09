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

# 图片文件保存目录
save_dir = "../resource/logo"

# 创建游标对象
cursor = conn.cursor()

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

html='''<div class="list" data-v-2c4c8cbb="" style="height: 425px;"><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">1</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/b8e287225a586dded6e24d183e17351eacec379c.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">OMG</div></div></div><div class="cell" data-v-02e20055="">2/<!--v-if-->0</div><div class="cell" data-v-02e20055="">2</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">1</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/bc9fb3f95f9f6e4cc5d2313d3d8e3b7b3086fa83.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">WE</div></div></div><div class="cell" data-v-02e20055="">2/<!--v-if-->0</div><div class="cell" data-v-02e20055="">2</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">3</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/256e6b6c4e57ca196feb0ff3661daed4686a0f72.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">RA</div></div></div><div class="cell" data-v-02e20055="">2/<!--v-if-->1</div><div class="cell" data-v-02e20055="">2</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">3</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/7a0bd8b0bcc77dd4a3b7687d7349dcb8d5d8ac9f.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">TT</div></div></div><div class="cell" data-v-02e20055="">2/<!--v-if-->1</div><div class="cell" data-v-02e20055="">2</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">5</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/9f0e6cf7c384d96a46720c86b43a45a60df5b9eb.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">EDG</div></div></div><div class="cell" data-v-02e20055="">1/<!--v-if-->2</div><div class="cell" data-v-02e20055="">1</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">5</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/3e875598eb939cbeca899f3322a4e5474f160561.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">IG</div></div></div><div class="cell" data-v-02e20055="">1/<!--v-if-->2</div><div class="cell" data-v-02e20055="">1</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">5</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/0a2f3ad91e14602eae4c10b70403681d9032f8c8.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">UP</div></div></div><div class="cell" data-v-02e20055="">1/<!--v-if-->2</div><div class="cell" data-v-02e20055="">1</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">8</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/cd94367914022022c4f0ea5eab3307a7c74df1ae.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">RNG</div></div></div><div class="cell" data-v-02e20055="">0/<!--v-if-->3</div><div class="cell" data-v-02e20055="">0</div><!--v-if--></div></div>'''

# 使用BeautifulSoup解析HTML
soup = BeautifulSoup(html, 'html.parser')

# 提取赛事信息
cards = soup.find_all('div', class_='points-record')
for card in cards:
    cells = card.find_all('div', class_='cell')
    if len(cells) == 4 or len(cells) == 5:
        rank = cells[0].text.strip()
        team_logo_url = "https:" + cells[1].find('img')['src']
        team_name = cells[1].find('div', class_='name-text').text.strip()
        win_lose = cells[2].text.strip().split('/')
        wins = int(win_lose[0])
        losses = int(win_lose[1])
        points = cells[3].text.strip()

        print('名次:', rank)
        print('战队:', team_name)
        print('战队logo URL:', team_logo_url)
        print('胜:', wins)
        print('负:', losses)
        print('积分:', points)
        print()

        # 下载并保存图片到本地
        download_and_save_image(team_logo_url, save_dir)

        team_logo_url = "http://121.36.26.12:8081/resource/logo/" + team_logo_url.split("/")[-1]

        # print('Url:',team_logo_url)
        # print()
        # 存储数据到数据库
        insert_data_query = '''
        INSERT INTO stage_container (ranking, team_name, team_logo, wins, losses, points, typeSF)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        '''
        values = (rank, team_name, team_logo_url, wins, losses, points, -1)
        cursor.execute(insert_data_query, values)

# 提交事务并关闭连接
conn.commit()
cursor.close()
conn.close()