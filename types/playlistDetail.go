package types

// PlaylistDetailData 获取歌单详细 API 的返回数据
type PlaylistDetailData struct {
	RawJson       string
	Code          int `json:"code"`
	RelatedVideos struct {
	} `json:"relatedVideos"`
	Playlist struct {
		Id                    int           `json:"id"`
		Name                  string        `json:"name"`
		CoverImgId            int           `json:"coverImgId"`
		CoverImgUrl           string        `json:"coverImgUrl"`
		CoverImgIdStr         string        `json:"coverImgId_str"`
		AdType                int           `json:"adType"`
		UserId                int           `json:"userId"`
		CreateTime            int           `json:"createTime"`
		Status                int           `json:"status"`
		OpRecommend           bool          `json:"opRecommend"`
		HighQuality           bool          `json:"highQuality"`
		NewImported           bool          `json:"newImported"`
		UpdateTime            int           `json:"updateTime"`
		TrackCount            int           `json:"trackCount"`
		SpecialType           int           `json:"specialType"`
		Privacy               int           `json:"privacy"`
		TrackUpdateTime       int           `json:"trackUpdateTime"`
		CommentThreadId       string        `json:"commentThreadId"`
		PlayCount             int           `json:"playCount"`
		TrackNumberUpdateTime int           `json:"trackNumberUpdateTime"`
		SubscribedCount       int           `json:"subscribedCount"`
		CloudTrackCount       int           `json:"cloudTrackCount"`
		Ordered               bool          `json:"ordered"`
		Description           interface{}   `json:"description"`
		Tags                  []interface{} `json:"tags"`
		UpdateFrequency       interface{}   `json:"updateFrequency"`
		BackgroundCoverId     int           `json:"backgroundCoverId"`
		BackgroundCoverUrl    interface{}   `json:"backgroundCoverUrl"`
		TitleImage            int           `json:"titleImage"`
		TitleImageUrl         interface{}   `json:"titleImageUrl"`
		EnglishTitle          interface{}   `json:"englishTitle"`
		OfficialPlaylistType  interface{}   `json:"officialPlaylistType"`
		Subscribers           []struct {
			DefaultAvatar       bool        `json:"defaultAvatar"`
			Province            int         `json:"province"`
			AuthStatus          int         `json:"authStatus"`
			Followed            bool        `json:"followed"`
			AvatarUrl           string      `json:"avatarUrl"`
			AccountStatus       int         `json:"accountStatus"`
			Gender              int         `json:"gender"`
			City                int         `json:"city"`
			Birthday            int         `json:"birthday"`
			UserId              int         `json:"userId"`
			UserType            int         `json:"userType"`
			Nickname            string      `json:"nickname"`
			Signature           string      `json:"signature"`
			Description         string      `json:"description"`
			DetailDescription   string      `json:"detailDescription"`
			AvatarImgId         int         `json:"avatarImgId"`
			BackgroundImgId     int         `json:"backgroundImgId"`
			BackgroundUrl       string      `json:"backgroundUrl"`
			Authority           int         `json:"authority"`
			Mutual              bool        `json:"mutual"`
			ExpertTags          interface{} `json:"expertTags"`
			Experts             interface{} `json:"experts"`
			DjStatus            int         `json:"djStatus"`
			VipType             int         `json:"vipType"`
			RemarkName          interface{} `json:"remarkName"`
			AuthenticationTypes int         `json:"authenticationTypes"`
			AvatarDetail        *struct {
				UserType        int    `json:"userType"`
				IdentityLevel   int    `json:"identityLevel"`
				IdentityIconUrl string `json:"identityIconUrl"`
			} `json:"avatarDetail"`
			Anchor             bool   `json:"anchor"`
			AvatarImgIdStr     string `json:"avatarImgIdStr"`
			BackgroundImgIdStr string `json:"backgroundImgIdStr"`
			AvatarImgIdStr1    string `json:"avatarImgId_str"`
		} `json:"subscribers"`
		Subscribed bool `json:"subscribed"`
		Creator    struct {
			DefaultAvatar       bool        `json:"defaultAvatar"`
			Province            int         `json:"province"`
			AuthStatus          int         `json:"authStatus"`
			Followed            bool        `json:"followed"`
			AvatarUrl           string      `json:"avatarUrl"`
			AccountStatus       int         `json:"accountStatus"`
			Gender              int         `json:"gender"`
			City                int         `json:"city"`
			Birthday            int         `json:"birthday"`
			UserId              int         `json:"userId"`
			UserType            int         `json:"userType"`
			Nickname            string      `json:"nickname"`
			Signature           string      `json:"signature"`
			Description         string      `json:"description"`
			DetailDescription   string      `json:"detailDescription"`
			AvatarImgId         int         `json:"avatarImgId"`
			BackgroundImgId     int         `json:"backgroundImgId"`
			BackgroundUrl       string      `json:"backgroundUrl"`
			Authority           int         `json:"authority"`
			Mutual              bool        `json:"mutual"`
			ExpertTags          interface{} `json:"expertTags"`
			Experts             interface{} `json:"experts"`
			DjStatus            int         `json:"djStatus"`
			VipType             int         `json:"vipType"`
			RemarkName          interface{} `json:"remarkName"`
			AuthenticationTypes int         `json:"authenticationTypes"`
			AvatarDetail        interface{} `json:"avatarDetail"`
			Anchor              bool        `json:"anchor"`
			AvatarImgIdStr      string      `json:"avatarImgIdStr"`
			BackgroundImgIdStr  string      `json:"backgroundImgIdStr"`
			AvatarImgIdStr1     string      `json:"avatarImgId_str"`
		} `json:"creator"`
		Tracks []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
			Pst  int    `json:"pst"`
			T    int    `json:"t"`
			Ar   []struct {
				Id    int           `json:"id"`
				Name  string        `json:"name"`
				Tns   []interface{} `json:"tns"`
				Alias []interface{} `json:"alias"`
			} `json:"ar"`
			Alia []string    `json:"alia"`
			Pop  float64     `json:"pop"`
			St   int         `json:"st"`
			Rt   *string     `json:"rt"`
			Fee  int         `json:"fee"`
			V    int         `json:"v"`
			Crbt interface{} `json:"crbt"`
			Cf   string      `json:"cf"`
			Al   struct {
				Id     int      `json:"id"`
				Name   string   `json:"name"`
				PicUrl string   `json:"picUrl"`
				Tns    []string `json:"tns"`
				PicStr string   `json:"pic_str,omitempty"`
				Pic    int      `json:"pic"`
			} `json:"al"`
			Dt int `json:"dt"`
			H  *struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"h"`
			M *struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"m"`
			L *struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"l"`
			A                    interface{}   `json:"a"`
			Cd                   string        `json:"cd"`
			No                   int           `json:"no"`
			RtUrl                interface{}   `json:"rtUrl"`
			Ftype                int           `json:"ftype"`
			RtUrls               []interface{} `json:"rtUrls"`
			DjId                 int           `json:"djId"`
			Copyright            int           `json:"copyright"`
			SId                  int           `json:"s_id"`
			Mark                 int           `json:"mark"`
			OriginCoverType      int           `json:"originCoverType"`
			OriginSongSimpleData *struct {
				SongId  int    `json:"songId"`
				Name    string `json:"name"`
				Artists []struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"artists"`
				AlbumMeta struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"albumMeta"`
			} `json:"originSongSimpleData"`
			Single          int `json:"single"`
			NoCopyrightRcmd *struct {
				Type     int     `json:"type"`
				TypeDesc string  `json:"typeDesc"`
				SongId   *string `json:"songId"`
			} `json:"noCopyrightRcmd"`
			Mst         int         `json:"mst"`
			Cp          int         `json:"cp"`
			Mv          int         `json:"mv"`
			Rtype       int         `json:"rtype"`
			Rurl        interface{} `json:"rurl"`
			PublishTime int         `json:"publishTime"`
			VideoInfo   struct {
				MoreThanOne bool `json:"moreThanOne"`
				Video       *struct {
					Vid         string      `json:"vid"`
					Type        int         `json:"type"`
					Title       *string     `json:"title"`
					PlayTime    int         `json:"playTime"`
					CoverUrl    *string     `json:"coverUrl"`
					PublishTime int         `json:"publishTime"`
					Artists     interface{} `json:"artists"`
					Alias       interface{} `json:"alias"`
				} `json:"video"`
			} `json:"videoInfo"`
			Tns []string `json:"tns,omitempty"`
			Pc  struct {
				Nickname string `json:"nickname"`
				Alb      string `json:"alb"`
				Cid      string `json:"cid"`
				Fn       string `json:"fn"`
				Br       int    `json:"br"`
				Uid      int    `json:"uid"`
				Ar       string `json:"ar"`
				Version  int    `json:"version"`
				Sn       string `json:"sn"`
			} `json:"pc,omitempty"`
		} `json:"tracks"`
		VideoIds interface{} `json:"videoIds"`
		Videos   interface{} `json:"videos"`
		TrackIds []struct {
			Id         int         `json:"id"`
			V          int         `json:"v"`
			T          int         `json:"t"`
			At         int         `json:"at"`
			Alg        interface{} `json:"alg"`
			Uid        int         `json:"uid"`
			RcmdReason string      `json:"rcmdReason"`
		} `json:"trackIds"`
		ShareCount         int         `json:"shareCount"`
		CommentCount       int         `json:"commentCount"`
		RemixVideo         interface{} `json:"remixVideo"`
		SharedUsers        interface{} `json:"sharedUsers"`
		HistorySharedUsers interface{} `json:"historySharedUsers"`
	} `json:"playlist"`
	Urls       interface{} `json:"urls"`
	Privileges []struct {
		Id        int  `json:"id"`
		Fee       int  `json:"fee"`
		Paied     int  `json:"payed"`
		RealPaied int  `json:"realPayed"`
		St        int  `json:"st"`
		Pl        int  `json:"pl"`
		Dl        int  `json:"dl"`
		Sp        int  `json:"sp"`
		Cp        int  `json:"cp"`
		Subp      int  `json:"subp"`
		Cs        bool `json:"cs"`
		Maxbr     int  `json:"maxbr"`
		Fl        int  `json:"fl"`
		Pc        *struct {
			Id                  int    `json:"id"`
			UserId              int    `json:"userId"`
			SongId              int    `json:"songId"`
			Md5                 string `json:"md5"`
			Song                string `json:"song"`
			Artist              string `json:"artist"`
			Album               string `json:"album"`
			Bitrate             int    `json:"bitrate"`
			FileName            string `json:"fileName"`
			SongDfsId           int    `json:"songDfsId"`
			Cover               int    `json:"cover"`
			Lyric               int    `json:"lyric"`
			Cue                 int    `json:"cue"`
			ConvertLyric        int    `json:"convertLyric"`
			Version             int    `json:"version"`
			AddTime             int    `json:"addTime"`
			FileSize            int    `json:"fileSize"`
			Status              int    `json:"status"`
			OriginalAudioSongId int    `json:"originalAudioSongId"`
			LrcType             string `json:"lrcType"`
		} `json:"pc"`
		Toast              bool        `json:"toast"`
		Flag               int         `json:"flag"`
		PaidBigBang        bool        `json:"paidBigBang"`
		PreSell            bool        `json:"preSell"`
		PlayMaxbr          int         `json:"playMaxbr"`
		DownloadMaxbr      int         `json:"downloadMaxbr"`
		Rscl               interface{} `json:"rscl"`
		FreeTrialPrivilege struct {
			ResConsumable  bool `json:"resConsumable"`
			UserConsumable bool `json:"userConsumable"`
		} `json:"freeTrialPrivilege"`
		ChargeInfoList []struct {
			Rate          int         `json:"rate"`
			ChargeUrl     interface{} `json:"chargeUrl"`
			ChargeMessage interface{} `json:"chargeMessage"`
			ChargeType    int         `json:"chargeType"`
		} `json:"chargeInfoList"`
	} `json:"privileges"`
	SharedPrivilege interface{} `json:"sharedPrivilege"`
	ResEntrance     interface{} `json:"resEntrance"`
}
