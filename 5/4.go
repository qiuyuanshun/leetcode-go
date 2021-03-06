/*
* @Author: qiuyu
* @Date:   2019-04-25 02:01:34
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-25 02:03:08
*/
func longestPalindrome(s string) string {

   slice := make([]string,0,4)
   for _ , char := range s {
       slice = append(slice,"#")
       slice = append(slice,string(char))
   }
   slice = append(slice,"#") 
   maxR := 0
   maxIndex := 0
   sliceLen := len(slice)
   for index , _ := range slice{
       if(index >=1){
           r := 0
           i := index - 1
           j := index + 1
           for {
               if i<0 || j >= sliceLen{
                   break
               }
               if(slice[i] == slice[j]){
                   r++
                   i--
                   j++
               }else{
                   break
               }
           }
           if r > maxR{
               maxR = r
               maxIndex = index
           }
       }
   }
   res := ""
   for index,str := range slice{
       if index >= (maxIndex-maxR) && index <= (maxIndex+maxR) && str != "#"{
           res += str
       } 
   }
   return res
}