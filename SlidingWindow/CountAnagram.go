package SlidingWindow

import (
	"fmt"
	"strings"
)

//public class Anagrams {
//    public static void main(StringTopic[] args) {
//        StringTopic str = "aabaabbabbabaa";
//        StringTopic ptr = "aaba";
//        int sum = 0;
//        int count = 0;
//        int k = ptr.length();
//
//        Map<Character,Integer> map = new HashMap<>();
//        for(int i=0;i<ptr.length();i++){
//            char ch = ptr.charAt(i);
//            if(map.containsKey(ch)){
//                int a = map.get(ch);
//                map.put(ch,++a);
//            } else{
//                map.put(ch,1);
//            }
//        }
//
//        count = map.size();
//        int i=0,j=0;
//
//        while(j<str.length()){
//            char ch = str.charAt(j);
//            if(map.containsKey(ch)){
//                int a = map.get(ch);
//                map.put(ch,--a);
//                if(a==0){
//                    count--;
//                }
//            }
//
//            if(j-i+1 < k){
//                j++;
//                continue;
//            }
//
//            if(j-i+1 == k){
//                if(count == 0){
//                    sum++;
//                }
//                if(map.containsKey(str.charAt(i))){
//                    int a = map.get(str.charAt(i));
//                    if(a==0){
//                        count++;
//                    }
//                    map.put(str.charAt(i), ++a);
//                }
//            }
//            i++;
//            j++;
//        }
//
//        System.out.println(sum); //3
//    }
//}

func CountAnagram(str1 string, str2 string) {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	arr := make([]int, 26)
	for i := 0; i < len(str1); i++ {
		arr[rune(str1[i])-97]++
	}
	start := 0
	var check string
	count := 0
	for end := 0; end < len(str2); end++ {
		check = check + string(rune(str2[end]))
		if end-start+1 == len(str1) {
			if isAnagram(arr, check) {
				count++
			}
			start++
			check = check[1:]
		}
	}
	fmt.Println(count)
}

func isAnagram(arr []int, str1 string) bool {
	count := make([]int, 26)
	for i := 0; i < len(str1); i++ {
		count[rune(str1[i])-97]++
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != count[i] {
			return false
		}
	}
	return true
}
