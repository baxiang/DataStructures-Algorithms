package com.github.array;

import java.util.ArrayList;

public class ArrayQueue<E> implements Queue<E> {

    private ArrayList<E> list;

    public ArrayQueue() {
        list = new ArrayList<>();
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
    public void enqueue(E e) {
        list.add(e);

    }

    @Override
    public E dequeue() {
        if (!isEmpty()) {
            return list.remove(list.size() - 1);
        }
        return null;
    }

    @Override
    public E getFront() {
        if (!isEmpty()) {
            return list.get(0);
        }
        return null;
    }

    @Override
    public String toString() {
        return list.toString();
    }
}
