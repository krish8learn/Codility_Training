import java.util.*;
public class Stonewall {
    public static void main(String[] args) {
        Stonewall st = new Stonewall();
        System.out.println(st.solution(new int[]{8,8,5,7,9,8,7,4,8}));
    }   
    
    public int solution(int[] H){
        Stack<Integer> stack = new Stack<Integer>();
        int count = 0;
        for(int i:H){
            while(!stack.isEmpty() && stack.peek()>i){
                stack.pop();
            }

            if(stack.isEmpty() || i > stack.peek()){
                stack.push(i);
                count++;
            }
        }
        return count;
    }
}
