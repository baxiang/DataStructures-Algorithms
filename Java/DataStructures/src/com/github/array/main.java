package com.github.array;

import javax.jnlp.ClipboardService;

public class main {
    public static ListNode removeElements(ListNode head, int val) {
            ListNode dummyHead = new ListNode(-1);
            ListNode currNode = dummyHead;
            while (head!=null){
                if (head.val!=val){
                    currNode.next = head;
                    currNode = currNode.next;
                }
                head = head.next;
            }
            currNode.next = null;
          return dummyHead.next;
        }



    public static void main(String[] args) {
       ListNode head = new ListNode(1);
       ListNode a = new ListNode(2);
       head.next = a;
        ListNode b = new ListNode(6);
        a.next = b;
        ListNode c = new ListNode(3);
        b.next = c;
        ListNode d = new ListNode(4);
        c.next = d;
        ListNode e = new ListNode(5);
        d.next = e;
        ListNode f = new ListNode(6);
        e.next = f;
        ListNode g = new ListNode(7);
        f.next = g;
        ListNode h = removeElements(head,6);
        while (h!=null){
            System.out.println(h.val);
            h = h.next;
        }
    }
}
