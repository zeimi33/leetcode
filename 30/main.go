package main

import "fmt"

func main(){
fmt.Println(findSubstring("wordgoodgoodgoodbestword",[]string{"word","good","best","good"}))
}
 "wordgoodgoodgoodbestword"
["word","good","best","good"]
// "barfoothefoobarman"
// ["foo","bar"]
// "barfoofoobarthefoobarman"
// ["bar","foo","the"]
func findSubstring(s string, words []string) []int {
	if len(words)==0{
		return nil
	}
	ret := []int{}
	length := len(words[0])
	num := len(words)
	m := map[string]int{}
	for _,v :=range words{
		m[v] =2
	}
	stack := []string{}
	jishu :=0
	index := 0
	for len(s) > 0{
		jishu++
	//	fmt.Println(s)
		key := s[:length]
	//	fmt.Println(key)
		s = s[length:]
		if m[key] ==2{
			stack = append(stack,key)
			m[key] =1
			index++

		}else{
			fmt.Println(key)
							if _,ok := m[key];!ok{
					for v :=range m{
						m[v] =2
						index--
					}
					continue
				}else{
			for {

				if len(stack) ==0{break}
				if stack[0] != key{
					m[stack[0]]=2
					stack = stack[1:]
					index--
					continue
				}
				if stack[0] ==key{
					m[stack[0]]=2
					stack = stack[1:]
					index--
					break
				}

			}
		}
			stack = append(stack,key)
			m[key] =1
			index++

		}

		if index ==num {
				ret = append(ret,jishu*length-num*length)
			}
				fmt.Println(key,s,m,ret,index,jishu,jishu*length,num*length)
	}
	return ret
}



func findSubstring(s string,words []string)[]int{
	var res []int
    var len_leng=len(words)
    if len_leng==0{
        return res
    }
	var len_word=len(words[0])
	var len_words=len_word*len_leng
	var len_s=len(s)
	if len_words>len_s{
        return res
    }
	var wordmap=make(map[string]int)
	
	for _,val:=range words{
		wordmap[val]++
	}
	
	
	for i:=0;i+len_words<=len_s;i++{
		//定义一个map
		var tarmap=make(map[string]int)
		var nums int
		s_arr:=s[i:i+len_words]
		
		//满足条件就写入tarmap中，避免超时
		if _,ok:=wordmap[s_arr[0:len_word]];ok{
			for k:=0;k<len_words;k+=len_word{
				tarmap[s_arr[k:k+len_word]]++
			}
		}
		
		if len(tarmap)==len(wordmap){
			nums=ss(wordmap,tarmap)
		}
		
		//如果两个tarmap与wordmap相等，则得到结果res
		if nums==len(wordmap){
			res=append(res,i)
			
		}
		
		
		
	}
	
	return res
	
}

func ss(s map[string]int,s1 map[string]int) int{
	var nums int
	for key,val:=range s{
		
		for key1,val1:=range s1{
			
			if key==key1 && val==val1{
				nums++
			}
			
		}
		
	}
	
	return nums
	
}
