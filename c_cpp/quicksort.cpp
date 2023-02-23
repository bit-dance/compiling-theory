#include<iostream>
int Partition(int *arr, int low, int high)
{
    int tmp = arr[low];
    while (low < high)
    {
        while (low < high && arr[high] > tmp)
        {
            high--;
        }
        if (low < high)
        {
            arr[low] = arr[high];
        }
        while (low < high && arr[low] <= tmp)
        {
            low++;
        }
        if (low < high)
        {
            arr[high] = arr[low];
        }
    }
    arr[low] = tmp;
    return low;
}
void Quick(int *arr, int low, int high)
{
    int par = Partition(arr, low, high);
    if (low < par - 1) 
    {
        Quick(arr, low, par - 1);
    }
    if (par + 1 < high) 
    {
        Quick(arr, par + 1, high);
    }
}
void QuickSort(int *arr, int len)
{
    Quick(arr, 0, len - 1); 
}

int main(){
    std::cout<<"Please input 5 numbers:"<<std::endl;
    int a[5];
    for (int i = 0; i < 5; i++)
    {
        scanf("%d",&a[i]);
    }
    QuickSort(a,5);
    for (int i = 0; i < 5; i++)
    {
        std::cout<<a[i]   ;
    }
}