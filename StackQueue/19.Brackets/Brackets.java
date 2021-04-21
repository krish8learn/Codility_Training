import java.util.Stack;

class Brackets{
    public static void main(String[] args) {
        Brackets bl = new Brackets();
        System.out.println(bl.solution("([)()]"));
        //System.out.println(bl.solution("{[()()]}"));
    }

    public int solution(String S){
        Stack<Character> stack = new Stack<Character>();
        for(int i=0; i<S.length(); i++){
            if(S.charAt(i) == '(' || S.charAt(i)=='{' || S.charAt(i) == '['){
                stack.push(S.charAt(i));
            }else{
                if(stack.isEmpty()){
                    return 0;
                }else{
                    if(S.charAt(i)==')' && stack.pop()!='('){
                        return 0;
                    }else if(S.charAt(i)=='}' && stack.pop()!='{'){
                        return 0;
                    }else if (S.charAt(i)==']' && stack.pop()!='['){
                        return 0;
                    }
                }
            }
        }
        if(stack.isEmpty()){
            return 1;
        }else{
            return 0;
        }
    }
}