class ListNode<T>{
    T _val;
    ListNode<T> _next;
    ListNode([this._val,this._next]);
}

class LinkList<T> {
    ListNode<T> _first;
    LinkList([this._first]);
    LinkList<T> get next => new LinkList<T>(_first._next);

    T get val => _first._val;
    set next(LinkList<T> l) => _first._next = l._first;
    set val(T t) => _first._val = t;

    void addAt(LinkList<T> pre ,T value){ // insert after pre node.
      var node = new ListNode<T>(value);
      if (pre == null){
        _first = node;
      }else{
        node._next = pre.next._first;
        pre.next = new LinkList<T>(node);
      }
    }
  
    void add(T value){
      ListNode<T> pre,l;
      var node = new ListNode<T>(value);
      pre = _first;
      for (l = _first;l != null;l=l._next){
        pre = l;
      }
      if (pre == null){
        _first = node;
      }else{
        pre._next = node;
      }
    }
  
    void remove(LinkList<T> p){
      var head = new ListNode<T>(); // Add a head node for removing operation.
      ListNode<T> pre;
      head._next = _first;
      pre = head;
      for (;pre._next != null && pre._next != p._first;pre = pre._next);
      if(pre._next != null && pre._next == p._first){
        if(pre._next == _first){
          head._next = _first._next;
          _first = head._next;
        }else{
          pre._next = pre._next._next;
        }
      }
    }
  
    void printList(){
      ListNode<T> p = _first;
      var l = [];
      while(p != null){
        l.add(p._val);
        p = p._next;
      }
      print(l);
    }
    bool IsEmpty() => _first == null;

    LinkList<T> reverseBetween(int m, int n) {
      var cnt = n - m + 1;
      ListNode<T> pre, p, q,L; // L is head node of list which is empty node.
      L = new ListNode<T>();
      pre = L;
      p = _first;
      for (var i = 1; i < m; i++) { // p point to the first node to be reversed.
        pre = p;
        p = p._next;
      }
      var reversedEndNode = p; // The last node of reversed subList.
      pre._next = null;      // pre is head node of subList to be reversed.
      while (cnt > 0) {        // insert nodes to head of list.
        var temp = pre._next;
        pre._next = p;
        q = p._next;
        p._next = temp;
        p = q;
        cnt--;
      }
      if (p != null) { // Connect reversed subList to origin list.
        reversedEndNode._next = p;
      }
      if (L != pre) { // Find the head node.
        L._next = _first;
      }
      return new LinkList<T>(L._next);
    }
}

void main(){
    var l = [1.2,2.3,3.4,4.5,5.6,4,0,5.0,2,8,9];
    print(l);
    LinkList L = new LinkList<num>();
    for (var i in l){
    	L.add(i);
    }
    L = L.reverseBetween(3,8);
    L.printList();
    L.remove(L);
    L.printList();
    l.sort((b,a) => a.compareTo(b));
    print(l);
}
