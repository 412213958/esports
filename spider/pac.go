package spider

import (
	"encoding/json"
	"esports/utils"
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func downloadAndConvertImage(url string) string {
	// 获取文件夹路径
	folderPath := "resource/wb_img"
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return ""
	}

	// 发送 HTTP GET 请求获取图片数据
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return ""
	}
	defer resp.Body.Close()

	// 获取文件名
	parts := strings.Split(url, "?")[0]
	fileName := filepath.Base(parts)
	filePath := filepath.Join(folderPath, fileName)

	// 创建一个文件来保存图片
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	defer file.Close()

	// 将图片数据写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return ""
	}

	// 转换链接格式
	newUrl := fmt.Sprintf("http://121.36.26.12:8081/resource/wb_img/%s", fileName)
	return newUrl
}

func downloadAndConvertImage1(url string) string {
	// 获取文件夹路径
	folderPath := "resource/wb_img"
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return ""
	}

	// 发送 HTTP GET 请求获取图片数据
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return ""
	}
	defer resp.Body.Close()
	// 获取文件名
	parts := strings.Split(url, "?")[0]
	// 构建文件路径
	fileName := filepath.Base(parts)
	filePath := filepath.Join(folderPath, fileName)

	// 创建一个文件来保存图片
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	defer file.Close()

	// 将图片数据写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return ""
	}

	// 转换链接格式
	newUrl := fmt.Sprintf("http://121.36.26.12:8081/resource/wb_img/%s", fileName)
	return newUrl

}

