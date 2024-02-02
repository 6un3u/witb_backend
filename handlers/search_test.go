package handlers_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/6un3u/witb_backend/handlers"
	"github.com/6un3u/witb_backend/utils"
	"github.com/joho/godotenv"
)

func TestMakeSearchResult(t *testing.T) {
	err := godotenv.Load("../.env")
	utils.HandleErr(err)

	books := handlers.MakeSearchResult("희랍어 시간")

	if len(books) == 0 {
		t.Error("There isn't any Result")
	}
}

func TestExtractBookInfo(t *testing.T) {
	searchText := `
      {
        "CMDTCODE": "9788954682152",
        "CMDT_NAME": "작별하지 않는다",
        "SALE_CMDTID": "S000000781116",
        "SALE_CMDT_DVSN_CODE_SUB": "KOR",
        "SCORE": "102040",
        "RELATE_HTML_LIST": "",
        "SALE_CMDT_GRP_DVSN_CODE": "SGK",
        "DQ_ID": "TOT_S000000781116",
        "SALE_CMDT_DVSN_CODE": "TOT",
        "TOT_RELATE_HTML_LIST": "9788954682152$@소설$@작별하지 않는다$@한강$@문학동네$@2021$@09$@14000.00$@12600.00$@10.00$@700$@N$@N$@KOR$@https://contents.kyobobook.co.kr/sih/fit-in/200x0/pdt/9788954682152.jpg$@0$@0$@0$@0$@0$@5.00$@2016년 『채식주의자』로 인터내셔널 부커상을 수상하고 2018년 『흰』으로 같은 상 최종 후보에 오른 한강 작가의 5년 만의 신작 장편소설 『작별하지 않는다』가 출간되었다. 2019년 겨울부터 이듬해 봄까지 계간 『문학동네』에 전반부를 연재하면서부터 큰 관심을 모았고, 그뒤 일 년여에 걸쳐 후반부를 집필하고 또 전체를 공들여 다듬는 지난한 과정을 거쳐 완성되었다. 본래 「눈 한 송이가 녹는 동안」(2015년 황순원문학상 수상작), 「작별」(2018년 김유정문학상 수상작)을 잇는 ‘눈’ 3부작의 마지막 작품으로 구상되었으나 그 자체 완결된 작품의 형태로 엮이게 된바, 한강 작가의 문학적 궤적에서 『작별하지 않는다』가 지니는 각별한 의미를 짚어볼 수 있다. 이로써 『소년이 온다』(2014), 『흰』(2016), ‘눈’ 연작(2015, 2017) 등 근작들을 통해 어둠 속에서도 한줄기 빛을 향해 나아가는 인간의 고투와 존엄을 그려온 한강 문학이 다다른 눈부신 현재를 또렷한 모습으로 확인할 수 있게 되었다. 오래지 않은 비극적 역사의 기억으로부터 길어올린, 그럼에도 인간을 끝내 인간이게 하는 간절하고 지극한 사랑의 이야기가 눈이 시리도록 선연한 이미지와 유려하고 시적인 문장에 실려 압도적인 아름다움으로 다가온다.$@$@0$@0$@0$@010101"
      }
	`
	var resultdoc handlers.ResultDocument
	err := json.Unmarshal([]byte(searchText), &resultdoc)
	utils.HandleErr(err)
	result := handlers.ExtractBookInfo(resultdoc)
	log.Println(result)
}
