package service

import (
	"flower/mysql"
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	err := mysql.Init("root:linym6303763!@tcp(123.207.1.119:3306)/flower?charset=utf8mb4")
	if err != nil {
		// TODO
	}
	os.Exit(m.Run())
}

func BenchmarkFrontProdService_GetProduct(b *testing.B) {

	var wg  sync.WaitGroup
	wg.Add(20)

	for i:=0; i<20;i++  {
		go func() {
			for j:=0; j<50 ; j++  {
				p, _, err := FrontProdSrv.GetProduct(13)
				if err != nil {
					fmt.Printf(err.Error())
				}else{
					println(p.Id)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

