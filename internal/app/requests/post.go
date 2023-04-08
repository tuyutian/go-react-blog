package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type PostRequest struct {
}

func (r *PostRequest) Messages() govalidator.MapData {
	message := govalidator.MapData{
		"title":    []string{"required:标题不能为空", "between:标题大小为1到150字符"},
		"content":  []string{"required:内容不能为空"},
		"subtitle": []string{"between:副标题大小为1到150字符"},
	}
	return message
}

func (r *PostRequest) Rules() govalidator.MapData {
	rules := govalidator.MapData{
		"title":    []string{"require", "between:1,150"},
		"subtitle": []string{"between:1,150"},
		"content":  []string{"required"},
	}
	return rules
}

type DetailRequest struct {
}

func (r *DetailRequest) Messages() govalidator.MapData {
	message := govalidator.MapData{
		"id": []string{"required:详情id不能为空", "numeric:必须是数字类型"},
	}
	return message
}

func (r *DetailRequest) Rules() govalidator.MapData {
	rules := govalidator.MapData{
		"id": []string{"require", "numberic"},
	}
	return rules
}
