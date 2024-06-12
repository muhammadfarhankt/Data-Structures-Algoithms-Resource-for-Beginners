# Comprehensive Data Structures and Algorithms Roadmap for Beginners
<details>
	<summary>
		<h2>Basics </h2>
	</summary>
Welcome to the Comprehensive Data Structures and Algorithms (DSA) Roadmap for beginners! This roadmap equips you with a detailed understanding of fundamental concepts in Data Structures (DS) and Algorithms (Algo) during your first week of learning. Each section delves into specific topics with ample sample questions and practical exercises.

## Focus Areas

- Foundational Concepts
- Arrays
- Linked Lists
- Strings
- Searching Algorithms
- Recursion
- JavaScript Built-in Data Structures

## Foundational Concepts

### Data Structures (DS) and Algorithms (Algo)

- Understand how DS organize data efficiently and Algo solve problems with step-by-step instructions.
- Explore real-world examples of DS usage (e.g., shopping lists - arrays, social network connections - graphs).
- **Sample Questions:**
  - What are the different types of Data Structures? Explain their advantages and disadvantages.
  - How do Algorithms help us solve computational problems? Provide examples of common algorithms used in daily life.

### Memory Allocation and Memory Leaks

- Grasp how programs manage memory during execution.
- Understand the concept of memory leaks (unreleased memory) and their impact on program performance.
- **Sample Questions:**
  - Explain the difference between static and dynamic memory allocation.
  - How do memory leaks occur in programs? Describe their consequences and prevention techniques.

### Complexity Analysis

- Learn how to measure the efficiency of algorithms based on time and space complexity.
- Focus on Big-O Notation for asymptotic analysis, understanding how input size affects algorithm performance.
- **Sample Questions:**
  - Define time complexity and space complexity. How do they differ?
  - Explain the concept of Big-O Notation. Analyze the time complexity of simple algorithms like finding the maximum element in an array.
  - Explore different time complexities (Constant, Linear, Logarithmic, Quadratic, Exponential) with code examples.

## Arrays

### Array Concepts

- Master the concept of arrays, their fixed size, and contiguous memory allocation.
- Understand common array operations:
  - Initialization: Creating an array with specific values.
  - Accessing elements using indices.
  - Modifying elements (Set, Update).
  - Traversing through all elements (iteration).
  - Inserting elements at specific positions.
  - Deleting elements from specific positions.
- **Sample Questions:**
  - Implement functions to initialize an array with user-defined values and display its contents.
  - Write code to find the sum or average of elements in an array.
  - Practice inserting an element at the beginning, middle, or end of an array (shifting elements if needed).
  - Implement a function to delete an element at a specific position and handle cases like deleting the first or last element.
  - Reverse the order of elements in an array.

### Leetcode Questions



## Linked Lists

### Introduction

- Understand linked lists, their dynamic nature, and non-contiguous memory allocation.
- Explore different types of linked lists: Singly Linked List (one pointer per node), Doubly Linked List (two pointers per node), Circular Linked List (tail points back to head).
- **Sample Questions:**
  - Differentiate between arrays and linked lists. Explain the advantages and disadvantages of each.
  - Illustrate the concept of nodes in a linked list with diagrams.
  - Describe the differences between Singly, Doubly, and Circular Linked Lists.

### Building Linked Lists

- Implement functions to create Singly and Doubly linked lists in your chosen programming language.
- Understand how nodes are connected through pointers.
- **Sample Questions:**
  - Write code to create a Singly Linked List with a head node containing a specific value.
  - Implement a function to insert a new node at the beginning of a Singly Linked List.
  - Practice creating a Doubly Linked List with functionalities to add a node at the end.

### Operations

- Master linked list operations for both Singly and Doubly Linked Lists:
  - Initialization (creating an empty list).
  - Accessing elements (consider limitations due to non-indexed nature).
  - Modifying elements (updating data within a node).
  - Traversing through the list (iterating using pointers).
  - Inserting elements at specific positions (handling edge cases like inserting at the beginning or end).
  - Deleting elements with a specific value or at a specific position.
- **Sample Questions:**
  - Implement a function to traverse a Singly Linked List and print the data of each node.
  - Write code to delete the head node, a specific node in the middle, or the last node in a Singly Linked List.
  - Implement a function to reverse a Singly Linked List (iterative and recursive approaches).
  - (For Doubly Linked Lists) Practice inserting a node before or after a specific node with a given value.
  - (For Doubly Linked Lists) Write code to delete a node by just its reference (without searching for its value).

### Conversion

- Implement functions to convert an array to a linked list and vice versa.
- Write code to take an array of integers and create a Singly Linked List with those elements.
- Practice converting a Singly Linked List back into an array, preserving the element order.

### Comparison

- **Sample Questions :**
  - Discuss the scenarios where arrays are preferable over linked lists and vice versa.
  - Analyze the time and space complexity of common operations (access, insertion, deletion) for both arrays and linked lists.
 
### [15 LeetCode problems to get good at Linked List:](https://www.linkedin.com/posts/ashishps1_15-leetcode-problems-to-get-good-at-linked-activity-7206294158251102208-xQqD?utm_source=share&utm_medium=member_desktop)

## Strings

### String Fundamentals

- Explore strings as data structures, understand primitive vs. object strings.
- Grasp common string operations:
  - Initialization: Creating a string with specific characters.
  - Accessing characters using indices.
  - Modifying characters (limited in most languages).
- Finding the length of a string.
  - Concatenation: Joining two or more strings.
  - Searching for substrings within a string.
  - Extracting substrings from a string.
  - String comparison (lexicographic order).
- **Sample Questions:**
  - Implement functions to initialize a string with user input and display its characters.
  - Write code to find the first or last occurrence of a specific character within a string.
  - Practice extracting a substring from a string based on starting and ending indices.
  - Implement a function to compare two strings lexicographically (alphabetical order).
  - Write code to reverse a string (iterative and recursive approaches).

### Sample Workouts

- Explore string manipulation techniques like replacing characters, finding the frequency of each character, etc.

## Searching Algorithms

### Linear Search

- Understand the concept of linear search, iterating through a list to find a specific element.
- **Sample Questions:**
  - Implement a function for linear search in arrays.
  - Analyze the time complexity of linear search (worst-case scenario).
  - Practice using linear search to find an element in a Singly Linked List (consider the limitations).

### Binary Search

- Learn the efficient binary search algorithm for sorted arrays, repeatedly halving the search space.
- **Sample Questions:**
  - Implement a function for binary search in sorted arrays.
  - Explain why binary search only works on sorted arrays.
  - Analyze the time complexity of binary search (logarithmic).
  - Practice using binary search to find an element in a sorted Singly Linked List (potentially converting it to an array first).

## Recursion

### Understanding Recursion

- Grasp the concept of recursive functions, where a function calls itself.
- Understand the importance of base cases to prevent infinite recursion.
- **Sample Questions:**
  - Explain the concept of recursion with a simple example (e.g., factorial calculation).
  - Identify potential issues with recursion (stack overflow errors) and how to avoid them.

### Sample Workouts

- Implement functions using recursion for problems like calculating factorial, finding Fibonacci numbers, performing a depth-first search on a tree (advanced).

## JavaScript Built-in Data Structures

### Arrays

- Explore built-in methods like:
  - push(): Add an element to the end of an array.
  - pop(): Remove the last element from an array.
  - shift(): Remove the first element from an array.
  - unshift(): Add an element to the beginning of an array.
  - forEach(): Execute a function for each element.
  - map(): Create a new array with elements transformed by a function.
  - filter(): Create a new array with elements that pass a test implemented by a function.
  - reduce(): Reduce an array to a single value using a provided function.
  - concat(): Merge two or more arrays.
  - slice(): Extract a section of an array.
  - splice(): Add/remove.
- **Sample Questions (Continued):**
  - Write code to use `forEach` to iterate through an array and print each element.
  - Practice using `map` to create a new array with squares of all elements in the original array.
  - Implement a function using `filter` to find all even numbers in an array.
  - Explore using `reduce` to find the sum or average of elements in an array.

### Objects

- Understand object operations:
  - Creating objects with key-value pairs.
  - Accessing properties using dot notation or bracket notation.
  - Modifying property values.
  - Adding or removing properties.
  - Checking if a property exists.
- Explore built-in methods like:
  - Object.keys(): Get an array of all object property names.
  - Object.values(): Get an array of all object property values.
  - Object.entries(): Get an array of key-value pairs as arrays.
- **Sample Questions:**
  - Implement code to create an object representing a person with properties like name, age, and city.
  - Write code to access and modify a specific property value within an object.
  - Practice using `Object.keys` to iterate through all properties of an object and print their values.
  - Explore using `Object.entries` to create a new array containing key-value pairs from an object.

