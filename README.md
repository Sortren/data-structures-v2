# Common Data Structures Implemented in Go

<img align = "center" alt = "Go" src = "https://user-images.githubusercontent.com/79079000/175411497-2073bf19-73a2-40d5-ab79-31858254a29d.png" />

### Technologies used:

<img align = "left" alt = "Go" width = "26px" src = "https://user-images.githubusercontent.com/79079000/175410045-72365b01-8b43-4fed-9bc1-dc0424848656.png" />

<br />

---

### Linked List

1. Specified in /src/pkg/models/linked_list.go
2. There is FIFO and LIFO functionality, you can push items to the front and append them to the end. So you can use it as a Stack and Queue depends on the need
3. Everything is based on generics so list might be any type
4. After detaching the node from list, the memory is going to be free by garbage collector
5. There are two ways of printing out the list, you can do it explicitly with method "PrintExplicitly()" which will return the every element of list with its bindings such as next and previous nodes (pointers). Also there is a method "Print()" which will print the list as it would normally do in format "[ 1 2 3 4 5 ...]"
6. The list is bidirectional, every node points to the next and previous node.

---
