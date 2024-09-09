from bs4 import BeautifulSoup
import mysql.connector
import requests
import os
import urllib3

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

# 创建表
create_table_query = """
CREATE TABLE IF NOT EXISTS weibo_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user VARCHAR(255),
    content TEXT,
    time VARCHAR(255),
    image_src VARCHAR(255),
    avatar_src VARCHAR(255),
    lp VARCHAR(255)
)
"""
cursor.execute(create_table_query)

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

html_content = """<ul class="index_left_weibo" data-v-24d379bc=""><li data-v-24d379bc=""><a href="/u/5605458" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5605458/p/177/3117055446140_100x100.jpg" class="fll" data-v-24d379bc="">
          管泽元
        </a> <a href="/p/1425935" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">为什么凌晨4点多<br>
楼里传来了电钻声。。。。<br>
把我钻醒了。。。。 &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img src="https://img.scoregg.com/m/5605458/p/247/2404152526118.png?imageMogr2/thumbnail/300x" data-v-24d379bc=""></span></a> <p class="index_left_weibo_info" data-v-24d379bc="">
          4小时前
          <a href="/p/1425935" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425935" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5694496" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/z/5694496/p/195/3114410788404_100x100.png" class="fll" data-v-24d379bc="">
          _Tian_____
        </a> <a href="/p/1425931" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">http://t.cn/A68q7kCQ &ZeroWidthSpace;</p> </a> <p class="index_left_weibo_info" data-v-24d379bc="">
          8小时前
          <a href="/p/1425931" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425931" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/6288915" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/z/6288915/p/211/1117520259305_100x100.png" class="fll" data-v-24d379bc="">
          TES_369
        </a> <a href="/p/1425912" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">那个，天神，恭喜你啊，24了@滔搏电子竞技俱乐部 视频速发 &ZeroWidthSpace;</p> </a> <p class="index_left_weibo_info" data-v-24d379bc="">
          11小时前
          <a href="/p/1425912" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425912" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5620389" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5620389/p/178/2113285249424_100x100.jpg" class="fll" data-v-24d379bc="">
          米勒桑苏
        </a> <a href="/p/1425900" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">tes纯在虐！nip这状态 想要进季后赛好像只能指望乐爷了[融化]#2024lpl夏季赛# &ZeroWidthSpace;</p> </a> <p class="index_left_weibo_info" data-v-24d379bc="">
          12小时前
          <a href="/p/1425900" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425900" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5620389" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5620389/p/178/2113285249424_100x100.jpg" class="fll" data-v-24d379bc="">
          米勒桑苏
        </a> <a href="/p/1425890" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">rookie致命探草，tes成功翻盘[good]#2024lpl夏季赛# &ZeroWidthSpace;</p> </a> <p class="index_left_weibo_info" data-v-24d379bc="">
          12小时前
          <a href="/p/1425890" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425890" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5605458" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5605458/p/177/3117055446140_100x100.jpg" class="fll" data-v-24d379bc="">
          管泽元
        </a> <a href="/p/1425868" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">兄弟们 这把Light选手摄像头内容比游戏里还有意思，像看真人秀[doge]#2024lpl# &ZeroWidthSpace;</p> </a> <p class="index_left_weibo_info" data-v-24d379bc="">
          13小时前
          <a href="/p/1425868" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425868" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5620389" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5620389/p/178/2113285249424_100x100.jpg" class="fll" data-v-24d379bc="">
          米勒桑苏
        </a> <a href="/p/1425840" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">blg第一盘被虐。。牢赞这豹女好狠#2024lpl夏季赛# &ZeroWidthSpace;</p> </a> <p class="index_left_weibo_info" data-v-24d379bc="">
          14小时前
          <a href="/p/1425840" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425840" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/6093203" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/6093203/p/185/2415545058895_100x100.jpg" class="fll" data-v-24d379bc="">
          余霜
        </a> <a href="/p/1425832" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">巴黎出行必备！🚨🚨🚨#余霜[超话]# &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img src="https://img.scoregg.com/m/6093203/p/247/2318002877150.png?imageMogr2/thumbnail/300x" data-v-24d379bc=""></span></a> <p class="index_left_weibo_info" data-v-24d379bc="">
          14小时前
          <a href="/p/1425832" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425832" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5620389" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5620389/p/178/2113285249424_100x100.jpg" class="fll" data-v-24d379bc="">
          米勒桑苏
        </a> <a href="/p/1425813" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">猴子ak！ &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img src="https://img.scoregg.com/m/5620389/p/247/2314453635882.png?imageMogr2/thumbnail/300x" data-v-24d379bc=""></span></a> <p class="index_left_weibo_info" data-v-24d379bc="">
          18小时前
          <a href="/p/1425813" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425813" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a href="/u/5652938" class="index_left_weibo_avatar" data-v-24d379bc=""><img src="https://img.scoregg.com/m/5652938/p/179/2620030748130_100x100.jpg" class="fll" data-v-24d379bc="">
          王多多1991
        </a> <a href="/p/1425811" class="index_left_weibo_main" data-v-24d379bc=""><p data-v-24d379bc="">晚上七点，解说TES：NIP。搭档瞳夕。<br>
TES目前2-2，仍然有可能面临抢位赛“骑士之路”。<br>
而NIP自组内赛以来，一胜难求，以0-5垫底。只能寄希望于抢位赛杀出重围。<br>
欢迎收看啊，欢迎收看。<br>
<br>
#2024lpl夏季赛# &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img src="https://img.scoregg.com/m/5652938/p/247/2314302631187.png?imageMogr2/thumbnail/300x" data-v-24d379bc=""></span></a> <p class="index_left_weibo_info" data-v-24d379bc="">
          18小时前
          <a href="/p/1425811" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__comment" data-v-24d379bc=""></i></a> <a href="/p/1425811" class="flr" data-v-24d379bc=""><i class="gray_2 iconfont icon-circle__good_Gray" data-v-24d379bc=""></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/5620389" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/5620389/p/178/2113285249424_100x100.jpg" class="fll">
          米勒桑苏
        </a> <a data-v-24d379bc="" href="/p/1425798" class="index_left_weibo_main"><p data-v-24d379bc="">外设一套一套的 技术嘛[doge] &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img data-v-24d379bc="" src="https://img.scoregg.com/m/5620389/p/247/2308451331137.png?imageMogr2/thumbnail/300x"></span></a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425798" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425798" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/6328684" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/z/6328684/p/196/0317420053113_100x100.png" class="fll">
          WE_Mark
        </a> <a data-v-24d379bc="" href="/p/1425796" class="index_left_weibo_main"><p data-v-24d379bc="">http://t.cn/A68bv5Wy &ZeroWidthSpace;</p> </a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425796" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425796" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/6093262" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/6093262/p/185/2416175070672_100x100.jpg" class="fll">
          解说雨童
        </a> <a data-v-24d379bc="" href="/p/1425794" class="index_left_weibo_main"><p data-v-24d379bc="">换上了[doge]，差点忘了都 &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img data-v-24d379bc="" src="https://img.scoregg.com/m/6093262/p/247/2300003084735.png?imageMogr2/thumbnail/300x"></span></a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425794" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425794" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/5608387" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/5608387/p/185/2223592119050_100x100.jpg" class="fll">
          解说Remember
        </a> <a data-v-24d379bc="" href="/p/1425791" class="index_left_weibo_main"><p data-v-24d379bc="">还没睡觉想看比赛的兄弟，<br>
虎牙528222为大家直播<br>
LEC夏季胜者组决赛G2vsFNC(Bo5)[心]<br>
#remember记得[超话]# &ZeroWidthSpace;</p> </a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425791" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425791" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/6093203" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/6093203/p/185/2415545058895_100x100.jpg" class="fll">
          余霜
        </a> <a data-v-24d379bc="" href="/p/1425778" class="index_left_weibo_main"><p data-v-24d379bc="">最近超级爱的两个美甲💅🏻今天刚刚改成渐变粉[春游家族]<br>
<br>
每次做完美甲回家管哥都会从游戏里抽空出来评价一句。今天是：“我来看看你的指甲，嗯！粉粉的好好看哦”<br>
<br>
哈哈哈哈哈哈哈哈哈，情绪价值满分💯的对象，谁懂啊[赢牛奶]#余霜[超话]# &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img data-v-24d379bc="" src="https://img.scoregg.com/m/6093203/p/247/2222152078581.png?imageMogr2/thumbnail/300x"></span></a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425778" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425778" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/1993026" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/1993026/p/185/2819165150046_100x100.png" class="fll">
          1996zzr
        </a> <a data-v-24d379bc="" href="/p/1425710" class="index_left_weibo_main"><p data-v-24d379bc="">今天是大暑，一定要吃大薯条。小时候常听长辈说“大暑吃大薯，解馋又消暑。”我们老祖宗认为大暑是一年之中最热的节气，所以要注意消暑。点一份大薯，再把它们一根一根消灭，这就是所谓的“消薯”，这是我们的宝贵传统，要好好保护不能丟掉，不过现在很多年轻人已经不知道这层寓意了，有些遗憾。 &ZeroWidthSpace;</p> </a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425710" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425710" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/2860086" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/2860086/p/183/1914530369625_100x100.png" class="fll">
          _Cool
        </a> <a data-v-24d379bc="" href="/p/1425686" class="index_left_weibo_main"><p data-v-24d379bc="">看《泳者之心》想起《爆裂鼓手》 <br>
和人生的一部分很像<br>
#电影推荐一哈子# &ZeroWidthSpace;</p> </a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425686" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425686" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/951589" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/zhubo/951589/picture/15/9/28/2015092813225418438_100x100.jpg" class="fll">
          JoKer
        </a> <a data-v-24d379bc="" href="/p/1425681" class="index_left_weibo_main"><p data-v-24d379bc="">根本不需要担心延迟退休，能不能活那么久都不好说呢，万一川普上台直接发疯按钮一按[二哈] &ZeroWidthSpace;</p> </a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425681" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425681" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/951589" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/zhubo/951589/picture/15/9/28/2015092813225418438_100x100.jpg" class="fll">
          JoKer
        </a> <a data-v-24d379bc="" href="/p/1425677" class="index_left_weibo_main"><p data-v-24d379bc="">看完唐诡第二个案子，我对这部剧的评价要下降一个大档次。<br>
制作还是精良的，这没问题，但是剧本故事上的问题比第一个还要大，还要离谱。<br>
第一个案子看完我觉得最大的问题在于反派的装神弄鬼的合理性，cos个壁画里的鬼出来杀人，这个行为到底有什么实际意义？<br>
当然这里还勉强可以硬用反派就是喜欢装神弄 &ZeroWidthSpace;</p> </a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425677" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425677" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li><li data-v-24d379bc=""><a data-v-24d379bc="" href="/u/5620389" class="index_left_weibo_avatar"><img data-v-24d379bc="" src="https://img.scoregg.com/m/5620389/p/178/2113285249424_100x100.jpg" class="fll">
          米勒桑苏
        </a> <a data-v-24d379bc="" href="/p/1425668" class="index_left_weibo_main"><p data-v-24d379bc="">本周解说安排[赢牛奶]#2024lpl夏季赛# &ZeroWidthSpace;</p> <span data-v-24d379bc=""><img data-v-24d379bc="" src="https://img.scoregg.com/m/5620389/p/247/2214153691678.png?imageMogr2/thumbnail/300x"></span></a> <p data-v-24d379bc="" class="index_left_weibo_info">
          1天前
          <a data-v-24d379bc="" href="/p/1425668" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__comment"></i></a> <a data-v-24d379bc="" href="/p/1425668" class="flr"><i data-v-24d379bc="" class="gray_2 iconfont icon-circle__good_Gray"></i></a></p></li></ul>"""