## Bonus Topics (for curious minds)

- Linear vs. Non-linear Data Structures (e.g., Arrays vs. Trees, Graphs)
- Contiguous vs. Non-contiguous Memory Allocation (related to Arrays vs. Linked Lists)
- Stack vs. Heap Memory (different memory management regions)
- Garbage Collection (automatic memory management in languages like JavaScript)
- Jagged Arrays (arrays of arrays)
- Pros and Cons of Recursion (efficiency considerations)
- Factorial, Fibonacci, Prime Number Calculations (with and without recursion)

## Remember

- Practice consistently!
- Experiment with code examples and test your understanding with additional problems.
- Refer to online tutorials and visualizations for better comprehension.
 
</details>

<details>
	<summary>
		<h2>Intermediate</h2>
	</summary>

#### This roadmap equips you with fundamental knowledge of sorting algorithms, stacks & queues, and hash tables in your second week of learning Data Structures (DS) and Algorithms (Algo).

## Sorting Algorithms

### Understanding Sorting: 
Grasp the concept of sorting algorithms for rearranging elements based on a specific order (ascending or descending). Explore various techniques for different scenarios.

### Sorting Techniques: 
We'll cover five common sorting algorithms, each with complexities, advantages, and disadvantages:
- **Bubble Sort**: Simple but inefficient for large datasets - quadratic time complexity.
- **Insertion Sort**: Iterative approach, good for partially sorted data - average case closer to linear time complexity.
- **Selection Sort**: Repeatedly finding minimum/maximum element and placing it - quadratic time complexity.

### Sample Questions: 
* Implement functions for Bubble Sort, Insertion Sort, and Selection Sort.
* Analyze the time complexity (best, average, worst case) of each sorting algorithm.
* Discuss the strengths and weaknesses of each sorting technique.

### Advanced Sorting Techniques: 
Explore more complex but efficient sorting algorithms for larger datasets:
- **Quick Sort**: Divide-and-conquer strategy, pivoting elements to partition the data - average near linear time complexity, but can degrade to worst-case quadratic.
- **Merge Sort**: Another divide-and-conquer approach, recursively dividing data until single elements, then merging sorted sub-lists - near linear time complexity.

### Sample Questions: 
* Implement functions for Quick Sort and Merge Sort.
* Analyze the time complexity of Quick Sort and Merge Sort.
* Compare the performance of Quick Sort and Merge Sort in different scenarios (stability, in-place vs. extra space).

### Sorting Algorithm Comparison: 
Understand when to choose specific sorting algorithms based on data size, desired order, and stability requirements.

## Stacks and Queues

### Stack and Queue Concepts: 
Grasp the concept of stacks (LIFO - Last In First Out) and queues (FIFO - First In First Out) as linear data structures. Explore their real-world applications.

### Stack Operations: 
Implement stacks using arrays and linked lists. Master stack operations:
- **Push**: Add an element to the top of the stack.
- **Pop**: Remove and return the top element from the stack.
- **Peek**: Access the top element without removing it.
- **IsEmpty**: Check if the stack is empty.
- **IsFull**: Check if the stack is full (relevant for array implementation).

### Sample Questions: 
* Write code to implement a stack using an array and perform push, pop, peek, and isEmpty operations.
* Implement a stack using a linked list with similar functionalities.

### Queue Operations: 
Implement queues using arrays and linked lists. Master queue operations:
- **Enqueue**: Add an element to the back of the queue.
- **Dequeue**: Remove and return the front element from the queue.
- **Peek**: Access the front element without removing it.
- **IsEmpty**: Check if the queue is empty.
- **IsFull**: Check if the queue is full (relevant for array implementation).

### Sample Questions: 
* Write code to implement a queue using an array and perform enqueue, dequeue, peek, and isEmpty operations.
* Implement a queue using a linked list with similar functionalities.

### Advanced Stack and Queue Concepts: 
Explore additional stack and queue concepts:
- **Circular Queue**: A queue where the last element wraps around to the beginning.
- **Priority Queue**: A queue where elements are prioritized based on a specific value (e.g., importance).

### Sample Questions: 
* Implement a circular queue using an array, handling the wrapping behavior.
* Discuss the use cases of priority queues and how they prioritize elements.

### Stack and Queue Applications: 
Understand the various applications of stacks and queues in real-world scenarios (e.g., function call stack, undo/redo functionality, task scheduling).

## Hash Tables

### Hash Table Fundamentals: 
Grasp the concept of hash tables, a data structure for efficient key-value lookup. Understand how hash tables use hash functions to map keys to unique indices in an array. Explore the advantages of hash tables for fast retrieval based on keys.

### Hash Functions: 
Learn about hash functions, algorithms that convert keys into unique indices within a specific range. Explore different hash function properties like efficiency and collision avoidance.

### Sample Questions: 
* Implement a simple hash function for strings (e.g., sum of character codes modulo table size).

### Hash Table Implementation: 
Explore ways to implement hash tables using arrays. Understand how to handle collisions (situations where multiple keys map to the same index).

### Sample Questions (continued): 
* Implement a basic hash table with separate chaining for collision handling (storing colliding elements in a linked list at the corresponding index).

### Collision Handling Methods:

Deep dive into various collision handling techniques:
- **Separate Chaining**: Storing colliding elements in a linked list at the index.
- **Open Addressing**: Probing for the next available slot in the array when a collision occurs (linear probing, quadratic probing, double hashing).

### Sample Questions (continued): 
* Implement a hash table with linear probing for collision handling.
* Discuss the trade-offs between separate chaining and open addressing techniques.

### Perfect Hashing (Optional):
Briefly explore the concept of perfect hashing, where a hash function guarantees no collisions (advanced topic).

### Re-Hashing: 
Understand the concept of re-hashing, resizing the hash table when the load factor (number of elements divided by table size) becomes too high.

### Hash Table vs. Set: 
Compare and contrast hash tables and sets, understanding their key differences and use cases.

### Sample Questions (continued): 
* Discuss scenarios where a hash table might be preferable over a set, and vice versa.

### Hash Key vs. Array Key: 
Differentiate between hash keys (used for lookup) and array indices (fixed positions in an array).

### Dynamic Restructuring: 
Explore how hash tables can dynamically resize themselves to maintain efficiency.

### Week Set, Week Map (Optional): 
Briefly discuss week set and week map as specialized hash table implementations (advanced topic).

### Collision Handling - Deep Dive (Optional): 
For interested learners, delve deeper into specific collision handling techniques like:
- **Linear Probing**: Probing for the next available slot in a linear fashion.
- **Quadratic Probing**: Probing with a quadratic function to reduce clustering of collided elements.
- **Double Hashing**: Using a secondary hash function to probe for a different set of indices in case of a collision.

### Clustering: 
Understand the concept of clustering in hash tables, where collisions tend to group together, impacting performance.

### Advanced Collision Handling Techniques (Optional): 
Explore advanced collision handling techniques like:
- **Cuckoo Hashing**: Utilizing two hash tables to resolve collisions.
- **Robin Hood Hashing**: Stealing elements from less loaded buckets to improve balance.

### SHA - Secure Hashing Algorithm (Optional): 
Briefly introduce the concept of secure hashing algorithms like SHA, used for data integrity and security purposes.

## Remember:

- Practice implementing hash tables with different collision handling techniques.
- Experiment with various hash functions and analyze their impact on performance.
- Refer to online resources for further exploration of advanced hashing concepts.
- This roadmap equips you with a solid foundation for understanding sorting algorithms, stacks & queues, and hash tables. Keep practicing and

 ### Sets

- Understand Sets, collections of unique values.
- Explore common Set operations:
  - Adding elements (add()).
  - Checking if an element exists (has()).
  - Removing elements (delete()).
  - Finding the size of the Set (size()).
  - Removing all elements (clear()).
- **Sample Questions:**
  - Write code to create a Set containing unique names from an array of strings.
  - Implement a function to check if a specific element exists in a Set.
  - Practice removing duplicate elements from an array using Sets.

### Maps

- Understand Maps, collections that use key-value pairs (like objects but can have any data type as keys).
- Explore common Map operations (similar to Sets):
  - Setting key-value pairs (set()).
  - Getting the value for a key (get()).
  - Checking if a key exists (has()).
  - Removing a key-value pair (delete()).
  - Finding the size of the Map (size()).
  - Removing all elements (clear()).
- **Sample Questions:**
  - Implement code to create a Map where keys are student IDs and values are student names.
  - Write a function to retrieve the name of a student given their ID (using a Map).
  - Practice using Maps to store configuration settings with key-value pairs.

### Comparison

- **Sample Questions:**
  - Discuss the use cases for Arrays vs. Sets and Objects vs. Maps.
  - Analyze the time complexity of common operations (add, remove, search) for Arrays, Sets, and Maps.


