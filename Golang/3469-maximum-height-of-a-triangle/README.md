<h2><a href="https://leetcode.com/problems/maximum-height-of-a-triangle">Maximum Height of a Triangle</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>You are given two integers <code>red</code> and <code>blue</code> representing the count of red and blue colored balls. You have to arrange these balls to form a triangle such that the 1<sup>st</sup> row will have 1 ball, the 2<sup>nd</sup> row will have 2 balls, the 3<sup>rd</sup> row will have 3 balls, and so on.</p>

<p>All the balls in a particular row should be the <strong>same</strong> color, and adjacent rows should have <strong>different</strong> colors.</p>

<p>Return the <strong>maximum</strong><em> height of the triangle</em> that can be achieved.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">red = 2, blue = 4</span></p>

<p><strong>Output:</strong> 3</p>

<p><strong>Explanation:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/16/brb.png" style="width: 300px; height: 240px; padding: 10px;" /></p>

<p>The only possible arrangement is shown above.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">red = 2, blue = 1</span></p>

<p><strong>Output:</strong> <span class="example-io">2</span></p>

<p><strong>Explanation:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/16/br.png" style="width: 150px; height: 135px; padding: 10px;" /><br />
The only possible arrangement is shown above.</p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">red = 1, blue = 1</span></p>

<p><strong>Output:</strong> <span class="example-io">1</span></p>
</div>

<p><strong class="example">Example 4:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">red = 10, blue = 1</span></p>

<p><strong>Output:</strong> <span class="example-io">2</span></p>

<p><strong>Explanation:</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/16/br.png" style="width: 150px; height: 135px; padding: 10px;" /><br />
The only possible arrangement is shown above.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= red, blue &lt;= 100</code></li>
</ul>
