#include <stdio.h>
#include <string.h>

#define N 20
#define ll __int64

ll a[(1 << N) + 1],aux[(1 << N) +1];
ll sum [N + 1][2];

ll LowerBound(ll a[],ll low, ll high ,ll value){
	ll l, h,pos;
	l = low;
	h = high;
	pos = low - 1;
	while(l <= h) {
		ll mid = (l + h) >> 1;
		if (value > a[mid]) {
			l = mid + 1;
			pos = mid;
		} else {
			h = mid - 1;
		}
	}
	return pos;
}
void Merge(ll a[],ll aux[],ll low, ll mid, ll high) {
	ll i,j,k;
	i = low,j = mid+1,k = low;
	for (; k <= high; k++) {
		if (j > high || i <= mid && aux[i] <= aux[j]) {
			a[k] = aux[i];
			i++;
		} else {
			a[k] = aux[j];
			j++;
		}
	}
}
void Inversions(ll a[],ll aux[], ll low, ll high, ll deep) {
	if (low >= high) {
		return;
	}
	ll mid = (low + high) >> 1;
	Inversions(aux,a,low, mid, deep-1);
	Inversions(aux,a,mid+1, high, deep-1);
	ll count = 0;
	ll i = low,j = mid;
	for (; i <= mid; i++) {
		j = LowerBound(aux,j+1, high, aux[i]);
		count += j - mid;
	}
	sum[deep][0] += count;
	count = 0;
	for (i= mid+1, j = low-1; i <= high; i++) {
		j = LowerBound(aux,j+1, mid, aux[i]);
		count += j - low + 1;
	}
	sum[deep][1] += count;
	if (aux[mid] <= aux[mid+1]) {
		memcpy(a+low,aux+low,sizeof(ll)*(high-low+1));
		return;
	}
	Merge(a,aux,low, mid, high);
}

int main() {
	ll n, m, que,i;
	ll pow[30];pow[0] = 1;
	for(i=1;i<30;i++)pow[i] = pow[i-1]*2;
	while(~scanf("%I64d",&n)){
		memset(sum,0,sizeof sum);
		ll length = pow[n];
		a[0] = length;
		for (i= 1; i <= length; i++){
			scanf("%I64d", &a[i]);
		}
		memcpy(aux,a,sizeof(ll)*(length+1));
		Inversions(a,aux,1, a[0], n);
		scanf("%I64d", &m);
		while(m--) {
			scanf("%I64d", &que);
			ll ans = 0,temp;
			while(que--){
				temp = sum[que+1][0];
				sum[que+1][0] = sum[que+1][1];
				sum[que+1][1] = temp;
			}
			for (i = 0; i <= n; i++) {
				ans += sum[i][0];
			}
			printf("%I64d\n",ans);
		}
	}
	return 0;
}
