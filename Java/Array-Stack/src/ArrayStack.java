import java.util.ArrayList;

public class ArrayStack<E> implements Stack<E> {

    private ArrayList<E> list;
    public ArrayStack(){
        list = new ArrayList<E>();
    }

    @Override
    public int getSize() {
        return list.size();
    }

    @Override
    public boolean isEmpty() {
        return list.isEmpty();
    }

    @Override
    public void push(E e) {
        list.add(list.size(),e);
    }

    @Override
    public E pop() {
        if (!isEmpty()){
            return list.remove(list.size()-1);
        }
        return null;
    }

    @Override
    public E Peek() {
        if (!isEmpty()){
            return list.get(list.size()-1);
        }
        return null;
    }

    @Override
    public String toString(){
       return list.toString();
    }
}
