package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BrickworkFile = "diamond_count.txt"
const NameFile = "names.txt"

const GiftFile = "gift.txt"

// 写入钻石数量到文件
func WriteDiamondCount(filename, person string, diamonds int) error {
	// 格式化内容
	content := fmt.Sprintf("Name: %s\nDiamonds: %d\n\n", person, diamonds)

	// 以追加模式打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 将内容写入文件
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}

// 从文件中读取某人的钻石数量
func ReadDiamondCount(filename, person string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var foundPerson bool
	for scanner.Scan() {
		line := scanner.Text()

		// 检查是否匹配目标人名
		if strings.HasPrefix(line, "Name: ") && strings.HasSuffix(line, person) {
			foundPerson = true
		}

		// 如果已经找到名字，则读取下一行的钻石数量
		if foundPerson && strings.HasPrefix(line, "Diamonds: ") {
			countStr := strings.TrimPrefix(line, "Diamonds: ")
			diamonds, err := strconv.Atoi(countStr)
			if err != nil {
				return 0, fmt.Errorf("无法解析钻石数量: %w", err)
			}
			return diamonds, nil
		}
	}

	// 检查扫描过程中的错误
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("文件读取错误: %w", err)
	}

	// 如果没有找到指定人名
	return 0, fmt.Errorf("未找到名为 %s 的人", person)
}

// 修改文件中某人的钻石数量
func ModifyDiamondCount(filename, person string, newDiamondCount int) error {
	// 打开文件进行读取
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 临时存储文件内容
	var lines []string
	var foundPerson bool

	// 逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 如果找到目标人的名字
		if strings.HasPrefix(line, "Name: ") && strings.HasSuffix(line, person) {
			foundPerson = true
			// 更新该行的钻石数量
			lines = append(lines, fmt.Sprintf("Name: %s", person))
			lines = append(lines, fmt.Sprintf("Diamonds: %d", newDiamondCount))
			continue
		}

		// 如果没有找到目标人，或已经处理过该人，直接将当前行保存
		if foundPerson && strings.HasPrefix(line, "Diamonds: ") {
			// 跳过这一行
			continue
		}

		lines = append(lines, line)
	}

	// 检查读取过程中是否有错误
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取文件错误: %w", err)
	}

	// 如果未找到目标人
	if !foundPerson {
		return fmt.Errorf("未找到名为 %s 的人", person)
	}

	// 打开文件进行写入（覆盖原文件）
	file, err = os.Create(filename)
	if err != nil {
		return fmt.Errorf("无法创建文件: %w", err)
	}
	defer file.Close()

	// 将修改后的内容写入文件
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("写入文件失败: %w", err)
		}
	}

	return nil
}

// 将人名写入文件
func WriteNamesToFile(filename string, names []string) error {

	// 打开文件，如果文件不存在则创建，使用追加模式（os.O_APPEND）
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 将人名逐个写入文件，每个名字后面加上换行符
	for _, name := range names {
		_, err := file.WriteString(name + "\n")
		if err != nil {
			return fmt.Errorf("写入文件失败: %w", err)
		}
	}

	return nil
}

// 从文件中检查是否存在某个人名
func CheckNameInFile(filename, person string) (bool, error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return false, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 创建扫描器
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 获取当前行的内容
		line := scanner.Text()
		// 检查是否与目标人名匹配
		if strings.TrimSpace(line) == person {
			return true, nil
		}
	}

	// 检查读取过程中的错误
	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("文件读取错误: %w", err)
	}

	// 如果文件中没有找到指定的人名
	return false, nil
}

// 将奖品写入文件，格式为 "姓名: 获奖奖品"，追加写入
func WritePrizeToFile(filename, name, prize string) error {
	// 打开文件，以追加模式打开文件
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 格式化内容，写入 "姓名: 获奖奖品"
	content := fmt.Sprintf("%s: %s\n", name, prize)
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}
