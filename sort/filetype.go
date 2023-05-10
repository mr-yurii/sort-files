package sort

import (
	"path/filepath"
	"strings"
)

var dirsForFiles = map[string]string{
	".mp3":  "Music",
	".flac": "Music",
	".wav":  "Music",
	".aiff": "Music",
	".m4a":  "Music",
	".wma":  "Music",
	".ogg":  "Music",
	".mp4":  "Videos",
	".mov":  "Videos",
	".avi":  "Videos",
	".mkv":  "Videos",
	".wmv":  "Videos",
	".flv":  "Videos",
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".gif":  "Images",
	".bmp":  "Images",
	".tiff": "Images",
	".svg":  "Images",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".xls":  "Documents",
	".xlsx": "Documents",
	".ppt":  "Documents",
	".pptx": "Documents",
	".txt":  "Documents",
	".rtf":  "Documents",
	".zip":  "Archives",
	".rar":  "Archives",
	".tar":  "Archives",
	".gz":   "Archives",
	".7z":   "Archives",
	".exe":  "Executables",
	".dll":  "Executables",
	".sh":   "Executables",
	".bin":  "Executables",
	".deb":  "Executables",
	".rpm":  "Executables",
}

func GetFolderName(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	folder, exists := dirsForFiles[ext]
	if exists {
		return folder
	}
	return ""
}
