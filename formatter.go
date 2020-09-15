package logger2

import (
	"encoding/json"
	"fmt"
)

type Formatter interface {
	Format(*Entity) ([]byte, error)
}

type JSON_Formatter struct {
}

func (j *JSON_Formatter) Format(entity *Entity) ([]byte, error) {
	b, err := json.Marshal(entity.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields %v", err)
	}
	return b, nil
}
