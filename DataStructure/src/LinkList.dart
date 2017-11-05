import 'dart:collection';

class _ListNode<T> {
  T _val;
  _ListNode<T> _next;
  _ListNode([this._val, this._next]);

  @override
  String toString() => '$_val';
}

class LinkList<T> extends Object with IterableMixin<T> {
  _ListNode<T> _first;
  LinkList([this._first]);


  Iterator<T> get iterator => new _LinkListIterator<T>(_first);

  bool insertAfter(LinkList<T> pre, T value) {
    var node = new _ListNode<T>(value);
    if (pre._first == null) {
      return false;
    }
    node._next = pre._first._next;
    pre._first._next = node;
    return true;
  }

  void add(T value) {
    var node = new _ListNode<T>(value);
    var pre = _first;
    for (var l = _first; l != null; l = l._next) {
      pre = l;
    }
    if (pre == null) {
      _first = node;
    } else {
      pre._next = node;
    }
  }

  void remove(LinkList<T> p) {
    var head = new _ListNode<T>(); // Add a head node for removing operation.
    head._next = _first;
    var pre = head;
    for (; pre._next != null && pre._next != p._first; pre = pre._next);
    if (pre._next != null && pre._next == p._first) {
      if (pre._next == _first) {
        head._next = _first._next;
        _first = head._next;
      } else {
        pre._next = pre._next._next;
      }
    }
  }

  LinkList<T> reverseBetween(int m, int n) {
    var cnt = n - m + 1;
    _ListNode<T> pre, p, q, L; // L is a new created head node.
    L = new _ListNode<T>();
    pre = L;
    p = _first;
    for (var i = 1; i < m; i++) {
      // p point to the first node to be reversed.
      pre = p;
      p = p._next;
    }
    var reversedEndNode = p; // The last node of reversed subList.
    pre._next = null; // pre is head node of subList to be reversed.
    while (cnt > 0) {
      // insert nodes to head of list.
      var temp = pre._next;
      pre._next = p;
      q = p._next;
      p._next = temp;
      p = q;
      cnt--;
    }
    if (p != null) {
      // Connect reversed subList to origin list.
      reversedEndNode._next = p;
    }
    if (L != pre) {
      // Find the head node.
      L._next = _first;
    }
    return new LinkList<T>(L._next);
  }
}

class _LinkListIterator<T> extends Iterator<T> {
  _ListNode<T> _curr;
  _LinkListIterator(_ListNode<T> L) {
    _curr = new _ListNode<T>();
    _curr._next = L;
  }
  bool moveNext() {
    if (_curr == null || _curr._next == null) {
      return false;
    }
    _curr = _curr._next;
    return true;
  }

  T get current => _curr == null ? null : _curr._val;
}

void main() {
  var l = [1.2, 2.3, 3.4, 4.5, 5.6, 4, 0, 5.0, 2, 8, 9];
  LinkList L = new LinkList<num>();
  print(L.isEmpty);
  for (var i in l) {
    L.add(i);
  }
  print(L);
  L = L.reverseBetween(3, 8);
  print(L);
  L.remove(L);
  print(L);
  print('first = ${L.first} last = ${L.last} length = ${L.length}');
  var nl = L.where((i) => i < 7).map((i) => i = i * 10).toList();
  print(nl);
  nl.sort((a, b) => a.compareTo(b));
  print(nl);
}
