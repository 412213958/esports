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


# 图片文件保存目录
save_dir = "../resource/logo"

# 创建游标对象
cursor = conn.cursor()

# 创建表
create_table_query = '''
CREATE TABLE IF NOT EXISTS stage_container (
  id INT AUTO_INCREMENT PRIMARY KEY,
  ranking INT,
  team_name VARCHAR(255),
  team_logo VARCHAR(255),
  wins INT,
  losses INT,
  points INT,
  typeSF INT
)
'''
cursor.execute(create_table_query)

html='''<div class="match-points-container" data-v-2c4c8cbb="" data-v-3f1e8d92=""><div class="tabs" data-v-2c4c8cbb=""><div class="tab active" data-v-2c4c8cbb="">登峰组</div><div class="tab" data-v-2c4c8cbb="">涅槃组</div></div><div class="header" data-v-2c4c8cbb=""><div class="th" data-v-2c4c8cbb="">名次</div><div class="th team" data-v-2c4c8cbb="">战队</div><div class="th" data-v-2c4c8cbb="">胜/负</div><div class="th" data-v-2c4c8cbb="">积分</div><!--v-if--></div><div class="list" data-v-2c4c8cbb="" style="height: 425px;"><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">1</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/176363723a29df1246030e3eba51e71d3be00373.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">LNG</div></div></div><div class="cell" data-v-02e20055="">3/<!--v-if-->0</div><div class="cell" data-v-02e20055="">3</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">2</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/7af596e4a35b7e287ff7ceb3bd42d4ba638d106b.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">AL</div></div></div><div class="cell" data-v-02e20055="">3/<!--v-if-->1</div><div class="cell" data-v-02e20055="">3</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">2</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/1eb01d43ce5069e04ff3df69e949cae2d9929915.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">JDG</div></div></div><div class="cell" data-v-02e20055="">3/<!--v-if-->1</div><div class="cell" data-v-02e20055="">3</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">4</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/737aeb7338739305c1f65c36eb1ef72f9ce3d098.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">WBG</div></div></div><div class="cell" data-v-02e20055="">2/<!--v-if-->1</div><div class="cell" data-v-02e20055="">2</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">5</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/2ca4893d7e2f2f85ba22fbd4ce1c2067713df81a.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">BLG</div></div></div><div class="cell" data-v-02e20055="">1/<!--v-if-->0</div><div class="cell" data-v-02e20055="">1</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">6</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/05a37988ad4e87fba7ca6b0eadb603982e41cad7.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">FPX</div></div></div><div class="cell" data-v-02e20055="">0/<!--v-if-->2</div><div class="cell" data-v-02e20055="">0</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">6</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/7dff90050ed9335b600a61239439fd379ddaa3ee.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">LGD</div></div></div><div class="cell" data-v-02e20055="">0/<!--v-if-->2</div><div class="cell" data-v-02e20055="">0</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">6</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/a1858308f7251261f8d3a97912abcdf39f2152b5.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">TES</div></div></div><div class="cell" data-v-02e20055="">0/<!--v-if-->2</div><div class="cell" data-v-02e20055="">0</div><!--v-if--></div><div class="points-record" data-v-02e20055="" data-v-2c4c8cbb=""><div class="cell" data-v-02e20055="">9</div><div class="cell team" data-v-02e20055=""><img src="//i1.hdslb.com/bfs/legacy/c86ac68213687d66a15fd3ff66668e6ba3c17328.png@38w_38h.avif" loading="lazy" class="logo" data-v-02e20055=""><div class="name" data-v-02e20055=""><div class="name-text" data-v-02e20055="">NIP</div></div></div><div class="cell" data-v-02e20055="">0/<!--v-if-->3</div><div class="cell" data-v-02e20055="">0</div><!--v-if--></div></div></div>'''



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
        # download_and_save_image(team_logo_url, save_dir)

        team_logo_url = "http://121.36.26.12:8081/resource/logo/" + team_logo_url.split("/")[-1]
        # 存储数据到数据库
        insert_data_query = '''
        INSERT INTO stage_container (ranking, team_name, team_logo, wins, losses, points, typeSF)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        '''
        values = (rank, team_name, team_logo_url, wins, losses, points, 1)
        cursor.execute(insert_data_query, values)

# 提交事务并关闭连接
conn.commit()
cursor.close()
conn.close()