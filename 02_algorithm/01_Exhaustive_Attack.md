<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Exhaustive Search穷举法(Brute-Force蛮力法)](#exhaustive-search%E7%A9%B7%E4%B8%BE%E6%B3%95brute-force%E8%9B%AE%E5%8A%9B%E6%B3%95)
  - [特点](#%E7%89%B9%E7%82%B9)
  - [应用](#%E5%BA%94%E7%94%A8)
  - [解决方式](#%E8%A7%A3%E5%86%B3%E6%96%B9%E5%BC%8F)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Exhaustive Search穷举法(Brute-Force蛮力法)

## 特点
![](.01_Exhaustive_Attack_images/advs_n_dis.png)

## 应用
1. 排序
2. 顺序查找
3. 字符串匹配
4. 最近点对问题
![](.01_Exhaustive_Attack_images/linear_position.png)
5. 凸包问题
![](.01_Exhaustive_Attack_images/tubao.png)
![](.01_Exhaustive_Attack_images/tubao1.png)
![](.01_Exhaustive_Attack_images/tubao2.png)

6. traveling salesman problem 
![](.01_Exhaustive_Attack_images/tsp.png)
![](.01_Exhaustive_Attack_images/hamilton.png)
![](.01_Exhaustive_Attack_images/hamilton2.png)
![](.01_Exhaustive_Attack_images/bf_tsp.png)
![](.01_Exhaustive_Attack_images/bf_tsp2.png)
![](.01_Exhaustive_Attack_images/strategy_tree.png)
![](.01_Exhaustive_Attack_images/bf_tsp3.png)

7. 背包问题--同理子集问题
![](.01_Exhaustive_Attack_images/bag.png)
![](.01_Exhaustive_Attack_images/bag_bf.png)
![](.01_Exhaustive_Attack_images/knapSack_code.png)

8. 迷宫问题
![](.01_Exhaustive_Attack_images/maze.png)

9. N皇后问题
![](.01_Exhaustive_Attack_images/n_queen.png)
![](.01_Exhaustive_Attack_images/n_queen1.png)

## 解决方式
多步决策-->空间状态搜索
![](.01_Exhaustive_Attack_images/make_decision.png)
![](.01_Exhaustive_Attack_images/state_search.png)
- 线型
![](.01_Exhaustive_Attack_images/state_space.png)
- 树型
![](.01_Exhaustive_Attack_images/tree_state.png)
- 图形
![](.01_Exhaustive_Attack_images/graph_state.png)

处理方式
![](.01_Exhaustive_Attack_images/bfs_dfs.png)
1. DFS 深度优先搜索（回溯法）
![](.01_Exhaustive_Attack_images/dfs.png)