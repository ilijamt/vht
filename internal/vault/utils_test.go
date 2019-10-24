package vault_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/vault/api"
	"math/rand"
	"strings"
	"time"
)

type dataCase struct {
	Test int
	Time time.Time
}

var writeRandomData = func(rootPath string, client *api.Client, total int) (fullPath string, err error) {
	lvl := 1 + rand.Intn(total)
	var paths []string
	for i := 0; i < lvl; i++ {
		paths = append(paths, uuid.New().String())
	}
	fullPath = fmt.Sprintf("%s/%s", rootPath, strings.Join(paths, "/"))
	data := map[string]interface{}{"data": dataCase{Test: rand.Intn(1000), Time: time.Now()}}
	_, err = client.Logical().Write(fullPath, data)
	return fullPath, err
}
