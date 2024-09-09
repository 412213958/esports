from bs4 import BeautifulSoup
import mysql.connector
from datetime import datetime

html1='''<div class="list-col">
          
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">第一轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152756/56c5bb71d14d13ec280908542ac8e1a9/0" alt=""></span>
          <a class="team-name">OMG</a>
          <a class="score">1</a>
        </p>
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103153050/5c22150153baa973f73adf4909a041ec.png/0" alt=""></span>
          <a class="team-name">西安WE</a>
          <a class="score">3</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">3月31日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59262&amp;bMatchId=10945" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10945" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143747/7a28f81c61de2aef3d4306cd97863094/0" alt=""></span>
          <a class="team-name">WBGTapTap</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152918/562affbb9f99b0e472644f93a19291e7/0" alt=""></span>
          <a class="team-name">IG</a>
          <a class="score">2</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月01日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59292&amp;bMatchId=10946" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10946" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">胜者组-第一轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104012/428b4f8509a330697b871cc6c53e36dc/0" alt=""></span>
          <a class="team-name">BLG星纪魅族</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152746/c8a8f59fadc0da0bda0c4997e3f62afb/0" alt=""></span>
          <a class="team-name">深圳NIP</a>
          <a class="score">1</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月06日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59452&amp;bMatchId=10951" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10951" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104232/ef51868e0de1d1516122e67de4237792/0" alt=""></span>
          <a class="team-name">TES</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143554/1e5189e914e40bfe18045bccbfc7ef82.png/0" alt=""></span>
          <a class="team-name">北京JDG英特尔</a>
          <a class="score">0</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月07日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59481&amp;bMatchId=10952" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10952" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">败者组-第一轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152746/c8a8f59fadc0da0bda0c4997e3f62afb/0" alt=""></span>
          <a class="team-name">深圳NIP</a>
          <a class="score">2</a>
        </p>
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143554/1e5189e914e40bfe18045bccbfc7ef82.png/0" alt=""></span>
          <a class="team-name">北京JDG英特尔</a>
          <a class="score">3</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月10日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59509&amp;bMatchId=10953" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10953" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        </div>
        <div class="list-col">
          
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">第二轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152746/c8a8f59fadc0da0bda0c4997e3f62afb/0" alt=""></span>
          <a class="team-name">深圳NIP</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103153050/5c22150153baa973f73adf4909a041ec.png/0" alt=""></span>
          <a class="team-name">西安WE</a>
          <a class="score">2</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月02日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59323&amp;bMatchId=10947" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10947" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
        </h3>
        <!--两支战队-->
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240528142427/93978e71ed61983aaae7689f95c1f07d/0" alt=""></span>
          <a class="team-name">苏州LNG九号电动</a>
          <a class="score">2</a>
        </p>
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143747/7a28f81c61de2aef3d4306cd97863094/0" alt=""></span>
          <a class="team-name">WBGTapTap</a>
          <a class="score">3</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月03日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59361&amp;bMatchId=10948" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10948" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">胜者组-第二轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104012/428b4f8509a330697b871cc6c53e36dc/0" alt=""></span>
          <a class="team-name">BLG星纪魅族</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104232/ef51868e0de1d1516122e67de4237792/0" alt=""></span>
          <a class="team-name">TES</a>
          <a class="score">2</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月11日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59545&amp;bMatchId=10954" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10954" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">败者组-第二轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143554/1e5189e914e40bfe18045bccbfc7ef82.png/0" alt=""></span>
          <a class="team-name">北京JDG英特尔</a>
          <a class="score">1</a>
        </p>
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104232/ef51868e0de1d1516122e67de4237792/0" alt=""></span>
          <a class="team-name">TES</a>
          <a class="score">3</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月14日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59584&amp;bMatchId=10955" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10955" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        </div>
        <div class="list-col">
          
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">第三轮</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152856/4afad949fdd8b7728b9936dc92483e50/0" alt=""></span>
          <a class="team-name">FPX</a>
          <a class="score">1</a>
        </p>
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230103152746/c8a8f59fadc0da0bda0c4997e3f62afb/0" alt=""></span>
          <a class="team-name">深圳NIP</a>
          <a class="score">3</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月04日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59395&amp;bMatchId=10949" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10949" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143554/1e5189e914e40bfe18045bccbfc7ef82.png/0" alt=""></span>
          <a class="team-name">北京JDG英特尔</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20230921143747/7a28f81c61de2aef3d4306cd97863094/0" alt=""></span>
          <a class="team-name">WBGTapTap</a>
          <a class="score">0</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月05日 18:00</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59448&amp;bMatchId=10950" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10950" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        <div class="list-col-one over">
        <!-- 比赛轮数 -->
        <h3 class="match-part-name">
          
          <span class="sheaves-num">决赛</span>
          <span class="status">已结束</span>
          
        </h3>
        <!--两支战队-->
        <p class="team win">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104012/428b4f8509a330697b871cc6c53e36dc/0" alt=""></span>
          <a class="team-name">BLG星纪魅族</a>
          <a class="score">3</a>
        </p>
        <p class="team ">
          <span class="img-wrap"><img src="//img.crawler.qq.com/lolwebvideo/20240423104232/ef51868e0de1d1516122e67de4237792/0" alt=""></span>
          <a class="team-name">TES</a>
          <a class="score">1</a>
        </p>
        <!--对局下方的游戏状态,时间等信息栏-->
        <p class="info">
          <span class="time">4月20日 16:30</span>
          
            <a class="herf-video sprites-icon" target="_blank" href="//lpl.qq.com/es/video_detail.shtml?nid=59635&amp;bMatchId=10956" onclick="PTTSendClick('event','video','查看视频')">视频</a>
            <a class="herf-data sprites-icon" target="_blank" href="//lpl.qq.com/es/stats.shtml?bmid=10956" onclick="PTTSendClick('event','gamedata','查看数据')">数据</a>
          
        </p>
      </div>
      
        </div>
        '''
