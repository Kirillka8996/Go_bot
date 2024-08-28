package httpjson

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DataFromUrl выполняет GET-запрос по указанному URL и возвращает ответ.
func dataFromUrl(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Проверка на успешный статус-код
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // Закрываем тело ответа, если статус не OK
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	return resp, nil
}

func jsonFromUrl(resp *http.Response) (map[string]any, error) {
	defer resp.Body.Close() // Закрываем тело ответа после использования

	jsonStr, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Распаковываем JSON-строку в карту
	data := make(map[string]any)
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return data, nil
}

func FetchJsonFromUrl(url string) (map[string]any, error) {
	resp, err := dataFromUrl(url)
	if err != nil {
		return nil, err
	}

	data, err := jsonFromUrl(resp)
	if err != nil {
		return nil, err
	}

	return data, nil
}
