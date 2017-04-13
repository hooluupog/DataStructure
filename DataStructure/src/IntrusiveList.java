/** Intrusive LinkedList implementation.
 *  Definition:
 *      An intrusive data structure is one where the data items contain the metadata needed,
 *      instead of the metadata pointing to the contained data.The intrusive linked list is 
 *      the data oriented(FYI,data oriented programming) version of the linked list container.
 *      For example,the external linked list looks like this:
 *
 *      class ListNode<E>{
 *           E val;
 *           ListNode<E> next;
 *       }
 *       class LinkList<E> {
 *           ListNode<E> head;
 *       }
 *        
 *             -> Item.next -> Item.next -> nil
 *                .obj         .obj
 *                  |            |
 *                  v            v
 *                 data         data    
 *                 
 *     Whereas, The Intrusive linked list:
 *     class ListNode<E extends ListNode<E>>{
 *         ListNode<E> prev;
 *         ListNode<E> next;
 *     }
 *     class LinkList<E extends ListNode<E>> extends Iterable<E>{
 *         ListNode<E> head;
 *     }
 *     
 *             -> data
 *                    .Item.next -> data
 *                                      .Item.next -> nil
 *                                      
 *     As you see, the 'intrusive' means that we store the list node inside the type E.
 *     The goal of this is to decrease the number of allocations; instead of `Item` 
 *     and `data` being separate chunks of memory, they are allocated as one.
 *     For more details,see http://www.codefarms.com/publications/intrusiv/intr.htm.
 * 
 *  Pro: 
 *     Decoupling memory allocation from the container itself to avoid unnecessary
 *     memory copies;
 *     Circular implementation simplifies list manipulation implementation.    
 * */

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;
import java.util.NoSuchElementException;
import java.util.stream.*;
import java.io.*;

class ListNode<E extends ListNode<E>>{
    LinkList<E> list;
    ListNode<E> prev;
    ListNode<E> next;
}

class LinkList<E extends ListNode<E>> implements Iterable<E>{
    private ListNode<E> head;
    private int len = 0;
    LinkList(){
        head = new ListNode<E>();
        head.list = this;
        head.prev = head.next = head;
    }

    public E getFirst(){
        return (E)head.next;
    }

    public E getLast(){
        return (E)head.prev;
    }

    public boolean isEmpty(){
        return len == 0;
    }

    public int length(){
        return len;
    }

    public void add(E e) {
        // append [e] into LinkList.
        insert(e, head.prev);
    }

    public void addFirst(E e) {
        // add [e] at the first.
        insert(e, head);
    }

    public void insertBefore(E e, E at) {
        //Insert [e] before [at].
        insert(e, at.prev);
    }

    public void insertAfter(E e, E at) {
        //Insert [e] after [at].
        insert(e, at);
    }

    private void insert(E e, ListNode<E> at) {
        if (e.list != null) {
            throw new UnsupportedOperationException("ListNode is already in a LinkList.");
        }
        e.list = this; // Link value into LinkList.
        e.prev = at;
        e.next = at.next;
        at.next = e;
        e.next.prev = e;
        len++;
    }

    public E remove(E e) {
        e.prev.next = e.next;
        e.next.prev = e.prev;
        e.list = null;
        e.prev = e.next = null;
        len--;
        return e;
    }

    public void reverseBetween(int m, int n) {
        if (m < 1 || n > len) {
            throw new IllegalArgumentException(String.format("Index out of bounds. Wrong range (%d,%d)",m,n));
        }
        int cnt = n - m; // 1 <= m <= n <= LinkList.Length
        E start = getFirst();
        for (int i = 1; i < m; i++) {
            //start points to the first node to be reversed.
            start = (E)start.next;
        }
        E end = start;
        while (cnt > 0) {
            // Insert nodes from m to n into list.
            E curr = (E)end.next;
            remove(curr);
            insertBefore(curr, start);
            start = curr;
            cnt--;
        }
    }
    public Iterator<E> iterator(){
        return new LinkListIterator<E>(this);
    }
    // inner class.
    static final class LinkListIterator<E extends ListNode<E>> implements Iterator<E> {
        final private LinkList<E> list;
        private ListNode<E> next;
        LinkListIterator(LinkList<E> L){
            list = L;
            next = L.head.next;
        }

        public boolean hasNext() {
            if(list.isEmpty() || next == list.head) {
                return false;
            }
            return true;
        }
        public E next(){
            if(!this.hasNext()){
                throw new NoSuchElementException();
            }
            E curr = (E)next;
            next = next.next;
            return curr;
        }
    }
    void printList() {
        ArrayList<E> l = new ArrayList<E>();
        Iterator<E> iter = this.iterator();
        iter.forEachRemaining(l::add);
        System.out.println(l);
    }
}

//  Here is how to use Intrusive LinkList. E wrapped to ListNode<E>
class Slist<E> extends ListNode {
    E val;
    Slist(E val){this.val = val;};
    @Override
    public String toString() {
        return val.toString();
    }
}

public class IntrusiveList{
    public static void main(String[]args) {
        LinkList L = new LinkList();
        Slist<Integer> e = new Slist<Integer>(0);
        L.add(e);
        Integer[] data = {1,2,3,4,5,6,4,0,5,2,8,9};
        ArrayList<Integer> l = new ArrayList<Integer>(Arrays.asList(data));
        for (int i : l.subList(0, l.size() / 2)) {
            L.add(new Slist<Integer>(i));
        }
        for (int i : l.subList(l.size() / 2, l.size())) {
            L.addFirst(new Slist<Integer>(i));
        }
        L.printList();
        L.reverseBetween(3, 8);
        L.printList();
        L.remove(e);
        L.printList();
        System.out.println(String.format("first = %s last = %s length = %d",L.getFirst(),L.getLast(),L.length()));
        ArrayList<Integer> nl = new ArrayList<>();
        Iterable<Slist> iterable = L::iterator;
        Stream<Slist> ls = StreamSupport.stream(iterable.spliterator(),false);
        ls.mapToInt(i -> Integer.parseInt(i.toString()))
            .filter(i -> i < 7)
            .forEach(i ->nl.add(i)); 
        System.out.println(nl);
        nl.sort((a,b) ->a.compareTo(b));
        System.out.println(nl);
        L.reverseBetween(1, L.length());
        L.printList();
        L.reverseBetween(0, L.length());
    }
}
