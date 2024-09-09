package model

type Matches struct {
	Id    int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title string `gorm:"column:title;default:NULL"`
	State string `gorm:"column:state;default:NULL"`
	Type  string `gorm:"column:type;default:NULL;comment:'1:英雄联盟，2:王者荣耀，3:第五人格，4:无畏契约，5:DOTA2，6:CSGO，7:绝地求生，8:和平精英，9:综合赛事，10:守望先锋'"`
}

func (m *Matches) TableName() string {
	return "matches"
}

type MatchInformation struct {
	Id             int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	MatchPartName  string `gorm:"column:match_part_name;default:NULL"`
	MatchPart1Name string `gorm:"column:match_part1_name;default:NULL"`
	Status         string `gorm:"column:status;default:NULL"`
	Team1          string `gorm:"column:team1;default:NULL"`
	Team1Score     int32  `gorm:"column:team1_score;default:NULL"`
	Team1ImgSrc    string `gorm:"column:team1_img_src;default:NULL"`
	Team2          string `gorm:"column:team2;default:NULL"`
	Team2Score     int32  `gorm:"column:team2_score;default:NULL"`
	Team2ImgSrc    string `gorm:"column:team2_img_src;default:NULL"`
	Time           string `gorm:"column:time;default:NULL"`
	VideoLink      string `gorm:"column:video_link;default:NULL"`
	DataLink       string `gorm:"column:data_link;default:NULL"`
	FormatName     string `gorm:"column:format_name;default:NULL"`
	Lp             string `gorm:"column:lp;default:NULL"`
	Lp1            string `gorm:"column:lp1;default:NULL"`
	Num1           string `gorm:"column:num1;default:NULL"`
}

func (m *MatchInformation) TableName() string {
	return "match_information"
}
