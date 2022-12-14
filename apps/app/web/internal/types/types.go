// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Id         uint64 `json:"id"`          //用户ID
	Username   string `json:"username"`    //用户名
	Password   string `json:"password"`    //用户密码，MD5加密
	Phone      string `json:"phone"`       //手机号
	Question   string `json:"question"`    //找回密码问题
	Answer     string `json:"answer"`      //找回密码答案
	CreateTime int64  `json:"create_time"` //创建时间
	UpdateTime int64  `json:"update_time"` //更新时间
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo UserInfo `json:"userInfo"`
}

type HomeBannerResponse struct {
	Banners []*Banner `json:"banners"`
}

type Banner struct {
	ID   int64  `json:"id"`
	Name string `json:"name"` // 名称
	URL  string `json:"url"`  // 图片地址
}

type RecommendRequest struct {
	Cursor int64 `json:"cursor"`
	Ps     int64 `form:"ps,default=20"` // 每页大小
}

type RecommendResponse struct {
	Products      []*Product `json:"products"`
	IsEnd         bool       `json:"is_end"`         // 是否最后一页
	RecommendTime int64      `json:"recommend_time"` // 商品列表最后一个商品的推荐时间
}

type Product struct {
	ID          int64    `json:"id"`          // 商品ID
	Name        string   `json:"name"`        // 产品名称
	Images      []string `json:"images"`      // 图片
	Description string   `json:"description"` // 商品描述
	Price       float64  `json:"price"`       // 商品价格
	Stock       int64    `json:"stock"`       // 库存
	Category    string   `json:"category"`    // 分类
	Status      int64    `json:"status"`      // 状态：1-正常，2-下架
	CreateTime  int64    `json:"create_time"` // 创建时间
	UpdateTime  int64    `json:"update_time"` // 更新时间
}
