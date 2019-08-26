package tag_service

import (
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"

	"gin-blog/models"
	"gin-blog/pkg/export"
	"gin-blog/pkg/file"
	//"gin-blog/pkg/gredis"
)

type Tag struct {
	ID       int
	Name     string
	CreateBy int
	UpdateBy int
	State    int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreateBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["update_by"] = t.UpdateBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (a *Tag) Get() (*models.Tag, error) {
	//var cacheArticle *models.Article

	// cache := cache_service.Article{ID: a.ID}
	// key := cache.GetArticleKey()
	// if gredis.Exists(key) {
	// 	data, err := gredis.Get(key)
	// 	if err != nil {
	// 		logging.Info(err)
	// 	} else {
	// 		json.Unmarshal(data, &cacheArticle)
	// 		return cacheArticle, nil
	// 	}
	// }

	tag, err := models.GetTag(a.ID)
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, article, 3600)
	return tag, nil
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	// var (
	// 	tags, cacheTags []models.Tag
	// )

	// cache := cache_service.Tag{
	// 	State: t.State,

	// 	PageNum:  t.PageNum,
	// 	PageSize: t.PageSize,
	// }
	// key := cache.GetTagsKey()
	// if gredis.Exists(key) {
	// 	data, err := gredis.Get(key)
	// 	if err != nil {
	// 		logging.Info(err)
	// 	} else {
	// 		json.Unmarshal(data, &cacheTags)
	// 		return cacheTags, nil
	// 	}
	// }

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			strconv.Itoa(v.CreateBy),
			//strconv.Itoa(v.CreateTime),
			v.CreateTime.String(),
			strconv.Itoa(v.UpdateBy),
			//strconv.Itoa(v.UpdateTime),
			v.UpdateTime.String(),
		}

		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}

	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows, err := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			create_by, _ := strconv.Atoi(data[2])
			models.AddTag(data[1], 1, create_by)
		}
	}

	return nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_flag"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
