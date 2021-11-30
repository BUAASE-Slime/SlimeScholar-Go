package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"

	"gitee.com/online-publish/slime-scholar-go/service"
	"github.com/olivere/elastic/v7"

	"golang.org/x/net/context"
	"io"
	"os"
)

const AUTHOR_DIR = "H:\\Scholar"
const PAPER_DIR = "E:\\Paper"
const FILE_NUM = 3
const AUTHOR_FILE_PREFIX = "aminer_authors_"
const PAPER_FILE_PREFIX = "s2-corpus-"
const BULK_SIZE = 100000

var fieldsMap map[string]int = make(map[string]int)
var success_num, fail_num = 0, 0

var max_citation_num = 0   // 看一下所有论文的最大引用数目
var max_references_num = 0 //
type pub struct {
	id           string `json:"id"`
	author_order int    `json:"author_order"`
}

type Author struct {
	id         string   `json:"id"`
	name       string   `json:"name"`
	n_pubs     int      `json:"n_pubs"`
	n_citation int      `json:"n_citation"`
	h_index    int      `json:"h_index"`
	papers     []string `json:"papers"`
}

type Paper struct {
	Id            string   `json:"id"`
	Title         string   `json:"title"`
	Abstract      string   `json:"abstract"`
	Url           string   `json:"url"`
	PdfUrls       []string `json:"pdf_urls"`
	S2PdfUrl      string   `json:"s2pdf_urls"`
	InCitations   []string `json:"in_citations"`
	OutCitations  []string `json:"out_citations"`
	FieldsOfStudy []string `json:"study_fields"`
	Year          int      `json:"year"`
	Venue         string   `json:"venue"`
	JournalName   string   `json:"journal_name"`
	JournalVolume string   `json:"journal_volume"`
	JournalPages  string   `json:"journal_pages"`
	Doi           string   `json:"doi"`
	DoiUrl        string   `json:"doi_url"`
	MagId         string   `json:"mag_id"`
	Authors       []Author `json:"authors"`
}

