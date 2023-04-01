def minimumDistances(a):
    n=len(a)
    minimum=n+1         #Initializing with n+1 because if no updates take place we should return -1 as there are no matching pairs
    # Note:- If any matching pairs is found the distance would be less than n
    index=dict()      #Creating an empty dictionary
    for i in range(n):
        if a[i] in index:
            if i-index[a[i]] < minimum :    #Check the distance between last occurence and current index of the element
                minimum = i-index[a[i]]
            index[a[i]]=i   #Update the last occurence of the element by current index
        else:
            index[a[i]]=i   #Initialize with first occurence of element
    if minimum==n+1:        #If no matching elements return -1
        return -1
    return minimum


a = [ 7, 1, 3, 4, 1, 7]
print(minimumDistances(a))