#include<bits/stdc++.h>
#define MOD 1000000007
#define MAXN 1000000
#define IO ios_base::sync_with_stdio(false);cin.tie(NULL);
#define forf(i,a,b) for(i=a;i<b;i++)
#define forr(i,a,b) for(i=a;i>b;i--)
#define mp make_pair
#define f first
#define s second
#define pb(x) push_back(x)
typedef  long long  int ll;
typedef std::vector<ll> vi;
 
#define input(vec,a,b) for(ll i =a;i<b;i++) cin>>vec[i]
#define print(vec,a,b) for(ll i=a;i<b;i++) cout<<vec[i]<<" " ;cout<<endl;
#define all(a) a.begin(),a.end()
using namespace std;
 

 
void solve(ll Case){
            ll n ,i;
            cin>>n;
            vi a(n);
            input(a,0,n);
            ll ans = 0;
            forf(i,0,n-1)
            {
                ans = max(ans , a[i]*a[i+1]);
            }
            cout<<ans<<endl;
}
 
int main()
{
      IO;
      ll t=1,i;
      cin>>t;
 
 
      for(i=1;i<=t;i++)
      {
          solve(i);
      }
    return 0;
}
