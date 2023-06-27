package utils

import "encoding/json"

type Callback struct {
  Call string `json:"call"`
  Data interface{} `json:"data"`
}

func NewCallback(call string, data interface{}) Callback {
  return Callback{Call: call, Data: data}
}

func NewCallbackFromCallbackData(data string) (Callback, error) {
  var callback Callback
  err := json.Unmarshal([]byte(data), &callback)
  if err != nil {
    return callback, err
  }

  return callback, nil
}

func MarshalCallback(callback Callback) (string, error) {
  data, err := json.Marshal(callback)
  if err != nil {
    return "", err
  }

  return string(data), nil
}

