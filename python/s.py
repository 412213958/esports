import requests

# 定义请求头
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 "
                  "Safari/537.36 Edg/125.0.0.0",
    "Accept-Language": "en-US,en;q=0.9"
}

# 示例用法
url = "https://api.bilibili.com/x/esports/component/contests/link?csrf=ea956ef0b38c43557fa560360ee2ff86&next=28466&sid=0"  # 替换为您要获取的网址




def get_html_content(url):
    try:
        # 发送请求，并传递请求头
        response = requests.get(url, headers=headers)
        # 如果请求成功（状态码 200）
        if response.status_code == 200:
            return response.text
        else:
            print(f"请求失败，状态码: {response.status_code}")
    except requests.exceptions.RequestException as e:
        print(f"请求时发生错误: {e}")


html_content = get_html_content(url)
if html_content:
    print(html_content)
