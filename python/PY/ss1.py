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

html_content = """<div title="热门赛事" data-v-0e96796e="" data-v-71580d08=""><div class="index_left_wrap" data-v-0e96796e=""><div class="index_slide_title" data-v-0e96796e=""><svg aria-hidden="true" class="icon" data-v-0e96796e=""><use xlink:href="#icon-remensaishi" data-v-0e96796e=""></use></svg>
      热门赛事
      <a href="/tournament_list" class="index_slide_more gray_3 flr" data-v-0e96796e="">更多</a></div> <div class="index_left_main" data-v-0e96796e=""><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=666" title="2024 LDL发展联赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/241/2915495332590.png" class="fll" data-v-0e96796e="">
            2024 LDL发展联赛
          </a>
          https://img.scoregg.com/z/2373870/p/241/2915495332590.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=719" title="2024 LCK夏季赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/245/2313473677028.png" class="fll" data-v-0e96796e="">
            2024 LCK夏季赛
          </a>
          https://img.scoregg.com/z/2373870/p/245/2313473677028.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=724" title="2024 LPL夏季赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/245/2715190395182.png" class="fll" data-v-0e96796e="">
            2024 LPL夏季赛
          </a>
          https://img.scoregg.com/z/2373870/p/245/2715190395182.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=726" title="2024 LCS 夏季赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/245/2813141468471.png" class="fll" data-v-0e96796e="">
            2024 LCS 夏季赛
          </a>
          https://img.scoregg.com/z/2373870/p/245/2813141468471.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=747" title="2024 VCS夏季赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/246/2111020043644.png" class="fll" data-v-0e96796e="">
            2024 VCS夏季赛
          </a>
          https://img.scoregg.com/z/2373870/p/246/2111020043644.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=748" title="2024 CBLOL夏季赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/246/2111063718889.png" class="fll" data-v-0e96796e="">
            2024 CBLOL夏季赛
          </a>
          https://img.scoregg.com/z/2373870/p/246/2111063718889.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=749" title="2024 LLA闭幕赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/246/2111143873754.png" class="fll" data-v-0e96796e="">
            2024 LLA闭幕赛
          </a>
          https://img.scoregg.com/z/2373870/p/246/2111143873754.png
        </div></div><div class="item" data-v-0e96796e=""><div class="index_left_border" data-v-0e96796e=""><a href="/match_pc?tournamentID=746" title="2024 PCS夏季赛" data-v-0e96796e=""><img src="https://img.scoregg.com/z/2373870/p/246/1916185898852.png" class="fll" data-v-0e96796e="">
            2024 PCS夏季赛
          </a>
          https://img.scoregg.com/z/2373870/p/246/1916185898852.png
        </div></div></div></div></div>
        """  # 这里替换为实际的 HTML 内容
# 创建表
cursor.execute("""
    CREATE TABLE IF NOT EXISTS tournaments (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255),
        image_src VARCHAR(255),
        type_s VARCHAR(255)
    )
""")



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
        cursor.execute("INSERT INTO tournaments (title, image_src, type_s) VALUES (%s, %s, %s)", (title, src, "热门赛事"))

conn.commit()
cursor.close()
conn.close()