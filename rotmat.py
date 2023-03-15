def rotate(lst,r):        #This function rotates the passed array by r elements clockwise
    length=len(lst)
    r=r%length            #If r is more than length,we can minimise the work and get the same result by taking mod by length of the array
    temp=lst[length-r:]+lst[:length-r]      #First take all elements from (length-r) to end of the array and then take all the elements before (length-r)
    return temp
 
m,n,r=[int(x) for x in input().split()]
mat=[]
for _ in range(m):
    mat.append([int(x) for x in input().split()])
 
pad=0               #Variable for maintaing the padding
depth=min(m,n)//2   #There would be min(m,n)//2 layers 
layers=[]           #List to store each rotated layer
 
#For each depth start by leaving the padding and go till the padding on other side for each direction down,right,up,left respectively
for index in range(depth):
    i,j=pad,pad
    arr=[]
    while i<(m-pad):
        arr.append(mat[i][j])
        i+=1
    i-=1            #To keep 'i' in range 
    j+=1            #This is necessary to avoid taking same elements again
    while j<(n-pad):
        arr.append(mat[i][j])
        j+=1
    j-=1            #To keep 'j' in range 
    i-=1            #This is necessary to avoid taking same elements again
    while i>=pad:
        arr.append(mat[i][j])
        i-=1
    i+=1            #To keep 'i' in range 
    j-=1            #This is necessary to avoid taking same elements again
    while j>pad:
        arr.append(mat[i][j])
        j-=1
    j+=1            #To keep 'j' in range 
    pad+=1
    layers.append(rotate(arr,r))        #Append the rotated array in 'layers' list
 
#Now repeat the same procedure but instead of taking the elements, we have to insert them in their required place
pad=0
for index in range(depth):
    pointer=0           #This variable is used to take elements one by one
    i,j=pad,pad
    arr=[]
    while i<(m-pad):
        mat[i][j]=layers[index][pointer]
        pointer+=1
        i+=1
    i-=1
    j+=1
    while j<(n-pad):
        mat[i][j]=layers[index][pointer]
        pointer+=1
        j+=1
    j-=1
    i-=1
    while i>=pad:
        mat[i][j]=layers[index][pointer]
        pointer+=1
        i-=1
    i+=1
    j-=1
    while j>pad:
        mat[i][j]=layers[index][pointer]
        pointer+=1
        j-=1
    j+=1
    pad+=1
for i in mat:
    for j in i:
        print(j,end=" ")
    print()
