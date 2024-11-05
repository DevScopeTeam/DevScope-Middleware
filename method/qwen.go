package method

import (
	"DevScope-Middleware/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func RequestQwen(system_prompt, user_prompt string) (string, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 构建请求体
	requestBody := RequestBody{
		// 模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
		Model: "qwen-plus",
		Messages: []Message{
			{
				Role:    "system",
				Content: system_prompt,
			},
			{
				Role:    "user",
				Content: user_prompt,
			},
		},
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	// 创建 POST 请求
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+config.Qwen_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	choices := gjson.Get(string(bodyText), "choices")
	return choices.Get("0.message.content").String(), nil
}