func JsonToPaper(jsonStr string) Paper {
	var item map[string]interface{} = make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &item)
	ok := false
	if err != nil {
		panic(err)
	}
	var paper Paper
	paper.Id = item["id"].(string)
	paper.Title = item["title"].(string)
	paper.Abstract, ok = item["abstract"].(string)
	if !ok {
		paper.Abstract = ""
	}
	paper.Url = item["s2Url"].(string)
	paper.S2PdfUrl = item["s2PdfUrl"].(string)
	year, ok := item["year"].(float64)
	if !ok {
		year = 0
	}
	paper.Year = int(year)
	paper.JournalPages = item["journalPages"].(string)
	paper.JournalName = item["journalName"].(string)
	paper.JournalVolume = item["journalVolume"].(string)
	paper.Doi = item["doi"].(string)
	paper.DoiUrl = item["doiUrl"].(string)
	pdf_urls := make([]string, 10000)
	for i, url := range item["pdfUrls"].([]interface{}) {
		pdf_urls[i] = url.(string)
	}
	paper.PdfUrls = pdf_urls
	in_citations := make([]string, 10000)
	for i, str := range item["inCitations"].([]interface{}) {
		in_citations[i] = str.(string)
	}
	paper.InCitations = in_citations
	out_citations := make([]string, 10000)
	for i, str := range item["outCitations"].([]interface{}) {
		out_citations[i] = str.(string)
	}
	paper.OutCitations = out_citations
	fields := make([]string, 10000)
	_, ok = item["FieldsOfStudy"].([]interface{})
	if !ok {
		item["FieldsOfStudy"] = make([]interface{}, 0)
	}
	for i, str := range item["FieldsOfStudy"].([]interface{}) {
		fields[i] = str.(string)
	}
	paper.FieldsOfStudy = fields
	authors := make([]Author, 10000)
	_, ok = item["authors"].([]map[string]interface{})
	if !ok {
		item["authors"] = make([]map[string]interface{}, 0)
	}
	for i, item_author := range item["authors"].([]map[string]interface{}) {
		author_new := Author{id: item_author["id"].(string), name: item_author["name"].(string)}
		authors[i] = author_new
	}
	paper.Authors = authors

	//author.position, ok = item["position"].(string)
	//if !ok {
	//	author.position = ""
	//}

	if err != nil {
		panic(err)
	}
	return paper
}
func proc_single_paper(m map[string]interface{}) map[string]interface{} {
	m["rank"],_ = strconv.ParseInt(m["rank"].(string),10,64)
	m["citation_count"],_ = strconv.ParseInt(m["citation_count"].(string),10,64)
	m["reference_count"],_ = strconv.ParseInt(m["reference_count"].(string),10,64)
	m["paper_id"],_ = strconv.ParseInt(m["paper_id"].(string),10,64)
	m["year"],_ = strconv.ParseInt(m["year"].(string),10,64)
	return m
}
func proc_file(file_path string, index string) {
	open, err := os.Open(file_path)
	if err != nil {
		fmt.Println(file_path + "打开失败")
		return
	}
	scanner := bufio.NewScanner(open)
	i := 0
	fin, error := os.OpenFile(file_path, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()
	client := service.ESClient
	bulkRequest := client.Bulk()
	//simpleBulkRequest := client.Bulk()
	reader := bufio.NewReader(fin)
	for {
		line, error_read := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}
		json_str := line
 
		var m map[string]interface{} = make(map[string]interface{})
		err = json.Unmarshal([]byte(json_str), &m)
		if err !=nil{panic(err)}
		reference_count,err := strconv.Atoi(m["reference_count"].(string))
		if reference_count > max_references_num {
			max_references_num = reference_count
		}
		//m = proc_single_paper(m)
		// 因为这些数据到es中已经超过了100G 由于io的限制会导致查询的特别慢。。于是杉树一些不必哟啊的属性。 将引用，被引用信息分开存储，减少paper 索引的数据量
		//m["comment_num"] ,m["download_num"],m["collect_num"],m["browser_num"]= 0,0,0,0
		//TODO 存到数据库中吧
		doc := elastic.NewBulkIndexRequest().Index(index).Id(m["paper_id"].(string)).Doc(m)
		//simpleBulkRequest.Add(elastic.NewBulkIndexRequest().Index("simple_paper").Id(m["id"].(string)).Doc(service.SimplifyPaper(m)))

		bulkRequest.Add(doc)
		if i%BULK_SIZE == 0 {
			response, err := bulkRequest.Do(context.Background())
			if err != nil {
				panic(err)
			}
			success_num += len(response.Succeeded())
			fail_num += len(response.Failed())
			fmt.Println("success_num", success_num, "fail_num", fail_num)

		}

		if error_read != nil {
			if err == io.EOF {
				fmt.Printf("%#v\n", line)
				break
			}
			panic(err)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	response, err := bulkRequest.Do(context.Background())

	if err != nil {
		panic(err)
	}

	success_num += len(response.Succeeded())
	fail_num += len(response.Failed())
	//response,err = simpleBulkRequest.Do(context.Background())
	if(len(response.Failed())>0){panic(err)}
	if fail_num > 0 {
		fmt.Println("error:")
	}
	for _, item := range response.Failed() {
		fmt.Println(item.Error)
	}
	fmt.Println("line sum", i)
	fmt.Println("success_num", success_num, "fail_num", fail_num, "max_citation_num", max_citation_num, "max_references_num", max_references_num)
	fmt.Println(fieldsMap)
}

func proc_author(file_path string, index string) {
	open, err := os.Open(file_path)
	if err != nil {
		fmt.Println(file_path + "打开失败")
		return
	}
	scanner := bufio.NewScanner(open)
	i := 0
	fin, error := os.OpenFile(file_path, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()
	client := service.ESClient
	bulkRequest := client.Bulk()
	reader := bufio.NewReader(fin)
	for {
		line, error_read := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}
		json_str := line

		//_ = JsonToPaper(json_str)
		//if(i<5){fmt.Println(paper)}
		var m map[string]interface{}
		_ = json.Unmarshal([]byte(json_str), &m)
		//if len(m["author_id"].(interface{})) == 0 {
		//	continue
		//} // 数据501行中存在"author_id": [],  过滤
		//m["author_id"] = m["author_id"].([]interface{})[0].(string)
		m["paper_count"] ,_= strconv.Atoi(m["paper_count"].(string))
		m["citation_count"] ,_= strconv.Atoi(m["citation_count"].(string))
		m["rank"] ,_= strconv.Atoi(m["rank"].(string))
		doc := elastic.NewBulkIndexRequest().Index(index).Id(m["author_id"].(string)).Doc(m)
		bulkRequest.Add(doc)
		if i%BULK_SIZE == 0 {
			response, err := bulkRequest.Do(context.Background())
			if err != nil {
				panic(err)
			}
			success_num += len(response.Succeeded())
			fail_num += len(response.Failed())
			fmt.Println("success_num", success_num, "fail_num", fail_num)

		}

		if error_read != nil {
			if err == io.EOF {
				fmt.Printf("%#v\n", line)
				break
			}
			panic(err)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	response, err := bulkRequest.Do(context.Background())
	if err != nil {
		panic(err)
	}
	success_num += len(response.Succeeded())
	fail_num += len(response.Failed())
	fmt.Println("line sum", i)
	fmt.Println("success_num", success_num, "fail_num", fail_num)
	fmt.Println(fieldsMap)
}
func proc_journal(file_path string, index string,main_id string) {
	open, err := os.Open(file_path)
	if err != nil {
		fmt.Println(file_path + "打开失败")
		return
	}
	scanner := bufio.NewScanner(open)
	i := 0
	fin, error := os.OpenFile(file_path, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()
	client := service.ESClient
	bulkRequest := client.Bulk()
	reader := bufio.NewReader(fin)
	for {
		line, error_read := reader.ReadString('\n')
		if len(line) == 0 {
			break
		}
		json_str := line

		//_ = JsonToPaper(json_str)
		//if(i<5){fmt.Println(paper)}
		var m map[string]interface{}
		_ = json.Unmarshal([]byte(json_str), &m)
		//if len(m["author_id"].([]interface{})) == 0{continue} // 数据501行中存在"author_id": [],  过滤
		//m["id"] = m["id"].([]interface{})[0].(string)
		doc := elastic.NewBulkIndexRequest().Index(index).Id(m[main_id].(string)).Doc(m)
		bulkRequest.Add(doc)
		if i%BULK_SIZE == 0 {
			response, err := bulkRequest.Do(context.Background())
			if err != nil {
				panic(err)
			}
			success_num += len(response.Succeeded())
			fail_num += len(response.Failed())
			if fail_num > 0{fmt.Println((response.Failed()[0].Error))}
			fmt.Println("success_num", success_num, "fail_num", fail_num)

		}

		if error_read != nil {
			if err == io.EOF {
				fmt.Printf("%#v\n", line)
				break
			}
			panic(err)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	response, err := bulkRequest.Do(context.Background())
	if err != nil {
		panic(err)
	}
	success_num += len(response.Succeeded())
	fail_num += len(response.Failed())

	fmt.Println("line sum", i)
	fmt.Println("success_num", success_num, "fail_num", fail_num)
	fmt.Println(fieldsMap)
}
func load_paper() {
	//cluster_block_exception index [paper] blocked by: [TOO_MANY_REQUESTS/12/disk usage exceeded flood-stage watermark, index has read-only-allow-delete block
	//考虑磁盘空间问题 ：方法：curl -XPUT -H "Content-Type: application/json" http://10.2.7.70:9204/_all/_settings -d '{"index.blocks.read_only_allow_delete": null}'
	//max_citation_num 220497 max_references_num 26676
	service.Init()
	proc_file(PAPER_DIR+"\\myPapers.txt", "paper")
	//for i := 0; i < 6000; i++ {
	//	var str string
	//	if i < 1000 {
	//		str = fmt.Sprintf("%03d", i)
	//	} else {
	//		str = strconv.Itoa(i)
	//	}
	//	fmt.Println(str)
	//	proc_file(PAPER_DIR+"\\"+PAPER_FILE_PREFIX+str, "paper")
	//
	//}
}
func load_authors() {
	service.Init()
	proc_author("H:\\myAuthors.txt", "author")
}
func load_journal() {
	service.Init()
	proc_journal("H:\\myJournals.txt", "journal","journal_id")
}
//func load_incitations() {
//	service.Init()
//	proc_journal("H:\\ScholarinCitations.txt", "incitations")
//}
func load_paper_author(){
	service.Init()
	proc_journal("H:\\myPaperAuthorAffiliations.txt","paper_author","paper_id")
}
func load_paper_rel(){
	service.Init()
	proc_journal("H:\\myPaperReferences.txt","reference","paper_id")
}
func load_conference(){
	service.Init()
	proc_journal("H:\\myConferenceInstances.txt","reference","conference_id")
}


func print1() {
	for i := 0; i < 1; i++ {
		fmt.Printf("%s\n", fmt.Sprintf("%04d", i))
	}
}
func main() {
	service.Init()
	//load_paper()

	//print1()
	//load_authors()

	//load_journal()
	//load_incitations()
	//load_paper_rel()
	//load_paper_a uthor()
	load_conference()
	load_journal()
}
