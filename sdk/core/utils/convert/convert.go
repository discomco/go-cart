package convert

import "encoding/json"

func Any2Data(p any) ([]byte, error) {
	result, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Data2Any(data []byte, p any) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}
