package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	INVALID_JSON_PARAMS:            "请求参数非json",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:           "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_GET_TAGS_FAIL:            "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:           "统计标签失败",
	ERROR_ADD_TAG_FAIL:             "新增标签失败",
	ERROR_EDIT_TAG_FAIL:            "修改标签失败",
	ERROR_DELETE_TAG_FAIL:          "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:          "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:          "导入标签失败",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:         "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:      "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL: "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:        "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:       "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:        "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:         "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:  "生成文章海报失败",

	ERROR_NOT_EXIST_COMMENT:        "该评论不存在",
	ERROR_ADD_COMMENT_FAIL:         "新增评论失败",
	ERROR_DELETE_COMMENT_FAIL:      "删除评论失败",
	ERROR_CHECK_EXIST_COMMENT_FAIL: "检查评论是否存在失败",
	ERROR_EDIT_COMMENT_FAIL:        "修改评论失败",
	ERROR_COUNT_COMMENT_FAIL:       "统计评论失败",
	ERROR_GET_COMMENTS_FAIL:        "获取多个评论失败",
	ERROR_GET_COMMENT_FAIL:         "获取单个评论失败",
	ERROR_GEN_COMMENT_POSTER_FAIL:  "生成评论海报失败",

	ERROR_NOT_EXIST_CATEGORY:        "该分类不存在",
	ERROR_ADD_CATEGORY_FAIL:         "新增分类失败",
	ERROR_DELETE_CATEGORY_FAIL:      "删除分类失败",
	ERROR_CHECK_EXIST_CATEGORY_FAIL: "检查分类是否存在失败",
	ERROR_EDIT_CATEGORY_FAIL:        "修改分类失败",
	ERROR_COUNT_CATEGORY_FAIL:       "统计分类失败",
	ERROR_GET_CATEGORYS_FAIL:        "获取多个分类失败",
	ERROR_GET_CATEGORY_FAIL:         "获取单个分类失败",
	ERROR_GEN_CATEGORY_POSTER_FAIL:  "生成分类海报失败",
	ERROR_ALREADY_EXIST_CATEGORY:    "改分类已经存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
