package app

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
)

func printLine(res string, AsciiShapes []string, char int, i int) string {
	str := strings.Split(AsciiShapes[char-32], "\n")
	res = res + str[i]
	return res
}

func CheckNewLines(arr []string) bool {
	for _, v := range arr {
		if v != "" {
			return false
		}
	}
	return true
}

func AsciiArt(input string, banner string) (string, error) {
	if len(input) == 0 || len(input) > 200 || !(banner == "standard" || banner == "shadow" || banner == "thinkertoy" ){
		return "", errors.New("input error")
	}
	fileName := banner + ".txt"

	info, err := os.Stat("app/" + fileName)
	if err == nil {
		if info.Mode() != 0o400 {
			erRemoveFile := exec.Command("rm", "-f", "app/"+fileName).Run()
			if erRemoveFile != nil {
				return "", erRemoveFile
			}
		}
	}

	_, errOpen := os.Open("app/" + fileName)
	if errOpen != nil {
		if os.IsNotExist(errOpen) {
			url := "https://learn.zone01oujda.ma/api/content/root/public/subjects/ascii-art/" + fileName

			// download the file
			errDownload := exec.Command("wget", "-q", "-P", "app/", url).Run()
			if errDownload != nil {
				return "", errDownload
			}

			// change the permission
			erPermission := exec.Command("chmod", "400", "app/"+fileName).Run()
			if erPermission != nil {
				return "", erPermission
			}
		}
	}

	file, er := os.Open("app/" + fileName)
	if er != nil {
		return "", er
	}
	defer file.Close()

	fileData, erR := io.ReadAll(file)
	if erR != nil {
		return "", erR
	}

	CleanFileData := strings.ReplaceAll(string(fileData), "\r", "")
	AsciiShapes := strings.Split(CleanFileData, "\n\n")

	input = strings.ReplaceAll(input, "\r", "")

	// validate user input
	for _, char := range input {
		if char != '\n' && (char < 32 || char > 126) {
			return "", errors.New("The string includes characters outside the ASCII range... " + string(char))
		}
	}
	inputLines := strings.Split(input, "\n")

	// if the input contain only new lines
	if CheckNewLines(inputLines) {
		inputLines = inputLines[1:]
	}
	res := ""
	for _, Inpline := range inputLines {
		if Inpline == "" {
			res += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range Inpline {
				if char != ' ' {
					res = printLine(res, AsciiShapes, int(char), i)
				} else {
					res = printLine(res, AsciiShapes, int(char), i+1)
				}
			}
			res += "\n"
		}
	}
	return res, nil
}
