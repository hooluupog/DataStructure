import 'dart:io';
import 'dart:convert';

List<int> findPivot(int n, List<int> src, List<int> dst) {
  dst.sort();
  int max = 0;
  List<int> ans = [];
  for (var i = 0; i < n; i++) {
    if (src[i] == dst[i] && src[i] > max) {
      ans.add(src[i]);
    }
    if (src[i] > max) {
      max = src[i];
    }
  }
  return ans;
}

void main() async {
  int n;
  var src = new List<int>();
  var s = (await stdin.transform(ASCII.decoder).toList())
      .join()
      .trim()
      .split(new RegExp(r'[\n\r\n\s+]+'));
  s.forEach((i) => src.add(int.parse(i)));
  n = src[0];
  src = src.sublist(1, n+1);
  var dst = new List<int>.from(src);
  var ans = findPivot(src.length, src, dst);
  print(ans.length);
  var buffer = new StringBuffer();
  ans.forEach((i) => buffer.write(' $i'));
  print(buffer.toString().trimLeft());
}