func pac(doc *goquery.Document, db *gorm.DB) {
	// comments 表
	{
		var tableNames []string
		result := db.Raw("SHOW TABLES").Scan(&tableNames)
		if result.Error != nil {
			fmt.Println("Error checking table existence for comments:", result.Error)
			return
		}

		tableExists := false
		for _, tableName := range tableNames {
			if tableName == "comments" {
				tableExists = true
				break
			}
		}

		// 如果表存在，清空表
		if tableExists {
			result = db.Exec("TRUNCATE TABLE comments")
			if result.Error != nil {
				fmt.Println("Error truncating table comments:", result.Error)
				return
			}
		}

		// 创建 comments 表（如果不存在）
		result = db.Exec(`
        CREATE TABLE IF NOT EXISTS comments (
            id INT PRIMARY KEY AUTO_INCREMENT,
            user_avatar VARCHAR(255),
            user_name VARCHAR(255),
            post_time VARCHAR(255),
            comment_title VARCHAR(255),
            comment_reply VARCHAR(255),
            lp VARCHAR(255)
        )
        `)
		if result.Error != nil {
			fmt.Println("Error creating table comments:", result.Error)
			return
		}
	}

	// match_data 表
	{
		var tableNames []string
		result := db.Raw("SHOW TABLES").Scan(&tableNames)
		if result.Error != nil {
			fmt.Println("Error checking table existence for match_data:", result.Error)
			return
		}

		tableExists := false
		for _, tableName := range tableNames {
			if tableName == "match_data" {
				tableExists = true
				break
			}
		}

		// 如果表存在，清空表
		if tableExists {
			result = db.Exec("TRUNCATE TABLE match_data")
			if result.Error != nil {
				fmt.Println("Error truncating table match_data:", result.Error)
				return
			}
		}

		// 创建 match_data 表（如果不存在）
		result = db.Exec(`
        CREATE TABLE IF NOT EXISTS match_data (
            id INT AUTO_INCREMENT PRIMARY KEY,
            match_info VARCHAR(255),
            image_src VARCHAR(255),
            deadline_time VARCHAR(50),
            team1 VARCHAR(50),
            team1_odds DECIMAL(10, 2),
            team2 VARCHAR(50),
            team2_odds DECIMAL(10, 2),
            lp VARCHAR(255)
        )
        `)
		if result.Error != nil {
			fmt.Println("Error creating table match_data:", result.Error)
			return
		}
	}

	// articles 表
	{
		var tableNames []string
		result := db.Raw("SHOW TABLES").Scan(&tableNames)
		if result.Error != nil {
			fmt.Println("Error checking table existence for articles:", result.Error)
			return
		}

		tableExists := false
		for _, tableName := range tableNames {
			if tableName == "articles" {
				tableExists = true
				break
			}
		}

		// 如果表存在，清空表
		if tableExists {
			result = db.Exec("TRUNCATE TABLE articles")
			if result.Error != nil {
				fmt.Println("Error truncating table articles:", result.Error)
				return
			}
		}

		// 创建 articles 表（如果不存在）
		result = db.Exec(`
        CREATE TABLE IF NOT EXISTS articles (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255),
            description TEXT,
            image_link VARCHAR(255),
            lp VARCHAR(255)
        )
        `)
		if result.Error != nil {
			fmt.Println("Error creating table articles:", result.Error)
			return
		}
	}

	// weibo_data 表
	{
		var tableNames []string
		result := db.Raw("SHOW TABLES").Scan(&tableNames)
		if result.Error != nil {
			fmt.Println("Error checking table existence for weibo_data:", result.Error)
			return
		}

		tableExists := false
		for _, tableName := range tableNames {
			if tableName == "weibo_data" {
				tableExists = true
				break
			}
		}

		// 如果表存在，清空表
		if tableExists {
			result = db.Exec("TRUNCATE TABLE weibo_data")
			if result.Error != nil {
				fmt.Println("Error truncating table weibo_data:", result.Error)
				return
			}
		}

		// 创建 weibo_data 表（如果不存在）
		result = db.Exec(`
        CREATE TABLE IF NOT EXISTS weibo_data (
            id INT AUTO_INCREMENT PRIMARY KEY,
            user VARCHAR(255),
            content TEXT,
            time VARCHAR(255),
            image_src VARCHAR(255),
            avatar_src VARCHAR(255),
            lp VARCHAR(255)
        )
        `)
		if result.Error != nil {
			fmt.Println("Error creating table weibo_data:", result.Error)
			return
		}
	}

	// tournaments 表
	{
		var tableNames []string
		result := db.Raw("SHOW TABLES").Scan(&tableNames)
		if result.Error != nil {
			fmt.Println("Error checking table existence:", result.Error)
			return
		}

		tableExists := false
		for _, tableName := range tableNames {
			if tableName == "tournaments" {
				tableExists = true
				break
			}
		}

		// 如果表存在，清空表
		if tableExists {
			result = db.Exec("TRUNCATE TABLE tournaments")
			if result.Error != nil {
				fmt.Println("Error truncating table:", result.Error)
				return
			}
		}

		// 创建表（如果不存在）
		result = db.Exec(`
        CREATE TABLE IF NOT EXISTS tournaments (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255),
            image_src VARCHAR(255),
            type_s VARCHAR(255),
            lp VARCHAR(255)
        )
        `)
		if result.Error != nil {
			fmt.Println("Error creating table:", result.Error)
			return
		}
	}

	// tournaments 表提取数据
	{
		// 筛选出 title="热门赛事" 的 div
		热门赛事Div := doc.Find("div[title='热门赛事']")
		// 提取赛事信息
		热门赛事Div.Find(".item").Each(func(i int, s *goquery.Selection) {
			linkAndTitle := s.Find("a[title]")
			imageUrl := s.Find("img").AttrOr("src", "")
			if linkAndTitle.Length() > 0 && imageUrl != "" {
				title := linkAndTitle.AttrOr("title", "")

				// 下载并转换图片链接
				newImageUrl := downloadAndConvertImage(imageUrl)

				// 保存数据到数据库
				result := db.Exec("INSERT INTO tournaments (title, image_src, type_s) VALUES (?,?,?)", title, newImageUrl, "热门赛事")
				if result.Error != nil {
					fmt.Println("Error inserting into database:", result.Error)
					return
				}
			}
		})
		热门圈子Div := doc.Find("div[title='热门圈子']")
		// 提取赛事信息
		热门圈子Div.Find(".item").Each(func(i int, s *goquery.Selection) {
			linkAndTitle := s.Find("a[title]")
			imageUrl := s.Find("img").AttrOr("src", "")
			if linkAndTitle.Length() > 0 && imageUrl != "" {
				title := linkAndTitle.AttrOr("title", "")

				// 下载并转换图片链接
				newImageUrl := downloadAndConvertImage1(imageUrl)

				// 保存数据到数据库
				result := db.Exec("INSERT INTO tournaments (title, image_src, type_s) VALUES (?,?,?)", title, newImageUrl, "热门圈子")
				if result.Error != nil {
					fmt.Println("Error inserting into database:", result.Error)
					return
				}
			}
		})
	}

	// weibo_data 表提取数据
	{
		// 筛选出 <li data-v-24d379bc> 的元素
		liElements := doc.Find("li[data-v-24d379bc]")

		liElements.Each(func(i int, s *goquery.Selection) {
			userAvatar, _ := s.Find("a.index_left_weibo_avatar img").Attr("src")
			imgSrc, _ := s.Find("a.index_left_weibo_main img").Attr("src")
			userName := strings.TrimSpace(s.Find("a.index_left_weibo_avatar").Text())
			weiboContent := strings.TrimSpace(s.Find("a.index_left_weibo_main p").Text())
			postTime := strings.TrimSpace(s.Find("p.index_left_weibo_info").Text())

			if userAvatar != "" {
				// 下载并转换图片链接
				userAvatar = downloadAndConvertImage(userAvatar)

			}
			if imgSrc != "" {
				// 下载并转换图片链接
				imgSrc = downloadAndConvertImage1(imgSrc)

			}

			// 保存数据到数据库
			result := db.Exec("INSERT INTO weibo_data (user, content, time, avatar_src, image_src) VALUES (?,?,?,?,?)", userName, weiboContent, postTime, userAvatar, imgSrc)
			if result.Error != nil {
				fmt.Println("Error inserting into database:", result.Error)
				return
			}
		})
	}

	// articles 表提取数据
	{
		// 筛选出 <li data-v-3713f460> 的元素
		liElements := doc.Find("li[data-v-3713f460]")

		liElements.Each(func(i int, s *goquery.Selection) {
			title := strings.TrimSpace(s.Find("h2").Text())
			title = title[6:]
			description := strings.TrimSpace(s.Find("p.gray_3").Text())
			image, _ := s.Find("img").Attr("src")

			if image != "" {
				// 下载并转换图片链接
				image = downloadAndConvertImage1(image)

			}

			// 保存数据到数据库
			result := db.Exec("INSERT INTO articles (title, description, image_link) VALUES (?,?,?)", title, description, image)
			if result.Error != nil {
				fmt.Println("Error inserting into database:", result.Error)
				return
			}
		})
	}

	// comments 表提取数据
	{
		ulElements := doc.Find("ul.index_right_comment")
		ulElements.Find("li").Each(func(i int, s *goquery.Selection) {
			avatarSrc, _ := s.Find("div.index_right_comment_info a.index_right_comment_avatar img").Attr("src")
			userName := strings.TrimSpace(s.Find("div.index_right_comment_info p.index_right_comment_name a").Text())
			postTime := strings.TrimSpace(s.Find("div.index_right_comment_info p.index_right_comment_name span").Text())
			commentTitle := strings.TrimSpace(s.Find("a.index_right_comment_main p.index_right_comment_title").Text())
			commentReply := strings.TrimSpace(s.Find("a.index_right_comment_main p.index_right_comment_reply").Text())

			if avatarSrc != "" {
				// 下载并转换图片链接
				avatarSrc = downloadAndConvertImage(avatarSrc)

			}

			result := db.Exec("INSERT INTO comments (user_avatar, user_name, post_time, comment_title, comment_reply) VALUES (?,?,?,?,?)", avatarSrc, userName, postTime, commentTitle, commentReply)
			if result.Error != nil {
				fmt.Println("Error inserting into database:", result.Error)
				return
			}
		})
	}
}

