package service

import (
	"errors"
	"github.com/yzx9/motion/command/domain/video/adapter/dto"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 19:52
*@Version: V1.0
 */
var collectDao = PO.NewCollectDB()

type CollectService interface {
	Collect(dto *dto.CollectDto) error
}

type CollectServiceImpl struct {
}

func (service *CollectServiceImpl) Collect(dto *dto.CollectDto) error {
	if dto == nil {
		return errors.New("无点赞信息")
	}
	var collectVideo PO.CollectVideo
	if err := dto.ToCollect(&collectVideo); err != nil {
		return err
	}
	return collectDao.InsertCollectVideo(&collectVideo)
}
