package main

import "testing"

func TestTriangle(t *testing.T) {
	tests:=[]struct{ a, b, c int}{
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
		{30000,40000,50000},
	}
	for _,tt:=range  tests{
		if acr:=calc(tt.a,tt.b);acr!=tt.c{
			t.Errorf(" %d的平方加 %d的平方测试 期望等于%d  函数实际等于%d",tt.a,tt.b,tt.c,acr)
		}
	}

}

func BenchmarkA(b *testing.B){
	var a1,b1,c1=3,4,5
		for i:=0;i<b.N;i++{
			act1:=calc(a1,b1)
			if act1!=c1{
				b.Errorf("%d平方加 %d平方 期望 %d 实际 %d",a1,b1,c1,act1)
			}
		}
}