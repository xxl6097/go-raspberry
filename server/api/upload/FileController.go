package upload

import (
	"crypto/md5"
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"go-raspberry/server/utils"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type FileController struct {
}

// NewController http://openai.clife.net:9010/v1/api/user/signin
func NewController() *FileController {
	return &FileController{}
}

func (this *FileController) Upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	case "POST":
		//ParseMultipartForm将请求的主体作为multipart/form-data解析。请求的整个主体都会被解析，得到的文件记录最多 maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。如果必要，ParseMultipartForm会自行调用 ParseForm。重复调用本方法是无意义的
		//设置内存大小
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		m := r.MultipartForm
		files := m.File["file"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dir := "./files/" + utils.GetTimeDir()
			if !utils.IsDirExists(dir) {
				err := utils.CreateMutiDir(dir)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			//create destination file making sure the path is writeable.
			filepath := dir + files[i].Filename
			dst, err := os.Create(filepath)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			glog.Info(filepath)
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		Respond(w, Ok())

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//curl -F "uploadfile=@C:\Users\wangs\Desktop\xueba_license.txt" localhost:8080/upload