</details>

<details>
	<summary>
		<h2>Advanced</h2>
	</summary>

## Trees

### Linear vs. Non-Linear vs. Hierarchical Data Structures
- Review linear data structures (arrays, linked lists) for sequential access.
- Introduce non-linear data structures (trees, graphs) for hierarchical relationships.
- Understand hierarchical structures with parent-child relationships.

### Tree Fundamentals
- Grasp the concept of trees, a collection of nodes connected by edges.
- Explore tree terminology: parent, child, root, leaf, sibling, ancestor, descendant, path, distance, degree, depth, height, edge, subtree.

### Types of Trees (by Nodes)
- Binary Tree: Each node has at most two children (left and right).
- Ternary Tree: Each node has at most three children.
- K-ary Tree: Each node has at most K children.
- Threaded Binary Tree: A space-efficient binary tree variation with implicit pointers.

### Types of Trees (by Structure)
- Complete Tree: All levels except possibly the last are completely filled.
- Full Tree: Every node except possibly leaves has two children.
- Perfect Tree: Every internal node has two children and all leaves are at the same level.
- Degenerate Tree: A tree where most nodes have only one child.
  - Left-Skew Tree: More nodes lean left than right.
  - Right-Skew Tree: More nodes lean right than left.

### Binary Search Tree (BST)
- BST vs. Binary Tree: Understand the additional properties of a BST.
- Uses of BST: Efficient searching and sorting of data with a specific ordering.
- Balanced vs. Unbalanced Tree: Explore the impact of balance on BST performance.
- Properties of BST: Ordering property (left subtree < root < right subtree).
- BST Operations:
  - Insertion: Maintain BST property while adding new elements.
  - Deletion: Remove an element while preserving BST order.
- Traversal:
  - Depth-First Search (DFS):
    - InOrder: Visit left subtree, root, then right subtree (sorted order for BST).
    - PreOrder: Visit root, then left subtree, then right subtree.
    - PostOrder: Visit left subtree, then right subtree, then root.
  - Breadth-First Search (BFS): Visit nodes level by level.

### Balanced Search Trees
- AVL Tree: A self-balancing BST with a height difference constraint (logarithmic search time).
- Red-Black Tree: Another self-balancing BST with specific node color properties (logarithmic search time).

### Prefix Tree (Trie)
- String vs. Trie: Explore how tries efficiently store and retrieve strings with a prefix search functionality.
- Trie Operations:
  - Initialization: Create an empty trie.
  - Insertion: Insert a new string into the trie.
  - Deletion: Delete a string from the trie (if it exists).
  - Search: Search for a specific string prefix in the trie.
- Prefix and Suffix Trees: Specialized tries for efficient prefix and suffix searches.
- Terminator character: A special character marking the end of a string in the trie.
- Compressed Trie: Techniques for reducing memory usage in tries (e.g., Radix Trie).

## Heaps

### Min Heap vs. Max Heap
- Understand heaps, tree-based structures where the root has the highest (max heap) or lowest (min heap) value compared to its children.
- Heap Operations:
  - Get Value of Children/Parent: Access child or parent node values based on their positions.
  - Initialization/Heapify: Convert an array into a valid heap structure.
  - Insertion: Add a new element to the heap while maintaining the heap property.
  - Deletion: Remove the root element (min/max value) from the heap and re-organize.
- Heapsort: Sorting algorithm utilizing a heap structure for efficient time complexity (average/near worst-case - n log n).

## Graphs

### Graph Fundamentals
- Understand graphs, data structures consisting of vertices (nodes) connected by edges (links) representing relationships.
- Explore graph terminology: vertex, edge, adjacency list, adjacency matrix.

### Types of Graphs
- Directed (Unidirectional): Edges have a direction (from one vertex to another).
- Undirected (Bidirectional): Edges have no direction (connect two vertices).
- Cyclic: A graph containing a closed loop (cycle) of vertices.
- Disconnected: A graph where some vertices are not reachable from others.
- Weighted Graph: Edges have associated weights (costs).
- Unweighted Graph: Edges have no weights (all connections are considered equal).
- Bipartite Graph: A graph where vertices can be divided into two sets such that no edges connect vertices within the same set.

### Graph Traversals
- Breadth-First Search (BFS): Explore vertices level by level, starting from a source vertex.
- Depth-First Search (DFS): Explore vertices along a path until a dead end is reached, then backtrack and explore another path.

### Applications of Graphs
- Modeling networks (social, computer, transportation).
- Route finding (GPS navigation).
- Task scheduling (dependency relationships).
- Minimum spanning tree (finding the most efficient set of connections).

### Additional Graph Concepts
- River Size Problem: Finding the size (number of nodes) of the connected component containing a given vertex.

## Algorithms

Some algorithms heavily utilize the data structures covered this week.

