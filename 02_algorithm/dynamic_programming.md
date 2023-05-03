# dynamic programming 动态规划
![](.dynamic_programming_images/dp_info.png)
![](.dynamic_programming_images/dp_info2.png)

优化前代码:递归自顶向下
![](.dynamic_programming_images/dp_code1.png)

优化后代码：递推自底向上
![](.dynamic_programming_images/dp_code2.png)
节省空间：交替滚动
![](.dynamic_programming_images/dp_code3.png)


## 应用
![](.dynamic_programming_images/dp_application.png)

1. 编辑距离-->转成插入insert,删除delete，替换switch
![](.dynamic_programming_images/edit_distance.png)

2. 爬台阶
![](.dynamic_programming_images/climb_stair.png)
要么n-1台阶上去，要么n-2台阶上去 

3. 找零钱

![](.dynamic_programming_images/get_change1.png)    

分治递归
![](.dynamic_programming_images/get_change2.png)
![](.dynamic_programming_images/get_change3.png)

递归实现  
![](.dynamic_programming_images/get_change_code.png)

递推实现
![](.dynamic_programming_images/get_change_code2.png)


4. 机器人走法
![](.dynamic_programming_images/robot_walk.png)
![](.dynamic_programming_images/robot_code.png)

递归算法：优化代码
![](.dynamic_programming_images/robot_code1.png)

递推算法
![](.dynamic_programming_images/robot_code2.png)
![](.dynamic_programming_images/robot_code3.png)
![](.dynamic_programming_images/robot_code4.png)