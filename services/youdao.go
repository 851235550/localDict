package services

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// https://github.com/creatcode/api/blob/master/YoudaoDic.md#associate
type YoudaoService struct {
	url         string
	jsonversion string
	xmlVersion  string
	dicts       string

	// maybe used in the future
	// client      string
	// keyFrom    string
	// model      string
	// mid        string
	// imei       string
	// vendor     string
	// screen     string
	// network    string
	// abTest     string
}

func NewYoudaoService() *YoudaoService {
	return &YoudaoService{
		url:         "https://dict.youdao.com/jsonapi",
		jsonversion: "2",
		xmlVersion:  "5.1",
		dicts:       "{\"count\":2,\"dicts\":[[\"ec\",\"ce\",\"simple\",\"wordform\",\"auth_sents_part\",\"ec21\",\"input\",\"rel_word\",\"ce_new\",\"blng_sents_part\"],[\"web_trans\"],[\"fanyi\"]]}",
	}
}

func (y *YoudaoService) buildQueryUrl(word string) string {
	u, err := url.Parse(y.url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	query := u.Query()
	query.Set("q", word)
	query.Set("jsonversion", y.jsonversion)
	query.Set("xmlVersion", y.xmlVersion)
	query.Set("dicts", y.dicts)

	u.RawQuery = query.Encode()

	return u.String()
}

type YoudaoResponse struct {
	Input string   `json:"input"`
	EC    YoudaoEC `json:"ec"`
}

type YoudaoEC struct {
	Word []YoudaoECWord `json:"word"`
}

type YoudaoECWord struct {
	Usphone string           `json:"usphone"`
	Ukphone string           `json:"ukphone"`
	Trs     []YoudaoECWordTr `json:"trs"`
	Wfs     []YoudaoECWordWf `json:"wfs"`
}

type YoudaoECWordTr struct {
	Tr []YoudaoECWordTrL `json:"tr"`
}

type YoudaoECWordTrL struct {
	L YoudaoECWordTrLI `json:"l"`
}

type YoudaoECWordTrLI struct {
	I []string `json:"i"`
}

type YoudaoECWordWf struct {
	Wf YoudaoECWordWfValue `json:"wf"`
}

type YoudaoECWordWfValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (y *YoudaoService) E2C(word string) {
	q := y.buildQueryUrl(word)
	resp, err := Get(q)
	if err != nil {
		panic(err)
	}

	youdaoResp := YoudaoResponse{}
	err = json.Unmarshal(resp, &youdaoResp)
	if err != nil {
		panic(err)
	}

	youdaoResp.Print()
}

func (r *YoudaoResponse) Print() {
	if len(r.EC.Word) == 0 {
		return
	}
	ECWord0 := r.EC.Word[0]
	fmt.Printf("美:[%s] 英:[%s]\n", ECWord0.Usphone, ECWord0.Ukphone)

	for _, tr := range ECWord0.Trs {
		for _, l := range tr.Tr {
			if len(l.L.I) == 0 {
				continue
			}
			fmt.Printf("%s\n", l.L.I[0])
		}
	}

	for _, wf := range ECWord0.Wfs {
		fmt.Printf("%s:%s  ", wf.Wf.Name, wf.Wf.Value)
	}
}
