package shellUtil

import (
	"log"
	"os/exec"
)

func RunShellCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	//fmt.Println("----cmd---",cmd)
	result, err := cmd.Output()
	if err != nil {
		log.Println("Execute Command Error:" + err.Error())
		return "", err
	}
	return string(result), nil
}

/*
** 执行脚本
 */
func RunScripts(fileName string) (result string, err error) {
	cmd := exec.Command("sh", "-c", fileName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

/*
** 执行脚本
 */
func RunScriptsWithParam(fileName string, params ...string) (result string, err error) {
	paramAry := []string{}
	paramAry = append(paramAry, params...)

	var cmd *exec.Cmd
	switch len(paramAry) {
	case 0:
		cmd = exec.Command("sh", fileName)
	case 1:
		cmd = exec.Command("sh", fileName, paramAry[0])
	case 2:
		cmd = exec.Command("sh", fileName, paramAry[0], paramAry[1])
	case 3:
		cmd = exec.Command("sh", fileName, paramAry[0], paramAry[1], paramAry[2])
	case 4:
		cmd = exec.Command("sh", fileName, paramAry[0], paramAry[1], paramAry[2], paramAry[3])
	case 5:
		cmd = exec.Command("sh", fileName, paramAry[0], paramAry[1], paramAry[2], paramAry[3], paramAry[4])
	case 6:
		cmd = exec.Command("sh", fileName, paramAry[0], paramAry[1], paramAry[2], paramAry[3], paramAry[4], paramAry[5])
	case 7:
		cmd = exec.Command("sh", fileName, paramAry[0], paramAry[1], paramAry[2], paramAry[3], paramAry[4], paramAry[5], paramAry[6])
	}
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
