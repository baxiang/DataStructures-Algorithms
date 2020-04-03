package com.github.tree;

public class TreeNode<T> {
    /**
     * @Description 数据域
     */
    private T data;
    /**
     * @Description 左子树
     */
    public TreeNode<T> leftChild;
    /**
     * @Description 右子树
     */
    public TreeNode<T> rightChild;


    public TreeNode(T data) {
        this(null, data, null);
    }

    public TreeNode(TreeNode leftChild, T data, TreeNode rightChild) {
        this.leftChild = leftChild;
        this.data = data;
        this.rightChild = rightChild;
    }

    public T getData() {
        return data;
    }

    public TreeNode<T> getLeftChild() {
        return leftChild;
    }

    public TreeNode<T> getRightChild() {
        return rightChild;
    }

    public void setData(T data) {
        this.data = data;
    }
}
