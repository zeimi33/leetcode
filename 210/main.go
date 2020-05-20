package main

func main(){


}

const (
	no  = 1
	ing = 2
	yes = 3
)


func findOrder(numCourses int, prerequisites [][]int) []int {
    var (
        edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        invalid bool
        dfs func(u int)
    )

    dfs = func(u int) {
        visited[u] = 1
        for _, v := range edges[u] {
            if visited[v] == no {
                dfs(v)
                if invalid {
                    return
                }
            } else if visited[v] == ing {
                invalid = true
                return
            }
        }
        visited[u] = yes
        result = append(result, u)
    }

    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
    }

    for i := 0; i < numCourses && !invalid; i++ {
        if visited[i] == no {
            dfs(i)
        }
    }
    if invalid {
        return []int{}
    }
    for i := 0; i < len(result)/2; i ++ {
        result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	
    return result
}
