// LeetCode 297.
// Serialize and Deserialize Binary Tree

class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;
  TreeNode(this.val);

  @override
  String toString() => val.toString();
}

// Encodes a tree to a single string.
String serialize(TreeNode root) {
  return root == null ? '#':
    root.val.toString() +',' + serialize(root.left) + ',' + serialize(root.right);
}

// Decodes your encoded data to tree.
TreeNode deserialize(String data) => buildTree(data.split(',').reversed.toList());

TreeNode buildTree(List<String> data) {
  var val = data.removeLast();
  if (val == '#') return null;
  var root = new TreeNode(int.parse(val));
  root.left = buildTree(data);
  root.right = buildTree(data);
  return root;
}

// LeetCode 331
// Verify Preorder Serialization of a Binary Tree

bool isValidSerialization(String preorder) {
  if (preorder == '#') return true;
  var str = preorder.split(',');
  List<String> stack = [];
  for (int i = 0; i < str.length; i++) {
    if (str[i] == '#') {
      if (stack.length == 1 && str.getRange(i, str.length).join() == '##') {
        return true;
      }
      if (stack.length == 0) {
        return false;
      }
      stack.removeLast();
    } else {
      stack.add(str[i]);
    }
  }
  return false;
}

// LeetCode 652
// Find Duplicate Subtrees
List<TreeNode> findDuplicateSubtrees(TreeNode root) {
  List<TreeNode> res = [];
  postOrder(root, new Map(), res);
  return res;
}

String postOrder(TreeNode root, Map<String, int> table, List<TreeNode> res) {
  if (root == null) return '#';
  var serial = postOrder(root.left, table, res) +
      postOrder(root.right, table, res) + root.val.toString();
  if (table.putIfAbsent(serial, () => 0) == 1) res.add(root);
  table[serial] += 1;
  return serial;
}

String preOrder(TreeNode root) {
  return root == null ? '' :
    root.val.toString() + ' ' + preOrder(root.left) + ' ' + preOrder(root.right);
}

void main() {
  var str = '9,3,4,#,#,1,#,#,2,#,6,#,#';
  var str2 = '1,2,4,#,#,#,3,2,4,#,#,#,4,#,#';
  print(isValidSerialization(str2));
  TreeNode t = deserialize(str2);
  var newstr = serialize(t);
  print(newstr);
  print(isValidSerialization(newstr));
  var res = findDuplicateSubtrees(t);
  res.forEach((i) => print(preOrder(i)));
  TreeNode x = new TreeNode(88);
}