- Greedy Method: An algorithmic approach that makes the optimal choice at each step with the aim of finding a near-optimal solution overall.
- Graph Algorithms:
  - Minimum Spanning Tree (MST) Algorithms:
    - Kruskal's Algorithm: A greedy algorithm to find a MST for a weighted graph.
    - Prim's Algorithm: Another greedy algorithm for finding a MST.
  - Shortest Path Algorithms:
    - Dijkstra's Algorithm: Finding the shortest path between a source vertex and all other reachable vertices in a weighted graph.
    - Bellman-Ford Algorithm: Can handle graphs with negative edge weights (Dijkstra's works for non-negative weights).
  - Topological Sorting: Ordering vertices in a directed acyclic graph (DAG) such that for every directed edge from u to v, u appears before v in the ordering.
  - Floyd-Warshall Algorithm: Finding the shortest paths between all pairs of vertices in a weighted graph.
  - Bipartite Graph Checking: Determining if a graph is a bipartite graph.
  - Max Flow Algorithm (Ford-Fulkerson Algorithm): Finding the maximum flow of data through a network.

### Week 3 Additional Topics

#### Questions and Discussions
- Graph vs. Tree: Understand the key differences and relationships between trees and graphs.
- Forest (in Tree): A collection of disconnected trees.
- Operators:
  - Binary Operators: Operations involving two operands (e.g., +, -, *).
  - Priority: Order of operations based on precedence rules (e.g., multiplication before addition).
  - Infix, Prefix (Polish Notation), Postfix (Reverse Polish Notation): Different ways to represent expressions.
- General Concepts:
  - Logarithms: Understand the concept of logarithms and their applications in computer science (e.g., time complexity analysis).
  - File Structure vs. Data Structure: Differentiate between file structures for data storage and data structures for in-memory data organization.
  - Data Structure Applications: Explore how data structures are used in various programming domains.
  - Void vs. Null: Understand the difference between void (absence of a value) and null (a special pointer value).
  - Dynamic Data Structures: Data structures that can grow or shrink in size at runtime.
  - Dynamic Memory Management/Allocation: Techniques for allocating and freeing memory during program execution.
  - Heap vs. Stack: Understand the differences between heaps (used for dynamic allocation) and stacks (used for function calls and local variables).
  - Pointers in Data Structures:
    - Mastering pointers is crucial for many data structures (especially trees and graphs).
    - Explore how pointers allow efficient memory management and navigation within data structures.
  - Recursive Algorithms:
    - Understand the concept of recursion, a function that calls itself.
    - Explore how recursion can be a powerful tool for solving problems that can be broken down into smaller, self-similar subproblems.
    - Be aware of potential drawbacks of recursion, such as stack overflow for very deep recursion.
  - Divide and Conquer on Recursion:
    - Understand divide-and-conquer, a common algorithmic paradigm that recursively divides a problem into smaller subproblems, solves those subproblems, and combines the solutions.
  - Which is the Fastest Sorting Algorithm Available?
    - The answer depends on factors like data size, pre-sortedness, and memory usage.
    - Heapsort (average/near worst-case - n log n) is often a good choice for general-purpose sorting.
    - Quicksort (average - n log n, but worst-case - n^2) can be faster on average but has a worse worst-case scenario.
    - Merge Sort (n log n) is generally slower than Heapsort or Quicksort on average but has a guaranteed n log n time complexity.
  - Multi-Linked Lists:
    - A data structure where each node can have multiple pointers to other nodes.
    - Useful for representing complex relationships between data elements.
  - Sparse Matrices:
    - Matrices where most elements are zero.
    - Special storage techniques can be used to efficiently represent and manipulate sparse matrices.
  - Disadvantages of Implementing Queues Using Arrays:
    - Fixed size: Arrays cannot grow or shrink dynamically, limiting flexibility.
    - Queue overflow/underflow: Handling these conditions can be complex with arrays.
  - Void Pointer:
    - A pointer that can point to any data type.
    - Useful for generic programming techniques.
  - Lexical Analysis:
    - The process of breaking down text into meaningful units (tokens) like keywords, identifiers, operators.
  - Lexeme:
    - A meaningful sequence of characters identified during lexical analysis (e.g., a keyword like "if").
  - Pattern Matching:
    - Finding specific patterns (sequences of characters or symbols) within text or data.
  - Closest Path (Graph):
    - Finding the shortest path between two vertices in a graph.
    - Often solved using graph traversal algorithms like Dijkstra's algorithm.
  - Degree of the Node (Graph):
    - The number of edges connected to a node in a graph.
  - Spanning Tree:
    - A subgraph of a graph that connects all its vertices without cycles.
  - Minimum Spanning Tree (MST):
    - A spanning tree with the minimum total edge weight.
    - Useful for finding the most efficient set of connections in a weighted graph (e.g., Kruskal's or Prim's algorithm).
  - AVL Tree:
    - A self-balancing binary search tree with a maximum height difference of 1 between subtrees, ensuring efficient search time (logarithmic).
  - B-Tree:
    - A self-balancing tree designed for efficient storage and retrieval of large datasets, particularly useful for databases.
  - Full Tree:
    - Every node except possibly leaves has two children.
  - Complete Tree:
    - All levels except possibly the last are completely filled.
  - Perfect Tree:
    - Every internal node has two children and all leaves are at the same level.
  - Heap Applications:
    - Priority queues (e.g., scheduling tasks based on priority).
    - Heapsort (efficient sorting algorithm).
  - BFS Complexity:
    - Breadth-First Search has a time complexity of O(V + E), where V is the number of vertices and E is the number of edges in the graph.
  - Shortest Path Algorithm:
    - Dijkstra's algorithm and Bellman-Ford algorithm are commonly used shortest path algorithms with varying complexities depending on the graph type (weighted/unweighted, presence of negative edge weights).
  - Dijkstra's Algorithm:
    - Finds the shortest paths from a source vertex to all reachable vertices in a weighted graph with non-negative edge weights. Time complexity: O(V^2) in the worst case, but often performs better in practice (average complexity depends on the graph structure).
  - Bellman-Ford Algorithm:
    - Can handle graphs with negative edge weights. Time complexity: O(V * E).
  - Topological Sorting:
    - Ordering vertices in a directed acyclic graph (DAG) such that for every directed edge from u to v, u appears before v in the ordering. Time complexity: O(V + E).
  - Acyclic Travel:
    - Traversing acyclic
    - Acyclic graphs (graphs without cycles) can be traversed efficiently using algorithms like topological sorting.
  - Graph vs. Tree:
    - Trees are hierarchical structures with a single root node and parent-child relationships.
    - Graphs can be more general, allowing for cycles and representing arbitrary relationships between nodes.
  - Additional Types of Graphs:
    - Complete Graph: Every pair of vertices is connected by an edge.
    - Graph Indexing: Techniques for efficiently searching and retrieving data within a graph structure.
    - Representing Graphs in Memory: Different approaches to store graphs in memory using adjacency lists or adjacency matrices.
  - Cycles Detection:
    - Algorithms like depth-first search (DFS) can be used to detect cycles in graphs.
  - Practical Questions Asked:
    - Be prepared for interview-style questions that test your understanding of data structures and algorithms covered in Week 3. This could involve implementing algorithms, analyzing time and space complexity, or explaining trade-offs between different data structures.

</details>

<details>
	<summary><h2>LeetCode Practice Questions by Week</h2></summary>

## Week 1: Arrays, Strings, Linked Lists

### Arrays (Easy):
1. [Two Sum](https://leetcode.com/problems/two-sum/) (1)
2. [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) (3)
3. [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) (121)
4. [Missing Number](https://leetcode.com/problems/missing-number/) (268)
5. [Move Zeroes](https://leetcode.com/problems/move-zeroes/) (283)

### Arrays (Medium):
6. [3Sum](https://leetcode.com/problems/3sum/) (15)
7. [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) (26)
8. [Rotate Array](https://leetcode.com/problems/rotate-array/) (189)
9. [Merge Intervals](https://leetcode.com/problems/merge-intervals/) (56)
10. [Majority Element](https://leetcode.com/problems/majority-element/) (169)

### Strings (Easy):
11. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) (3)
12. [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) (13)
13. [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) (20)
14. [Reverse String](https://leetcode.com/problems/reverse-string/) (344)
15. [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) (14)

### Strings (Medium):
16. [Group Anagrams](https://leetcode.com/problems/group-anagrams/) (49)
17. [Edit Distance](https://leetcode.com/problems/edit-distance/) (72)

## Week 2: Sorting Algorithms, Stacks & Queues, Hash Tables

### Sorting Algorithms (Easy):
18. [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) (21)
19. [Majority Element](https://leetcode.com/problems/majority-element/) (169)
20. [Merge Sorted Arrays](https://leetcode.com/problems/merge-sorted-array/) (88)

### Sorting Algorithms (Medium):
21. [Sort Colors](https://leetcode.com/problems/sort-colors/) (75)
22. [Sort an Array](https://leetcode.com/problems/sort-an-array/) (912)
23. [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) (215)

### Stacks (Easy):
24. [Min Stack](https://leetcode.com/problems/min-stack/) (155)
25. [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) (232)
26. [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) (102)

### Queues (Easy):
27. [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) (232)
28. [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) (102)
29. [Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/) (225)

### Hash Tables (Easy):
30. [Two Sum](https://leetcode.com/problems/two-sum/) (1)
31. [Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/) (205)
32. [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) (347)

### Hash Tables (Medium):
33. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) (3)
34. [Group Anagrams](https://leetcode.com/problems/group-anagrams/) (49)
35. [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) (208)

## Week 3: Trees, Binary Search Trees, Heaps, Tries, Graphs

### Trees (Easy):
36. [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) (94)
37. [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) (101)
38. [Path Sum](https://leetcode.com/problems/path-sum/) (112)

### Binary Search Trees (Easy):
39. [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) (98)
40. [Same Tree](https://leetcode.com/problems/same-tree/) (100)
41. [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) (230)

### Binary Search Trees (Medium):
42. [Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/) (700)
43. [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) (235)

### Heaps (Easy):
44. [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) (215)
45. [Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/) (371)

### Heaps (Medium):
46. [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)
47. [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)

### Tries (Easy):
48. [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) (208)
49. [Word Search II](https://leetcode.com/problems/word-search-ii/)

### Graphs (Easy):
50. [Clone Graph](https://leetcode.com/problems/clone-graph/) (133)
51. [Number of Islands](https://leetcode.com/problems/number-of-islands/)

### Graphs (Medium):
52. [Word Ladder](https://leetcode.com/problems/word-ladder/)
53. [Is Graph Bipartite?](https://leetcode.com/problems/is-graph-bipartite/)

</details>

# LeetCode Grind Resource for Beginners!

### Hey folks! 👋 If you're diving into the world of LeetCode and looking for a structured path to improve your skills, I've curated a list of questions in the order of increasing complexity. 📈

### This resource is specifically designed for beginners, making your LeetCode journey a bit smoother. 💡 Feel free to check out the list and let me know if you have any questions or suggestions! 🚀

### REMEMBER Nothing worth having comes easy! CONSISTENCY is the Key. Do atlest one question every day							
							

<details>
	<summary> <strong> Math </strong> </summary>	
	
1. [`2235. Add Two Integers`](./Golang/Leetcode%202235%20Add%20Two%20Integers.go) : Simplest Leetcode Question
2. [`412. Fizz Buzz`](./Golang/Leetcode%20412%20Fizz%20Buzz%20Golang.go)
3. [`2469 Convert the Temperature`](./Golang/Leetcode%202469%20Convert%20the%20Temperature%20Golang%20Solution.go)
4. [`1952. Three Divisors`](./Golang/Leetcode%201952.%20Three%20Divisors.go)
5. [`2455. Average Value of Even Numbers That Are Divisible by Three`](./Golang/Leetcode%202455.%20Average%20Value%20of%20Even%20Numbers%20That%20Are%20Divisible%20by%20Three.go)
6. [`3028. Ant on the Boundary`](./Golang/Leetccode%203028.%20Ant%20on%20the%20Boundary.go)
7. [`1313. Decompress Run-Length Encoded List`](./Golang/Leetcode%201313.%20Decompress%20Run-Length%20Encoded%20List.go)
8. [`3099. Harshad Number`](./Golang/Leetcode%203099.%20Harshad%20Number.go)
9. [`507. Perfect Number`](./Golang/Leetcode%20507.%20Perfect%20Number.go)
10. [`1614. Maximum Nesting Depth of the Parentheses`](./Golang/Leetcode%201614.%20Maximum%20Nesting%20Depth%20of%20the%20Parentheses.go)
11. [`657. Robot Return to Origin`](./Golang/Leetcode%20657.%20Robot%20Return%20to%20Origin.go)
12. [`367. Valid Perfect Square`](./Golang/Leetcode%20367.%20Valid%20Perfect%20Square.go)
13. [`561. Array Partition`](./Golang/Leetcode%20561.%20Array%20Partition.go)
14. [`2833. Furthest Point From Origin`](./Golang/Leetcode%202833.%20Furthest%20Point%20From%20Origin.go) : You can use if else condition if didn't know hashmaps
15. [`2427. Number of Common Factors`](./Golang/Leetcode%202427%20Number%20of%20Common%20Factors.go)
16. [`1979. Find Greatest Common Divisor of Array`](./Golang/Leetcode%201979.%20Find%20Greatest%20Common%20Divisor%20of%20Array.go)
17. [`2974. Minimum Number Game`](./Golang/Leetcode%202974.%20Minimum%20Number%20Game.go)
18. [`9. Palindrome Number`](./Golang/Leetcode%209%20Palindrome%20Number.go)
19. [`1281. Subtract the Product and Sum of Digits of an Integer`](./Golang/Leetcode%201281%20Subtract%20the%20Product%20and%20Sum%20of%20Digits%20of%20an%20Integer.go)
20.  [`2413. Smallest Even Multiple`](./Golang/Leetcode%202413%20Smallest%20Even%20Multiple.go)
21.  [`1431. Kids With the Greatest Number of Candies`](./Golang/Leetcode%201431.%20Kids%20With%20the%20Greatest%20Number%20of%20Candies.go)
22.  [`2706. Buy Two Chocolates`](./Golang/Leetcode%202706%20Buy%20Two%20Chocolates.go)
23.  [`268. Missing Number`](./Golang/Leetcode%20268.%20Missing%20Number.go)
24.  [`383. Ransom Note`](./Golang/Leetcode%20383.%20Ransom%20Note.go)
25.  [`896. Monotonic Array`](./Golang/Leetcode%20896.%20Monotonic%20Array.go)
26.  [`2965. Find Missing and Repeated Values`](./Golang/Leetcode%202965.%20Find%20Missing%20and%20Repeated%20Values.go)
27.  [`2894. Divisible and Non-divisible Sums Difference`](./Golang/Leetcode%202894%20Divisible%20and%20Non-divisible%20Sums%20Difference.go)
28.  [`2769. Find the Maximum Achievable Number`](./Golang/Leetcode%202769%20Find%20the%20Maximum%20Achievable%20Number.go)
29.  [`2535. Difference Between Element Sum and Digit Sum of an Array`](./Golang/Leetcode%202535%20Difference%20Between%20Element%20Sum%20and%20Digit%20Sum%20of%20an%20Array.go)
30.  [`2544. Alternating Digit Sum`](./Golang/Leetcode%202544%20Alternating%20Digit%20Sum.go)
31.  [`2154. Keep Multiplying Found Values by Two`](./Golang/Leetcode%202154.%20Keep%20Multiplying%20Found%20Values%20by%20Two.go)
32.  [`1351. Count Negative Numbers in a Sorted Matrix`](./Golang/Leetcode%201351.%20Count%20Negative%20Numbers%20in%20a%20Sorted%20Matrix.go)
33.  [`1317. Convert Integer to the Sum of Two No-Zero Integers`](./Golang/Leetcode%201317.%20Convert%20Integer%20to%20the%20Sum%20of%20Two%20No-Zero%20Integers.go)
34.  [`1720. Decode XORed Array`](./Golang/Leetcode%201720.%20Decode%20XORed%20Array.go)
35.  [`2574. Left and Right Sum Differences`](./Golang/Leetcode%202574.%20Left%20and%20Right%20Sum%20Differences.go)
36.  [`3000. Maximum Area of Longest Diagonal Rectangle`](./Golang/Leetcode%203000.%20Maximum%20Area%20of%20Longest%20Diagonal%20Rectangle.go)
37.  [`191. Number of 1 Bits`](./Golang/Leetcode%20191.%20Number%20of%201%20Bits.go)
38.  [`2859. Sum of Values at Indices With K Set Bits`](./Golang/Leetcode%202859.%20Sum%20of%20Values%20at%20Indices%20With%20K%20Set%20Bits.go)
39.  [`509. Fibonacci Number`](./Golang/Leetcode%20509.%20Fibonacci%20Number.go)
40.  [`70. Climbing Stairs`](./Golang/Leetcode%2070.%20Climbing%20Stairs.go) : Similiar to Fibonacci
41.  [`231. Power of Two`](./Golang/Leetcode%20231.%20Power%20of%20Two.go)
42.  [`326. Power of Three`](./Golang/Leetcode%20326.%20Power%20of%20Three.go)
43.  [`342. Power of Four`](./Golang/Leetcode%20342.%20Power%20of%20Four.go)
44.  [`35. Search Insert Position`](./Golang/Leetcode%2035%20Search%20Insert%20Position.go) : Binary Search Implementation
45.  [`455. Assign Cookies`](./Golang/Leetcode%20455%20Assign%20Cookies.go)
46.  [`1385. Find the Distance Value Between Two Arrays`](./Golang/Leetcode%201385.%20Find%20the%20Distance%20Value%20Between%20Two%20Arrays.go)
47.  [`121. Best Time to Buy and Sell Stock`](./Golang/Leetcode%20121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock.go)
48.  [`1588. Sum of All Odd Length Subarrays`](./Golang/Leetcode%201588%20Sum%20of%20All%20Odd%20Length%20Subarrays.go)
49.  [`645. Set Mismatch`](./Golang/Leetcode%20645%20Set%20Mismatch.go)
50.  [`977. Squares of a Sorted Array`](./Golang/Leetcode%20977.%20Squares%20of%20a%20Sorted%20Array.go)
51.  [`628. Maximum Product of Three Numbers`](./Golang/Leetcode%20628%20Maximum%20Product%20of%20Three%20Numbers.go)
52.  [`414. Third Maximum Number`](./Golang/Leetcode%20414.%20Third%20Maximum%20Number.go)
53.  [`2119. A Number After a Double Reversal`](./Golang/Leetcode%202119%20A%20Number%20After%20a%20Double%20Reversal.go)
54. [`1304. Find N Unique Integers Sum up to Zero`](./Golang/Leetcode%201304%20Find%20N%20Unique%20Integers%20Sum%20up%20to%20Zero.go)
55. [`2475. Number of Unequal Triplets in Array`](./Golang/Leetcode%202475%20Number%20of%20Unequal%20Triplets%20in%20Array.go)
56. [`1688. Count of Matches in Tournament`](./Golang/Leetcode%201688%20Count%20of%20Matches%20in%20Tournament.go)
57. [`389. Find the Difference`](./Golang/Leetcode%20389%20Find%20the%20Difference%20Golang%20Solution.go)
58. [`1512. Number of Good Pairs`](./Golang/Leetcode%201512%20Number%20of%20Good%20Pairs.go)
59.  [`2180. Count Integers With Even Digit Sum`](./Golang/Leetcode%202180%20Count%20Integers%20With%20Even%20Digit%20Sum.go)
60.  [`7. Reverse Integer`](./Golang/Leetcode%207%20Reverse%20Integer.go)
61.  [`1710. Maximum Units on a Truck`](./Golang/Leetcode%201710.%20Maximum%20Units%20on%20a%20Truck.go)
62.  [`66. Plus One`](./Golang/Leetcode%2066%20Plus%20One.go)
63.  [`2824. Count Pairs Whose Sum is Less than Target`](./Golang/Leetcode%202824%20Count%20Pairs%20Whose%20Sum%20is%20Less%20than%20Target.go)
64.  [`2540. Minimum Common Value`](./Golang/Leetcode%202540.%20Minimum%20Common%20Value.go) : Two pointer approach
65.  [`442. Find All Duplicates in an Array`](./Golang/Leetcode%20442.%20Find%20All%20Duplicates%20in%20an%20Array.go) : Medium - Easy level
66.  [`2807. Insert Greatest Common Divisors in Linked List`](./Golang/Leetcode%202807%20Insert%20Greatest%20Common%20Divisors%20in%20Linked%20List.go) : Medium Question but Medium - Easy level
67.  [`2125. Number of Laser Beams in a Bank`](./Golang/Leetcode%202125%20Number%20of%20Laser%20Beams%20in%20a%20Bank.go) : Medium - Easy level
68.  [`2870. Minimum Number of Operations to Make Array Empty`](./Golang/Leetcode%202870%20Minimum%20Number%20of%20Operations%20to%20Make%20Array%20Empty.go) : Medium - Easy level
69.  [`2396. Strictly Palindromic Number.go`](./Golang/Leetcode%202396.%20Strictly%20Palindromic%20Number.go)
70.   [`2610. Convert an Array Into a 2D Array With Conditions`](./Golang/Leetcode%202610%20Convert%20an%20Array%20Into%20a%202D%20Array%20With%20Conditions.go) : Medium
71.   [`380. Insert Delete GetRandom O(1)`](./Golang/Leetcode%20380.%20Insert%20Delete%20GetRandom%20O(1).go) : Medium
72.   [`46. Permutations`](./Golang/Leetcode%2046.%20Permutations.go) : Medium (Recursion)
73.   [`1481. Least Number of Unique Integers after K Removals`](./Golang/Leetcode%201481.%20Least%20Number%20of%20Unique%20Integers%20after%20K%20Removals.go) : Medium O(N) Solution
74.   [`1291. Sequential Digits`](./Golang/Leetcode%201291.%20Sequential%20Digits.go) : Medium

</details>

<details>
	<summary> <strong> Array </strong> </summary>		

1. [`2455. Average Value of Even Numbers That Are Divisible by Three`](./Golang/Leetcode%202455.%20Average%20Value%20of%20Even%20Numbers%20That%20Are%20Divisible%20by%20Three.go)
2. [`3028. Ant on the Boundary`](./Golang/Leetccode%203028.%20Ant%20on%20the%20Boundary.go)
3. [`961. N-Repeated Element in Size 2N Array`](./Golang/Leetcode%20961.%20N-Repeated%20Element%20in%20Size%202N%20Array.go)
4. [`561. Array Partition`](./Golang/Leetcode%20561.%20Array%20Partition.go)
5. [`1313. Decompress Run-Length Encoded List`](./Golang/Leetcode%201313.%20Decompress%20Run-Length%20Encoded%20List.go)
6. [`1614. Maximum Nesting Depth of the Parentheses`](./Golang/Leetcode%201614.%20Maximum%20Nesting%20Depth%20of%20the%20Parentheses.go)
7. [`2089. Find Target Indices After Sorting Array`](./Golang/Leetcode%202089%20Find%20Target%20Indices%20After%20Sorting%20Array.go)
8. [`2974. Minimum Number Game`](./Golang/Leetcode%202974.%20Minimum%20Number%20Game.go)
9. [`2215. Find the Difference of Two Arrays`](./Golang/Leetcode%202215.%20Find%20the%20Difference%20of%20Two%20Arrays.go)
10. [`2798. Number of Employees Who Met the Target`](./Golang/Leetcode%202798%20Number%20of%20Employees%20Who%20Met%20the%20Target.go)
11. [`1431. Kids With the Greatest Number of Candies`](./Golang/Leetcode%201431.%20Kids%20With%20the%20Greatest%20Number%20of%20Candies.go)
12. [`2706. Buy Two Chocolates`](./Golang/Leetcode%202706%20Buy%20Two%20Chocolates.go)
13. [`383. Ransom Note`](./Golang/Leetcode%20383.%20Ransom%20Note.go)
14. [`3000. Maximum Area of Longest Diagonal Rectangle`](./Golang/Leetcode%203000.%20Maximum%20Area%20of%20Longest%20Diagonal%20Rectangle.go)
15. [`191. Number of 1 Bits`](./Golang/Leetcode%20191.%20Number%20of%201%20Bits.go)
16. [`2864. Maximum Odd Binary Number`](./Golang/Leetcode%202864.%20Maximum%20Odd%20Binary%20Number.go)
17. [`2859. Sum of Values at Indices With K Set Bits`](./Golang/Leetcode%202859.%20Sum%20of%20Values%20at%20Indices%20With%20K%20Set%20Bits.go)
18. [`1672. Richest Customer Wealth`](./Golang/Leetcode%201672%20Richest%20Customer%20Wealth.go)
19. [`2441. Largest Positive Integer That Exists With Its Negative`](./Golang/Leetcode%202441%20Largest%20Positive%20Integer%20That%20Exists%20With%20Its%20Negative.go)
20. [`2544. Alternating Digit Sum`](./Golang/Leetcode%202544%20Alternating%20Digit%20Sum.go)
21. [`1720. Decode XORed Array`](./Golang/Leetcode%201720.%20Decode%20XORed%20Array.go)
22. [`268. Missing Number`](./Golang/Leetcode%20268.%20Missing%20Number.go)
23. [`2965. Find Missing and Repeated Values`](./Golang/Leetcode%202965.%20Find%20Missing%20and%20Repeated%20Values.go)
24. [`1207. Unique Number of Occurrences`](./Golang/Leetcode%201207.%20Unique%20Number%20of%20Occurrences.go)
25. [`2574. Left and Right Sum Differences`](./Golang/Leetcode%202574.%20Left%20and%20Right%20Sum%20Differences.go)
26. [`455. Assign Cookies`](./Golang/Leetcode%20455%20Assign%20Cookies.go)
27. [`3005. Count Elements With Maximum Frequency`](./Golang/Leetcode%203005.%20Count%20Elements%20With%20Maximum%20Frequency.go)
28. [`896. Monotonic Array`](./Golang/Leetcode%20896.%20Monotonic%20Array.go)
29. [`977. Squares of a Sorted Array`](./Golang/Leetcode%20977.%20Squares%20of%20a%20Sorted%20Array.go)
30. [`1385. Find the Distance Value Between Two Arrays](./Golang/Leetcode%201385.%20Find%20the%20Distance%20Value%20Between%20Two%20Arrays.go)
31. [`121. Best Time to Buy and Sell Stock`](./Golang/Leetcode%20121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock.go)
32. [`2475. Number of Unequal Triplets in Array`](./Golang/Leetcode%202475%20Number%20of%20Unequal%20Triplets%20in%20Array.go)
33. [`1913. Maximum Product Difference Between Two Pairs`](./Golang/Leetcode%201913%20Maximum%20Product%20Difference%20Between%20Two%20Pairs.go)
34. [`2176. Count Equal and Divisible Pairs in an Array`](./Golang/Leetcode%202176%20Count%20Equal%20and%20Divisible%20Pairs%20in%20an%20Array.go)
35. [`26. Remove Duplicates from Sorted Array`](./Golang/Leetcode%2026%20Remove%20Duplicates%20from%20Sorted%20Array.go)
36. [`2540. Minimum Common Value`](./Golang/Leetcode%202540.%20Minimum%20Common%20Value.go) : Two pointer approach
37. [`349. Intersection of Two Arrays`](./Golang/Leetcode%20349.%20Intersection%20of%20Two%20Arrays.go)
38. [`350. Intersection of Two Arrays II`](./Golang/Leetcode%20350.%20Intersection%20of%20Two%20Arrays%20II.go)
39. [`643. Maximum Average Subarray I`](./Golang/Leetcode%20643.%20Maximum%20Average%20Subarray%20I.go)
40. [`1089. Duplicate Zeros`](./Golang/Leetcode%201089.%20Duplicate%20Zeros.go):  Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
41. [`2006. Count Number of Pairs With Absolute Difference K`](./Golang/Leetcode%202006%20Count%20Number%20of%20Pairs%20With%20Absolute%20Difference%20K.go)
42. [`628. Maximum Product of Three Numbers`](./Golang/Leetcode%20628%20Maximum%20Product%20of%20Three%20Numbers.go)
43. [`1710. Maximum Units on a Truck`](./Golang/Leetcode%201710.%20Maximum%20Units%20on%20a%20Truck.go)
44. [`66. Plus One`](./Golang/Leetcode%2066%20Plus%20One.go)
45. [`2433. Find The Original Array of Prefix Xor`](./Golang/Leetcode%202433%20Find%20The%20Original%20Array%20of%20Prefix%20Xor.go)
46. [`2824. Count Pairs Whose Sum is Less than Target`](./Golang/Leetcode%202824%20Count%20Pairs%20Whose%20Sum%20is%20Less%20than%20Target.go)
47. [`1588. Sum of All Odd Length Subarrays`](./Golang/Leetcode%201588%20Sum%20of%20All%20Odd%20Length%20Subarrays.go)
48. [`3090. Maximum Length Substring With Two Occurrences](./Golang/Leetcode%201588%20Sum%20of%20All%20Odd%20Length%20Subarrays.go)
49. [`442. Find All Duplicates in an Array`](./Golang/Leetcode%20442.%20Find%20All%20Duplicates%20in%20an%20Array.go) : Medium - Easy level
50. [`2125. Number of Laser Beams in a Bank`](./Golang/Leetcode%202125%20Number%20of%20Laser%20Beams%20in%20a%20Bank.go) : Medium - Easy level
51. [`2870. Minimum Number of Operations to Make Array Empty`](./Golang/Leetcode%202870%20Minimum%20Number%20of%20Operations%20to%20Make%20Array%20Empty.go) : Medium - Easy level
52. [`2396. Strictly Palindromic Number.go`](./Golang/Leetcode%202396.%20Strictly%20Palindromic%20Number.go)
53. [`2610. Convert an Array Into a 2D Array With Conditions`](./Golang/Leetcode%202610%20Convert%20an%20Array%20Into%20a%202D%20Array%20With%20Conditions.go) : Medium
54. [`380. Insert Delete GetRandom O(1)`](./Golang/Leetcode%20380.%20Insert%20Delete%20GetRandom%20O(1).go) : Medium
55. [`46. Permutations`](./Golang/Leetcode%2046.%20Permutations.go) : Medium (Recursion)
56. [`1481. Least Number of Unique Integers after K Removals`](./Golang/Leetcode%201481.%20Least%20Number%20of%20Unique%20Integers%20after%20K%20Removals.go) : Medium O(N) Solution
</details>

<details>
	<summary> <strong> Strings </strong> </summary>	
	
1. [`1108. Defanging an IP Address`](./Golang/Leetcode%201108%20Defanging%20an%20IP%20Address%20Golang%20Solution.go)
2. [`657. Robot Return to Origin`](./Golang/Leetcode%20657.%20Robot%20Return%20to%20Origin.go)
3. [`2833. Furthest Point From Origin`](./Golang/Leetcode%202833.%20Furthest%20Point%20From%20Origin.go) : You can use if else condition if didn't know hashmaps
4. [`2351. First Letter to Appear Twice`](./Golang/Leetcode%202351%20First%20Letter%20to%20Appear%20Twice.go)
5. [`387. First Unique Character in a String`](./Golang/Leetcode%20387.%20First%20Unique%20Character%20in%20a%20String.go)
6. [`383. Ransom Note`](./Golang/Leetcode%20383.%20Ransom%20Note.go)
7. [`1704. Determine if String Halves Are Alike`](./Golang/Leetcode%201704.%20Determine%20if%20String%20Halves%20Are%20Alike.go)
8. [`2108. Find First Palindromic String in the Array`](./Golang/Leetcode%202108.%20Find%20First%20Palindromic%20String%20in%20the%20Array.go)
9. [`744. Find Smallest Letter Greater Than Target`](./Golang/Leetcode%20744%20Find%20Smallest%20Letter%20Greater%20Than%20Target.go)
10. [`1816. Truncate Sentence`](./Golang/Leetcode%201816.%20Truncate%20Sentence.go)
11. [`1528. Shuffle String`](./Golang/Leetcode%201528.%20Shuffle%20String.go)
12. [`191. Number of 1 Bits`](./Golang/Leetcode%20191.%20Number%20of%201%20Bits.go)
13. [`1773. Count Items Matching a Rule`](./Golang/Leetcode%201773.%20Count%20Items%20Matching%20a%20Rule.go)
14. [`2114. Maximum Number of Words Found in Sentences`](./Golang/Leetcode%202114.%20Maximum%20Number%20of%20Words%20Found%20in%20Sentences.go)
15. [`1662. Check If Two String Arrays are Equivalent`](./Golang/Leetcode%201662.%20Check%20If%20Two%20String%20Arrays%20are%20Equivalent.go)
16. [`1678. Goal Parser Interpretation`](./Golang/Leetcode%201678%20Goal%20Parser%20Interpretation.go)
17. [`2273. Find Resultant Array After Removing Anagrams`](./Golang/Leetcode%202273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams.go)
18. [`2828. Check if a String Is an Acronym of Words`](./Golang/Leetcode%202828%20Check%20if%20a%20String%20Is%20an%20Acronym%20of%20Words.go)
19. [`2942. Find Words Containing Character`](./Golang/Leetcode%202942%20Find%20Words%20Containing%20Character.go)
20. [`1624. Largest Substring Between Two Equal Characters`](./Golang/Leetcode%201624%20Largest%20Substring%20Between%20Two%20Equal%20Characters.go)
21. [`1689. Partitioning Into Minimum Number Of Deci-Binary Numbers`](./Golang/Leetcode%201689%20Partitioning%20Into%20Minimum%20Number%20Of%20Deci-Binary%20Numbers.go)
22. [`3090. Maximum Length Substring With Two Occurrences](./Golang/Leetcode%201588%20Sum%20of%20All%20Odd%20Length%20Subarrays.go)
23. [`1347. Minimum Number of Steps to Make Two Strings Anagram`](./Golang/Leetcode%201347.%20Minimum%20Number%20of%20Steps%20to%20Make%20Two%20Strings%20Anagram.go): Medium - Easy
24. [`2186. Minimum Number of Steps to Make Two Strings Anagram II`](./Golang/Leetcode%202186.%20Minimum%20Number%20of%20Steps%20to%20Make%20Two%20Strings%20Anagram%20II.go): Medium
25. [`1657. Determine if Two Strings Are Close`](./Golang/Leetcode%201657.%20Determine%20if%20Two%20Strings%20Are%20Close.go): Medium
26. [`567. Permutation in String`](./Golang/Leetcode%20567.%20Permutation%20in%20String.go) : Sliding Window Approach
27. [`438. Find All Anagrams in a String`](./Golang/Leetcode%20438.%20Find%20All%20Anagrams%20in%20a%20String.go) : Sliding Window
</details>

<details>
	<summary> <strong> Hash Maps / Hash Table / Set /Dictionary </strong> </summary>	
	
1. [`1. Two Sum`](./Golang/Leetcode%201%20Two%20Sum.go)
2. [`217. Contains Duplicate`](./Golang/Leetcode%20217%20Contains%20Duplicate.go): Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
3. [`961. N-Repeated Element in Size 2N Array`](./Golang/Leetcode%20961.%20N-Repeated%20Element%20in%20Size%202N%20Array.go)
4. [`2833. Furthest Point From Origin`](./Golang/Leetcode%202833.%20Furthest%20Point%20From%20Origin.go)
5. [`1748. Sum of Unique Elements`](./Golang/Leetcode%201748%20Sum%20of%20Unique%20Elements.go)
6. [`1207. Unique Number of Occurrences`](./Golang/Leetcode%201207.%20Unique%20Number%20of%20Occurrences.go)
7. [`2351. First Letter to Appear Twice`](./Golang/Leetcode%202351%20First%20Letter%20to%20Appear%20Twice.go)
8. [`387. First Unique Character in a String`](./Golang/Leetcode%20387.%20First%20Unique%20Character%20in%20a%20String.go)
9. [`2215. Find the Difference of Two Arrays`](./Golang/Leetcode%202215.%20Find%20the%20Difference%20of%20Two%20Arrays.go)
10. [`1941. Check if All Characters Have Equal Number of Occurrences`](./Golang/Leetcode%201941%20Check%20if%20All%20Characters%20Have%20Equal%20Number%20of%20Occurrences.go)
11. [`287. Find the Duplicate Number`](./Golang/Leetcode%20287%20Find%20the%20Duplicate%20Number.go)
12. [`2154. Keep Multiplying Found Values by Two`](./Golang/Leetcode%202154.%20Keep%20Multiplying%20Found%20Values%20by%20Two.go)
13. [`575. Distribute Candies`](./Golang/Leetcode%20575%20Distribute%20Candies.go)
14. [`3005. Count Elements With Maximum Frequency`](./Golang/Leetcode%203005.%20Count%20Elements%20With%20Maximum%20Frequency.go)
15. [`1512. Number of Good Pairs`](./Golang/Leetcode%201512%20Number%20of%20Good%20Pairs.go)
16. [`169. Majority Element`](./Golang/Leetcode%20169%20Majority%20Element.go)
17. [`2190. Most Frequent Number Following Key In an Array`](./Golang/Leetcode%202190.%20Most%20Frequent%20Number%20Following%20Key%20In%20an%20Array.go)
18. [`383. Ransom Note`](./Golang/Leetcode%20383.%20Ransom%20Note.go)
19. [`1624. Largest Substring Between Two Equal Characters`](./Golang/Leetcode%201624%20Largest%20Substring%20Between%20Two%20Equal%20Characters.go)
20. [`205. Isomorphic Strings`](./Golang/Leetcode%20205%20Isomorphic%20Strings.go)
21. [`242. Valid Anagram`](./Golang/Leetcode%20242%20Valid%20Anagram.go)
22. [`1832. Check if the Sentence Is Pangram`](./Golang/Leetcode%201832%20Check%20if%20the%20Sentence%20Is%20Pangram.go)
23. [`771. Jewels and Stones`](./Golang/Leetcode%20771%20Jewels%20and%20Stones.go)
24. [`202. Happy Number`](./Golang/Leetcode%20202%20Happy%20Number.go)
25. [`2965. Find Missing and Repeated Values`](./Golang/Leetcode%202965.%20Find%20Missing%20and%20Repeated%20Values.go)
26. [`1282. Group the People Given the Group Size They Belong To`](./Golang/Leetcode%201282%20Group%20the%20People%20Given%20the%20Group%20Size%20They%20Belong%20To.go)
27. [`349. Intersection of Two Arrays`](./Golang/Leetcode%20349.%20Intersection%20of%20Two%20Arrays.go)
28. [`350. Intersection of Two Arrays II`](./Golang/Leetcode%20350.%20Intersection%20of%20Two%20Arrays%20II.go)
29. [`2357. Make Array Zero by Subtracting Equal Amounts`](./Golang/Leetcode%202357%20Make%20Array%20Zero%20by%20Subtracting%20Equal%20Amounts.go)
30. [`1370. Increasing Decreasing String`](./Golang/Leetcode%201370%20Increasing%20Decreasing%20String.go)
31. [`2367. Number of Arithmetic Triplets`](./Golang/Leetcode%202367%20Number%20of%20Arithmetic%20Triplets.go)
32. [`219. Contains Duplicate II`](./Golang/Leetcode%20219.%20Contains%20Duplicate%20II.go)
33. [`2404. Most Frequent Even Element`](./Golang/Leetcode%202404.%20Most%20Frequent%20Even%20Element.go)
34. [`3090. Maximum Length Substring With Two Occurrences](./Golang/Leetcode%201588%20Sum%20of%20All%20Odd%20Length%20Subarrays.go)
35. [`442. Find All Duplicates in an Array`](./Golang/Leetcode%20442.%20Find%20All%20Duplicates%20in%20an%20Array.go) : Medium - Easy level
36. [`1347. Minimum Number of Steps to Make Two Strings Anagram`](./Golang/Leetcode%201347.%20Minimum%20Number%20of%20Steps%20to%20Make%20Two%20Strings%20Anagram.go): Medium - Easy
37. [`2186. Minimum Number of Steps to Make Two Strings Anagram II`](./Golang/Leetcode%202186.%20Minimum%20Number%20of%20Steps%20to%20Make%20Two%20Strings%20Anagram%20II.go): Medium
38. [`1657. Determine if Two Strings Are Close`](./Golang/Leetcode%201657.%20Determine%20if%20Two%20Strings%20Are%20Close.go): Medium
39. [`380. Insert Delete GetRandom O(1)`](./Golang/Leetcode%20380.%20Insert%20Delete%20GetRandom%20O(1).go) : Medium
40. [`49. Group Anagrams`](./Golang/Leetcode%2049.%20Group%20Anagrams.go) : Medium
</details>

<details>
	<summary> <strong> Linked List </strong> </summary>	
	
1. [`1290. Convert Binary Number in a Linked List to Integer`](./Golang/Leetcode%201290%20Convert%20Binary%20Number%20in%20a%20Linked%20List%20to%20Integer.go):  Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number. Return the decimal value of the number in the linked list.
2. [`876. Middle of the Linked List`](./Golang/Leetcode%20876%20Middle%20of%20the%20Linked%20List.go): Given the head of a singly linked list, return the middle node of the linked list. If there are two middle nodes, return the second middle node.
3. [`206. Reverse Linked List`](./Golang/Leetcode%20206.%20Reverse%20Linked%20List.go)
4. [`234. Palindrome Linked List`](./Golang/Leetcode%20234.%20Palindrome%20Linked%20List.go)
5. [`160. Intersection of Two Linked Lists`](./Golang/Leetcode%20160%20Intersection%20of%20Two%20Linked%20Lists.go): Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.
6. [`141. Linked List Cycle`](./Golang/Leetcode%20141%20Linked%20List%20Cycle.go): Given head, the head of a linked list, determine if the linked list has a cycle in it.
7. [`19. Remove Nth Node From End of List`](./Golang/Leetcode%2019%20Remove%20Nth%20Node%20From%20End%20of%20List.go): Given the head of a linked list, remove the nth node from the end of the list and return its head.
8. [`2095. Delete the Middle Node of a Linked List`](./Golang/Leetcode%202095%20Delete%20the%20Middle%20Node%20of%20a%20Linked%20List.go): You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
9.  [`2807. Insert Greatest Common Divisors in Linked List`](./Golang/Leetcode%202807%20Insert%20Greatest%20Common%20Divisors%20in%20Linked%20List.go) : Medium Question but Medium - Easy level
10. [`707. Design Linked List`](./Golang/Leetcode%20707%20Design%20Linked%20List.go): (Medium) Design your implementation of the linked list.
</details>


<details>
<summary> <strong> SQL </strong> </summary>	
	
1. [`1757. Recyclable and Low Fat Products`](./SQL/1757.%20Recyclable%20and%20Low%20Fat%20Products.sql)
2. [`584. Find Customer Referee`](./SQL/584.%20Find%20Customer%20Referee.go)
3. [`595. Big Countries`](./SQL/595.%20Big%20Countries.sql)
4. [`1148. Article Views I`](./SQL/1148.%20Article%20Views%20I.sql)
5. [`1683. Invalid Tweets`](./SQL/1683.%20Invalid%20Tweets.sql)
6. [`1378. Replace Employee ID With The Unique Identifier`](./SQL/Leetcode%201378.%20Replace%20Employee%20ID%20With%20The%20Unique%20Identifier.sql)
7. [`1068. Product Sales Analysis I`](./SQL/1068.%20Product%20Sales%20Analysis%20I.sql)
8. [`2356. Number of Unique Subjects Taught by Each Teacher`](./SQL/2356.%20Number%20of%20Unique%20Subjects%20Taught%20by%20Each%20Teacher.sql)
9. [`1581. Customer Who Visited but Did Not Make Any Transactions`](./SQL/1581.%20Customer%20Who%20Visited%20but%20Did%20Not%20Make%20Any%20Transactions.sql)
</details>


<details>
	<summary> <strong> Sorting </strong> </summary>	
	
1. [`1089. Duplicate Zeros`](./Golang/Leetcode%201089.%20Duplicate%20Zeros.go):  Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
</details>

<details>
	<summary> <strong> Recursion </strong> </summary>	
	
1. [`144 Binary Tree Preorder Traversal`](./Golang/Leetcode%20144%20Binary%20Tree%20Preorder%20Traversal.go)
2. [`94 Binary Tree Inorder Traversal`](./Golang/Leetcode%2094%20Binary%20Tree%20Inorder%20Traversal.go)
3. [`145 Binary Tree Postorder Traversal`](./Golang/Leetcode%20145%20Binary%20Tree%20Postorder%20Traversal.go)
4. [`231. Power of Two`](./Golang/Leetcode%20231.%20Power%20of%20Two.go)
5. [`326. Power of Three`](./Golang/Leetcode%20326.%20Power%20of%20Three.go)
6. [`342. Power of Four`](./Golang/Leetcode%20342.%20Power%20of%20Four.go)
7. [`46. Permutations`](./Golang/Leetcode%2046.%20Permutations.go) : Medium (Recursion)
8. [`1302. Deepest Leaves Sum`](./Golang/Leetcode%201302.%20Deepest%20Leaves%20Sum.go) : Medium
   
</details>


<details>
	<summary> <strong> Tree </strong> </summary>	
	
1. [`144 Binary Tree Preorder Traversal`](./Golang/Leetcode%20144%20Binary%20Tree%20Preorder%20Traversal.go)
2. [`94 Binary Tree Inorder Traversal`](./Golang/Leetcode%2094%20Binary%20Tree%20Inorder%20Traversal.go)
3. [`145 Binary Tree Postorder Traversal`](./Golang/Leetcode%20145%20Binary%20Tree%20Postorder%20Traversal.go)
4. [`938. Range Sum of BST`](./Golang/Leetcode%20938%20Range%20Sum%20of%20BST.go)
5. [`872. Leaf-Similar Trees`](./Golang/Leetcode%20872%20Leaf-Similar%20Trees.go)
6. [`1302. Deepest Leaves Sum`](./Golang/Leetcode%201302.%20Deepest%20Leaves%20Sum.go) : Medium
</details>

<details>
	<summary> <strong> Stack </strong> </summary>	
	
1. [`1089. Duplicate Zeros`](./Golang/Leetcode%201089.%20Duplicate%20Zeros.go):  Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
</details>

<details>
	<summary> <strong> Queue </strong> </summary>	
	
1. [`1089. Duplicate Zeros`](./Golang/Leetcode%201089.%20Duplicate%20Zeros.go):  Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
</details>

<details>
	<summary> <strong> Heap </strong> </summary>	
	
1. [`1089. Duplicate Zeros`](./Golang/Leetcode%201089.%20Duplicate%20Zeros.go):  Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
</details>

<details>
	<summary> <strong> Trie </strong> </summary>	
	
1. [`1089. Duplicate Zeros`](./Golang/Leetcode%201089.%20Duplicate%20Zeros.go):  Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
</details>
