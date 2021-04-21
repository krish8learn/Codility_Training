import java.util.ArrayList;
import java.util.Collections;

class GenomicRangeQuery{
    public static void main(String[] args) {
        GenomicRangeQuery gq = new GenomicRangeQuery();
        int[] abc = gq.solution("CAGCCTA", new int[]{2,5,0}, new int[]{4,5,6});
        for(int i =0; i<abc.length; i++){
            System.out.println(abc[i]);
        }

        /*int[] abc = gq.solution("AC", new int[]{0,0,1}, new int[]{0,1,1});
        for(int i =0; i<abc.length; i++){
            System.out.println(abc[i]);
        }*/
    }

    public int[] solution(String S, int[] P, int[] Q){
        int[] result = new int[P.length];
        int k=0;
        char[] strand = S.toCharArray();
        for(int i=0; i<P.length; i++){
            ArrayList<Integer> al = new ArrayList<Integer>();
            for(int j=P[i]; j<=Q[i]; j++){
                //System.out.println("k--<"+k);
                //ArrayList<Integer> al = new ArrayList<Integer>();
                if(strand[j]=='A'){
                    al.add(1);
                }else if(strand[j]=='C'){
                    al.add(2);
                }else if(strand[j]=='G'){
                    al.add(3);
                }else if(strand[j]=='T'){
                    al.add(4);
                }
            }
            result[k] = Collections.min(al);
            k++;
        }
        return result;
    }
}