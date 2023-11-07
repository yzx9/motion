package util

/**
*@Description:
*@Author: BZ
*@date: 2023/11/3 16:41
*@Version: V1.0
 */

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func NewID() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, fmt.Errorf("fails to create id generator")
	}
	return node.Generate().Int64(), nil

}
