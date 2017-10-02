import 'dart:io';
import 'dart:convert';

int partition(List<int> list, int low, int high) {
  var mid = low + (high - low) ~/ 2;
  var temp = list[low];
  list[low] = list[mid];
  list[mid] = temp; // 枢轴元素交换到首位
  var j = low + 1; // j指向第一个大于枢轴元素的位置
  for (var i = low; i <= high; i++) {
    if (list[low] > list[i]) {
      // 比枢轴元素小的元素交换到前面
      temp = list[i];
      list[i] = list[j];
      list[j++] = temp;
    }
  }
  temp = list[low];
  list[low] = list[j - 1];
  list[j - 1] = temp; // 枢轴元素归位
  return j - 1;
}

/// 快速排序非递归算法实现
void nonrecursiveQSort(List<int> list) {
  List<int> stack = [];
  stack.addAll([0, list.length - 1]);
  while (stack.isNotEmpty) {
    // 栈不空
    var high = stack.removeLast();
    var low = stack.removeLast();
    var pivot = partition(list, low, high); // 一趟划分，返回枢轴下标索引
    if (pivot - 1 > low) {
      // 子序列长度大于1，入栈
      stack.addAll([low, pivot - 1]);
    }
    if (pivot + 1 < high) {
      // 子序列长度大于1，入栈
      stack.addAll([pivot + 1, high]);
    }
  }
}

void main() async {
  List<int> a = [];
  var s = (await stdin.transform(ASCII.decoder).toList())
      .join()
      .trim()
      .split(new RegExp(r'[\n\r\n\s+]+'));
  s.forEach((i) => a.add(int.parse(i)));
  nonrecursiveQSort(a);
  var buffer = new StringBuffer();
  a.forEach((i) => buffer.write(' $i'));
  print(buffer.toString().trimLeft());
}
