<h2><a href="https://leetcode.com/problems/ant-on-the-boundary">Ant on the Boundary</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>An ant is on a boundary. It sometimes goes <strong>left</strong> and sometimes <strong>right</strong>.</p>

<p>You are given an array of <strong>non-zero</strong> integers <code>nums</code>. The ant starts reading <code>nums</code> from the first element of it to its end. At each step, it moves according to the value of the current element:</p>

<ul>
	<li>If <code>nums[i] &lt; 0</code>, it moves <strong>left</strong> by<!-- notionvc: 55fee232-4fc9-445f-952a-f1b979415864 --> <code>-nums[i]</code> units.</li>
	<li>If <code>nums[i] &gt; 0</code>, it moves <strong>right</strong> by <code>nums[i]</code> units.</li>
</ul>

<p>Return <em>the number of times the ant <strong>returns</strong> to the boundary.</em></p>

<p><strong>Notes:</strong></p>

<ul>
	<li>There is an infinite space on both sides of the boundary.</li>
	<li>We check whether the ant is on the boundary only after it has moved <code>|nums[i]|</code> units. In other words, if the ant crosses the boundary during its movement, it does not count.<!-- notionvc: 5ff95338-8634-4d02-a085-1e83c0be6fcd --></li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,3,-5]
<strong>Output:</strong> 1
<strong>Explanation:</strong> After the first step, the ant is 2 steps to the right of the boundary<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->.
After the second step, the ant is 5 steps to the right of the boundary<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->.
After the third step, the ant is on the boundary.
So the answer is 1.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,2,-3,-4]
<strong>Output:</strong> 0
<strong>Explanation:</strong> After the first step, the ant is 3 steps to the right of the boundary<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->.
After the second step, the ant is 5 steps to the right of the boundary<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->.
After the third step, the ant is 2 steps to the right of the boundary<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->.
After the fourth step, the ant is 2 steps to the left of the boundary<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->.
The ant never returned to the boundary, so the answer is 0.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>-10 &lt;= nums[i] &lt;= 10</code></li>
	<li><code>nums[i] != 0</code></li>
</ul>
