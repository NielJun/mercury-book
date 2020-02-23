package filter

import (
	"testing"
)

func TestReplace(t *testing.T) {

	err:= Init("./data/filter.dat.txt")
	if err != nil {
		t.Errorf("敏感词库文件加载失败 %#v",err)
	}

	dataStr := `金澤文子王八蛋看见女司机带你飞离开谁都看被逼是打开裸体暗示无毛的博客把叫床，叫阿巴斯嗲随便丢吧淫荡那可是当男生看到就能看到看那快结束的`
	result,isReplaced :=  Replace(dataStr,"***")
	if isReplaced == false {
		return
	}
	t.Logf(result)
}
