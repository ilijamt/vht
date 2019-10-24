package vault
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/hashicorp/vault/api"
//	"golang.org/x/sync/semaphore"
//	"io"
//	"strings"
//	"sync/atomic"
//)
//
//// DeletePaths removes all the paths in the list from Vault
//func CopyPath(src string, dst string, concurrency int8, client *api.Client, w io.Writer) (written uint64, err error) {
//	if client == nil {
//		return written, errors.New(ErrMissingVaultClient)
//	}
//
//	if src == "" || dst == "" {
//		return written, errors.New(ErrMissingPath)
//	}
//
//	var foundPaths []string
//	foundPaths, err = Tree(src, client, concurrency)
//	if err != nil {
//		return written, err
//	}
//
//	type Path struct {
//		Source, Destination string
//	}
//
//	var path Path
//	var sem = semaphore.NewWeighted(int64(concurrency))
//
//	for _, fp := range FilterOnlyDataPaths(foundPaths) {
//		path = Path{
//			Source:      fp,
//			Destination: strings.ReplaceAll(fp, src, dst),
//		}
//
//		if err = sem.Acquire(context.Background(), 1); err != nil {
//			return written, err
//		}
//
//		go func(p Path) {
//			defer sem.Release(1)
//			secret, err := client.Logical().Read(p.Source)
//			if err != nil {
//				_, _ = fmt.Fprintf(w, "Failed to read with %v\n", err)
//				return
//			}
//			fmt.Println(p.Destination, secret.Data)
//			_, err = client.Logical().Write(p.Destination, secret.Data)
//			if err != nil {
//				_, _ = fmt.Fprintf(w, "Failed to write with %v\n", err)
//				return
//			}
//			atomic.AddUint64(&written, 1)
//		}(path)
//	}
//
//	// Iterate over the nodes so we can get the new data
//	//for _, path := range nodes {
//	//	newPath := strings.ReplaceAll(fmt.Sprintf("%s", path), "//", "/")
//	//	if err = sem.Acquire(context.Background(), 1); err != nil {
//	//		return written, err
//	//	}
//	//	go func(path string) {
//	//		defer sem.Release(1)
//	//		pths, err := Tree(newPath, client, concurrency)
//	//		if err == nil {
//	//			mu.Lock()
//	//			paths = append(paths, pths...)
//	//			mu.Unlock()
//	//		}
//	//	}(newPath)
//	//}
//
//	return written, sem.Acquire(context.Background(), int64(concurrency))
//}
