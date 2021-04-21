
import java.util.Stack;

class Fish{
    public static void main(String[] args) {
        Fish fi = new Fish();
        //System.out.println(fi.solution(new int[]{4,3}, new int[]{0,1}));
        System.out.println(fi.solution(new int[]{1}, new int[]{0}));
        //System.out.println(fi.solution(new int[]{4,3,2,1,5}, new int[]{0,1,0,0,0}));
        //System.out.println(fi.solution(new int[]{2,3,1,5,4}, new int[]{0,1,0,1,1}));
    }

    public int solution(int[] A, int[] B){
        Stack<Integer> stack = new Stack<Integer>();
        int alive = B.length;
        for(int p=0; p<A.length; p++){
            if(B[p]==1){
                stack.push(A[p]);
            }else if(B[p]==0){
                while(!stack.isEmpty()){
                    if(stack.peek()>A[p]){
                        alive--;
                        break;
                    }else if(stack.peek()<A[p]){
                        stack.pop();
                        //stack.push(A[p]);
                        alive--;
                        //break;
                    }
                }
            }
        }

        return alive;
    }
}