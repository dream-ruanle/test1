package main

func findRedundantConnection(edges [][]int) []int {
	n:=len(edges)
	parent:=make([]int,n)
	for i:=range parent{
		parent[i]=i
	}
	var find func(int)int
	find= func(x int) int {
		if x!=parent[x]{
			parent[x]=find(parent[x])
		}
		return parent[x]
	}
	var union func(int,int)
	union= func(form int, to int) {
		parent[find(form)]=parent[to]
	}
	var res []int
	for _,row:=range edges{
		if find(row[0])==find(row[1]){
			res=row
		}
		union(row[0],row[1])
	}
	return res
}
