<!--
  菜单管理页面
  菜单树形结构展示和管理
-->
<template>
  <div class="menu-container">
    <el-card>
      <div class="card-header">
        <span>菜单列表</span>
        <el-button type="primary" @click="handleAdd">
          <el-icon><Plus /></el-icon>
          新增菜单
        </el-button>
      </div>
      
      <el-table
        :data="menuList"
        row-key="menuId"
        border
        stripe
        default-expand-all
      >
        <el-table-column prop="menuName" label="菜单名称" width="200" />
        <el-table-column prop="icon" label="图标" width="100" />
        <el-table-column prop="orderNum" label="排序" width="80" />
        <el-table-column prop="perms" label="权限标识" />
        <el-table-column prop="path" label="路由地址" />
        <el-table-column prop="component" label="组件路径" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 0 ? 'success' : 'danger'">
              {{ row.status === 0 ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="{ row }">
            <el-button type="primary" link>编辑</el-button>
            <el-button type="danger" link>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const menuList = ref([])

const handleAdd = () => {
  console.log('新增菜单')
}
</script>

<style lang="scss" scoped>
.menu-container {
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
}
</style>
