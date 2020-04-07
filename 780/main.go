package main

import "fmt"

func main(){
	fmt.Println(reachingPoints(6,5,11,16))
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for ty >= sy && tx >=sx{
		if ty  == tx{
			break
		}
		if ty > tx{
			if tx >sx{
				ty = ty%tx
			}else{
				return ((ty-sy)%tx ==0)
            }
		}else{
			
				if ty > sy{
				tx = tx%ty
			}else{
				
				return ((tx-sx)%ty ==0)
			
		}
	}
	fmt.Println(tx,ty,sy)
}
	return (tx == sx && ty == sy);
}	