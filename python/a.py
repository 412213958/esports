from bs4 import BeautifulSoup
import mysql.connector
from datetime import datetime

# 创建数据库连接
conn = mysql.connector.connect(
    host="192.168.5.151",
    user="esports",
    password="123456",
    database="esports"
)

# 创建游标对象
cursor = conn.cursor()

# 创建赛事日历表
create_table_query = '''
CREATE TABLE race_calendar (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255),
  session VARCHAR(255),
  date VARCHAR(255),
  time VARCHAR(255),
  team1_name VARCHAR(255),
  team1_logo VARCHAR(255),
  team2_name VARCHAR(255),
  team2_logo VARCHAR(255)
)
'''

cursor.execute(create_table_query)

# 将HTML文本赋值给变量html
'''

# 使用BeautifulSoup解析HTML
soup = BeautifulSoup(html, 'html.parser')

# 提取赛事信息
cards = soup.find_all('div', class_='card')
for card in cards:
    # 提取赛事标题
    title = card.find('div', class_='title').text.strip()
    # 提取日期和时间
    des = card.find('div', class_='des').text.strip().split('\xa0\xa0')
    match = des[0]
    des1 = des[1].split(' ')
    date = des1[1]
    time = des1[2]
    # 提取参赛队伍或选手
    con = card.find_all('div', class_='con')
    print("con：",con)
    team1_name = con[0].find('span', class_='line-ellipsis').text.strip()
    team1_logo = con[0].find('img')['src']
    team2_name = con[1].find('span', class_='line-ellipsis').text.strip()
    team2_logo = con[1].find('img')['src']

    # print('des:', des)
    # print('des1:', des1)
    print('match:', match)
    print('日期:', date)
    print('时间:', time)
    print()
    # 插入赛事数据到数据库
    insert_data_query = '''
    INSERT INTO race_calendar (title, session, date, time, team1_name, team1_logo, team2_name, team2_logo)
    VALUES (%s, %s, %s, %s, %s, %s, %s, %s)
    '''
    values = (title, match, date, time, team1_name, "http://121.36.26.12:8081/resource/logo/" + team1_logo.split("/")[-1], team2_name, "http://121.36.26.12:8081/resource/logo/" + team2_logo.split("/")[-1])
    cursor.execute(insert_data_query, values)

# 提交事务
conn.commit()

# 关闭游标和连接
cursor.close()
conn.close()

# # 打印赛事信息
# print('赛事标题:', title)
# print('日期:', date)
# print('时间:', time)
# print('参赛队伍1:', team1_name)
# print('参赛队伍1的logo:', team1_logo)
# print('参赛队伍2:', team2_name)
# print('参赛队伍2的logo:', team2_logo)