type AutoGenerated struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Count string `json:"count"`
		List  []struct {
			V                string `json:"v"`
			Key              string `json:"key"`
			NewDay           string `json:"new_day"`
			DateTxt          string `json:"date_txt"`
			ViewType         int    `json:"view_type"`
			TournamentID     string `json:"tournament_id"`
			TournamentName   string `json:"tournament_name"`
			TournamentEnName string `json:"tournament_en_name"`
			TournamentImage  string `json:"tournament_image"`
			RoundName        string `json:"round_name"`
			RoundSonName     string `json:"round_son_name"`
			MatchID          string `json:"match_id"`
			TeamImageThumbA  string `json:"team_image_thumb_a"`
			TeamImageThumbB  string `json:"team_image_thumb_b"`
			MatchDate        string `json:"match_date"`
			MatchDate1       string `json:"match_date1"`
			MatchTeamA       string `json:"match_team_a"`
			MatchTeamB       string `json:"match_team_b"`
			GameCount        string `json:"game_count"`
			StartID          string `json:"start_id"`
			BetList          []struct {
				BetID         string `json:"bet_id"`
				Title         string `json:"title"`
				CategoryName  string `json:"category_name"`
				BetEndTime    string `json:"bet_end_time"`
				BetEndTimeTxt string `json:"bet_end_time_txt"`
				Status        string `json:"status"`
				StatusTxt     string `json:"status_txt"`
				TotalPrice    string `json:"total_price"`
				PeopleNum     string `json:"people_num"`
				ResultItemID  string `json:"result_item_id"`
				DynamicID     int    `json:"dynamic_id"`
				Items         []struct {
					BetItemID    string `json:"bet_item_id"`
					InitPrice    string `json:"init_price"`
					Price        string `json:"price"`
					MemberMaxBet string `json:"member_max_bet"`
					ItemName     string `json:"item_name"`
					ItemNameEn   string `json:"item_name_en"`
					ItemNameTw   string
					WinRate      string `json:"win_rate"`
					Odds         string `json:"odds"`
				} `json:"items"`
				TeamAWin      string `json:"team_a_win"`
				TeamBWin      string `json:"team_b_win"`
				MatchStatus   string `json:"match_status"`
				MatchBetCount string `json:"match_bet_count"`
			} `json:"bet_list"`
		} `json:"list"`
	} `json:"data"`
	TaskData struct{}      `json:"task_data"`
	Badge    []interface{} `json:"badge"`
	Event    []interface{} `json:"event"`
}

