package git

import (
	"fmt"
	"testing"
)

func TestGitClone(t *testing.T) {
	s := Clone("ssh://git@gitlab.hellobike.cn:10022/MopedApp/JYBaseModel.git")
	if s.Error != nil {
		fmt.Println(s.Error)
	}

}