html2=''''''
html3=''''''

conn = mysql.connector.connect(
    host="192.168.30.46",
    user="esports",
    password="123456",
    database="esports"
)

# 创建游标对象
cursor = conn.cursor()

# 使用BeautifulSoup解析HTML
soup = BeautifulSoup(html1, 'html.parser')

# 建表语句
create_table_sql = """
CREATE TABLE IF NOT EXISTS match_information (
    id INT AUTO_INCREMENT PRIMARY KEY,
    match_part_name VARCHAR(255),
    status VARCHAR(255),
    team1 VARCHAR(255),
    team1_score INT,
    team1_img_src VARCHAR(255),
    team2 VARCHAR(255),
    team2_score INT,
    team2_img_src VARCHAR(255),
    time VARCHAR(255),
    video_link VARCHAR(255),
    data_link VARCHAR(255),
    format_name VARCHAR(255)
);
"""
cursor.execute(create_table_sql)  # 执行建表语句

match_data = []

def extract_info(list_col_one):
    try:
        match_part_name = list_col_one.find('h3', class_='match-part-name').find('span', class_='sheaves-num').text.strip() if list_col_one.find('h3', class_='match-part-name').find('span', class_='sheaves-num') else "未找到比赛轮数"
        status = list_col_one.find('h3', class_='match-part-name').find('span', class_='status').text.strip() if list_col_one.find('h3', class_='match-part-name').find('span', class_='status') else "未找到状态"
        team1 = list_col_one.find('p', class_='team').find('a', class_='team-name').text.strip() if list_col_one.find('p', class_='team').find('a', class_='team-name') else "未找到队伍 1"
        team1_score = list_col_one.find('p', class_='team').find('a', class_='score').text.strip() if list_col_one.find('p', class_='team').find('a', class_='score') else "未找到队伍 1 比分"
        team1_img_src = list_col_one.find('p', class_='team').find('span', class_='img-wrap').find('img')['src'] if list_col_one.find('p', class_='team').find('span', class_='img-wrap').find('img') else "未找到队伍 1 头像"
        team2 = list_col_one.find('p', class_='team win').find('a', class_='team-name').text.strip() if list_col_one.find('p', class_='team win').find('a', class_='team-name') else "未找到队伍 2"
        team2_score = list_col_one.find('p', class_='team win').find('a', class_='score').text.strip() if list_col_one.find('p', class_='team win').find('a', class_='score') else "未找到队伍 2 比分"
        team2_img_src = list_col_one.find('p', class_='team win').find('span', class_='img-wrap').find('img')['src'] if list_col_one.find('p', class_='team win').find('span', class_='img-wrap').find('img') else "未找到队伍 2 头像"
        time = list_col_one.find('p', class_='info').find('span', class_='time').text.strip() if list_col_one.find('p', class_='info').find('span', class_='time') else "未找到时间"
        video_link = list_col_one.find('p', class_='info').find('a', class_='herf-video')['href'] if list_col_one.find('p', class_='info').find('a', class_='herf-video') else "未找到视频链接"
        data_link = list_col_one.find('p', class_='info').find('a', class_='herf-data')['href'] if list_col_one.find('p', class_='info').find('a', class_='herf-data') else "未找到数据链接"

        return {
            '比赛轮数': match_part_name,
            '状态': status,
            '队伍 1': team1,
            '队伍 1 比分': team1_score,
            '队伍 1 头像': team1_img_src,
            '队伍 2': team2,
            '队伍 2 比分': team2_score,
            '队伍 2 头像': team2_img_src,
            '时间': time,
            '视频链接': video_link,
            '数据链接': data_link
        }
    except AttributeError as e:
        print(f"在处理当前元素时发生错误: {e}")
        return None

for list_col_one in soup.find_all('div', class_='list-col-one over'):
    data = extract_info(list_col_one)
    print(data)
    if data:
        match_data.append(data)

# 构建插入语句
sql = "INSERT INTO match_information (match_part_name, status, team1, team1_score, team1_img_src, team2, team2_score, team2_img_src, time, video_link, data_link) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"

# 执行插入操作
for item in match_data:
    values = (item['比赛轮数'], item['状态'], item['队伍 1'], item['队伍 1 比分'], item['队伍 1 头像'], item['队伍 2'], item['队伍 2 比分'], item['队伍 2 头像'], item['时间'], item['视频链接'], item['数据链接'])
    cursor.execute(sql, values)

# 提交更改
conn.commit()

# 关闭游标和连接
cursor.close()
conn.close()