func adjustURLStartScore(url string) string {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	newStartScore := fmt.Sprintf("%d%02d%02d", year, month, day)
	return fmt.Sprintf("%s&start_score=%s", url, newStartScore)
}

func performDatabaseOperations(data AutoGenerated, doc *goquery.Document) {
	pac(doc, utils.DB)

	url := "https://www.scoregg.com/services/bet/bet_single_list.php?type=1&tournament_id=0"

	url = adjustURLStartScore(url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 使用 GORM 创建和操作表
	// 检查表 1 是否存在
	var tableNames []string
	result := utils.DB.Raw("SHOW TABLES").Scan(&tableNames)
	if result.Error != nil {
		fmt.Println("Error checking table 1 existence:", result.Error)
		return
	}
	var table1Exists bool
	for _, tableName := range tableNames {
		if tableName == "table1" {
			table1Exists = true
			break
		}
	}
	// 如果表存在，清空表 1 的数据
	if table1Exists {
		result = utils.DB.Exec("TRUNCATE TABLE table1")
		if result.Error != nil {
			fmt.Println("Error truncating table 1:", result.Error)
			return
		}
	}
	// 创建表 1（修改了 count 字段名为 bet_count ）
	if !table1Exists {
		err := utils.DB.Table("table1").AutoMigrate(&struct {
			V                string
			BetKey           string
			NewDay           string
			DateTxt          string
			ViewType         int
			TournamentID     string
			TournamentName   string
			TournamentEnName string
			TournamentImage  string
			RoundName        string
			RoundSonName     string
			MatchID          string
			TeamImageThumbA  string
			TeamImageThumbB  string
			MatchDate        string
			MatchDate1       string
			MatchTeamA       string
			MatchTeamB       string
			GameCount        string
			StartID          string
		}{})
		if err != nil {
			fmt.Println("Error creating table 1:", err)
			return
		}
	}

	// 检查表 2 是否存在
	var table2Exists bool
	for _, tableName := range tableNames {
		if tableName == "table2" {
			table2Exists = true
			break
		}
	}
	// 如果表存在，清空表 2 的数据
	if table2Exists {
		result = utils.DB.Exec("TRUNCATE TABLE table2")
		if result.Error != nil {
			fmt.Println("Error truncating table 2:", result.Error)
			return
		}
	}
	// 创建表 2
	if !table2Exists {
		err := utils.DB.Table("table2").AutoMigrate(&struct {
			BetID         string
			Title         string
			CategoryName  string
			BetEndTime    string
			BetEndTimeTxt string
			Status        string
			StatusTxt     string
			TotalPrice    string
			PeopleNum     string
			ResultItemID  string
			DynamicID     int
		}{})
		if err != nil {
			fmt.Println("Error creating table 2:", err)
			return
		}
	}

	// 检查表 3 是否存在
	var table3Exists bool
	for _, tableName := range tableNames {
		if tableName == "table3" {
			table3Exists = true
			break
		}
	}
	// 如果表存在，清空表 3 的数据
	if table3Exists {
		result = utils.DB.Exec("TRUNCATE TABLE table3")
		if result.Error != nil {
			fmt.Println("Error truncating table 3:", result.Error)
			return
		}
	}
	// 创建表 3
	if !table3Exists {
		err := utils.DB.Table("table3").AutoMigrate(&struct {
			BetItemID     string
			InitPrice     string
			Price         string
			MemberMaxBet  string
			ItemName      string
			ItemNameEn    string
			ItemNameTw    string
			WinRate       string
			Odds          string
			TeamAWin      string
			TeamBWin      string
			MatchStatus   string
			MatchBetCount string
		}{})
		if err != nil {
			fmt.Println("Error creating table 3:", err)
			return
		}
	}

	// 插入数据到第一张表
	for _, item := range data.Data.List {
		item.TournamentImage = downloadAndConvertImage(item.TournamentImage)
		item.TeamImageThumbA = downloadAndConvertImage(item.TeamImageThumbA)
		item.TeamImageThumbB = downloadAndConvertImage(item.TeamImageThumbB)
		result := utils.DB.Table("table1").Create(&struct {
			V                string
			BetKey           string
			NewDay           string
			DateTxt          string
			ViewType         int
			TournamentID     string
			TournamentName   string
			TournamentEnName string
			TournamentImage  string
			RoundName        string
			RoundSonName     string
			MatchID          string
			TeamImageThumbA  string
			TeamImageThumbB  string
			MatchDate        string
			MatchDate1       string
			MatchTeamA       string
			MatchTeamB       string
			GameCount        string
			StartID          string
		}{
			V:                item.V,
			BetKey:           item.Key,
			NewDay:           item.NewDay,
			DateTxt:          item.DateTxt,
			ViewType:         item.ViewType,
			TournamentID:     item.TournamentID,
			TournamentName:   item.TournamentName,
			TournamentEnName: item.TournamentEnName,
			TournamentImage:  item.TournamentImage,
			RoundName:        item.RoundName,
			RoundSonName:     item.RoundSonName,
			MatchID:          item.MatchID,
			TeamImageThumbA:  item.TeamImageThumbA,
			TeamImageThumbB:  item.TeamImageThumbB,
			MatchDate:        item.MatchDate,
			MatchDate1:       item.MatchDate1,
			MatchTeamA:       item.MatchTeamA,
			MatchTeamB:       item.MatchTeamB,
			GameCount:        item.GameCount,
			StartID:          item.StartID,
		})
		if result.Error != nil {
			fmt.Printf("Error inserting data for table 1: %v\n", result.Error)
			continue
		}
		if result.RowsAffected == 0 {
			fmt.Println("No rows were affected for table 1")
		}
	}

	// 插入数据到第二张表
	for _, item := range data.Data.List {
		for _, bet := range item.BetList {
			result := utils.DB.Table("table2").Create(&struct {
				BetID         string
				Title         string
				CategoryName  string
				BetEndTime    string
				BetEndTimeTxt string
				Status        string
				StatusTxt     string
				TotalPrice    string
				PeopleNum     string
				ResultItemID  string
				DynamicID     int
			}{
				BetID:         bet.BetID,
				Title:         bet.Title,
				CategoryName:  bet.CategoryName,
				BetEndTime:    bet.BetEndTime,
				BetEndTimeTxt: bet.BetEndTimeTxt,
				Status:        bet.Status,
				StatusTxt:     bet.StatusTxt,
				TotalPrice:    bet.TotalPrice,
				PeopleNum:     bet.PeopleNum,
				ResultItemID:  bet.ResultItemID,
				DynamicID:     bet.DynamicID,
			})
			if result.Error != nil {
				fmt.Printf("Error inserting data for table 2: %v\n", result.Error)
				continue
			}
			if result.RowsAffected == 0 {
				fmt.Println("No rows were affected for table 2")
			}
		}
	}

	// 插入数据到第三张表
	for _, item := range data.Data.List {
		for _, bet := range item.BetList {
			for _, item := range bet.Items {
				result := utils.DB.Table("table3").Create(&struct {
					BetItemID     string
					InitPrice     string
					Price         string
					MemberMaxBet  string
					ItemName      string
					ItemNameEn    string
					ItemNameTw    string
					WinRate       string
					Odds          string
					TeamAWin      string
					TeamBWin      string
					MatchStatus   string
					MatchBetCount string
				}{
					BetItemID:     item.BetItemID,
					InitPrice:     item.InitPrice,
					Price:         item.Price,
					MemberMaxBet:  item.MemberMaxBet,
					ItemName:      item.ItemName,
					ItemNameEn:    item.ItemNameEn,
					ItemNameTw:    item.ItemNameTw,
					WinRate:       item.WinRate,
					Odds:          item.Odds,
					TeamAWin:      bet.TeamAWin,
					TeamBWin:      bet.TeamBWin,
					MatchStatus:   bet.MatchStatus,
					MatchBetCount: bet.MatchBetCount,
				})
				if result.Error != nil {
					fmt.Printf("Error inserting data for table 3: %v\n", result.Error)
					continue
				}
				if result.RowsAffected == 0 {
					fmt.Println("No rows were affected for table 3")
				}
			}
		}
	}

	fmt.Println("Data insertion completed!")
}

func Transcation() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.scoregg.com/", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0")

	// 设置 Cookie
	cookieStr := "FAMULEIMEMBER=2c187cce22614b26681b7d8e5f56d2d3; member_id=6568946; f_token=ef4698d0aa34450dbc94a3587901d296; PHPSESSID=g0jn70df9mhciiom1bmvthq4q6; gameId=1; Hm_lvt_7d06c05f2674b5f5fffa6500f4e9da89=1721782657,1721803648,1721868732,1721954262; HMACCOUNT=550D0CC3EA12292D; Hm_lpvt_7d06c05f2674b5f5fffa6500f4e9da89=1721963763"
	req.Header.Add("Cookie", cookieStr)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	htmlContent := string(body)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	var data AutoGenerated
	performDatabaseOperations(data, doc)
}
