import java.util.Stack;
public class Nested {
    public static void main(String[] args) {
        Nested ns = new Nested();
        System.out.println(ns.solution("(()(())())"));
        System.out.println(ns.solution("())"));
    }

    public int solution(String S){
        Stack<Character> stack = new Stack<Character>();
        char[] ch = S.toCharArray();
        for(int i=0; i<ch.length; i++){
            if(ch[i]=='('){
                stack.push(ch[i]);
            }else if(!stack.isEmpty()){
                if(ch[i] == ')' && stack.peek()=='('){
                    stack.pop();
                }
            }else if (stack.isEmpty()){
                stack.push(ch[i]);
            } 
        }

        if(stack.isEmpty()){
            return 1;
        }
        return 0;
    }
}