# 确保图片保存文件夹存在
image_folder = './resource/wb_img'
if not os.path.exists(image_folder):
    os.makedirs(image_folder)

def download_and_save_image(image_src, save_path):
    response = requests.get(image_src)
    if response.status_code == 200:
        with open(save_path, 'wb') as file:
            file.write(response.content)
        print(f"图片下载成功并保存到 {save_path}")
    else:
        print(f"下载图片失败，状态码: {response.status_code}")

# 头像下载与保存函数
def download_and_save_avatar(avatar_src, save_path):
    response = requests.get(avatar_src)
    if response.status_code == 200:
        with open(save_path, 'wb') as file:
            file.write(response.content)
        print(f"头像下载成功并保存到 {save_path}")
    else:
        print(f"下载头像失败，状态码: {response.status_code}")


soup = BeautifulSoup(html_content, 'html.parser')

weibo_items = soup.find_all('li')

for item in weibo_items:
    user_avatar = item.find('a', class_='index_left_weibo_avatar').find('img')['src']
    user_name = item.find('a', class_='index_left_weibo_avatar').text.strip()
    weibo_content = item.find('a', class_='index_left_weibo_main').find('p').get_text(strip=True)
    post_time = item.find('p', class_='index_left_weibo_info').text.strip()
    print(post_time)
    try:
        image_src = item.find('a', class_='index_left_weibo_main').find('span').find('img')['src']
        image_name = os.path.basename(image_src.split('?')[0])  # 提取图片文件名
        image_save_path = os.path.join(image_folder, image_name)
        # download_and_save_image(image_src, image_save_path)
        image_src = "http://121.36.26.12:8081/resource/wb_img/"+image_name  # 修改图片链接为本地保存路径
    except AttributeError:
        image_src = None

    # 下载并保存头像
    avatar_name = f"{user_name}_avatar.jpg"  # 假设头像文件名基于用户名
    avatar_save_path = os.path.join(image_folder, avatar_name)  # 头像保存路径
    # download_and_save_avatar(user_avatar, avatar_save_path)
    avatar_save_path = "http://121.36.26.12:8081/resource/wb_img/"+avatar_name  # 修改图片链接为本地保存路径

    insert_query = """
    INSERT INTO weibo_data (user, content, time, image_src, avatar_src)
    VALUES (%s, %s, %s, %s, %s)
    """
    values = (user_name, weibo_content, post_time, image_src, avatar_save_path)
    cursor.execute(insert_query, values)

conn.commit()
conn.close()